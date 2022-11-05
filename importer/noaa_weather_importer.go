package importer

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/segmentio/fasthash/fnv1a"
	"golang.org/x/sync/errgroup"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/forecast"
	"metar.gg/ent/metar"
	"metar.gg/ent/skycondition"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
	"metar.gg/environment"
	"metar.gg/logging"
	"metar.gg/utils"
	"os"
	"strconv"
	"time"
)

type NoaaWeatherImporter struct {
	db     *ent.Client
	logger *logging.Logger
	stats  *ImportStatistics
}

func NewNoaaWeatherImporter(db *ent.Client, logger *logging.Logger) *NoaaWeatherImporter {
	return &NoaaWeatherImporter{
		db:     db,
		logger: logger,
	}
}

func (i *NoaaWeatherImporter) ImportMetars(url string, ctx context.Context) error {
	i.stats = NewImportStatistics("METAR", i.logger)

	i.stats.Start()

	filepath := fmt.Sprintf("metars-%s.xml", time.Now().Format("2006-01-02-15-04-05"))
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

	wg, ctx := errgroup.WithContext(ctx)

	wg.SetLimit(environment.Global.MaxConcurrentImports)

	xmlMetarResponse := XmlMetarResponse{}

	// Parse xml file
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(fileContent, &xmlMetarResponse)
	if err != nil {
		i.logger.Error(fmt.Sprintf("[IMPORT] Failed to parse METARs: %s", err))
		return err
	}

	for _, xmlMetar := range xmlMetarResponse.Data.METAR {
		xmlMetar := xmlMetar

		i.stats.AddTotal()
		wg.Go(func() error {
			return i.importMetar(&xmlMetar, ctx)
		})
	}

	err = wg.Wait()
	if err != nil {
		i.logger.Error(fmt.Sprintf("[IMPORT] Failed to import METARs: %s", err))
	}

	// Delete file
	err = os.Remove(filepath)
	if err != nil {
		return err
	}

	i.stats.End()

	return nil
}

func (i *NoaaWeatherImporter) ImportTafs(url string, ctx context.Context) error {
	i.stats = NewImportStatistics("TAF", i.logger)

	i.stats.Start()

	filepath := fmt.Sprintf("taf-%s.xml", time.Now().Format("2006-01-02-15-04-05"))
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

	wg, ctx := errgroup.WithContext(ctx)

	wg.SetLimit(environment.Global.MaxConcurrentImports)

	xmlTafResponse := XmlTafResponse{}

	// Parse xml file
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(fileContent, &xmlTafResponse)
	if err != nil {
		i.logger.Error(fmt.Sprintf("[IMPORT] Failed to parse TAFs: %s", err))
		return err
	}

	for _, xmlTaf := range xmlTafResponse.Data.TAF {
		xmlTaf := xmlTaf

		i.stats.AddTotal()
		wg.Go(func() error {
			return i.importTaf(&xmlTaf, ctx)
		})
	}

	err = wg.Wait()
	if err != nil {
		i.logger.Error(fmt.Sprintf("[IMPORT] Failed to import TAFs: %s", err))
	}

	// Delete file
	err = os.Remove(filepath)
	if err != nil {
		return err
	}

	i.stats.End()

	return nil
}

