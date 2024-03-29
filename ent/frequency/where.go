// Code generated by ent, DO NOT EDIT.

package frequency

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"metar.gg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Frequency {
	return predicate.Frequency(sql.FieldLTE(FieldID, id))
}

// ImportID applies equality check predicate on the "import_id" field. It's identical to ImportIDEQ.
func ImportID(v int) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldImportID, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldHash, v))
}

// ImportFlag applies equality check predicate on the "import_flag" field. It's identical to ImportFlagEQ.
func ImportFlag(v bool) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldImportFlag, v))
}

// LastUpdated applies equality check predicate on the "last_updated" field. It's identical to LastUpdatedEQ.
func LastUpdated(v time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldLastUpdated, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldType, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldDescription, v))
}

// Frequency applies equality check predicate on the "frequency" field. It's identical to FrequencyEQ.
func Frequency(v float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldFrequency, v))
}

// ImportIDEQ applies the EQ predicate on the "import_id" field.
func ImportIDEQ(v int) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldImportID, v))
}

// ImportIDNEQ applies the NEQ predicate on the "import_id" field.
func ImportIDNEQ(v int) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldImportID, v))
}

// ImportIDIn applies the In predicate on the "import_id" field.
func ImportIDIn(vs ...int) predicate.Frequency {
	return predicate.Frequency(sql.FieldIn(FieldImportID, vs...))
}

// ImportIDNotIn applies the NotIn predicate on the "import_id" field.
func ImportIDNotIn(vs ...int) predicate.Frequency {
	return predicate.Frequency(sql.FieldNotIn(FieldImportID, vs...))
}

// ImportIDGT applies the GT predicate on the "import_id" field.
func ImportIDGT(v int) predicate.Frequency {
	return predicate.Frequency(sql.FieldGT(FieldImportID, v))
}

// ImportIDGTE applies the GTE predicate on the "import_id" field.
func ImportIDGTE(v int) predicate.Frequency {
	return predicate.Frequency(sql.FieldGTE(FieldImportID, v))
}

// ImportIDLT applies the LT predicate on the "import_id" field.
func ImportIDLT(v int) predicate.Frequency {
	return predicate.Frequency(sql.FieldLT(FieldImportID, v))
}

// ImportIDLTE applies the LTE predicate on the "import_id" field.
func ImportIDLTE(v int) predicate.Frequency {
	return predicate.Frequency(sql.FieldLTE(FieldImportID, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.Frequency {
	return predicate.Frequency(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.Frequency {
	return predicate.Frequency(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldContainsFold(FieldHash, v))
}

// ImportFlagEQ applies the EQ predicate on the "import_flag" field.
func ImportFlagEQ(v bool) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldImportFlag, v))
}

// ImportFlagNEQ applies the NEQ predicate on the "import_flag" field.
func ImportFlagNEQ(v bool) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldImportFlag, v))
}

// LastUpdatedEQ applies the EQ predicate on the "last_updated" field.
func LastUpdatedEQ(v time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldLastUpdated, v))
}

// LastUpdatedNEQ applies the NEQ predicate on the "last_updated" field.
func LastUpdatedNEQ(v time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldLastUpdated, v))
}

// LastUpdatedIn applies the In predicate on the "last_updated" field.
func LastUpdatedIn(vs ...time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldIn(FieldLastUpdated, vs...))
}

// LastUpdatedNotIn applies the NotIn predicate on the "last_updated" field.
func LastUpdatedNotIn(vs ...time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldNotIn(FieldLastUpdated, vs...))
}

// LastUpdatedGT applies the GT predicate on the "last_updated" field.
func LastUpdatedGT(v time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldGT(FieldLastUpdated, v))
}

// LastUpdatedGTE applies the GTE predicate on the "last_updated" field.
func LastUpdatedGTE(v time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldGTE(FieldLastUpdated, v))
}

// LastUpdatedLT applies the LT predicate on the "last_updated" field.
func LastUpdatedLT(v time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldLT(FieldLastUpdated, v))
}

// LastUpdatedLTE applies the LTE predicate on the "last_updated" field.
func LastUpdatedLTE(v time.Time) predicate.Frequency {
	return predicate.Frequency(sql.FieldLTE(FieldLastUpdated, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Frequency {
	return predicate.Frequency(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Frequency {
	return predicate.Frequency(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldContainsFold(FieldType, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Frequency {
	return predicate.Frequency(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Frequency {
	return predicate.Frequency(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Frequency {
	return predicate.Frequency(sql.FieldContainsFold(FieldDescription, v))
}

// FrequencyEQ applies the EQ predicate on the "frequency" field.
func FrequencyEQ(v float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldEQ(FieldFrequency, v))
}

// FrequencyNEQ applies the NEQ predicate on the "frequency" field.
func FrequencyNEQ(v float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldNEQ(FieldFrequency, v))
}

// FrequencyIn applies the In predicate on the "frequency" field.
func FrequencyIn(vs ...float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldIn(FieldFrequency, vs...))
}

// FrequencyNotIn applies the NotIn predicate on the "frequency" field.
func FrequencyNotIn(vs ...float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldNotIn(FieldFrequency, vs...))
}

// FrequencyGT applies the GT predicate on the "frequency" field.
func FrequencyGT(v float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldGT(FieldFrequency, v))
}

// FrequencyGTE applies the GTE predicate on the "frequency" field.
func FrequencyGTE(v float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldGTE(FieldFrequency, v))
}

// FrequencyLT applies the LT predicate on the "frequency" field.
func FrequencyLT(v float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldLT(FieldFrequency, v))
}

// FrequencyLTE applies the LTE predicate on the "frequency" field.
func FrequencyLTE(v float64) predicate.Frequency {
	return predicate.Frequency(sql.FieldLTE(FieldFrequency, v))
}

// HasAirport applies the HasEdge predicate on the "airport" edge.
func HasAirport() predicate.Frequency {
	return predicate.Frequency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AirportTable, AirportColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAirportWith applies the HasEdge predicate on the "airport" edge with a given conditions (other predicates).
func HasAirportWith(preds ...predicate.Airport) predicate.Frequency {
	return predicate.Frequency(func(s *sql.Selector) {
		step := newAirportStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Frequency) predicate.Frequency {
	return predicate.Frequency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Frequency) predicate.Frequency {
	return predicate.Frequency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Frequency) predicate.Frequency {
	return predicate.Frequency(func(s *sql.Selector) {
		p(s.Not())
	})
}
