package importer

import (
	"context"
	"encoding/csv"
	"github.com/segmentio/fasthash/fnv1a"
	"golang.org/x/sync/errgroup"
	"io"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/frequency"
	"metar.gg/ent/runway"
	"metar.gg/environment"
	"metar.gg/logging"
	"metar.gg/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

type Importer struct {
	db     *ent.Client
	logger *logging.Logger
	stats  *ImportStatistics
}

func NewImporter(db *ent.Client, logger *logging.Logger) *Importer {
	return &Importer{
		db:     db,
		logger: logger,
		stats:  NewImportStatistics("", logger),
	}
}

type ImportLineFunction func(ctx context.Context, data []string) error
type CleanupImportFunction func(ctx context.Context) error

func (i *Importer) ImportAirports(ctx context.Context, url string) error {
	i.stats = NewImportStatistics("AIRPORTS", i.logger)

	err := i.importModelType(ctx, url, i.importAirportLine, i.cleanupAirports)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) ImportRunways(ctx context.Context, url string) error {
	i.stats = NewImportStatistics("RUNWAYS", i.logger)

	err := i.importModelType(ctx, url, i.importRunwayLine, i.cleanupRunways)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) ImportFrequencies(ctx context.Context, url string) error {
	i.stats = NewImportStatistics("FREQUENCIES", i.logger)

	err := i.importModelType(ctx, url, i.importFrequencyLine, i.cleanupFrequencies)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) importModelType(ctx context.Context, url string, importFunction ImportLineFunction, cleanupFunction CleanupImportFunction) error {
	i.stats.Start()

	filepath := "file.csv"
	err := utils.DownloadFile(url, filepath)
	if err != nil {
		return err
	}

	// Read csv file and map to Airport struct.
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	csvReader := csv.NewReader(f)

	maxGoroutines := environment.Global.MaxConcurrentImports
	guard := make(chan struct{}, maxGoroutines)

	wg := errgroup.Group{}

	lineCounter := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		lineCounter++

		if lineCounter == 1 {
			continue
		}

		guard <- struct{}{} // would block if guard channel is already filled
		i.stats.AddTotal()
		wg.Go(func() error {
			record := record

			defer func() { <-guard }()
			return importFunction(ctx, record)
		})
	}

	err = wg.Wait()
	if err != nil {
		return err
	}

	err = cleanupFunction(ctx)
	if err != nil {
		return err
	}

	// Delete the file after import.
	err = os.Remove(filepath)

	i.stats.End()

	return err
}

