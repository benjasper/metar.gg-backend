package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"metar.gg/graph/generated"
)

// Airport returns generated.AirportResolver implementation.
func (r *Resolver) Airport() generated.AirportResolver { return &airportResolver{r} }

// WeatherStation returns generated.WeatherStationResolver implementation.
func (r *Resolver) WeatherStation() generated.WeatherStationResolver {
	return &weatherStationResolver{r}
}

type airportResolver struct{ *Resolver }
type weatherStationResolver struct{ *Resolver }
