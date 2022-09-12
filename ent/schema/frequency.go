package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Frequency holds the schema definition for the Frequency entity.
type Frequency struct {
	ent.Schema
}

// Fields of the Frequency.
func (Frequency) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("A code for the frequency type. Some common values are \"TWR\" (tower), \"ATF\" or \"CTAF\" (common traffic frequency), \"GND\" (ground control), \"RMP\" (ramp control), \"ATIS\" (automated weather), \"RCO\" (remote radio outlet), \"ARR\" (arrivals), \"DEP\" (departures), \"UNICOM\" (monitored ground station), and \"RDO\" (a flight-service station)."),
		field.String("description").Comment("A description of the frequency."),
		field.Float("frequency").Comment("Radio frequency in megahertz. Note that the same frequency may appear multiple times for an airport, serving different functions"),
	}
}

// Edges of the Frequency.
func (Frequency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("airport", Airport.Type).
			Ref("frequencies").
			Unique(),
	}
}

// Indexes of the Frequency.
func (Frequency) Indexes() []ent.Index {
	return nil
}

// Mixin of the Frequency.
func (Frequency) Mixin() []ent.Mixin {
	return []ent.Mixin{ImportMixin{}, IDMixin{}}
}
