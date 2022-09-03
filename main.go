package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"metar.gg/ent"
	"metar.gg/importer"
)

func main() {
	client, err := ent.Open("mysql", "root:123@tcp(localhost:3306)/metar.gg?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	RunImport(client)
	//RunServer()
}

/*
func RunServer() {

	schema, _ := graphql.NewSchema(testutil.TestSchemaConfig)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
		return
	}
}
*/

func RunImport(db *ent.Client) {
	imp := importer.NewImporter(db)

	/*
		err := imp.ImportAirports("https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airports.csv")
		if err != nil {
			panic(err)
			return
		}

	*/

	err := imp.ImportRunways("https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/runways.csv")
	if err != nil {
		panic(err)
		return
	}
}
