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

// Nodes is the resolver for the nodes field.
func (r *airportConnectionResolver) Nodes(ctx context.Context, obj *ent.AirportConnection) ([]*ent.Airport, error) {
	// Destructure the edges
	var airports []*ent.Airport
	for _, edge := range obj.Edges {
		airports = append(airports, edge.Node)
	}

	return airports, nil
}

// Nodes is the resolver for the nodes field.
func (r *metarConnectionResolver) Nodes(ctx context.Context, obj *ent.MetarConnection) ([]*ent.Metar, error) {
	// Destructure the edges
	var metars []*ent.Metar
	for _, edge := range obj.Edges {
		metars = append(metars, edge.Node)
	}

	return metars, nil
}

// GetAirports is the resolver for the getAirports field.
func (r *queryResolver) GetAirports(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, identifier *string, hasWeather *bool) (*ent.AirportConnection, error) {
	first, last = BoundsForPagination(first, last)

	var where []predicate.Airport
	if identifier != nil {
		where = append(where, airport.IdentifierEqualFold(*identifier))
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

// MetarsVicinity is the resolver for the metarsVicinity field.
func (r *stationResolver) MetarsVicinity(ctx context.Context, obj *ent.Station, first *int, radius *float64) ([]*model.MetarWithDistance, error) {
	panic(fmt.Errorf("not implemented: MetarsVicinity - metarsVicinity"))
}

// AirportConnection returns generated.AirportConnectionResolver implementation.
func (r *Resolver) AirportConnection() generated.AirportConnectionResolver {
	return &airportConnectionResolver{r}
}

// MetarConnection returns generated.MetarConnectionResolver implementation.
func (r *Resolver) MetarConnection() generated.MetarConnectionResolver {
	return &metarConnectionResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type airportConnectionResolver struct{ *Resolver }
type metarConnectionResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
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
