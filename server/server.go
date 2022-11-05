package server

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/cenkalti/backoff/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
	"metar.gg/environment"
	"metar.gg/graph"
	"metar.gg/importer"
	"metar.gg/logging"
	"net/http"
	"strings"
	"sync"
	"time"
)

var weatherImportMutex = &sync.Mutex{}

type Server struct {
	db                 *ent.Client
	logger             *logging.Logger
	sitemap            *stm.Sitemap
	sitemapLastUpdated time.Time
}

func NewServer(db *ent.Client, logger *logging.Logger) *Server {
	return &Server{
		db:                 db,
		logger:             logger,
		sitemap:            nil,
		sitemapLastUpdated: time.Time{},
	}
}

func (s *Server) Run() error {
	if err := s.db.Schema.Create(
		context.Background(),
	); err != nil {
		s.logger.Fatal(err)
	}

	port := environment.Global.Port

	srv := handler.NewDefaultServer(graph.NewSchema(s.db))
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
			s.RunWeatherImport(context.Background())
		}()

		c.Status(http.StatusNoContent)
	})

	r.POST("/import/airports", func(c *gin.Context) {
		if !isAuthorized(c.Writer, c.Request) {
			return
		}

		go func() {
			s.RunAirportImport(context.Background())
		}()

		c.Status(http.StatusNoContent)
	})

	r.POST("/clean", func(c *gin.Context) {
		if !isAuthorized(c.Writer, c.Request) {
			return
		}

		go func() {
			s.DeleteOldData(context.Background())
		}()

		c.Status(http.StatusNoContent)
	})

	r.GET("/sitemap.xml", s.respondWithSitemap)

	s.logger.Info("Starting server on port " + port)

	r.Run(":" + port)

	return nil
}

func (s *Server) RunAirportImport(ctx context.Context) {
	imp := importer.NewImporter(s.db, s.logger)

	err := imp.ImportCountries(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/countries.csv")
	if err != nil {
		s.logger.Error(fmt.Sprintf("[IMPORT] Failed to import countries: %s", err))
	}

	err = imp.ImportRegions(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/regions.csv")
	if err != nil {
		s.logger.Error(fmt.Sprintf("[IMPORT] Failed to import regions: %s", err))
	}

	err = imp.ImportAirports(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airports.csv")
	if err != nil {
		s.logger.Error(fmt.Sprintf("[IMPORT] Failed to import airports: %s", err))
	}

	err = imp.ImportRunways(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/runways.csv")
	if err != nil {
		s.logger.Error(fmt.Sprintf("[IMPORT] Failed to import runways: %s", err))
	}

	err = imp.ImportFrequencies(ctx, "https://raw.githubusercontent.com/davidmegginson/ourairports-data/main/airport-frequencies.csv")
	if err != nil {
		s.logger.Error(fmt.Sprintf("[IMPORT] Failed to import frequencies: %s", err))
	}
}

func (s *Server) RunWeatherImport(ctx context.Context) {
	// We use try lock here, because we don't want to queue up multiple imports
	if !weatherImportMutex.TryLock() {
		s.logger.Warn("[IMPORT] Weather import already running")
		return
	}

	defer weatherImportMutex.Unlock()

	metarImporter := importer.NewNoaaWeatherImporter(s.db, s.logger)

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 5 * time.Minute

	err := backoff.Retry(func() error {
		err := metarImporter.ImportMetars("https://www.aviationweather.gov/adds/dataserver_current/current/metars.cache.xml", ctx)
		return err
	}, b)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[IMPORT] Failed to import METARs: %s", err))
	}

	b.Reset()

	err = backoff.Retry(func() error {
		tafImporter := importer.NewNoaaWeatherImporter(s.db, s.logger)
		err = tafImporter.ImportTafs("https://www.aviationweather.gov/adds/dataserver_current/current/tafs.cache.xml", ctx)
		return err
	}, b)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[IMPORT] Failed to import TAFs: %s", err))
	}
}

func (s *Server) DeleteOldData(ctx context.Context) {

	cutoff := time.Now().Add(-24 * time.Hour * time.Duration(environment.Global.WeatherDataRetentionDays))
	result, err := s.db.Metar.Delete().Where(metar.ObservationTimeLT(cutoff)).Exec(ctx)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to delete old METARs: %s", err))
	}

	s.logger.Info(fmt.Sprintf("Deleted %d old METARs, observed before %s", result, cutoff.Format(time.RFC1123Z)))

	cutoff = time.Now().Add(-24 * time.Hour * time.Duration(environment.Global.WeatherDataRetentionDays))
	result, err = s.db.Taf.Delete().Where(taf.IssueTimeLT(cutoff)).Exec(ctx)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to delete old TAFs: %s", err))
	}

	s.logger.Info(fmt.Sprintf("Deleted %d old TAFs issued before %s", result, cutoff.Format(time.RFC1123Z)))
}

