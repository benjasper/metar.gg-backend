package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"metar.gg/ent"
	"metar.gg/ent/airport"
	"metar.gg/ent/country"
	"metar.gg/ent/predicate"
	"metar.gg/ent/region"
	"metar.gg/ent/weatherstation"
	"metar.gg/graph/generated"
)

// GetAirports is the resolver for the getAirports field.
func (r *queryResolver) GetAirports(ctx context.Context, first *int, after *ent.Cursor, before *ent.Cursor, last *int, identifier *string, icao *string, iata *string, typeArg *airport.Type, search *string, hasWeather *bool, order []*ent.AirportOrder) (*ent.AirportConnection, error) {
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
			airport.IcaoCodeContainsFold(*search),
			airport.IataCodeEqualFold(*search),
			airport.IataCodeContainsFold(*search),
			airport.MunicipalityContainsFold(*search),
			airport.HasCountryWith(country.NameContainsFold(*search)),
			airport.LocalCodeEqualFold(*search),
			airport.LocalCodeContainsFold(*search),
			airport.HasRegionWith(region.NameContainsFold(*search)),
		))
	}

	if hasWeather != nil && *hasWeather {
		where = append(where, airport.HasStation())
	} else if hasWeather != nil && !*hasWeather {
		where = append(where, airport.Not(airport.HasStation()))
	}

	airportPaginateOptions := []ent.AirportPaginateOption{
		ent.WithAirportOrder(&ent.AirportOrder{Field: ent.AirportOrderFieldImportance, Direction: ent.OrderDirectionDesc}),
	}

	if len(order) > 0 {
		airportPaginateOptions = nil
		for _, airportOrder := range order {
			airportPaginateOptions = append(airportPaginateOptions, ent.WithAirportOrder(airportOrder))
		}
	}

	connection, err := r.client.Airport.Query().Where(
		airport.And(where...),
	).Paginate(ctx, after, first, before, last, airportPaginateOptions...)
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
