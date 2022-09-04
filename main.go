package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
	"metar.gg/ent"
	"metar.gg/importer"
	"metar.gg/logging"
	"metar.gg/server"
)

func main() {
	logger := logging.NewLogger()

	logger.Info("Starting up...")

	logger.Info("Connecting to database...")

	client, err := ent.Open("mysql", "root:123@tcp(localhost:3306)/metargg?parseTime=True")
	if err != nil {
		logger.Fatal(err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	logger.Info("Connected to database")

	logger.Info("Running migrations...")
	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		logger.Fatal(err)
	}

	logger.Info("Migrations ran successfully")

	ctx := context.TODO()

	wg := errgroup.Group{}
	wg.Go(func() error {
		return RunServer(client, logger)
	})

	//RunImport(ctx, client, logger)
	RunWeatherImport(ctx, client, logger)

	err = wg.Wait()
	if err != nil {
		logger.Fatal(err)
	}
}

func RunServer(db *ent.Client, logger *logging.Logger) error {
	err := server.NewServer().Run(db, logger)
	if err != nil {
		return err
	}

	return nil
}

func RunImport(ctx context.Context, db *ent.Client, logger *logging.Logger) {
	imp := importer.NewImporter(db)

	logger.Info("Importing airports...")

	err := imp.ImportAirports(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airports.csv")
	if err != nil {
		logger.Fatal(err)
		return
	}

	logger.Info("Finished importing airports")

	logger.Info("Importing runways...")

	err = imp.ImportRunways(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/runways.csv")
	if err != nil {
		logger.Fatal(err)
		return
	}

	logger.Info("Finished importing runways")

	logger.Info("Importing frequencies...")

	err = imp.ImportFrequencies(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airport-frequencies.csv")
	if err != nil {
		logger.Fatal(err)
		return
	}

	logger.Info("Finished importing frequencies")
}

func RunWeatherImport(ctx context.Context, db *ent.Client, logger *logging.Logger) {
	logger.Info("Importing weather...")

	metarImporter := importer.NewNoaaMetarImporter(db)
	err := metarImporter.ImportMetars("https://www.aviationweather.gov/adds/dataserver_current/current/metars.cache.xml", ctx)
	if err != nil {
		logger.Fatal(err)
		return
	}

	logger.Info("Finished importing weather")
}