// CSV file format: "id","ident","type","name","latitude_deg","longitude_deg","elevation_ft","continent","iso_country","iso_region","municipality","scheduled_service","gps_code","iata_code","local_code","home_link","wikipedia_link","keywords"
func (i *Importer) importAirportLine(ctx context.Context, data []string) error {
	// Hash the current line via md5
	line := strings.Join(data, "")
	hash := strconv.FormatUint(fnv1a.HashString64(line), 10)
	ourAirportID, _ := strconv.ParseInt(data[0], 10, 64)

	found, err := i.db.Airport.Update().Where(
		airport.Hash(hash),
		airport.ID(int(ourAirportID)),
	).
		SetImportFlag(true).
		Save(ctx)
	if err != nil {
		return err
	}

	if found == 1 {
		return nil
	}

	// Upsert airport, because we know it doesn't exist yet, or it changed

	lat, _ := strconv.ParseFloat(data[4], 64)
	lon, _ := strconv.ParseFloat(data[5], 64)

	elevation, _ := strconv.ParseInt(data[6], 10, 64)

	scheduledService := false
	if data[9] == "yes" {
		scheduledService = true
	}

	// Explode keywords into array.
	keywords := make([]string, 0)
	if data[17] != "" {
		keywords = strings.Split(data[17], ", ")
	}

	airportType := airport.TypeSmallAirport
	switch data[2] {
	case airport.TypeSmallAirport.String():
		airportType = airport.TypeSmallAirport
		break
	case airport.TypeMediumAirport.String():
		airportType = airport.TypeMediumAirport
		break
	case airport.TypeLargeAirport.String():
		airportType = airport.TypeLargeAirport
		break
	case airport.TypeSeaplaneBase.String():
		airportType = airport.TypeSeaplaneBase
		break
	case airport.TypeHeliport.String():
		airportType = airport.TypeHeliport
		break
	case airport.TypeClosedAirport.String():
		airportType = airport.TypeClosedAirport
		break
	}

	continent := airport.ContinentEurope
	switch data[11] {
	case airport.ContinentAfrica.String():
		continent = airport.ContinentAfrica
		break
	case airport.ContinentAntarctica.String():
		continent = airport.ContinentAntarctica
		break
	case airport.ContinentAsia.String():
		continent = airport.ContinentAsia
		break
	case airport.ContinentOceania.String():
		continent = airport.ContinentOceania
		break
	case airport.ContinentSouthAmerica.String():
		continent = airport.ContinentSouthAmerica
		break
	case airport.ContinentNorthAmerica.String():
		continent = airport.ContinentNorthAmerica
		break
	}

	err = i.db.Airport.Create().
		SetImportFlag(true).
		SetHash(hash).
		SetID(int(ourAirportID)).
		SetIdentifier(data[1]).
		SetType(airportType).
		SetName(data[3]).
		SetLatitude(lat).
		SetLongitude(lon).
		SetNillableElevation(utils.NillableWithInput(data[6], int(elevation))).
		SetContinent(continent).
		SetCountry(data[8]).
		SetRegion(data[9]).
		SetNillableMunicipality(utils.NillableString(data[10])).
		SetScheduledService(scheduledService).
		SetNillableGpsCode(utils.NillableString(data[12])).
		SetIataCode(data[13]).
		SetNillableLocalCode(utils.NillableString(data[14])).
		SetNillableWebsite(utils.NillableString(data[15])).
		SetNillableWikipedia(utils.NillableString(data[16])).
		SetKeywords(keywords).
		SetLastUpdated(time.Now()).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	i.stats.AddUpdated()

	return nil
}

