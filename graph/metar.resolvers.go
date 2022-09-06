package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
	"metar.gg/ent/runway"
	"metar.gg/graph/generated"
	"metar.gg/graph/model"
	"metar.gg/utils"
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
func (r *airportResolver) Metars(ctx context.Context, obj *ent.Airport, first *int) ([]*ent.Metar, error) {
	if first == nil {
		*first = 1
	}

	if (*first) > 12 {
		*first = 12
	}

	return obj.QueryMetars().Order(ent.Desc(metar.FieldObservationTime)).Limit(*first).All(ctx)
}

// MetarsVicinity is the resolver for the metarsVicinity field.
func (r *airportResolver) MetarsVicinity(ctx context.Context, obj *ent.Airport, first *int) ([]*model.MetarWithDistance, error) {
	const R = 6371.0 // km

	radius := 40.0

	maxLat := obj.Latitude + utils.RadiansToDegrees(radius/R)
	minLat := obj.Latitude - utils.RadiansToDegrees(radius/R)
	maxLon := obj.Longitude + utils.RadiansToDegrees(math.Asin(radius/R)/math.Cos(utils.DegreesToRadians(obj.Latitude)))
	minLon := obj.Longitude - utils.RadiansToDegrees(math.Asin(radius/R)/math.Cos(utils.DegreesToRadians(obj.Latitude)))

	var metarsWithIDAndDistance []*MetarWithDistance

	err := r.client.Metar.Query().Where(
		metar.LatitudeLTE(maxLat), metar.LatitudeGTE(minLat), metar.LongitudeLTE(maxLon), metar.LongitudeGTE(minLon),
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(metar.Table).As("metars")
		t2 := sql.Table(metar.Table).As("m2")

		s.From(t1)
		s.AppendSelect(fmt.Sprintf("acos(sin(radians(%f)) * sin(radians(metars.latitude)) + cos(radians(%f)) * cos(radians(metars.latitude)) * cos(radians(metars.longitude) - radians(%f))) * %f as distance", obj.Latitude, obj.Latitude, obj.Longitude, R))
		s.Where(sql.ExprP(fmt.Sprintf("m2.id is null AND acos(sin(radians(%f)) * sin(radians(metars.latitude)) + cos(radians(%f)) * cos(radians(metars.latitude)) * cos(radians(metars.longitude) - radians(%f))) * %f < %f", obj.Latitude, obj.Latitude, obj.Longitude, R, radius)))
		s.OrderExpr(sql.Expr("distance ASC"))
		// TODO change to metar identifier
		s.LeftJoin(t2).OnP(sql.ExprP("(metars.station_id = m2.station_id AND metars.observation_time > m2.observation_time)"))
	}).Limit(*first).Select("id").Scan(ctx, &metarsWithIDAndDistance)
	if err != nil {
		return nil, err
	}

	// Get all ids of metarsWithIDAndDistance
	var ids []int
	for _, metar := range metarsWithIDAndDistance {
		ids = append(ids, metar.ID)
	}

	metars, err := r.client.Metar.Query().Where(metar.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}

	// Create a map of id -> metar
	metarsMap := make(map[int]*ent.Metar)
	for _, metar := range metars {
		metarsMap[metar.ID] = metar
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
	if first == nil && last == nil {
		*first = 10
	}

	if first != nil && *first > 10 {
		*first = 10
	}

	if last != nil && *last > 10 {
		*last = 10
	}

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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type MetarWithDistance struct {
	ID            int     `json:"id"`
	Distance      float64 `json:"distance"`
	AirportMetars int     `json:"airport_metars"`
}
