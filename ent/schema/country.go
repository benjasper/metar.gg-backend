package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country. CSV fields: "id","code","name","continent","wikipedia_link","keywords"
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Immutable().Unique().Comment("The ISO 3166-1 alpha-2 code of the country. A handful of unofficial, non-ISO codes are also in use, such as \"XK\" for Kosovo."),
		field.String("name").Comment("The name of the country."),
		field.Enum("continent").NamedValues("Africa", "AF", "Antarctica", "AN", "Asia", "AS", "Europe", "EU", "North America", "NA", "South America", "SA", "Oceania", "OC").Comment("Where the airport is (primarily) located."),
		field.String("wikipedia_link").Comment("The wikipedia link of the country."),
		field.Strings("keywords").Comment("Keywords that can be used to search for the country."),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("airports", Airport.Type).Annotations(entgql.Skip()),
	}
}

// Mixin of the Country.
func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		ImportMixin{},
	}
}
