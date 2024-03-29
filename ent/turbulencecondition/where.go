// Code generated by ent, DO NOT EDIT.

package turbulencecondition

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"metar.gg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLTE(FieldID, id))
}

// Intensity applies equality check predicate on the "intensity" field. It's identical to IntensityEQ.
func Intensity(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldIntensity, v))
}

// MinAltitude applies equality check predicate on the "min_altitude" field. It's identical to MinAltitudeEQ.
func MinAltitude(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldMinAltitude, v))
}

// MaxAltitude applies equality check predicate on the "max_altitude" field. It's identical to MaxAltitudeEQ.
func MaxAltitude(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldMaxAltitude, v))
}

// IntensityEQ applies the EQ predicate on the "intensity" field.
func IntensityEQ(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldIntensity, v))
}

// IntensityNEQ applies the NEQ predicate on the "intensity" field.
func IntensityNEQ(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNEQ(FieldIntensity, v))
}

// IntensityIn applies the In predicate on the "intensity" field.
func IntensityIn(vs ...string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldIn(FieldIntensity, vs...))
}

// IntensityNotIn applies the NotIn predicate on the "intensity" field.
func IntensityNotIn(vs ...string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNotIn(FieldIntensity, vs...))
}

// IntensityGT applies the GT predicate on the "intensity" field.
func IntensityGT(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGT(FieldIntensity, v))
}

// IntensityGTE applies the GTE predicate on the "intensity" field.
func IntensityGTE(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGTE(FieldIntensity, v))
}

// IntensityLT applies the LT predicate on the "intensity" field.
func IntensityLT(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLT(FieldIntensity, v))
}

// IntensityLTE applies the LTE predicate on the "intensity" field.
func IntensityLTE(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLTE(FieldIntensity, v))
}

// IntensityContains applies the Contains predicate on the "intensity" field.
func IntensityContains(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldContains(FieldIntensity, v))
}

// IntensityHasPrefix applies the HasPrefix predicate on the "intensity" field.
func IntensityHasPrefix(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldHasPrefix(FieldIntensity, v))
}

// IntensityHasSuffix applies the HasSuffix predicate on the "intensity" field.
func IntensityHasSuffix(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldHasSuffix(FieldIntensity, v))
}

// IntensityEqualFold applies the EqualFold predicate on the "intensity" field.
func IntensityEqualFold(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEqualFold(FieldIntensity, v))
}

// IntensityContainsFold applies the ContainsFold predicate on the "intensity" field.
func IntensityContainsFold(v string) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldContainsFold(FieldIntensity, v))
}

// MinAltitudeEQ applies the EQ predicate on the "min_altitude" field.
func MinAltitudeEQ(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldMinAltitude, v))
}

// MinAltitudeNEQ applies the NEQ predicate on the "min_altitude" field.
func MinAltitudeNEQ(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNEQ(FieldMinAltitude, v))
}

// MinAltitudeIn applies the In predicate on the "min_altitude" field.
func MinAltitudeIn(vs ...int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldIn(FieldMinAltitude, vs...))
}

// MinAltitudeNotIn applies the NotIn predicate on the "min_altitude" field.
func MinAltitudeNotIn(vs ...int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNotIn(FieldMinAltitude, vs...))
}

// MinAltitudeGT applies the GT predicate on the "min_altitude" field.
func MinAltitudeGT(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGT(FieldMinAltitude, v))
}

// MinAltitudeGTE applies the GTE predicate on the "min_altitude" field.
func MinAltitudeGTE(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGTE(FieldMinAltitude, v))
}

// MinAltitudeLT applies the LT predicate on the "min_altitude" field.
func MinAltitudeLT(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLT(FieldMinAltitude, v))
}

// MinAltitudeLTE applies the LTE predicate on the "min_altitude" field.
func MinAltitudeLTE(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLTE(FieldMinAltitude, v))
}

// MaxAltitudeEQ applies the EQ predicate on the "max_altitude" field.
func MaxAltitudeEQ(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldEQ(FieldMaxAltitude, v))
}

// MaxAltitudeNEQ applies the NEQ predicate on the "max_altitude" field.
func MaxAltitudeNEQ(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNEQ(FieldMaxAltitude, v))
}

// MaxAltitudeIn applies the In predicate on the "max_altitude" field.
func MaxAltitudeIn(vs ...int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldIn(FieldMaxAltitude, vs...))
}

// MaxAltitudeNotIn applies the NotIn predicate on the "max_altitude" field.
func MaxAltitudeNotIn(vs ...int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldNotIn(FieldMaxAltitude, vs...))
}

// MaxAltitudeGT applies the GT predicate on the "max_altitude" field.
func MaxAltitudeGT(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGT(FieldMaxAltitude, v))
}

// MaxAltitudeGTE applies the GTE predicate on the "max_altitude" field.
func MaxAltitudeGTE(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldGTE(FieldMaxAltitude, v))
}

// MaxAltitudeLT applies the LT predicate on the "max_altitude" field.
func MaxAltitudeLT(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLT(FieldMaxAltitude, v))
}

// MaxAltitudeLTE applies the LTE predicate on the "max_altitude" field.
func MaxAltitudeLTE(v int) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(sql.FieldLTE(FieldMaxAltitude, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TurbulenceCondition) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TurbulenceCondition) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(func(s *sql.Selector) {
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
func Not(p predicate.TurbulenceCondition) predicate.TurbulenceCondition {
	return predicate.TurbulenceCondition(func(s *sql.Selector) {
		p(s.Not())
	})
}
