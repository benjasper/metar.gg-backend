package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// IcingCondition holds the schema definition for the IcingCondition entity.
type IcingCondition struct {
	ent.Schema
}

// Fields of the IcingCondition.
func (IcingCondition) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("intensity").Comment("The intensity of the icing."),
		field.Int("min_altitude").Optional().Nillable().Comment("The minimum altitude in feet that the icing is present."),
		field.Int("max_altitude").Optional().Nillable().Comment("The maximum altitude in feet that the icing is present."),
	}
}

// Edges of the IcingCondition.
func (IcingCondition) Edges() []ent.Edge {
	return nil
}
