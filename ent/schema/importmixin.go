package schema

import (
	"entgo.io/contrib/entgql"
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
		field.String("hash").Annotations(entgql.Skip()),
		field.Bool("import_flag").Default(false).Annotations(entgql.Skip()),
	}
}

// Indexes of the ImportMixin.
func (ImportMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash"),
	}
}