func (i *NoaaWeatherImporter) importMetar(x *XmlMetar, ctx context.Context) error {
	s, err := i.getStation(ctx, x.StationId, utils.Nillable(x.Latitude), utils.Nillable(x.Longitude), utils.Nillable(x.Elevation))
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	hash := x.Hash()

	// Check if m already exists
	_, err = i.db.Metar.Query().Where(metar.Hash(hash)).First(ctx)
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

	importTime := time.Now()

	if x.MetarType == "" {
		i.logger.Error(fmt.Sprintf("Empty metar type for station %s \n", x.StationId))
		return nil
	}

	metarType := metar.MetarTypeMETAR
	if x.MetarType == "SPECI" {
		metarType = metar.MetarTypeSPECI
	}

	// Check accuracy of last prediction
	lastMetar, err := i.db.Metar.Query().Where(metar.HasStationWith(weatherstation.StationID(x.StationId))).Order(ent.Desc(metar.FieldImportTime)).First(ctx)
	if err != nil {
		// No last metar found
	}

	if lastMetar != nil && lastMetar.NextImportTimePrediction != nil && *lastMetar.NextImportTimePrediction != (time.Time{}) {
		// Calculate difference between last prediction and actual import time
		diff := importTime.Sub(*lastMetar.NextImportTimePrediction)
		i.logger.CustomEvent(fmt.Sprintf("Import time prediction diff"), "import_time_prediction_diff", map[string]interface{}{
			"station_id":  x.StationId,
			"diff":        diff.String(),
			"diff_in_min": diff.Minutes(),
			"prediction":  lastMetar.NextImportTimePrediction.String(),
			"actual":      importTime.String(),
		})
	}

	// Log time difference between import and observation time
	diff := importTime.Sub(x.ObservationTime)
	i.logger.CustomEvent(fmt.Sprintf("Import observation time diff"), "import_observation_time_diff", map[string]interface{}{
		"station_id":  x.StationId,
		"diff":        diff.String(),
		"diff_in_min": diff.Minutes(),
		"import":      importTime.String(),
		"observation": x.ObservationTime.String(),
	})

	prediction, err := i.MakeNextImportPrediction(ctx, s.StationID, &importTime, &x.ObservationTime)
	if err != nil {
		// Ignore error, because it could be that we don't have enough data for a prediction
		i.logger.Warn(fmt.Sprintf("Failed to make next import prediction: %s", err))
	}

	t := transaction.Metar.Create().
		SetStation(s).
		SetRawText(x.RawText).
		SetImportTime(importTime).
		SetObservationTime(x.ObservationTime).
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
		SetNillableNextImportTimePrediction(prediction).
		SetMetarType(metarType).
		SetHash(hash)

	for _, condition := range x.SkyCondition {
		skyCover, err := getSkyCoverFromString(condition.Coverage)
		if err != nil {
			return err
		}

		sky, err := transaction.SkyCondition.Create().
			SetSkyCover(skyCover).
			SetNillableCloudBase(condition.Altitude).
			Save(ctx)
		if err != nil {
			return err
		}

		t.AddSkyConditions(sky)
	}

	err = t.Exec(ctx)
	if err != nil {
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	i.stats.AddUpdated()

	return err
}

func (i *NoaaWeatherImporter) importTaf(x *XmlTaf, ctx context.Context) error {
	s, err := i.getStation(ctx, x.StationId, utils.Nillable(x.Latitude), utils.Nillable(x.Longitude), utils.Nillable(x.Elevation))
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	hash := x.Hash()

	// Check if t already exists
	_, err = i.db.Taf.Query().Where(taf.Hash(hash)).First(ctx)
	if err == nil {
		// Metar already exists
		return nil
	}

	transaction, err := i.db.Tx(ctx)
	if err != nil {
		return err
	}

	tx := transaction.Taf.Create().
		SetStation(s).
		SetRawText(x.RawText).
		SetIssueTime(x.IssueTime).
		SetBulletinTime(x.BulletinTime).
		SetValidFromTime(x.ValidTimeFrom).
		SetValidToTime(x.ValidTimeTo).
		SetRemarks(x.Remarks).
		SetHash(hash)

	for _, xmlForecast := range x.Forecasts {

		fc := transaction.Forecast.Create().
			SetFromTime(xmlForecast.TimeFrom).
			SetToTime(xmlForecast.TimeTo).
			SetNillableChangeTime(xmlForecast.TimeBecoming).
			SetNillableChangeProbability(xmlForecast.Probability).
			SetNillableWindDirection(xmlForecast.WindDir).
			SetNillableWindSpeed(xmlForecast.WindSpeed).
			SetNillableWindGust(xmlForecast.WindGust).
			SetNillableWindShearHeight(xmlForecast.WindShear).
			SetNillableWindShearDirection(xmlForecast.WindShearDir).
			SetNillableWindShearSpeed(xmlForecast.WindShearSpd).
			SetNillableVisibilityHorizontal(xmlForecast.Visibility).
			SetNillableAltimeter(xmlForecast.Altimeter).
			SetNillableVisibilityVertical(xmlForecast.VertVis).
			SetWeather(xmlForecast.Weather).
			SetNotDecoded(xmlForecast.NotDecoded)

		// Forecast change indicator to enum
		if xmlForecast.Change != nil {
			switch *xmlForecast.Change {
			case "TEMPO":
				fc.SetChangeIndicator(forecast.ChangeIndicatorTEMPO)
				break
			case "BECMG":
				fc.SetChangeIndicator(forecast.ChangeIndicatorBECMG)
				break
			case "FM":
				fc.SetChangeIndicator(forecast.ChangeIndicatorFM)
				break
			case "PROB":
				fc.SetChangeIndicator(forecast.ChangeIndicatorPROB)
				break
			default:
				i.logger.Error(fmt.Sprintf("unknown forecast change indicator %s", *xmlForecast.Change))
				break
			}
		}

		for _, condition := range xmlForecast.SkyCondition {
			skyCover, err := getSkyCoverFromString(condition.Coverage)
			if err != nil {
				return err
			}

			skyC, err := transaction.SkyCondition.Create().
				SetSkyCover(skyCover).
				SetNillableCloudBase(condition.Altitude).
				Save(ctx)
			if err != nil {
				return err
			}

			fc.AddSkyConditions(skyC)
		}

		for _, xmlTurbulence := range xmlForecast.TurbulenceCondition {
			if xmlTurbulence.Intensity == "" && xmlTurbulence.MinAlt == 0 && xmlTurbulence.MaxAlt == 0 {
				continue
			}

			turb, err := transaction.TurbulenceCondition.Create().
				SetIntensity(xmlTurbulence.Intensity).
				SetMinAltitude(xmlTurbulence.MinAlt).
				SetMaxAltitude(xmlTurbulence.MaxAlt).
				Save(ctx)

			if err != nil {
				return err
			}

			fc.AddTurbulenceConditions(turb)
		}

		for _, xmlIcing := range xmlForecast.IcingCondition {
			if xmlIcing.Intensity == "" && xmlIcing.MinAlt == 0 && xmlIcing.MaxAlt == 0 {
				continue
			}

			icing, err := transaction.IcingCondition.Create().
				SetIntensity(xmlIcing.Intensity).
				SetMinAltitude(xmlIcing.MinAlt).
				SetMaxAltitude(xmlIcing.MaxAlt).
				Save(ctx)

			if err != nil {
				return err
			}

			fc.AddIcingConditions(icing)
		}

		for _, xmlTemperature := range xmlForecast.Temperature {
			if xmlTemperature.SurfaceTempC == 0 && (xmlTemperature.ValidTime == time.Time{}) && xmlTemperature.MinTempC == nil && xmlTemperature.MaxTempC == nil {
				continue
			}

			temperature, err := transaction.TemperatureData.Create().
				SetTemperature(xmlTemperature.SurfaceTempC).
				SetNillableMinTemperature(xmlTemperature.MinTempC).
				SetNillableMaxTemperature(xmlTemperature.MaxTempC).
				SetValidTime(xmlTemperature.ValidTime).
				Save(ctx)

			if err != nil {
				return err
			}

			fc.AddTemperatureData(temperature)
		}

		f, err := fc.Save(ctx)
		if err != nil {
			return err
		}

		tx.AddForecast(f)
	}

	err = tx.Exec(ctx)
	if err != nil {
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	i.stats.AddUpdated()

	return err
}

func getSkyCoverFromString(input string) (skycondition.SkyCover, error) {
	var skyCover skycondition.SkyCover = ""

	switch input {
	case "SCT":
		skyCover = skycondition.SkyCoverScattered
		break
	case "CAVOK":
		skyCover = skycondition.SkyCoverCeilingAndVisibilityOK
		break
	case "NSC":
		skyCover = skycondition.SkyCoverNoSignificantClouds
		break
	case "SKC":
		skyCover = skycondition.SkyCoverSkyClear
		break
	case "CLR":
		skyCover = skycondition.SkyCoverClear
		break
	case "BKN":
		skyCover = skycondition.SkyCoverBroken
		break
	case "FEW":
		skyCover = skycondition.SkyCoverFew
		break
	case "OVC":
		skyCover = skycondition.SkyCoverOvercast
		break
	case "OVX":
		skyCover = skycondition.SkyCoverVerticalVisibility
		break
	default:
		return "", fmt.Errorf("unknown sky cover %s", input)
	}

	return skyCover, nil
}

func (i *NoaaWeatherImporter) getStation(ctx context.Context, stationID string, latitude float64, longitude float64, elevation float64) (*ent.WeatherStation, error) {
	// Check if we have an airport with this station ID
	a, _ := i.db.Airport.Query().Where(airport.Identifier(stationID)).Only(ctx)

	line := fmt.Sprintf("%s%f%f%f", stationID, latitude, longitude, elevation)

	if a != nil {
		line += fmt.Sprintf("%d", a.ID)
	}

	hash := strconv.FormatUint(fnv1a.HashString64(line), 10)

	// Check if we already have this s
	s, _ := i.db.WeatherStation.Query().Where(weatherstation.StationID(stationID)).Only(ctx)
	if s != nil {
		if s.Hash == hash {
			return s, nil
		}

		// Update the station
		update := s.Update().SetHash(hash).SetLatitude(latitude).SetLongitude(longitude).SetElevation(elevation)
		if a != nil {
			update.SetAirport(a)
		}

		err := update.Exec(ctx)
		if err != nil {
			return nil, err
		}

		return s, nil
	}

	var err error

	stationCreation := i.db.WeatherStation.Create().SetHash(hash).SetLatitude(latitude).SetLongitude(longitude).SetElevation(elevation).SetStationID(stationID)
	if a != nil {
		stationCreation.SetAirport(a)
	}

	s, err = stationCreation.Save(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}
