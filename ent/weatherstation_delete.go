// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/predicate"
	"metar.gg/ent/weatherstation"
)

// WeatherStationDelete is the builder for deleting a WeatherStation entity.
type WeatherStationDelete struct {
	config
	hooks    []Hook
	mutation *WeatherStationMutation
}

// Where appends a list predicates to the WeatherStationDelete builder.
func (wsd *WeatherStationDelete) Where(ps ...predicate.WeatherStation) *WeatherStationDelete {
	wsd.mutation.Where(ps...)
	return wsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wsd *WeatherStationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, WeatherStationMutation](ctx, wsd.sqlExec, wsd.mutation, wsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wsd *WeatherStationDelete) ExecX(ctx context.Context) int {
	n, err := wsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wsd *WeatherStationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(weatherstation.Table, sqlgraph.NewFieldSpec(weatherstation.FieldID, field.TypeUUID))
	if ps := wsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wsd.mutation.done = true
	return affected, err
}

// WeatherStationDeleteOne is the builder for deleting a single WeatherStation entity.
type WeatherStationDeleteOne struct {
	wsd *WeatherStationDelete
}

// Where appends a list predicates to the WeatherStationDelete builder.
func (wsdo *WeatherStationDeleteOne) Where(ps ...predicate.WeatherStation) *WeatherStationDeleteOne {
	wsdo.wsd.mutation.Where(ps...)
	return wsdo
}

// Exec executes the deletion query.
func (wsdo *WeatherStationDeleteOne) Exec(ctx context.Context) error {
	n, err := wsdo.wsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{weatherstation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdo *WeatherStationDeleteOne) ExecX(ctx context.Context) {
	if err := wsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
