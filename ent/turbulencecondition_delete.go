// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/predicate"
	"metar.gg/ent/turbulencecondition"
)

// TurbulenceConditionDelete is the builder for deleting a TurbulenceCondition entity.
type TurbulenceConditionDelete struct {
	config
	hooks    []Hook
	mutation *TurbulenceConditionMutation
}

// Where appends a list predicates to the TurbulenceConditionDelete builder.
func (tcd *TurbulenceConditionDelete) Where(ps ...predicate.TurbulenceCondition) *TurbulenceConditionDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TurbulenceConditionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tcd.hooks) == 0 {
		affected, err = tcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TurbulenceConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcd.mutation = mutation
			affected, err = tcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcd.hooks) - 1; i >= 0; i-- {
			if tcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TurbulenceConditionDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TurbulenceConditionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: turbulencecondition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: turbulencecondition.FieldID,
			},
		},
	}
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TurbulenceConditionDeleteOne is the builder for deleting a single TurbulenceCondition entity.
type TurbulenceConditionDeleteOne struct {
	tcd *TurbulenceConditionDelete
}

// Exec executes the deletion query.
func (tcdo *TurbulenceConditionDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{turbulencecondition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TurbulenceConditionDeleteOne) ExecX(ctx context.Context) {
	tcdo.tcd.ExecX(ctx)
}
