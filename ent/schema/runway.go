package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Runway holds the schema definition for the Runway entity.
type Runway struct {
	ent.Schema
}

// Fields of the Runway.
func (Runway) Fields() []ent.Field {
	return []ent.Field{
		field.String("airport_identifier"),
		field.Int("length"),
		field.Int("width"),
		field.String("surface"),
		field.Bool("lighted"),
		field.Bool("closed"),

		field.String("low_runway_identifier").Comment("Low numbered runway identifier, like 18R"),
		field.Float("low_runway_latitude").Optional().Nillable(),
		field.Float("low_runway_longitude").Optional().Nillable(),
		field.Int("low_runway_elevation").Optional().Nillable(),
		field.Int("low_runway_heading").Optional().Nillable(),
		field.Int("low_runway_displaced").Optional().Nillable(),

		field.String("high_runway_identifier").Comment("High numbered runway identifier, like 18R"),
		field.Float("high_runway_latitude").Optional().Nillable(),
		field.Float("high_runway_longitude").Optional().Nillable(),
		field.Int("high_runway_elevation").Optional().Nillable(),
		field.Int("high_runway_heading").Optional().Nillable(),
		field.Int("high_runway_displaced").Optional().Nillable(),
	}
}

// Edges of the Runway.
func (Runway) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("airport", Airport.Type).
			Ref("runways").
			Unique(),
	}
}

// Indexes of the Runway.
func (Runway) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of the Runway.
func (Runway) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ImportMixin{},
		//mixin.Time{},
	}
}
