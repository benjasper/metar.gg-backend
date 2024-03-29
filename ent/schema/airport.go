package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.String("icao_code").Optional().MaxLen(4).NotEmpty().Annotations(entgql.OrderField("ICAO_CODE")).Comment("The four-letter ICAO code of the airport."),
		field.String("iata_code").Optional().Nillable().Comment("The three-letter IATA code for the airport."),
		field.String("identifier").Unique().Comment("This will be the ICAO code if available. Otherwise, it will be a local airport code (if no conflict), or if nothing else is available, an internally-generated code starting with the ISO2 country code, followed by a dash and a four-digit number."),
		field.Enum("type").Values("large_airport", "medium_airport", "small_airport", "closed_airport", "heliport", "seaplane_base").Comment("Type of airport."),
		field.Int("importance").Default(0).Comment("Importance of the airport.").Annotations(entgql.OrderField("IMPORTANCE")),
		field.String("name").Comment("The official airport name, including \"Airport\", \"Airstrip\", etc."),
		field.Float("latitude").Comment("Latitude of the airport in decimal degrees (positive is north)."),
		field.Float("longitude").Comment("Longitude of the airport in decimal degrees (positive is east)."),
		field.String("timezone").Optional().Nillable().Comment("The timezone of the airport."),
		field.Int("elevation").Optional().Nillable().Comment("Elevation of the airport, in feet."),
		field.String("municipality").Optional().Nillable().Comment("The primary municipality that the airport serves (when available). Note that this is not necessarily the municipality where the airport is physically located."),
		field.Bool("scheduled_service").Comment("Whether the airport has scheduled airline service."),
		field.String("gps_code").Optional().Nillable().Comment("The code that an aviation GPS database (such as Jeppesen's or Garmin's) would normally use for the airport. This will always be the ICAO code if one exists. Note that, unlike the ident column, this is not guaranteed to be globally unique."),
		field.String("local_code").Optional().Nillable().Comment("The local country code for the airport, if different from the gps_code and iata_code fields (used mainly for US airports)."),
		field.String("website").Optional().Nillable().Comment("The URL of the airport's website."),
		field.String("wikipedia").Optional().Nillable().Comment("The URL of the airport's Wikipedia page."),
		field.Strings("keywords").Comment("Extra keywords/phrases to assist with search. May include former names for the airport, alternate codes, names in other languages, nearby tourist destinations, etc."),
	}
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("region", Region.Type).Ref("airports").Unique().Comment("The region that the airport is located in."),
		edge.From("country", Country.Type).Ref("airports").Unique().Comment("The country that the airport is located in."),
		edge.To("runways", Runway.Type).Comment("Runways at the airport.").Annotations(entgql.Skip(), entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("frequencies", Frequency.Type).Comment("Frequencies at the airport.").Annotations(
			entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("station", WeatherStation.Type).Unique().Comment("Weather station at the airport."),
	}
}

// Indexes of the Airport.
func (Airport) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "municipality", "icao_code", "iata_code", "local_code", "identifier").
			Annotations(
				entsql.IndexTypes(map[string]string{
					dialect.MySQL: "FULLTEXT",
				}),
			).StorageKey("fulltext"),
		index.Fields("identifier"),
		index.Fields("name"),
		index.Fields("municipality"),
		index.Fields("local_code"),
		index.Fields("icao_code"),
		index.Fields("iata_code"),
	}
}

// Mixin of the Airport.
func (Airport) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ImportMixin{},
		IDMixin{},
	}
}

// Annotations of the Airport.
func (Airport) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_unicode_520_ci",
		},
	}
}
