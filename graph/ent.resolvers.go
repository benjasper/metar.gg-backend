package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"metar.gg/graph/generated"
)

// Airport returns generated.AirportResolver implementation.
func (r *Resolver) Airport() generated.AirportResolver { return &airportResolver{r} }

// Forecast returns generated.ForecastResolver implementation.
func (r *Resolver) Forecast() generated.ForecastResolver { return &forecastResolver{r} }

// IcingCondition returns generated.IcingConditionResolver implementation.
func (r *Resolver) IcingCondition() generated.IcingConditionResolver {
	return &icingConditionResolver{r}
}

// Metar returns generated.MetarResolver implementation.
func (r *Resolver) Metar() generated.MetarResolver { return &metarResolver{r} }

// Runway returns generated.RunwayResolver implementation.
func (r *Resolver) Runway() generated.RunwayResolver { return &runwayResolver{r} }

// SkyCondition returns generated.SkyConditionResolver implementation.
func (r *Resolver) SkyCondition() generated.SkyConditionResolver { return &skyConditionResolver{r} }

// TemperatureData returns generated.TemperatureDataResolver implementation.
func (r *Resolver) TemperatureData() generated.TemperatureDataResolver {
	return &temperatureDataResolver{r}
}

// TurbulenceCondition returns generated.TurbulenceConditionResolver implementation.
func (r *Resolver) TurbulenceCondition() generated.TurbulenceConditionResolver {
	return &turbulenceConditionResolver{r}
}

// WeatherStation returns generated.WeatherStationResolver implementation.
func (r *Resolver) WeatherStation() generated.WeatherStationResolver {
	return &weatherStationResolver{r}
}

type airportResolver struct{ *Resolver }
type forecastResolver struct{ *Resolver }
type icingConditionResolver struct{ *Resolver }
type metarResolver struct{ *Resolver }
type runwayResolver struct{ *Resolver }
type skyConditionResolver struct{ *Resolver }
type temperatureDataResolver struct{ *Resolver }
type turbulenceConditionResolver struct{ *Resolver }
type weatherStationResolver struct{ *Resolver }
