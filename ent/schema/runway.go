package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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

		field.String("low_numbered_runway_end_identifier"),
		field.Float("low_numbered_runway_end_latitude"),
		field.Float("low_numbered_runway_end_longitude"),
		field.Int("low_numbered_runway_end_elevation"),
		field.Int("low_numbered_runway_end_heading"),
		field.Int("low_numbered_runway_end_displaced"),

		field.String("high_numbered_runway_end_identifier"),
		field.Float("high_numbered_runway_end_latitude"),
		field.Float("high_numbered_runway_end_longitude"),
		field.Int("high_numbered_runway_end_elevation"),
		field.Int("high_numbered_runway_end_heading"),
		field.Int("high_numbered_runway_end_displaced"),
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
		mixin.Time{},
	}
}
