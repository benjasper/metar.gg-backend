package importer

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/cnf/structhash"
	"golang.org/x/sync/errgroup"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/skycondition"
	"metar.gg/logging"
	"metar.gg/utils"
	"os"
	"time"
)

type XmlQualityControlFlags struct {
	Auto                    bool `xml:"auto"`
	Corrected               bool `xml:"corrected"`
	MaintenanceIndicatorOn  bool `xml:"maintenance_indicator"`
	NoSignal                bool `xml:"no_signal"`
	LightningSensorOff      bool `xml:"lightning_sensor_off"`
	FreezingRainSensorOff   bool `xml:"freezing_rain_sensor_off"`
	PresentWeatherSensorOff bool `xml:"present_weather_sensor_off"`
}

type XmlSkyCondition struct {
	Coverage string `xml:"sky_cover,attr"`
	Altitude *int   `xml:"cloud_base_ft_agl,attr"`
}

type XmlMetar struct {
	RawText                   string                 `xml:"raw_text"`
	StationId                 string                 `xml:"station_id"`
	ObservationTime           time.Time              `xml:"observation_time"`
	Latitude                  *float64               `xml:"latitude"`
	Longitude                 *float64               `xml:"longitude"`
	TempC                     *float64               `xml:"temp_c"`
	DewpointC                 *float64               `xml:"dewpoint_c"`
	WindDirDegrees            *int                   `xml:"wind_dir_degrees"`
	WindSpeedKt               *int                   `xml:"wind_speed_kt"`
	WindGustKt                *int                   `xml:"wind_gust_kt"`
	VisibilityStatuteMi       *float64               `xml:"visibility_statute_mi"`
	AltimeterInHg             *float64               `xml:"altim_in_hg"`
	SeaLevelPressureMb        *float64               `xml:"sea_level_pressure_mb"`
	QualityControlFlags       XmlQualityControlFlags `xml:"quality_control_flags"`
	WxString                  *string                `xml:"wx_string"`
	SkyCondition              []XmlSkyCondition      `xml:"sky_condition"`
	FlightCategory            string                 `xml:"flight_category"`
	ThreeHrPressureTendencyMb *float64               `xml:"three_hr_pressure_tendency_mb"`
	MaxTempC                  *float64               `xml:"maxT_c"`
	MinTempC                  *float64               `xml:"minT_c"`
	MaxTemp24HrC              *float64               `xml:"maxT24hr_c"`
	MinTemp24HrC              *float64               `xml:"minT24hr_c"`
	PrecipIn                  *float64               `xml:"precip_in"`
	Precip3HrIn               *float64               `xml:"precip_3hr_in"`
	Precip6HrIn               *float64               `xml:"precip_6hr_in"`
	Precip24HrIn              *float64               `xml:"precip24hr_in"`
	SnowIn                    *float64               `xml:"snow_in"`
	VertVisFt                 *float64               `xml:"vert_vis_ft"`
	MetarType                 string                 `xml:"metar_type"`
	Elevation                 *float64               `xml:"elevation_m"`
}

func (x *XmlMetar) Hash() string {
	return fmt.Sprintf("%x", structhash.Md5(x, 1))
}

type NoaaMetarImporter struct {
	db     *ent.Client
	logger *logging.Logger
	stats  *ImportStatistics
}

func NewNoaaMetarImporter(db *ent.Client, logger *logging.Logger) *NoaaMetarImporter {
	return &NoaaMetarImporter{
		db:     db,
		logger: logger,
		stats:  NewImportStatistics("METAR", logger),
	}
}

func (i *NoaaMetarImporter) ImportMetars(url string, ctx context.Context) error {
	i.stats.Start()

	filepath := "metars.xml"
	err := utils.DownloadFile(url, filepath)
	if err != nil {
		return err
	}

	// Read xml file and parse it
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	maxGoroutines := 4
	guard := make(chan struct{}, maxGoroutines)

	wg := errgroup.Group{}

	// Parse xml file
	decoder := xml.NewDecoder(f)
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch se := token.(type) {

		// We have the start of an element.
		// However, we have the complete token in t
		case xml.StartElement:
			switch se.Name.Local {
			case "METAR":
				var metar XmlMetar
				err = decoder.DecodeElement(&metar, &se)
				if err != nil {
					return err
				}

				guard <- struct{}{} // would block if guard channel is already filled
				i.stats.AddTotal()
				wg.Go(func() error {
					defer func() { <-guard }()
					return i.importMetar(&metar, ctx)
				})
			}
		}
	}

	err = wg.Wait()
	if err != nil {
		return err
	}

	// Delete file
	err = os.Remove(filepath)
	if err != nil {
		return err
	}

	i.stats.End()

	return nil
}

