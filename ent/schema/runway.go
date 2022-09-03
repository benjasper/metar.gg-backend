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
		field.Int("length").Comment("Length of the runway in feet."),
		field.Int("width").Comment("Width of the runway surface in feet."),
		field.String("surface").Optional().Nillable().Comment("Code for the runway surface type. This is not yet a controlled vocabulary, but probably will be soon. Some common values include \"ASP\" (asphalt), \"TURF\" (turf), \"CON\" (concrete), \"GRS\" (grass), \"GRE\" (gravel), \"WATER\" (water), and \"UNK\" (unknown)."),
		field.Bool("lighted").Comment("Whether the runway is lighted at night or not."),
		field.Bool("closed").Comment("Whether the runway is currently closed or not."),

		field.String("low_runway_identifier").Comment("Low numbered runway identifier, like 18R."),
		field.Float("low_runway_latitude").Optional().Nillable().Comment("Latitude of the low numbered runway end, in decimal degrees (positive is north)."),
		field.Float("low_runway_longitude").Optional().Nillable().Comment("Longitude of the low numbered runway end, in decimal degrees (positive is east)."),
		field.Int("low_runway_elevation").Optional().Nillable().Comment("Elevation of the low numbered runway end, in feet."),
		field.Float("low_runway_heading").Optional().Nillable().Comment("True (not magnetic) heading of the lower numbered runway."),
		field.Int("low_runway_displaced_threshold").Optional().Nillable().Comment("Displaced threshold length of the lower numbered runway end, in feet."),

		field.String("high_runway_identifier").Comment("High numbered runway identifier, like 01L."),
		field.Float("high_runway_latitude").Optional().Nillable().Comment("Latitude of the high numbered runway end, in decimal degrees (positive is north)."),
		field.Float("high_runway_longitude").Optional().Nillable().Comment("Longitude of the high numbered runway end, in decimal degrees (positive is east)."),
		field.Int("high_runway_elevation").Optional().Nillable().Comment("Elevation of the high numbered runway end, in feet."),
		field.Float("high_runway_heading").Optional().Nillable().Comment("True (not magnetic) heading of the higher numbered runway."),
		field.Int("high_runway_displaced_threshold").Optional().Nillable().Comment("Displaced threshold length of the higher numbered runway end, in feet."),
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
