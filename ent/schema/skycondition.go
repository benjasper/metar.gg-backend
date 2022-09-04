package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SkyCondition holds the schema definition for the SkyCondition entity.
type SkyCondition struct {
	ent.Schema
}

// Fields of the SkyCondition.
func (SkyCondition) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Annotations(entgql.Skip()),
		field.Enum("sky_cover").NamedValues("Sky Clear", "SKC", "Few", "FEW", "Scattered", "SCT", "Clear", "CLR", "Broken", "BKN", "Overcast", "OVC", "Vertical Visibility", "OVX", "Ceiling and Visibility OK", "CAVOK"),
		field.Int("cloud_base").Optional().Nillable().Comment("Cloud base in feet."),
	}
}

// Edges of the SkyCondition.
func (SkyCondition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metar", Metar.Type).Ref("sky_conditions").Unique().Required(),
	}
}
