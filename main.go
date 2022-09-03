package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"metar.gg/ent"
	"metar.gg/importer"
	"metar.gg/server"
)

func main() {
	client, err := ent.Open("mysql", "root:123@tcp(localhost:3306)/metargg?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	RunImport(client)
	RunServer(client)
}

func RunServer(db *ent.Client) {
	err := server.NewServer().Run(db)
	if err != nil {
		panic(err)
		return
	}
}

func RunImport(db *ent.Client) {
	imp := importer.NewImporter(db)

	err := imp.ImportAirports("https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airports.csv")
	if err != nil {
		panic(err)
		return
	}

	err = imp.ImportRunways("https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/runways.csv")
	if err != nil {
		panic(err)
		return
	}

	err = imp.ImportFrequencies("https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airport-frequencies.csv")
	if err != nil {
		panic(err)
		return
	}
}
