package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TurbulenceCondition holds the schema definition for the TurbulenceCondition entity.
type TurbulenceCondition struct {
	ent.Schema
}

// Fields of the TurbulenceCondition.
func (TurbulenceCondition) Fields() []ent.Field {
	return []ent.Field{
		field.String("intensity").Comment("The intensity of the turbulence."),
		field.Int("min_altitude").Comment("The minimum altitude in feet that the turbulence is present."),
		field.Int("max_altitude").Comment("The maximum altitude in feet that the turbulence is present."),
	}
}

// Edges of the TurbulenceCondition.
func (TurbulenceCondition) Edges() []ent.Edge {
	return nil
}
