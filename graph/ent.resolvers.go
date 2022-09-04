package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"metar.gg/graph/generated"
)

// Airport returns generated.AirportResolver implementation.
func (r *Resolver) Airport() generated.AirportResolver { return &airportResolver{r} }

type airportResolver struct{ *Resolver }
