// Code generated by ent, DO NOT EDIT.

package skycondition

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"metar.gg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CloudBase applies equality check predicate on the "cloud_base" field. It's identical to CloudBaseEQ.
func CloudBase(v int) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCloudBase), v))
	})
}

// SkyCoverEQ applies the EQ predicate on the "sky_cover" field.
func SkyCoverEQ(v SkyCover) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkyCover), v))
	})
}

// SkyCoverNEQ applies the NEQ predicate on the "sky_cover" field.
func SkyCoverNEQ(v SkyCover) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkyCover), v))
	})
}

// SkyCoverIn applies the In predicate on the "sky_cover" field.
func SkyCoverIn(vs ...SkyCover) predicate.SkyCondition {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSkyCover), v...))
	})
}

// SkyCoverNotIn applies the NotIn predicate on the "sky_cover" field.
func SkyCoverNotIn(vs ...SkyCover) predicate.SkyCondition {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSkyCover), v...))
	})
}

// CloudBaseEQ applies the EQ predicate on the "cloud_base" field.
func CloudBaseEQ(v int) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCloudBase), v))
	})
}

// CloudBaseNEQ applies the NEQ predicate on the "cloud_base" field.
func CloudBaseNEQ(v int) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCloudBase), v))
	})
}

// CloudBaseIn applies the In predicate on the "cloud_base" field.
func CloudBaseIn(vs ...int) predicate.SkyCondition {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCloudBase), v...))
	})
}

// CloudBaseNotIn applies the NotIn predicate on the "cloud_base" field.
func CloudBaseNotIn(vs ...int) predicate.SkyCondition {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCloudBase), v...))
	})
}

// CloudBaseGT applies the GT predicate on the "cloud_base" field.
func CloudBaseGT(v int) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCloudBase), v))
	})
}

// CloudBaseGTE applies the GTE predicate on the "cloud_base" field.
func CloudBaseGTE(v int) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCloudBase), v))
	})
}

// CloudBaseLT applies the LT predicate on the "cloud_base" field.
func CloudBaseLT(v int) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCloudBase), v))
	})
}

// CloudBaseLTE applies the LTE predicate on the "cloud_base" field.
func CloudBaseLTE(v int) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCloudBase), v))
	})
}

// CloudBaseIsNil applies the IsNil predicate on the "cloud_base" field.
func CloudBaseIsNil() predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCloudBase)))
	})
}

// CloudBaseNotNil applies the NotNil predicate on the "cloud_base" field.
func CloudBaseNotNil() predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCloudBase)))
	})
}

// CloudTypeEQ applies the EQ predicate on the "cloud_type" field.
func CloudTypeEQ(v CloudType) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCloudType), v))
	})
}

// CloudTypeNEQ applies the NEQ predicate on the "cloud_type" field.
func CloudTypeNEQ(v CloudType) predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCloudType), v))
	})
}

// CloudTypeIn applies the In predicate on the "cloud_type" field.
func CloudTypeIn(vs ...CloudType) predicate.SkyCondition {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCloudType), v...))
	})
}

// CloudTypeNotIn applies the NotIn predicate on the "cloud_type" field.
func CloudTypeNotIn(vs ...CloudType) predicate.SkyCondition {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCloudType), v...))
	})
}

// CloudTypeIsNil applies the IsNil predicate on the "cloud_type" field.
func CloudTypeIsNil() predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCloudType)))
	})
}

// CloudTypeNotNil applies the NotNil predicate on the "cloud_type" field.
func CloudTypeNotNil() predicate.SkyCondition {
	return predicate.SkyCondition(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCloudType)))
	})
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
