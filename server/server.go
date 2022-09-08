package server

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/julienschmidt/httprouter"
	"metar.gg/ent"
	"metar.gg/environment"
	"metar.gg/graph"
	"metar.gg/importer"
	"metar.gg/logging"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(db *ent.Client, logger *logging.Logger) error {
	if err := db.Schema.Create(
		context.Background(),
	); err != nil {
		logger.Fatal(err)
	}

	port := environment.Global.Port

	router := httprouter.New()

	router.POST("/import/weather", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if !isAuthorized(w, r) {
			return
		}

		go func() {
			RunWeatherImport(context.Background(), db, logger)
		}()

		w.WriteHeader(http.StatusOK)
	})

	router.POST("/import/airports", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if !isAuthorized(w, r) {
			return
		}

		go func() {
			RunAirportImport(context.Background(), db, logger)
		}()

		w.WriteHeader(http.StatusOK)
	})

	srv := handler.NewDefaultServer(graph.NewSchema(db))
	router.Handler(http.MethodPost, "/graphql", srv)
	router.Handler(http.MethodGet, "/graphql", srv)

	logger.Info(fmt.Sprintf("Server started and listening on port %s", port))
	if err := http.ListenAndServe(":"+port, router); err != nil {
		return err
	}

	return nil
}

func RunAirportImport(ctx context.Context, db *ent.Client, logger *logging.Logger) {
	imp := importer.NewImporter(db, logger)

	err := imp.ImportAirports(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airports.csv")
	if err != nil {
		logger.Fatal(err)
		return
	}

	err = imp.ImportRunways(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/runways.csv")
	if err != nil {
		logger.Fatal(err)
		return
	}

	err = imp.ImportFrequencies(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airport-frequencies.csv")
	if err != nil {
		logger.Fatal(err)
		return
	}
}

func RunWeatherImport(ctx context.Context, db *ent.Client, logger *logging.Logger) {
	metarImporter := importer.NewNoaaMetarImporter(db, logger)
	err := metarImporter.ImportMetars("https://www.aviationweather.gov/adds/dataserver_current/current/metars.cache.xml", ctx)
	if err != nil {
		logger.Fatal(err)
		return
	}
}

func isAuthorized(w http.ResponseWriter, r *http.Request) bool {
	authorization := r.Header.Get("Authorization")
	if authorization != environment.Global.AdminSecret {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}
