// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/frequency"
	"metar.gg/ent/predicate"
)

// FrequencyDelete is the builder for deleting a Frequency entity.
type FrequencyDelete struct {
	config
	hooks    []Hook
	mutation *FrequencyMutation
}

// Where appends a list predicates to the FrequencyDelete builder.
func (fd *FrequencyDelete) Where(ps ...predicate.Frequency) *FrequencyDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FrequencyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FrequencyDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FrequencyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(frequency.Table, sqlgraph.NewFieldSpec(frequency.FieldID, field.TypeUUID))
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FrequencyDeleteOne is the builder for deleting a single Frequency entity.
type FrequencyDeleteOne struct {
	fd *FrequencyDelete
}

// Where appends a list predicates to the FrequencyDelete builder.
func (fdo *FrequencyDeleteOne) Where(ps ...predicate.Frequency) *FrequencyDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FrequencyDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{frequency.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FrequencyDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
