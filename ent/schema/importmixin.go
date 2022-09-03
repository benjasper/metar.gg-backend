package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// ImportMixin holds the schema definition for the ImportMixin.
type ImportMixin struct {
	mixin.Schema
}

// Fields of the ImportMixin.
func (ImportMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Uint64("hash"),
		field.Bool("import_flag").Default(false),
	}
}

// Indexes of the ImportMixin.
func (ImportMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash").Unique(),
	}
}
