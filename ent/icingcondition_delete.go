// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/icingcondition"
	"metar.gg/ent/predicate"
)

// IcingConditionDelete is the builder for deleting a IcingCondition entity.
type IcingConditionDelete struct {
	config
	hooks    []Hook
	mutation *IcingConditionMutation
}

// Where appends a list predicates to the IcingConditionDelete builder.
func (icd *IcingConditionDelete) Where(ps ...predicate.IcingCondition) *IcingConditionDelete {
	icd.mutation.Where(ps...)
	return icd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (icd *IcingConditionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(icd.hooks) == 0 {
		affected, err = icd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IcingConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			icd.mutation = mutation
			affected, err = icd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(icd.hooks) - 1; i >= 0; i-- {
			if icd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, icd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (icd *IcingConditionDelete) ExecX(ctx context.Context) int {
	n, err := icd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (icd *IcingConditionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: icingcondition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: icingcondition.FieldID,
			},
		},
	}
	if ps := icd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, icd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// IcingConditionDeleteOne is the builder for deleting a single IcingCondition entity.
type IcingConditionDeleteOne struct {
	icd *IcingConditionDelete
}

// Exec executes the deletion query.
func (icdo *IcingConditionDeleteOne) Exec(ctx context.Context) error {
	n, err := icdo.icd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{icingcondition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (icdo *IcingConditionDeleteOne) ExecX(ctx context.Context) {
	icdo.icd.ExecX(ctx)
}
