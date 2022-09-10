package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	srv := handler.NewDefaultServer(graph.NewSchema(db))
	srv.Use(extension.FixedComplexityLimit(80))

	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	r.Use(cors.New(config))

	r.POST("/graphql", gin.WrapH(srv))
	r.GET("/graphql", gin.WrapH(srv))

	r.POST("/import/weather", func(c *gin.Context) {
		if !isAuthorized(c.Writer, c.Request) {
			return
		}

		go func() {
			RunWeatherImport(context.Background(), db, logger)
		}()

		c.Status(http.StatusNoContent)
	})

	r.POST("/import/airports", func(c *gin.Context) {
		if !isAuthorized(c.Writer, c.Request) {
			return
		}

		go func() {
			RunAirportImport(context.Background(), db, logger)
		}()

		c.Status(http.StatusNoContent)
	})

	logger.Info("Starting server on port " + port)

	r.Run(":" + port)

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
