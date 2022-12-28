package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"metar.gg/ent"
	"metar.gg/graph/model"
)

// Altimeter is the resolver for the altimeter field.
func (r *forecastResolver) Altimeter(ctx context.Context, obj *ent.Forecast, unit model.PressureUnit) (*float64, error) {
	if obj.Altimeter == nil {
		return nil, nil
	}

	pressure := PressureFromInchesOfMercuryToUnit(*obj.Altimeter, unit)

	return &pressure, nil
}

// WindSpeed is the resolver for the windSpeed field.
func (r *forecastResolver) WindSpeed(ctx context.Context, obj *ent.Forecast, unit model.SpeedUnit) (*float64, error) {
	if obj.WindSpeed == nil {
		return nil, nil
	}

	speed := SpeedFromKnotsToUnit(*obj.WindSpeed, unit)

	return &speed, nil
}

// WindGust is the resolver for the windGust field.
func (r *forecastResolver) WindGust(ctx context.Context, obj *ent.Forecast, unit model.SpeedUnit) (*float64, error) {
	if obj.WindGust == nil {
		return nil, nil
	}

	speed := SpeedFromKnotsToUnit(*obj.WindGust, unit)

	return &speed, nil
}

// VisibilityHorizontal is the resolver for the visibilityHorizontal field.
func (r *forecastResolver) VisibilityHorizontal(ctx context.Context, obj *ent.Forecast, unit model.LengthUnit) (*float64, error) {
	if obj.VisibilityHorizontal == nil {
		return nil, nil
	}

	length := LengthFromStatuteMileToUnit(*obj.VisibilityHorizontal, unit)

	return &length, nil
}

// VisibilityVertical is the resolver for the visibilityVertical field.
func (r *forecastResolver) VisibilityVertical(ctx context.Context, obj *ent.Forecast, unit model.LengthUnit) (*float64, error) {
	if obj.VisibilityVertical == nil {
		return nil, nil
	}

	length := LengthFromFeetToUnit(float64(*obj.VisibilityVertical), unit)

	return &length, nil
}

// WindShearHeight is the resolver for the windShearHeight field.
func (r *forecastResolver) WindShearHeight(ctx context.Context, obj *ent.Forecast, unit model.LengthUnit) (*float64, error) {
	if obj.WindShearHeight == nil {
		return nil, nil
	}

	length := LengthFromFeetToUnit(float64(*obj.WindShearHeight), unit)

	return &length, nil
}

// WindShearSpeed is the resolver for the windShearSpeed field.
func (r *forecastResolver) WindShearSpeed(ctx context.Context, obj *ent.Forecast, unit model.SpeedUnit) (*float64, error) {
	if obj.WindShearSpeed == nil {
		return nil, nil
	}

	speed := SpeedFromKnotsToUnit(*obj.WindShearSpeed, unit)

	return &speed, nil
}

// MinAltitude is the resolver for the minAltitude field.
func (r *icingConditionResolver) MinAltitude(ctx context.Context, obj *ent.IcingCondition, unit model.LengthUnit) (*float64, error) {
	if obj.MinAltitude == nil {
		return nil, nil
	}

	length := LengthFromFeetToUnit(float64(*obj.MinAltitude), unit)

	return &length, nil
}

// MaxAltitude is the resolver for the maxAltitude field.
func (r *icingConditionResolver) MaxAltitude(ctx context.Context, obj *ent.IcingCondition, unit model.LengthUnit) (*float64, error) {
	if obj.MaxAltitude == nil {
		return nil, nil
	}

	length := LengthFromFeetToUnit(float64(*obj.MaxAltitude), unit)

	return &length, nil
}

// Altimeter is the resolver for the altimeter field.
func (r *metarResolver) Altimeter(ctx context.Context, obj *ent.Metar, unit model.PressureUnit) (float64, error) {
	return PressureFromInchesOfMercuryToUnit(obj.Altimeter, unit), nil
}

// Temperature is the resolver for the temperature field.
func (r *metarResolver) Temperature(ctx context.Context, obj *ent.Metar, unit model.TemperatureUnit) (float64, error) {
	return TemperatureFromCelsiusToUnit(obj.Temperature, unit), nil
}

// Dewpoint is the resolver for the dewpoint field.
func (r *metarResolver) Dewpoint(ctx context.Context, obj *ent.Metar, unit model.TemperatureUnit) (float64, error) {
	return TemperatureFromCelsiusToUnit(obj.Dewpoint, unit), nil
}

// WindSpeed is the resolver for the windSpeed field.
func (r *metarResolver) WindSpeed(ctx context.Context, obj *ent.Metar, unit model.SpeedUnit) (float64, error) {
	return SpeedFromKnotsToUnit(obj.WindSpeed, unit), nil
}

