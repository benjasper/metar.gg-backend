// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/predicate"
	"metar.gg/ent/skycondition"
)

// SkyConditionDelete is the builder for deleting a SkyCondition entity.
type SkyConditionDelete struct {
	config
	hooks    []Hook
	mutation *SkyConditionMutation
}

// Where appends a list predicates to the SkyConditionDelete builder.
func (scd *SkyConditionDelete) Where(ps ...predicate.SkyCondition) *SkyConditionDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SkyConditionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(scd.hooks) == 0 {
		affected, err = scd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkyConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scd.mutation = mutation
			affected, err = scd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scd.hooks) - 1; i >= 0; i-- {
			if scd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SkyConditionDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SkyConditionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: skycondition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: skycondition.FieldID,
			},
		},
	}
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SkyConditionDeleteOne is the builder for deleting a single SkyCondition entity.
type SkyConditionDeleteOne struct {
	scd *SkyConditionDelete
}

// Exec executes the deletion query.
func (scdo *SkyConditionDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{skycondition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SkyConditionDeleteOne) ExecX(ctx context.Context) {
	scdo.scd.ExecX(ctx)
}
