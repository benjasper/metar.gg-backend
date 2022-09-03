package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Airport holds the schema definition for the Airport entity.
type Airport struct {
	ent.Schema
}

// Fields of the Airport.
func (Airport) Fields() []ent.Field {
	return []ent.Field{
		field.String("identifier"),
		field.Enum("type").Values("large_airport", "medium_airport", "small_airport", "closed_airport", "heliport", "seaplane_base"),
		field.String("name"),
		field.Float("latitude"),
		field.Float("longitude"),
		field.Int("elevation").Optional().Nillable(),
		field.String("continent"),
		field.String("country"),
		field.String("region"),
		field.String("municipality"),
		field.Bool("scheduled_service"),
		field.String("gps_code").Optional().Nillable(),
		field.String("iata_code").Optional().Nillable(),
		field.String("local_code").Optional().Nillable(),
		field.String("website").Optional().Nillable(),
		field.String("wikipedia").Optional().Nillable(),
		field.Strings("keywords"),
	}
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("runways", Runway.Type).Required(),
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
		//mixin.Time{},
	}
}
