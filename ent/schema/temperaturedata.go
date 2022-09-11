package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TemperatureData holds the schema definition for the TemperatureData entity.
type TemperatureData struct {
	ent.Schema
}

// Fields of the TemperatureData.
func (TemperatureData) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Time("valid_time").Comment("The time the temperature data is valid."),
		field.Float("temperature").Comment("The surface temperature in degrees Celsius."),
		field.Float("min_temperature").Optional().Nillable().Comment("The minimum temperature in degrees Celsius."),
		field.Float("max_temperature").Optional().Nillable().Comment("The maximum temperature in degrees Celsius."),
	}
}

// Edges of the TemperatureData.
func (TemperatureData) Edges() []ent.Edge {
	return nil
}
