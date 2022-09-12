package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"time"
)

// ImportMixin holds the schema definition for the ImportMixin.
type ImportMixin struct {
	mixin.Schema
}

func now() time.Time {
	return time.Now()
}

// Fields of the ImportMixin.
func (ImportMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("import_id").Comment("The unique identifier of the import."),
		field.String("hash").Annotations(entgql.Skip()),
		field.Bool("import_flag").Default(false).Annotations(entgql.Skip()),
		field.Time("last_updated").Default(now).Comment("The last time the record was updated/created."),
	}
}

// Indexes of the ImportMixin.
func (ImportMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash"),
		index.Fields("import_id"),
	}
}
