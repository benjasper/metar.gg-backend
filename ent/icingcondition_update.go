// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/icingcondition"
	"metar.gg/ent/predicate"
)

// IcingConditionUpdate is the builder for updating IcingCondition entities.
type IcingConditionUpdate struct {
	config
	hooks     []Hook
	mutation  *IcingConditionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the IcingConditionUpdate builder.
func (icu *IcingConditionUpdate) Where(ps ...predicate.IcingCondition) *IcingConditionUpdate {
	icu.mutation.Where(ps...)
	return icu
}

// SetIntensity sets the "intensity" field.
func (icu *IcingConditionUpdate) SetIntensity(s string) *IcingConditionUpdate {
	icu.mutation.SetIntensity(s)
	return icu
}

// SetMinAltitude sets the "min_altitude" field.
func (icu *IcingConditionUpdate) SetMinAltitude(i int) *IcingConditionUpdate {
	icu.mutation.ResetMinAltitude()
	icu.mutation.SetMinAltitude(i)
	return icu
}

// SetNillableMinAltitude sets the "min_altitude" field if the given value is not nil.
func (icu *IcingConditionUpdate) SetNillableMinAltitude(i *int) *IcingConditionUpdate {
	if i != nil {
		icu.SetMinAltitude(*i)
	}
	return icu
}

// AddMinAltitude adds i to the "min_altitude" field.
func (icu *IcingConditionUpdate) AddMinAltitude(i int) *IcingConditionUpdate {
	icu.mutation.AddMinAltitude(i)
	return icu
}

// ClearMinAltitude clears the value of the "min_altitude" field.
func (icu *IcingConditionUpdate) ClearMinAltitude() *IcingConditionUpdate {
	icu.mutation.ClearMinAltitude()
	return icu
}

// SetMaxAltitude sets the "max_altitude" field.
func (icu *IcingConditionUpdate) SetMaxAltitude(i int) *IcingConditionUpdate {
	icu.mutation.ResetMaxAltitude()
	icu.mutation.SetMaxAltitude(i)
	return icu
}

// SetNillableMaxAltitude sets the "max_altitude" field if the given value is not nil.
func (icu *IcingConditionUpdate) SetNillableMaxAltitude(i *int) *IcingConditionUpdate {
	if i != nil {
		icu.SetMaxAltitude(*i)
	}
	return icu
}

// AddMaxAltitude adds i to the "max_altitude" field.
func (icu *IcingConditionUpdate) AddMaxAltitude(i int) *IcingConditionUpdate {
	icu.mutation.AddMaxAltitude(i)
	return icu
}

// ClearMaxAltitude clears the value of the "max_altitude" field.
func (icu *IcingConditionUpdate) ClearMaxAltitude() *IcingConditionUpdate {
	icu.mutation.ClearMaxAltitude()
	return icu
}

// Mutation returns the IcingConditionMutation object of the builder.
func (icu *IcingConditionUpdate) Mutation() *IcingConditionMutation {
	return icu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (icu *IcingConditionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(icu.hooks) == 0 {
		affected, err = icu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IcingConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			icu.mutation = mutation
			affected, err = icu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(icu.hooks) - 1; i >= 0; i-- {
			if icu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, icu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (icu *IcingConditionUpdate) SaveX(ctx context.Context) int {
	affected, err := icu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (icu *IcingConditionUpdate) Exec(ctx context.Context) error {
	_, err := icu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icu *IcingConditionUpdate) ExecX(ctx context.Context) {
	if err := icu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (icu *IcingConditionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IcingConditionUpdate {
	icu.modifiers = append(icu.modifiers, modifiers...)
	return icu
}

func (icu *IcingConditionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   icingcondition.Table,
			Columns: icingcondition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: icingcondition.FieldID,
			},
		},
	}
	if ps := icu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icu.mutation.Intensity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: icingcondition.FieldIntensity,
		})
	}
	if value, ok := icu.mutation.MinAltitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMinAltitude,
		})
	}
	if value, ok := icu.mutation.AddedMinAltitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMinAltitude,
		})
	}
	if icu.mutation.MinAltitudeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: icingcondition.FieldMinAltitude,
		})
	}
	if value, ok := icu.mutation.MaxAltitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMaxAltitude,
		})
	}
	if value, ok := icu.mutation.AddedMaxAltitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMaxAltitude,
		})
	}
	if icu.mutation.MaxAltitudeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: icingcondition.FieldMaxAltitude,
		})
	}
	_spec.Modifiers = icu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, icu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{icingcondition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// IcingConditionUpdateOne is the builder for updating a single IcingCondition entity.
type IcingConditionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *IcingConditionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIntensity sets the "intensity" field.
func (icuo *IcingConditionUpdateOne) SetIntensity(s string) *IcingConditionUpdateOne {
	icuo.mutation.SetIntensity(s)
	return icuo
}

// SetMinAltitude sets the "min_altitude" field.
func (icuo *IcingConditionUpdateOne) SetMinAltitude(i int) *IcingConditionUpdateOne {
	icuo.mutation.ResetMinAltitude()
	icuo.mutation.SetMinAltitude(i)
	return icuo
}

// SetNillableMinAltitude sets the "min_altitude" field if the given value is not nil.
func (icuo *IcingConditionUpdateOne) SetNillableMinAltitude(i *int) *IcingConditionUpdateOne {
	if i != nil {
		icuo.SetMinAltitude(*i)
	}
	return icuo
}

// AddMinAltitude adds i to the "min_altitude" field.
func (icuo *IcingConditionUpdateOne) AddMinAltitude(i int) *IcingConditionUpdateOne {
	icuo.mutation.AddMinAltitude(i)
	return icuo
}

// ClearMinAltitude clears the value of the "min_altitude" field.
func (icuo *IcingConditionUpdateOne) ClearMinAltitude() *IcingConditionUpdateOne {
	icuo.mutation.ClearMinAltitude()
	return icuo
}

// SetMaxAltitude sets the "max_altitude" field.
func (icuo *IcingConditionUpdateOne) SetMaxAltitude(i int) *IcingConditionUpdateOne {
	icuo.mutation.ResetMaxAltitude()
	icuo.mutation.SetMaxAltitude(i)
	return icuo
}

// SetNillableMaxAltitude sets the "max_altitude" field if the given value is not nil.
func (icuo *IcingConditionUpdateOne) SetNillableMaxAltitude(i *int) *IcingConditionUpdateOne {
	if i != nil {
		icuo.SetMaxAltitude(*i)
	}
	return icuo
}

// AddMaxAltitude adds i to the "max_altitude" field.
func (icuo *IcingConditionUpdateOne) AddMaxAltitude(i int) *IcingConditionUpdateOne {
	icuo.mutation.AddMaxAltitude(i)
	return icuo
}

// ClearMaxAltitude clears the value of the "max_altitude" field.
func (icuo *IcingConditionUpdateOne) ClearMaxAltitude() *IcingConditionUpdateOne {
	icuo.mutation.ClearMaxAltitude()
	return icuo
}

// Mutation returns the IcingConditionMutation object of the builder.
func (icuo *IcingConditionUpdateOne) Mutation() *IcingConditionMutation {
	return icuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (icuo *IcingConditionUpdateOne) Select(field string, fields ...string) *IcingConditionUpdateOne {
	icuo.fields = append([]string{field}, fields...)
	return icuo
}

// Save executes the query and returns the updated IcingCondition entity.
func (icuo *IcingConditionUpdateOne) Save(ctx context.Context) (*IcingCondition, error) {
	var (
		err  error
		node *IcingCondition
	)
	if len(icuo.hooks) == 0 {
		node, err = icuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IcingConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			icuo.mutation = mutation
			node, err = icuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(icuo.hooks) - 1; i >= 0; i-- {
			if icuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, icuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*IcingCondition)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from IcingConditionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (icuo *IcingConditionUpdateOne) SaveX(ctx context.Context) *IcingCondition {
	node, err := icuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (icuo *IcingConditionUpdateOne) Exec(ctx context.Context) error {
	_, err := icuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icuo *IcingConditionUpdateOne) ExecX(ctx context.Context) {
	if err := icuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (icuo *IcingConditionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *IcingConditionUpdateOne {
	icuo.modifiers = append(icuo.modifiers, modifiers...)
	return icuo
}

func (icuo *IcingConditionUpdateOne) sqlSave(ctx context.Context) (_node *IcingCondition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   icingcondition.Table,
			Columns: icingcondition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: icingcondition.FieldID,
			},
		},
	}
	id, ok := icuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IcingCondition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := icuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, icingcondition.FieldID)
		for _, f := range fields {
			if !icingcondition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != icingcondition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := icuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icuo.mutation.Intensity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: icingcondition.FieldIntensity,
		})
	}
	if value, ok := icuo.mutation.MinAltitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMinAltitude,
		})
	}
	if value, ok := icuo.mutation.AddedMinAltitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMinAltitude,
		})
	}
	if icuo.mutation.MinAltitudeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: icingcondition.FieldMinAltitude,
		})
	}
	if value, ok := icuo.mutation.MaxAltitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMaxAltitude,
		})
	}
	if value, ok := icuo.mutation.AddedMaxAltitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: icingcondition.FieldMaxAltitude,
		})
	}
	if icuo.mutation.MaxAltitudeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: icingcondition.FieldMaxAltitude,
		})
	}
	_spec.Modifiers = icuo.modifiers
	_node = &IcingCondition{config: icuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, icuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{icingcondition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