func (i *Importer) cleanupAirports(ctx context.Context) error {
	deleted, err := i.db.Airport.Delete().Where(
		airport.ImportFlag(false),
	).Exec(ctx)
	if err != nil {
		return err
	}

	i.stats.AddMultipleDeleted(deleted)

	_, err = i.db.Airport.Update().Where(
		airport.ImportFlag(true),
	).SetImportFlag(false).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// The raw data is in the following order: "id","airport_ref","airport_ident","length_ft","width_ft","surface",
// "lighted","closed","le_ident","le_latitude_deg","le_longitude_deg","le_elevation_ft","le_heading_degT",
// "le_displaced_threshold_ft","he_ident","he_latitude_deg","he_longitude_deg","he_elevation_ft","he_heading_degT",
// "he_displaced_threshold_ft"
func (i *Importer) importRunwayLine(ctx context.Context, data []string) error {
	// Hash the current line via md5
	line := strings.Join(data, "")
	hash := strconv.FormatUint(fnv1a.HashString64(line), 10)
	runwayID, _ := strconv.ParseInt(data[0], 10, 64)

	found, err := i.db.Runway.Update().Where(
		runway.Hash(hash),
		runway.ID(int(runwayID)),
	).
		SetImportFlag(true).
		Save(ctx)
	if err != nil {
		return err
	}

	if found == 1 {
		return nil
	}

	// Upsert airport, because we know it doesn't exist yet, or it changed
	airportID, err := strconv.ParseInt(data[1], 10, 64)
	if err != nil {
		return err
	}

	length, _ := strconv.ParseInt(data[3], 10, 64)
	width, _ := strconv.ParseInt(data[4], 10, 64)
	isLighted, _ := strconv.ParseBool(data[6])
	isClosed, _ := strconv.ParseBool(data[7])

	leElevation, _ := strconv.ParseInt(data[11], 10, 64)
	leHeading, _ := strconv.ParseFloat(data[12], 64)
	leDisplacedThreshold, _ := strconv.ParseInt(data[13], 10, 64)
	leLatitude, _ := strconv.ParseFloat(data[9], 64)
	leLongitude, _ := strconv.ParseFloat(data[10], 64)

	heElevation, _ := strconv.ParseInt(data[17], 10, 64)
	heHeading, _ := strconv.ParseFloat(data[18], 64)
	heDisplacedThreshold, _ := strconv.ParseInt(data[19], 10, 64)
	heLatitude, _ := strconv.ParseFloat(data[15], 64)
	heLongitude, _ := strconv.ParseFloat(data[16], 64)

	err = i.db.Runway.Create().
		SetImportFlag(true).
		SetHash(hash).
		SetID(int(runwayID)).
		SetAirportID(int(airportID)).
		SetLength(int(length)).
		SetWidth(int(width)).
		SetNillableSurface(utils.NillableString(data[5])).
		SetLighted(isLighted).
		SetClosed(isClosed).
		SetLowRunwayIdentifier(data[8]).
		SetNillableLowRunwayLatitude(utils.NillableWithInput(data[9], leLatitude)).
		SetNillableLowRunwayLongitude(utils.NillableWithInput(data[10], leLongitude)).
		SetNillableLowRunwayElevation(utils.NillableWithInput(data[11], int(leElevation))).
		SetNillableLowRunwayHeading(utils.NillableWithInput(data[12], leHeading)).
		SetNillableLowRunwayDisplacedThreshold(utils.NillableWithInput(data[13], int(leDisplacedThreshold))).
		SetHighRunwayIdentifier(data[14]).
		SetNillableHighRunwayLatitude(utils.NillableWithInput(data[15], heLatitude)).
		SetNillableHighRunwayLongitude(utils.NillableWithInput(data[16], heLongitude)).
		SetNillableHighRunwayElevation(utils.NillableWithInput(data[17], int(heElevation))).
		SetNillableHighRunwayHeading(utils.NillableWithInput(data[18], heHeading)).
		SetNillableHighRunwayDisplacedThreshold(utils.NillableWithInput(data[19], int(heDisplacedThreshold))).
		SetLastUpdated(time.Now()).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	i.stats.AddUpdated()

	return nil
}

func (i *Importer) cleanupRunways(ctx context.Context) error {
	deleted, err := i.db.Runway.Delete().Where(
		runway.ImportFlag(false),
	).Exec(ctx)
	if err != nil {
		return err
	}

	i.stats.AddMultipleDeleted(deleted)

	_, err = i.db.Runway.Update().Where(
		runway.ImportFlag(true),
	).SetImportFlag(false).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// The raw data is in the following order: "id","airport_ref","airport_ident","type","description","frequency_mhz"
func (i *Importer) importFrequencyLine(ctx context.Context, data []string) error {
	// Hash the current line via md5
	line := strings.Join(data, "")
	hash := strconv.FormatUint(fnv1a.HashString64(line), 10)
	frequencyID, _ := strconv.ParseInt(data[0], 10, 64)

	found, err := i.db.Frequency.Update().Where(
		frequency.Hash(hash),
		frequency.ID(int(frequencyID)),
	).
		SetImportFlag(true).
		Save(ctx)
	if err != nil {
		return err
	}

	if found == 1 {
		return nil
	}

	// Upsert airport, because we know it doesn't exist yet, or it changed
	airportID, err := strconv.ParseInt(data[1], 10, 64)
	if err != nil {
		return err
	}

	f, _ := strconv.ParseFloat(data[5], 64)

	err = i.db.Frequency.Create().
		SetImportFlag(true).
		SetHash(hash).
		SetID(int(frequencyID)).
		SetAirportID(int(airportID)).
		SetType(data[3]).
		SetDescription(data[4]).
		SetFrequency(f).
		SetLastUpdated(time.Now()).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	i.stats.AddUpdated()

	return nil
}

func (i *Importer) cleanupFrequencies(ctx context.Context) error {
	deleted, err := i.db.Frequency.Delete().Where(
		frequency.ImportFlag(false),
	).Exec(ctx)
	if err != nil {
		return err
	}

	i.stats.AddMultipleDeleted(deleted)

	_, err = i.db.Frequency.Update().Where(
		frequency.ImportFlag(true),
	).SetImportFlag(false).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
