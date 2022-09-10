package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
	"metar.gg/ent"
	"metar.gg/ent/migrate"
	"metar.gg/environment"
	"metar.gg/logging"
	"metar.gg/server"
)

func main() {
	logger := logging.NewLogger()

	environment.Initialize(logger)

	logger.Info("Starting up...")

	logger.Info("Connecting to database...")

	client, err := ent.Open("mysql", environment.Global.Database)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	logger.Info("Connected to database")

	logger.Info("Running migrations...")
	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background(), migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		logger.Fatal(err)
	}

	logger.Info("Migrations ran successfully")

	wg := errgroup.Group{}
	wg.Go(func() error {
		return RunServer(client, logger)
	})

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
