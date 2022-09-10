package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Station holds the schema definition for the Metar entity.
type Station struct {
	ent.Schema
}

// Fields of the Metar.
func (Station) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Annotations(entgql.Skip()),
		field.String("station_id").Immutable().Unique().Comment("The ICAO identifier of the station that provided the weather data or identifier of the weather station."),
		field.Float("latitude").Optional().Nillable().Comment("The latitude in decimal degrees of the station."),
		field.Float("longitude").Optional().Nillable().Comment("The longitude in decimal degrees of the station."),
		field.Float("elevation").Optional().Nillable().Comment("The elevation in meters of the station."),

		field.String("hash").Annotations(entgql.Skip()),
	}
}

// Edges of the Metar.
func (Station) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("airport", Airport.Type).Ref("station").Unique().Comment("The airport that hosts this station. This can also be empty if the metar is from a weather station outside an airport."),
		edge.To("metars", Metar.Type).Annotations(entgql.Skip()).Comment("The metars that were reported by this station."),
		edge.To("tafs", Taf.Type).Annotations(entgql.Skip()).Comment("The tafs that were reported by this station."),
	}
}

// Indexes of the Metar.
func (Station) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("station_id").Unique(),
		index.Fields("latitude"),
		index.Fields("longitude"),
	}
}

// Mixin of the Metar.
func (Station) Mixin() []ent.Mixin {
	return nil
}
