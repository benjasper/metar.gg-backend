package schema

import "entgo.io/ent"

// Frequency holds the schema definition for the Frequency entity.
type Frequency struct {
	ent.Schema
}

// Fields of the Frequency.
func (Frequency) Fields() []ent.Field {
	return nil
}

// Edges of the Frequency.
func (Frequency) Edges() []ent.Edge {
	return nil
}
