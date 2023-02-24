// Code generated by ent, DO NOT EDIT.

package skycondition

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"metar.gg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldLTE(FieldID, id))
}

// CloudBase applies equality check predicate on the "cloud_base" field. It's identical to CloudBaseEQ.
func CloudBase(v int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldEQ(FieldCloudBase, v))
}

// SkyCoverEQ applies the EQ predicate on the "sky_cover" field.
func SkyCoverEQ(v SkyCover) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldEQ(FieldSkyCover, v))
}

// SkyCoverNEQ applies the NEQ predicate on the "sky_cover" field.
func SkyCoverNEQ(v SkyCover) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNEQ(FieldSkyCover, v))
}

// SkyCoverIn applies the In predicate on the "sky_cover" field.
func SkyCoverIn(vs ...SkyCover) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldIn(FieldSkyCover, vs...))
}

// SkyCoverNotIn applies the NotIn predicate on the "sky_cover" field.
func SkyCoverNotIn(vs ...SkyCover) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNotIn(FieldSkyCover, vs...))
}

// CloudBaseEQ applies the EQ predicate on the "cloud_base" field.
func CloudBaseEQ(v int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldEQ(FieldCloudBase, v))
}

// CloudBaseNEQ applies the NEQ predicate on the "cloud_base" field.
func CloudBaseNEQ(v int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNEQ(FieldCloudBase, v))
}

// CloudBaseIn applies the In predicate on the "cloud_base" field.
func CloudBaseIn(vs ...int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldIn(FieldCloudBase, vs...))
}

// CloudBaseNotIn applies the NotIn predicate on the "cloud_base" field.
func CloudBaseNotIn(vs ...int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNotIn(FieldCloudBase, vs...))
}

// CloudBaseGT applies the GT predicate on the "cloud_base" field.
func CloudBaseGT(v int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldGT(FieldCloudBase, v))
}

// CloudBaseGTE applies the GTE predicate on the "cloud_base" field.
func CloudBaseGTE(v int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldGTE(FieldCloudBase, v))
}

// CloudBaseLT applies the LT predicate on the "cloud_base" field.
func CloudBaseLT(v int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldLT(FieldCloudBase, v))
}

// CloudBaseLTE applies the LTE predicate on the "cloud_base" field.
func CloudBaseLTE(v int) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldLTE(FieldCloudBase, v))
}

// CloudBaseIsNil applies the IsNil predicate on the "cloud_base" field.
func CloudBaseIsNil() predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldIsNull(FieldCloudBase))
}

// CloudBaseNotNil applies the NotNil predicate on the "cloud_base" field.
func CloudBaseNotNil() predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNotNull(FieldCloudBase))
}

// CloudTypeEQ applies the EQ predicate on the "cloud_type" field.
func CloudTypeEQ(v CloudType) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldEQ(FieldCloudType, v))
}

// CloudTypeNEQ applies the NEQ predicate on the "cloud_type" field.
func CloudTypeNEQ(v CloudType) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNEQ(FieldCloudType, v))
}

// CloudTypeIn applies the In predicate on the "cloud_type" field.
func CloudTypeIn(vs ...CloudType) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldIn(FieldCloudType, vs...))
}

// CloudTypeNotIn applies the NotIn predicate on the "cloud_type" field.
func CloudTypeNotIn(vs ...CloudType) predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNotIn(FieldCloudType, vs...))
}

// CloudTypeIsNil applies the IsNil predicate on the "cloud_type" field.
func CloudTypeIsNil() predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldIsNull(FieldCloudType))
}

// CloudTypeNotNil applies the NotNil predicate on the "cloud_type" field.
func CloudTypeNotNil() predicate.SkyCondition {
	return predicate.SkyCondition(sql.FieldNotNull(FieldCloudType))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SkyCondition) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SkyCondition) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
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
func Not(p predicate.SkyCondition) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		p(s.Not())
	})
}
