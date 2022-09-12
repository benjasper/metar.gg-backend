package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// IDMixin holds the schema definition for the IDMixin entity.
type IDMixin struct {
	mixin.Schema
}

// Fields of the IDMixin.
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Comment("The unique identifier of the record."),
	}
}

// Edges of the IDMixin.
func (IDMixin) Edges() []ent.Edge {
	return nil
}
