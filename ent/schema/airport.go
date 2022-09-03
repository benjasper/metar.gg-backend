package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Airport holds the schema definition for the Airport entity.
type Airport struct {
	ent.Schema
}

// Fields of the Airport.
func (Airport) Fields() []ent.Field {
	return []ent.Field{
		field.String("identifier"),
		field.String("type"),
		field.String("name"),
		field.Float("latitude"),
		field.Float("longitude"),
		field.Int("elevation"),
		field.String("continent"),
		field.String("country"),
		field.String("region"),
		field.String("municipality"),
		field.Bool("scheduled_service"),
		field.String("gps_code"),
		field.String("iata_code"),
		field.String("local_code"),
		field.String("website"),
		field.String("wikipedia"),
		field.Strings("keywords"),
	}
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("runways", Runway.Type),
	}
}

// Indexes of the Airport.
func (Airport) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("identifier"),
	}
}

// Mixin of the Airport.
func (Airport) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ImportMixin{},
		mixin.Time{},
	}
}
