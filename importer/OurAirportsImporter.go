package importer

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/segmentio/fasthash/fnv1a"
	"golang.org/x/sync/errgroup"
	"io"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/runway"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Importer struct {
	db *ent.Client
}

func NewImporter(db *ent.Client) *Importer {
	return &Importer{
		db: db,
	}
}

type ImportLineFunction func(data []string) error
type CleanupImportFunction func() error

// Download file and save it to disk.
func (i *Importer) downloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not download file from %s: %s", url, resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) ImportAirports(url string) error {
	err := i.importModelType(url, i.importAirportLine, i.cleanupAirports)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) ImportRunways(url string) error {
	err := i.importModelType(url, i.importRunwayLine, i.cleanupRunways)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) importModelType(url string, importFunction ImportLineFunction, cleanupFunction CleanupImportFunction) error {
	fmt.Println("Importing ", url)

	filepath := "file.csv"
	err := i.downloadFile(url, filepath)
	if err != nil {
		return err
	}

	// Read csv file and map to Airport struct.
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	maxGoroutines := 10
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
		wg.Go(func() error {
			record := record

			defer func() { <-guard }()
			return importFunction(record)
		})
	}

	err = wg.Wait()
	if err != nil {
		return err
	}

	err = cleanupFunction()
	if err != nil {
		return err
	}

	// Delete the file after import.
	err = os.Remove(filepath)

	return err
}

func (i *Importer) importAirportLine(data []string) error {
	// Hash the current line via md5
	line := strings.Join(data, "")
	hash := fnv1a.HashString64(line)
	ourAirportID, _ := strconv.ParseInt(data[0], 10, 64)

	ctx := context.TODO()

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

	scheduledService, _ := strconv.ParseBool(data[9])

	// Explode keywords into array.
	keywords := make([]string, 0)
	if data[14] != "" {
		keywords = strings.Split(data[14], ",")
	}

	err = i.db.Airport.Create().
		SetImportFlag(true).
		SetHash(hash).
		SetID(int(ourAirportID)).
		SetIdentifier(data[1]).
		SetType(data[2]).
		SetName(data[3]).
		SetLatitude(lat).
		SetLongitude(lon).
		SetElevation(int(elevation)).
		SetContinent(data[7]).
		SetCountry(data[8]).
		SetRegion(data[9]).
		SetMunicipality(data[10]).
		SetScheduledService(scheduledService).
		SetGpsCode(data[12]).
		SetIataCode(data[13]).
		SetLocalCode(data[14]).
		SetWebsite(data[15]).
		SetWikipedia(data[16]).
		SetKeywords(keywords).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) cleanupAirports() error {
	ctx := context.TODO()
	deleted, err := i.db.Airport.Delete().Where(
		airport.ImportFlag(false),
	).Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Deleted ", deleted, " rows from airports")

	saved, err := i.db.Airport.Update().Where(
		airport.ImportFlag(true),
	).SetImportFlag(false).Save(ctx)
	if err != nil {
		return err
	}

	println("Total airports ", saved, " rows")

	return nil
}

// The raw data is in the following order: "id","airport_ref","airport_ident","length_ft","width_ft","surface",
//"lighted","closed","le_ident","le_latitude_deg","le_longitude_deg","le_elevation_ft","le_heading_degT",
//"le_displaced_threshold_ft","he_ident","he_latitude_deg","he_longitude_deg","he_elevation_ft","he_heading_degT",
//"he_displaced_threshold_ft"
func (i *Importer) importRunwayLine(data []string) error {
	// Hash the current line via md5
	line := strings.Join(data, "")
	hash := fnv1a.HashString64(line)
	runwayID, _ := strconv.ParseInt(data[0], 10, 64)

	ctx := context.TODO()

	found, err := i.db.Airport.Update().Where(
		airport.Hash(hash),
		airport.ID(int(runwayID)),
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

	leElevation, _ := strconv.ParseInt(data[11], 10, 64)
	leHeading, _ := strconv.ParseInt(data[12], 10, 64)
	leDisplacedThreshold, _ := strconv.ParseInt(data[13], 10, 64)
	leLatitude, _ := strconv.ParseFloat(data[9], 64)
	leLongitude, _ := strconv.ParseFloat(data[10], 64)

	heElevation, _ := strconv.ParseInt(data[17], 10, 64)
	heHeading, _ := strconv.ParseInt(data[18], 10, 64)
	heDisplacedThreshold, _ := strconv.ParseInt(data[19], 10, 64)
	heLatitude, _ := strconv.ParseFloat(data[15], 64)
	heLongitude, _ := strconv.ParseFloat(data[16], 64)

	err = i.db.Runway.Create().
		SetImportFlag(true).
		SetHash(hash).
		SetID(int(runwayID)).
		SetAirportID(int(airportID)).
		SetAirportIdentifier(data[2]).
		SetLength(int(length)).
		SetWidth(int(width)).
		SetSurface(data[5]).
		SetLighted(isLighted).
		SetClosed(false).
		SetLowNumberedRunwayEndIdentifier(data[8]).
		SetLowNumberedRunwayEndLatitude(leLatitude).
		SetLowNumberedRunwayEndLongitude(leLongitude).
		SetLowNumberedRunwayEndElevation(int(leElevation)).
		SetLowNumberedRunwayEndHeading(int(leHeading)).
		SetLowNumberedRunwayEndDisplaced(int(leDisplacedThreshold)).
		SetHighNumberedRunwayEndIdentifier(data[14]).
		SetHighNumberedRunwayEndLatitude(heLatitude).
		SetHighNumberedRunwayEndLongitude(heLongitude).
		SetHighNumberedRunwayEndElevation(int(heElevation)).
		SetHighNumberedRunwayEndHeading(int(heHeading)).
		SetHighNumberedRunwayEndDisplaced(int(heDisplacedThreshold)).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (i *Importer) cleanupRunways() error {
	ctx := context.TODO()
	deleted, err := i.db.Runway.Delete().Where(
		runway.ImportFlag(false),
	).Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Deleted ", deleted, " rows from runways")

	saved, err := i.db.Runway.Update().Where(
		runway.ImportFlag(true),
	).SetImportFlag(false).Save(ctx)
	if err != nil {
		return err
	}

	println("Total runways ", saved, " rows")

	return nil
}
