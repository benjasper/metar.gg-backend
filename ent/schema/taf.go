package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Taf holds the schema definition for the Metar entity.
type Taf struct {
	ent.Schema
}

// Fields of the Metar.
func (Taf) Fields() []ent.Field {
	return []ent.Field{
		field.Text("raw_text").Comment("The raw TAF text."),
		field.Time("issue_time").Comment("The time the TAF was issued."),
		field.Time("import_time").Comment("The time the TAF was imported.").Default(time.Now),
		field.Time("bulletin_time").Comment("TAF bulletin time."),
		field.Time("valid_from_time").Comment("The start time of the TAF validity period.").Annotations(entgql.OrderField("valid_from_time")),
		field.Time("valid_to_time").Comment("The end time of the TAF validity period."),
		field.String("remarks").Comment("Remarks."),
		field.String("hash").Annotations(entgql.Skip()),
	}
}

// Edges of the Metar.
func (Taf) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("station", WeatherStation.Type).Ref("tafs").Unique().Required().Comment("The station that issued this taf."),
		edge.To("forecast", Forecast.Type).Comment("The forecasts").Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			},
		),
	}
}

// Indexes of the Metar.
func (Taf) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("issue_time"),
	}
}

// Mixin of the Metar.
func (Taf) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
