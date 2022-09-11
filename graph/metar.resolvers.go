package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
	"metar.gg/ent/runway"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
	"metar.gg/graph/generated"
	"metar.gg/graph/model"
)

// Runways is the resolver for the runways field.
func (r *airportResolver) Runways(ctx context.Context, obj *ent.Airport, closed *bool) ([]*ent.Runway, error) {
	var where []predicate.Runway
	if closed != nil {
		where = append(where, runway.Closed(*closed))
	}

	return obj.QueryRunways().Where(where...).All(ctx)
}

// StationsVicinity is the resolver for the stationsVicinity field.
func (r *airportResolver) StationsVicinity(ctx context.Context, obj *ent.Airport, first *int, radius *float64) ([]*model.StationWithDistance, error) {
	if radius == nil || *radius == 0 {
		*radius = 50.0
	}

	if radius != nil && *radius > 500.0 {
		return nil, fmt.Errorf("radius must be less than 500km")
	}

	minLat, maxLat, minLon, maxLon := GetMinMaxLatLongForGeoQuery(obj.Latitude, obj.Longitude, *radius)

	geoSQLQuery := GetGeoQuerySQL(obj.Latitude, obj.Longitude, "latitude", "longitude")

	var stationsWithIDAndDistance []*model.StationWithDistanceUnstructured

	err := r.client.WeatherStation.Query().Where(
		weatherstation.LatitudeLTE(maxLat), weatherstation.LatitudeGTE(minLat), weatherstation.LongitudeLTE(maxLon), weatherstation.LongitudeGTE(minLon),
	).Modify(func(s *sql.Selector) {
		s.AppendSelect(fmt.Sprintf("%s as distance", geoSQLQuery))
		s.Where(sql.ExprP(fmt.Sprintf("%s < %f", geoSQLQuery, *radius)))
		s.OrderExpr(sql.Expr("distance ASC"))
	}).Limit(*first).Select("id").Scan(ctx, &stationsWithIDAndDistance)
	if err != nil {
		return nil, err
	}

	// Get all ids of stationsWithIDAndDistance
	var ids []uuid.UUID
	for _, m := range stationsWithIDAndDistance {
		ids = append(ids, m.ID)
	}

	stations, err := r.client.WeatherStation.Query().Where(weatherstation.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}

	// Create a map of id -> m
	stationsMap := make(map[uuid.UUID]*ent.WeatherStation)
	for _, m := range stations {
		stationsMap[m.ID] = m
	}

	// Map the results to the model
	var results []*model.StationWithDistance
	for _, stationDistance := range stationsWithIDAndDistance {
		results = append(results, &model.StationWithDistance{
			Station:  stationsMap[stationDistance.ID],
			Distance: stationDistance.Distance,
		})
	}

	return results, nil
}

// GetAirports is the resolver for the getAirports field.
func (r *queryResolver) GetAirports(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, identifier *string, icao *string, iata *string, hasWeather *bool) (*ent.AirportConnection, error) {
	first, last = BoundsForPagination(first, last)

	var where []predicate.Airport
	if identifier != nil {
		where = append(where, airport.IdentifierEqualFold(*identifier))
	}

	if icao != nil {
		where = append(where, airport.IcaoCodeEqualFold(*icao))
	}

	if iata != nil {
		where = append(where, airport.IataCodeEqualFold(*iata))
	}

	if hasWeather != nil && *hasWeather {
		where = append(where, airport.HasStation())
	} else if hasWeather != nil && !*hasWeather {
		where = append(where, airport.Not(airport.HasStation()))
	}

	connection, err := r.client.Airport.Query().Where(
		where...,
	).Paginate(ctx, after, first, before, last)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// Metars is the resolver for the metars field.
func (r *weatherStationResolver) Metars(ctx context.Context, obj *ent.WeatherStation, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.MetarConnection, error) {
	first, last = BoundsForPagination(first, last)

	connection, err := obj.QueryMetars().Order(ent.Desc(metar.FieldObservationTime)).Paginate(ctx, after, first, before, last)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// Tafs is the resolver for the tafs field.
func (r *weatherStationResolver) Tafs(ctx context.Context, obj *ent.WeatherStation, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.TafConnection, error) {
	first, last = BoundsForPagination(first, last)

	connection, err := obj.QueryTafs().Order(ent.Desc(taf.FieldIssueTime)).Paginate(ctx, after, first, before, last)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
