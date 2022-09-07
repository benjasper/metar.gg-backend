package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
	"metar.gg/ent/runway"
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

// Metars is the resolver for the metars field.
func (r *airportResolver) Metars(ctx context.Context, obj *ent.Airport, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.MetarConnection, error) {
	first, last = BoundsForPagination(first, last)

	m, err := obj.QueryMetars().Order(ent.Desc(metar.FieldObservationTime)).Paginate(ctx, after, first, before, last)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// MetarsVicinity is the resolver for the metarsVicinity field.
func (r *airportResolver) MetarsVicinity(ctx context.Context, obj *ent.Airport, first *int, radius *float64) ([]*model.MetarWithDistance, error) {
	if radius == nil || *radius == 0 {
		*radius = 50.0
	}

	if radius != nil && *radius > 500.0 {
		return nil, fmt.Errorf("radius must be less than 500km")
	}

	minLat, maxLat, minLon, maxLon := GetMinMaxLatLongForGeoQuery(obj.Latitude, obj.Longitude, *radius)

	geoSQLQuery := GetGeoQuerySQL(obj.Latitude, obj.Longitude, "metars.latitude", "metars.longitude")

	var metarsWithIDAndDistance []*model.MetarWithDistanceUnstructured

	err := r.client.Metar.Query().Where(
		metar.LatitudeLTE(maxLat), metar.LatitudeGTE(minLat), metar.LongitudeLTE(maxLon), metar.LongitudeGTE(minLon),
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(metar.Table).As("metars")
		t2 := sql.Table(metar.Table).As("m2")

		s.From(t1)
		s.AppendSelect(fmt.Sprintf("%s as distance", geoSQLQuery))
		s.Where(sql.ExprP(fmt.Sprintf("m2.id is null AND %s < %f", geoSQLQuery, *radius)))
		s.OrderExpr(sql.Expr("distance ASC"))
		s.LeftJoin(t2).OnP(sql.ExprP("(metars.station_id = m2.station_id AND metars.observation_time > m2.observation_time)"))
	}).Limit(*first).Select("id").Scan(ctx, &metarsWithIDAndDistance)
	if err != nil {
		return nil, err
	}

	// Get all ids of metarsWithIDAndDistance
	var ids []int
	for _, m := range metarsWithIDAndDistance {
		ids = append(ids, m.ID)
	}

	metars, err := r.client.Metar.Query().Where(metar.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}

	// Create a map of id -> m
	metarsMap := make(map[int]*ent.Metar)
	for _, m := range metars {
		metarsMap[m.ID] = m
	}

	// Map the results to the model
	var results []*model.MetarWithDistance
	for _, metarDistance := range metarsWithIDAndDistance {
		results = append(results, &model.MetarWithDistance{
			Metar:    metarsMap[metarDistance.ID],
			Distance: &metarDistance.Distance,
		})
	}

	return results, nil
}

// GetAirports is the resolver for the getAirports field.
func (r *queryResolver) GetAirports(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, identifier *string, hasWeather *bool) (*ent.AirportConnection, error) {
	first, last = BoundsForPagination(first, last)

	var where []predicate.Airport
	if identifier != nil {
		where = append(where, airport.IdentifierEqualFold(*identifier))
	}

	if hasWeather != nil {
		where = append(where, airport.HasWeather(*hasWeather))
	}

	connection, err := r.client.Airport.Query().Where(
		where...,
	).Paginate(ctx, after, first, before, last)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
