// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
)

// MetarDelete is the builder for deleting a Metar entity.
type MetarDelete struct {
	config
	hooks    []Hook
	mutation *MetarMutation
}

// Where appends a list predicates to the MetarDelete builder.
func (md *MetarDelete) Where(ps ...predicate.Metar) *MetarDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MetarDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MetarMutation](ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MetarDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MetarDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(metar.Table, sqlgraph.NewFieldSpec(metar.FieldID, field.TypeUUID))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MetarDeleteOne is the builder for deleting a single Metar entity.
type MetarDeleteOne struct {
	md *MetarDelete
}

// Where appends a list predicates to the MetarDelete builder.
func (mdo *MetarDeleteOne) Where(ps ...predicate.Metar) *MetarDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MetarDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{metar.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MetarDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
