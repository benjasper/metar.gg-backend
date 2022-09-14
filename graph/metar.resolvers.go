package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"metar.gg/ent/region"

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
func (r *queryResolver) GetAirports(ctx context.Context, first *int, after *ent.Cursor, before *ent.Cursor, last *int, identifier *string, icao *string, iata *string, typeArg *airport.Type, search *string, hasWeather *bool) (*ent.AirportConnection, error) {
	first, last = BoundsForPagination(first, last)

	var where []predicate.Airport
	if identifier != nil {
		where = append(where, airport.Or(airport.IdentifierEqualFold(*identifier), airport.IdentifierHasPrefix(*identifier)))
	}

	if icao != nil {
		where = append(where, airport.Or(airport.IcaoCodeEqualFold(*icao), airport.IcaoCodeHasPrefix(*icao)))
	}

	if iata != nil {
		where = append(where, airport.Or(airport.IataCodeEqualFold(*iata), airport.IataCodeHasPrefix(*iata)))
	}

	if typeArg != nil {
		where = append(where, airport.TypeEQ(*typeArg))
	}

	if search != nil {
		// Search the airport by its name, ICAO, IATA, GPS code, municipality, local code and keywords.
		where = append(where, airport.Or(
			airport.NameContainsFold(*search),
			airport.IcaoCodeEqualFold(*search),
			airport.IcaoCodeHasPrefix(*search),
			airport.IataCodeEqualFold(*search),
			airport.IataCodeHasPrefix(*search),
			airport.GpsCodeEqualFold(*search),
			airport.GpsCodeHasPrefix(*search),
			airport.MunicipalityContains(*search),
			airport.LocalCodeEqualFold(*search),
			airport.LocalCodeHasPrefix(*search),
			airport.HasRegionWith(region.NameContainsFold(*search)),
		))
	}

	if hasWeather != nil && *hasWeather {
		where = append(where, airport.HasStation())
	} else if hasWeather != nil && !*hasWeather {
		where = append(where, airport.Not(airport.HasStation()))
	}

	connection, err := r.client.Airport.Query().Where(
		airport.Or(where...),
	).Order(ent.Asc(airport.FieldName)).Paginate(ctx, after, first, before, last)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// GetAirport is the resolver for the getAirport field.
func (r *queryResolver) GetAirport(ctx context.Context, id *string, identifier *string, icao *string, iata *string) (*ent.Airport, error) {
	var where []predicate.Airport
	if id != nil {
		uid, err := uuid.Parse(*id)
		if err != nil {
			return nil, err
		}

		where = append(where, airport.ID(uid))
	}

	if identifier != nil {
		where = append(where, airport.IdentifierEqualFold(*identifier))
	}

	if icao != nil {
		where = append(where, airport.IcaoCodeEqualFold(*icao))
	}

	if iata != nil {
		where = append(where, airport.IataCodeEqualFold(*iata))
	}

	a, err := r.client.Airport.Query().Where(airport.Or(where...)).First(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// GetStations is the resolver for the getStations field.
func (r *queryResolver) GetStations(ctx context.Context, first *int, after *ent.Cursor, before *ent.Cursor, last *int, identifier *string) (*ent.WeatherStationConnection, error) {
	first, last = BoundsForPagination(first, last)

	var where []predicate.WeatherStation
	if identifier != nil {
		where = append(where, weatherstation.Or(weatherstation.StationIDEqualFold(*identifier), weatherstation.StationIDHasPrefix(*identifier)))
	}

	connection, err := r.client.WeatherStation.Query().Where(
		where...,
	).Paginate(ctx, after, first, before, last)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// GetStation is the resolver for the getStation field.
func (r *queryResolver) GetStation(ctx context.Context, id *string, identifier *string) (*ent.WeatherStation, error) {
	var orPredicates []predicate.WeatherStation

	if id != nil {
		uid, err := uuid.Parse(*id)
		if err != nil {
			return nil, err
		}

		orPredicates = append(orPredicates, weatherstation.ID(uid))
	}

	if identifier != nil {
		orPredicates = append(orPredicates, weatherstation.StationIDEqualFold(*identifier))
	}

	return r.client.WeatherStation.Query().Where(weatherstation.Or(orPredicates...)).First(ctx)
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
