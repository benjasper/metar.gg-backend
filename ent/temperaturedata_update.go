// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/predicate"
	"metar.gg/ent/temperaturedata"
)

// TemperatureDataUpdate is the builder for updating TemperatureData entities.
type TemperatureDataUpdate struct {
	config
	hooks     []Hook
	mutation  *TemperatureDataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TemperatureDataUpdate builder.
func (tdu *TemperatureDataUpdate) Where(ps ...predicate.TemperatureData) *TemperatureDataUpdate {
	tdu.mutation.Where(ps...)
	return tdu
}

// SetValidTime sets the "valid_time" field.
func (tdu *TemperatureDataUpdate) SetValidTime(t time.Time) *TemperatureDataUpdate {
	tdu.mutation.SetValidTime(t)
	return tdu
}

// SetTemperature sets the "temperature" field.
func (tdu *TemperatureDataUpdate) SetTemperature(f float64) *TemperatureDataUpdate {
	tdu.mutation.ResetTemperature()
	tdu.mutation.SetTemperature(f)
	return tdu
}

// AddTemperature adds f to the "temperature" field.
func (tdu *TemperatureDataUpdate) AddTemperature(f float64) *TemperatureDataUpdate {
	tdu.mutation.AddTemperature(f)
	return tdu
}

// SetMinTemperature sets the "min_temperature" field.
func (tdu *TemperatureDataUpdate) SetMinTemperature(f float64) *TemperatureDataUpdate {
	tdu.mutation.ResetMinTemperature()
	tdu.mutation.SetMinTemperature(f)
	return tdu
}

// SetNillableMinTemperature sets the "min_temperature" field if the given value is not nil.
func (tdu *TemperatureDataUpdate) SetNillableMinTemperature(f *float64) *TemperatureDataUpdate {
	if f != nil {
		tdu.SetMinTemperature(*f)
	}
	return tdu
}

// AddMinTemperature adds f to the "min_temperature" field.
func (tdu *TemperatureDataUpdate) AddMinTemperature(f float64) *TemperatureDataUpdate {
	tdu.mutation.AddMinTemperature(f)
	return tdu
}

// ClearMinTemperature clears the value of the "min_temperature" field.
func (tdu *TemperatureDataUpdate) ClearMinTemperature() *TemperatureDataUpdate {
	tdu.mutation.ClearMinTemperature()
	return tdu
}

// SetMaxTemperature sets the "max_temperature" field.
func (tdu *TemperatureDataUpdate) SetMaxTemperature(f float64) *TemperatureDataUpdate {
	tdu.mutation.ResetMaxTemperature()
	tdu.mutation.SetMaxTemperature(f)
	return tdu
}

// SetNillableMaxTemperature sets the "max_temperature" field if the given value is not nil.
func (tdu *TemperatureDataUpdate) SetNillableMaxTemperature(f *float64) *TemperatureDataUpdate {
	if f != nil {
		tdu.SetMaxTemperature(*f)
	}
	return tdu
}

// AddMaxTemperature adds f to the "max_temperature" field.
func (tdu *TemperatureDataUpdate) AddMaxTemperature(f float64) *TemperatureDataUpdate {
	tdu.mutation.AddMaxTemperature(f)
	return tdu
}

// ClearMaxTemperature clears the value of the "max_temperature" field.
func (tdu *TemperatureDataUpdate) ClearMaxTemperature() *TemperatureDataUpdate {
	tdu.mutation.ClearMaxTemperature()
	return tdu
}

// Mutation returns the TemperatureDataMutation object of the builder.
func (tdu *TemperatureDataUpdate) Mutation() *TemperatureDataMutation {
	return tdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tdu *TemperatureDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TemperatureDataMutation](ctx, tdu.sqlSave, tdu.mutation, tdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tdu *TemperatureDataUpdate) SaveX(ctx context.Context) int {
	affected, err := tdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tdu *TemperatureDataUpdate) Exec(ctx context.Context) error {
	_, err := tdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdu *TemperatureDataUpdate) ExecX(ctx context.Context) {
	if err := tdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tdu *TemperatureDataUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemperatureDataUpdate {
	tdu.modifiers = append(tdu.modifiers, modifiers...)
	return tdu
}

func (tdu *TemperatureDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(temperaturedata.Table, temperaturedata.Columns, sqlgraph.NewFieldSpec(temperaturedata.FieldID, field.TypeUUID))
	if ps := tdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tdu.mutation.ValidTime(); ok {
		_spec.SetField(temperaturedata.FieldValidTime, field.TypeTime, value)
	}
	if value, ok := tdu.mutation.Temperature(); ok {
		_spec.SetField(temperaturedata.FieldTemperature, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedTemperature(); ok {
		_spec.AddField(temperaturedata.FieldTemperature, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.MinTemperature(); ok {
		_spec.SetField(temperaturedata.FieldMinTemperature, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedMinTemperature(); ok {
		_spec.AddField(temperaturedata.FieldMinTemperature, field.TypeFloat64, value)
	}
	if tdu.mutation.MinTemperatureCleared() {
		_spec.ClearField(temperaturedata.FieldMinTemperature, field.TypeFloat64)
	}
	if value, ok := tdu.mutation.MaxTemperature(); ok {
		_spec.SetField(temperaturedata.FieldMaxTemperature, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedMaxTemperature(); ok {
		_spec.AddField(temperaturedata.FieldMaxTemperature, field.TypeFloat64, value)
	}
	if tdu.mutation.MaxTemperatureCleared() {
		_spec.ClearField(temperaturedata.FieldMaxTemperature, field.TypeFloat64)
	}
	_spec.AddModifiers(tdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{temperaturedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tdu.mutation.done = true
	return n, nil
}

// TemperatureDataUpdateOne is the builder for updating a single TemperatureData entity.
type TemperatureDataUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TemperatureDataMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetValidTime sets the "valid_time" field.
func (tduo *TemperatureDataUpdateOne) SetValidTime(t time.Time) *TemperatureDataUpdateOne {
	tduo.mutation.SetValidTime(t)
	return tduo
}

// SetTemperature sets the "temperature" field.
func (tduo *TemperatureDataUpdateOne) SetTemperature(f float64) *TemperatureDataUpdateOne {
	tduo.mutation.ResetTemperature()
	tduo.mutation.SetTemperature(f)
	return tduo
}

// AddTemperature adds f to the "temperature" field.
func (tduo *TemperatureDataUpdateOne) AddTemperature(f float64) *TemperatureDataUpdateOne {
	tduo.mutation.AddTemperature(f)
	return tduo
}

// SetMinTemperature sets the "min_temperature" field.
func (tduo *TemperatureDataUpdateOne) SetMinTemperature(f float64) *TemperatureDataUpdateOne {
	tduo.mutation.ResetMinTemperature()
	tduo.mutation.SetMinTemperature(f)
	return tduo
}

// SetNillableMinTemperature sets the "min_temperature" field if the given value is not nil.
func (tduo *TemperatureDataUpdateOne) SetNillableMinTemperature(f *float64) *TemperatureDataUpdateOne {
	if f != nil {
		tduo.SetMinTemperature(*f)
	}
	return tduo
}

// AddMinTemperature adds f to the "min_temperature" field.
func (tduo *TemperatureDataUpdateOne) AddMinTemperature(f float64) *TemperatureDataUpdateOne {
	tduo.mutation.AddMinTemperature(f)
	return tduo
}

// ClearMinTemperature clears the value of the "min_temperature" field.
func (tduo *TemperatureDataUpdateOne) ClearMinTemperature() *TemperatureDataUpdateOne {
	tduo.mutation.ClearMinTemperature()
	return tduo
}

// SetMaxTemperature sets the "max_temperature" field.
func (tduo *TemperatureDataUpdateOne) SetMaxTemperature(f float64) *TemperatureDataUpdateOne {
	tduo.mutation.ResetMaxTemperature()
	tduo.mutation.SetMaxTemperature(f)
	return tduo
}

// SetNillableMaxTemperature sets the "max_temperature" field if the given value is not nil.
func (tduo *TemperatureDataUpdateOne) SetNillableMaxTemperature(f *float64) *TemperatureDataUpdateOne {
	if f != nil {
		tduo.SetMaxTemperature(*f)
	}
	return tduo
}

// AddMaxTemperature adds f to the "max_temperature" field.
func (tduo *TemperatureDataUpdateOne) AddMaxTemperature(f float64) *TemperatureDataUpdateOne {
	tduo.mutation.AddMaxTemperature(f)
	return tduo
}

// ClearMaxTemperature clears the value of the "max_temperature" field.
func (tduo *TemperatureDataUpdateOne) ClearMaxTemperature() *TemperatureDataUpdateOne {
	tduo.mutation.ClearMaxTemperature()
	return tduo
}

// Mutation returns the TemperatureDataMutation object of the builder.
func (tduo *TemperatureDataUpdateOne) Mutation() *TemperatureDataMutation {
	return tduo.mutation
}

// Where appends a list predicates to the TemperatureDataUpdate builder.
func (tduo *TemperatureDataUpdateOne) Where(ps ...predicate.TemperatureData) *TemperatureDataUpdateOne {
	tduo.mutation.Where(ps...)
	return tduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tduo *TemperatureDataUpdateOne) Select(field string, fields ...string) *TemperatureDataUpdateOne {
	tduo.fields = append([]string{field}, fields...)
	return tduo
}

// Save executes the query and returns the updated TemperatureData entity.
func (tduo *TemperatureDataUpdateOne) Save(ctx context.Context) (*TemperatureData, error) {
	return withHooks[*TemperatureData, TemperatureDataMutation](ctx, tduo.sqlSave, tduo.mutation, tduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tduo *TemperatureDataUpdateOne) SaveX(ctx context.Context) *TemperatureData {
	node, err := tduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tduo *TemperatureDataUpdateOne) Exec(ctx context.Context) error {
	_, err := tduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tduo *TemperatureDataUpdateOne) ExecX(ctx context.Context) {
	if err := tduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tduo *TemperatureDataUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemperatureDataUpdateOne {
	tduo.modifiers = append(tduo.modifiers, modifiers...)
	return tduo
}

func (tduo *TemperatureDataUpdateOne) sqlSave(ctx context.Context) (_node *TemperatureData, err error) {
	_spec := sqlgraph.NewUpdateSpec(temperaturedata.Table, temperaturedata.Columns, sqlgraph.NewFieldSpec(temperaturedata.FieldID, field.TypeUUID))
	id, ok := tduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TemperatureData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, temperaturedata.FieldID)
		for _, f := range fields {
			if !temperaturedata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != temperaturedata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tduo.mutation.ValidTime(); ok {
		_spec.SetField(temperaturedata.FieldValidTime, field.TypeTime, value)
	}
	if value, ok := tduo.mutation.Temperature(); ok {
		_spec.SetField(temperaturedata.FieldTemperature, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedTemperature(); ok {
		_spec.AddField(temperaturedata.FieldTemperature, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.MinTemperature(); ok {
		_spec.SetField(temperaturedata.FieldMinTemperature, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedMinTemperature(); ok {
		_spec.AddField(temperaturedata.FieldMinTemperature, field.TypeFloat64, value)
	}
	if tduo.mutation.MinTemperatureCleared() {
		_spec.ClearField(temperaturedata.FieldMinTemperature, field.TypeFloat64)
	}
	if value, ok := tduo.mutation.MaxTemperature(); ok {
		_spec.SetField(temperaturedata.FieldMaxTemperature, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedMaxTemperature(); ok {
		_spec.AddField(temperaturedata.FieldMaxTemperature, field.TypeFloat64, value)
	}
	if tduo.mutation.MaxTemperatureCleared() {
		_spec.ClearField(temperaturedata.FieldMaxTemperature, field.TypeFloat64)
	}
	_spec.AddModifiers(tduo.modifiers...)
	_node = &TemperatureData{config: tduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{temperaturedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tduo.mutation.done = true
	return _node, nil
}
