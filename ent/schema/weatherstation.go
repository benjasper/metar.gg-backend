package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// WeatherStation holds the schema definition for the WeatherStation entity.
type WeatherStation struct {
	ent.Schema
}

// Fields of the Metar.
func (WeatherStation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("station_id").Immutable().Unique().Comment("The ICAO identifier of the station that provided the weather data or identifier of the weather station."),
		field.Float("latitude").Optional().Nillable().Comment("The latitude in decimal degrees of the station."),
		field.Float("longitude").Optional().Nillable().Comment("The longitude in decimal degrees of the station."),
		field.Float("elevation").Optional().Nillable().Comment("The elevation in meters of the station."),

		field.String("hash").Annotations(entgql.Skip()),
	}
}

// Edges of the Metar.
func (WeatherStation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("airport", Airport.Type).Ref("station").Unique().Comment("The airport that hosts this station. This can also be empty if the metar is from a weather station outside an airport."),
		edge.To("metars", Metar.Type).Annotations(entgql.Skip()).Comment("The metars that were reported by this station."),
		edge.To("tafs", Taf.Type).Annotations(entgql.Skip()).Comment("The tafs that were reported by this station."),
	}
}

// Indexes of the Metar.
func (WeatherStation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("station_id").Unique(),
		index.Fields("latitude"),
		index.Fields("longitude"),
		index.Fields("hash"),
	}
}

// Mixin of the Metar.
func (WeatherStation) Mixin() []ent.Mixin {
	return nil
}
