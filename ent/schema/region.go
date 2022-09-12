package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Region holds the schema definition for the Region entity.
type Region struct {
	ent.Schema
}

// Fields of the Region. CSV fields: "id","code","local_code","name","continent","iso_country","wikipedia_link","keywords"
func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Comment("local_code prefixed with the country code to make a globally-unique identifier."),
		field.String("local_code").Comment("The local code for the administrative subdivision. Whenever possible, these are official ISO 3166:2, at the highest level available, but in some cases OurAirports has to use unofficial codes. There is also a pseudo code \"U-A\" for each country, which means that the airport has not yet been assigned to a region (or perhaps can't be, as in the case of a deep-sea oil platform)."),
		field.String("name"),
		field.String("wikipedia_link").Comment("The wikipedia link of the region."),
		field.Strings("keywords").Comment("Keywords that can be used to search for the region."),
	}
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("airports", Airport.Type).Annotations(entgql.Skip()),
	}
}

// Mixin of the Region.
func (Region) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		ImportMixin{},
	}
}