func isAuthorized(w http.ResponseWriter, r *http.Request) bool {
	authorization := r.Header.Get("Authorization")
	if authorization != environment.Global.AdminSecret {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func (s *Server) respondWithSitemap(c *gin.Context) {
	if environment.Global.SitemapBase == "" || environment.Global.SitemapAirportsPath == "" {
		c.Data(http.StatusPreconditionFailed, "text/plain", []byte("SitemapBase or SitemapAirportsPath not set"))

		return
	}

	if s.sitemap == nil || s.sitemapLastUpdated.Before(time.Now().Add(-24*time.Hour)) {
		s.sitemap = s.generateSitemap(c)
		s.sitemapLastUpdated = time.Now()
	}

	_, err := c.Writer.Write(s.sitemap.XMLContent())
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error while generating sitemap: %s", err.Error()))
		return
	}
}

func (s *Server) generateSitemap(c *gin.Context) *stm.Sitemap {
	sm := stm.NewSitemap(environment.Global.MaxConcurrentImports)

	sm.SetDefaultHost(environment.Global.SitemapBase)

	// Create method must be called first before adding entries to
	// the sitemap.
	sm.Create()

	// Query an airport with maximum importance
	maxImportantAirport, err := s.db.Airport.Query().Where(airport.HasStationWith(weatherstation.HasMetars())).Order(ent.Desc(airport.FieldImportance)).Select(airport.FieldImportance).First(c.Request.Context())
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error while generating sitemap: %s", err.Error()))
		return nil
	}

	maxImportance := float64(maxImportantAirport.Importance)

	// Get all airports that have a METAR with their latest METAR
	airportsCount, err := s.db.Airport.Query().
		Where(airport.HasStationWith(weatherstation.HasMetars())).
		WithStation(func(q *ent.WeatherStationQuery) {
			q.WithMetars(func(q *ent.MetarQuery) {
				q.Order(ent.Desc(metar.FieldObservationTime))
			})
		},
		).Count(c.Request.Context())
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error while generating sitemap: %s", err.Error()))
		return nil
	}

	for i := 0; i < airportsCount; i += 100 {
		airportsPage, err := s.db.Airport.Query().
			Select(airport.FieldIdentifier, airport.FieldImportance).
			Where(airport.HasStationWith(weatherstation.HasMetars())).
			WithStation(func(q *ent.WeatherStationQuery) {
				q.Select()
				q.WithMetars(func(q *ent.MetarQuery) {
					q.Select(metar.FieldObservationTime)
					q.Order(ent.Desc(metar.FieldObservationTime))
				})
			},
			).Offset(i).Limit(100).All(c.Request.Context())
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error while generating sitemap: %s", err.Error()))
			return nil
		}

		for _, a := range airportsPage {
			priority := float64(a.Importance) / maxImportance
			sm.Add(stm.URL{{"loc", fmt.Sprintf(environment.Global.SitemapAirportsPath, a.Identifier)}, {"changefreq", "always"}, {"priority", fmt.Sprintf("%.1f", priority)}, {"lastmod", a.Edges.Station.Edges.Metars[0].ObservationTime}})
		}
	}

	if environment.Global.SitemapAdditionalUrls != "" {
		urls := strings.Split(environment.Global.SitemapAdditionalUrls, ",")
		for _, u := range urls {
			sm.Add(stm.URL{{"loc", u}, {"changefreq", "always"}, {"priority", "1.0"}})
		}
	}

	sm.Add(stm.URL{{"loc", environment.Global.SitemapBase}, {"changefreq", "always"}, {"priority", "1.0"}})

	return sm
}
