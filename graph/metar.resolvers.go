package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
	"metar.gg/ent/runway"
	"metar.gg/graph/generated"
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