// WindGust is the resolver for the windGust field.
func (r *metarResolver) WindGust(ctx context.Context, obj *ent.Metar, unit model.SpeedUnit) (float64, error) {
	return SpeedFromKnotsToUnit(obj.WindGust, unit), nil
}

// Visibility is the resolver for the visibility field.
func (r *metarResolver) Visibility(ctx context.Context, obj *ent.Metar, unit model.LengthUnit) (float64, error) {
	return LengthFromStatuteMileToUnit(obj.Visibility, unit), nil
}

// VerticalVisibility is the resolver for the verticalVisibility field.
func (r *metarResolver) VerticalVisibility(ctx context.Context, obj *ent.Metar, unit model.LengthUnit) (*float64, error) {
	if obj.VertVis == nil {
		return nil, nil
	}

	length := LengthFromStatuteMileToUnit(*obj.VertVis, unit)

	return &length, nil
}

// SnowDepth is the resolver for the snowDepth field.
func (r *metarResolver) SnowDepth(ctx context.Context, obj *ent.Metar, unit model.SmallLengthUnit) (*float64, error) {
	if obj.SnowDepth == nil {
		return nil, nil
	}

	length := SmallLengthFromInchesToUnit(*obj.SnowDepth, unit)

	return &length, nil
}

// SeaLevelPressure is the resolver for the seaLevelPressure field.
func (r *metarResolver) SeaLevelPressure(ctx context.Context, obj *ent.Metar, unit model.PressureUnit) (*float64, error) {
	if obj.SeaLevelPressure == nil {
		return nil, nil
	}

	pressure := PressureFromHectopascalsToUnit(*obj.SeaLevelPressure, unit)

	return &pressure, nil
}

// PressureTendency is the resolver for the pressureTendency field.
func (r *metarResolver) PressureTendency(ctx context.Context, obj *ent.Metar, unit model.PressureUnit) (*float64, error) {
	if obj.PressureTendency == nil {
		return nil, nil
	}

	pressure := PressureFromHectopascalsToUnit(*obj.PressureTendency, unit)

	return &pressure, nil
}

// Length is the resolver for the length field.
func (r *runwayResolver) Length(ctx context.Context, obj *ent.Runway, unit model.LengthUnit) (*float64, error) {
	if obj.Length == 0 {
		return nil, nil
	}

	length := LengthFromFeetToUnit(float64(obj.Length), unit)

	return &length, nil
}

// Width is the resolver for the width field.
func (r *runwayResolver) Width(ctx context.Context, obj *ent.Runway, unit model.LengthUnit) (*float64, error) {
	if obj.Length == 0 {
		return nil, nil
	}

	width := LengthFromFeetToUnit(float64(obj.Width), unit)

	return &width, nil
}

// CloudBase is the resolver for the cloudBase field.
func (r *skyConditionResolver) CloudBase(ctx context.Context, obj *ent.SkyCondition, unit model.LengthUnit) (*float64, error) {
	if obj.CloudBase == nil {
		return nil, nil
	}

	length := LengthFromFeetToUnit(float64(*obj.CloudBase), unit)

	return &length, nil
}

// Temperature is the resolver for the temperature field.
func (r *temperatureDataResolver) Temperature(ctx context.Context, obj *ent.TemperatureData, unit model.TemperatureUnit) (float64, error) {
	return TemperatureFromCelsiusToUnit(obj.Temperature, unit), nil
}

// MinTemperature is the resolver for the minTemperature field.
func (r *temperatureDataResolver) MinTemperature(ctx context.Context, obj *ent.TemperatureData, unit model.TemperatureUnit) (*float64, error) {
	if obj.MinTemperature == nil {
		return nil, nil
	}

	minTemperature := TemperatureFromCelsiusToUnit(*obj.MinTemperature, unit)

	return &minTemperature, nil
}

// MaxTemperature is the resolver for the maxTemperature field.
func (r *temperatureDataResolver) MaxTemperature(ctx context.Context, obj *ent.TemperatureData, unit model.TemperatureUnit) (*float64, error) {
	if obj.MaxTemperature == nil {
		return nil, nil
	}

	maxTemperature := TemperatureFromCelsiusToUnit(*obj.MaxTemperature, unit)

	return &maxTemperature, nil
}

// MinAltitude is the resolver for the minAltitude field.
func (r *turbulenceConditionResolver) MinAltitude(ctx context.Context, obj *ent.TurbulenceCondition, unit model.LengthUnit) (*float64, error) {
	altitude := LengthFromFeetToUnit(float64(obj.MinAltitude), unit)

	return &altitude, nil
}

// MaxAltitude is the resolver for the maxAltitude field.
func (r *turbulenceConditionResolver) MaxAltitude(ctx context.Context, obj *ent.TurbulenceCondition, unit model.LengthUnit) (*float64, error) {
	altitude := LengthFromFeetToUnit(float64(obj.MaxAltitude), unit)

	return &altitude, nil
}