func (i *NoaaMetarImporter) importMetar(x *XmlMetar, ctx context.Context) error {
	a, _ := i.db.Airport.Query().Where(airport.Identifier(x.StationId)).Only(ctx)
	if x.StationId == "" {
		i.logger.Error(fmt.Sprintf("Could not find airport with identifier %s and metar %s", x.StationId, x.RawText))
		return nil
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	hash := x.Hash()

	// Check if m already exists
	_, err := i.db.Metar.Query().Where(metar.Hash(hash)).First(ctx)
	if err == nil {
		// Metar already exists
		return nil
	}

	transaction, err := i.db.Tx(ctx)
	if err != nil {
		return err
	}

	var flightCategory metar.FlightCategory
	switch x.FlightCategory {
	case "VFR":
		flightCategory = metar.FlightCategoryVFR
	case "MVFR":
		flightCategory = metar.FlightCategoryMVFR
	case "IFR":
		flightCategory = metar.FlightCategoryIFR
	case "LIFR":
		flightCategory = metar.FlightCategoryLIFR
	}

	if x.MetarType == "" {
		i.logger.Error(fmt.Sprintf("Empty metar type for station %s \n", x.StationId))
		return nil
	}

	metarType := metar.MetarTypeMETAR
	if x.MetarType == "SPECI" {
		metarType = metar.MetarTypeSPECI
	}

	t := transaction.Metar.Create().
		SetStationID(x.StationId).
		SetRawText(x.RawText).
		SetObservationTime(x.ObservationTime).
		SetNillableLatitude(x.Latitude).
		SetNillableLongitude(x.Longitude).
		SetTemperature(utils.Nillable(x.TempC)).
		SetDewpoint(utils.Nillable(x.DewpointC)).
		SetWindDirection(utils.Nillable(x.WindDirDegrees)).
		SetWindSpeed(utils.Nillable(x.WindSpeedKt)).
		SetWindGust(utils.Nillable(x.WindGustKt)).
		SetVisibility(utils.Nillable(x.VisibilityStatuteMi)).
		SetAltimeter(utils.Nillable(x.AltimeterInHg)).
		SetSeaLevelPressure(utils.Nillable(x.SeaLevelPressureMb)).
		SetQualityControlAutoStation(x.QualityControlFlags.Auto).
		SetQualityControlCorrected(x.QualityControlFlags.Corrected).
		SetQualityControlMaintenanceIndicatorOn(x.QualityControlFlags.MaintenanceIndicatorOn).
		SetQualityControlNoSignal(x.QualityControlFlags.NoSignal).
		SetQualityControlLightningSensorOff(x.QualityControlFlags.LightningSensorOff).
		SetQualityControlFreezingRainSensorOff(x.QualityControlFlags.FreezingRainSensorOff).
		SetQualityControlPresentWeatherSensorOff(x.QualityControlFlags.PresentWeatherSensorOff).
		SetNillablePresentWeather(x.WxString).
		SetNillableFlightCategory(utils.NillableWithInput(x.FlightCategory, flightCategory)).
		SetPressureTendency(utils.Nillable(x.ThreeHrPressureTendencyMb)).
		SetNillableMaxTemp6(x.MaxTempC).
		SetNillableMinTemp6(x.MinTempC).
		SetNillableMaxTemp24(x.MaxTemp24HrC).
		SetNillableMinTemp24(x.MinTemp24HrC).
		SetNillablePrecipitation(x.PrecipIn).
		SetNillablePrecipitation3(x.Precip3HrIn).
		SetNillablePrecipitation6(x.Precip6HrIn).
		SetNillablePrecipitation24(x.Precip24HrIn).
		SetNillableSnowDepth(x.SnowIn).
		SetNillableVertVis(x.VertVisFt).
		SetMetarType(metarType).
		SetNillableElevation(x.Elevation).
		SetHash(hash)

	if a != nil {
		t.SetAirportID(a.ID)
	}

	m, err := t.Save(ctx)
	if err != nil {
		return err
	}

	for _, condition := range x.SkyCondition {
		skyCover := skycondition.SkyCoverSkyClear

		switch condition.Coverage {
		case "SKC":
			skyCover = skycondition.SkyCoverSkyClear
			break
		case "CLR":
			skyCover = skycondition.SkyCoverClear
			break
		case "FEW":
			skyCover = skycondition.SkyCoverFew
			break
		case "SCT":
			skyCover = skycondition.SkyCoverScattered
			break
		case "BKN":
			skyCover = skycondition.SkyCoverBroken
			break
		case "OVC":
			skyCover = skycondition.SkyCoverOvercast
			break
		case "OVX":
			skyCover = skycondition.SkyCoverVerticalVisibility
			break
		case "CAVOK":
			skyCover = skycondition.SkyCoverCeilingAndVisibilityOK
			break
		default:
			return fmt.Errorf("unknown sky cover %s", condition.Coverage)
		}

		err = transaction.SkyCondition.Create().
			SetMetar(m).
			SetSkyCover(skyCover).
			SetNillableCloudBase(condition.Altitude).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	if a != nil && a.HasWeather == false {
		err = transaction.Airport.Update().Where(airport.ID(a.ID)).SetHasWeather(true).Exec(ctx)
		if err != nil {
			return err
		}
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	i.stats.AddUpdated()

	return err
}
