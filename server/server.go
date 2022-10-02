package server

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"metar.gg/ent"
	"metar.gg/ent/metar"
	"metar.gg/ent/taf"
	"metar.gg/environment"
	"metar.gg/graph"
	"metar.gg/importer"
	"metar.gg/logging"
	"net/http"
	"time"
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
	srv.Use(extension.FixedComplexityLimit(environment.Global.GraphQLQueryComplexityLimit))

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

	r.POST("/clean", func(c *gin.Context) {
		if !isAuthorized(c.Writer, c.Request) {
			return
		}

		go func() {
			DeleteOldData(context.Background(), db, logger)
		}()

		c.Status(http.StatusNoContent)
	})

	logger.Info("Starting server on port " + port)

	r.Run(":" + port)

	return nil
}

func RunAirportImport(ctx context.Context, db *ent.Client, logger *logging.Logger) {
	imp := importer.NewImporter(db, logger)

	err := imp.ImportCountries(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/countries.csv")
	if err != nil {
		logger.Error(fmt.Sprintf("[IMPORT] Failed to import countries: %s", err))
	}

	err = imp.ImportRegions(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/regions.csv")
	if err != nil {
		logger.Error(fmt.Sprintf("[IMPORT] Failed to import regions: %s", err))
	}

	err = imp.ImportAirports(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airports.csv")
	if err != nil {
		logger.Error(fmt.Sprintf("[IMPORT] Failed to import airports: %s", err))
	}

	err = imp.ImportRunways(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/runways.csv")
	if err != nil {
		logger.Error(fmt.Sprintf("[IMPORT] Failed to import runways: %s", err))
	}

	err = imp.ImportFrequencies(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airport-frequencies.csv")
	if err != nil {
		logger.Error(fmt.Sprintf("[IMPORT] Failed to import frequencies: %s", err))
	}

	imp.Cleanup()
}

func RunWeatherImport(ctx context.Context, db *ent.Client, logger *logging.Logger) {
	metarImporter := importer.NewNoaaWeatherImporter(db, logger)
	err := metarImporter.ImportMetars("https://www.aviationweather.gov/adds/dataserver_current/current/metars.cache.xml", ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("[IMPORT] Failed to import METARs: %s", err))
	}

	tafImporter := importer.NewNoaaWeatherImporter(db, logger)
	err = tafImporter.ImportTafs("https://www.aviationweather.gov/adds/dataserver_current/current/tafs.cache.xml", ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("[IMPORT] Failed to import TAFs: %s", err))
	}
}

func DeleteOldData(ctx context.Context, db *ent.Client, logger *logging.Logger) {

	cutoff := time.Now().Add(-24 * time.Hour)
	result, err := db.Metar.Delete().Where(metar.ObservationTimeLT(cutoff)).Exec(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to delete old METARs: %s", err))
	}

	logger.Info(fmt.Sprintf("Deleted %d old METARs, observed before %s", result, cutoff.Format(time.RFC1123Z)))

	keepDataFor := time.Duration(environment.Global.WeatherDataRetentionDays)

	cutoff = time.Now().Add(-24 * keepDataFor * time.Hour)
	result, err = db.Taf.Delete().Where(taf.IssueTimeLT(cutoff)).Exec(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to delete old TAFs: %s", err))
	}

	logger.Info(fmt.Sprintf("Deleted %d old TAFs issued before %s", result, cutoff.Format(time.RFC1123Z)))
}

func isAuthorized(w http.ResponseWriter, r *http.Request) bool {
	authorization := r.Header.Get("Authorization")
	if authorization != environment.Global.AdminSecret {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}
