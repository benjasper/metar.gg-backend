// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
	"metar.gg/ent/skycondition"
)

// SkyConditionUpdate is the builder for updating SkyCondition entities.
type SkyConditionUpdate struct {
	config
	hooks     []Hook
	mutation  *SkyConditionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SkyConditionUpdate builder.
func (scu *SkyConditionUpdate) Where(ps ...predicate.SkyCondition) *SkyConditionUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetSkyCover sets the "sky_cover" field.
func (scu *SkyConditionUpdate) SetSkyCover(sc skycondition.SkyCover) *SkyConditionUpdate {
	scu.mutation.SetSkyCover(sc)
	return scu
}

// SetCloudBase sets the "cloud_base" field.
func (scu *SkyConditionUpdate) SetCloudBase(i int) *SkyConditionUpdate {
	scu.mutation.ResetCloudBase()
	scu.mutation.SetCloudBase(i)
	return scu
}

// SetNillableCloudBase sets the "cloud_base" field if the given value is not nil.
func (scu *SkyConditionUpdate) SetNillableCloudBase(i *int) *SkyConditionUpdate {
	if i != nil {
		scu.SetCloudBase(*i)
	}
	return scu
}

// AddCloudBase adds i to the "cloud_base" field.
func (scu *SkyConditionUpdate) AddCloudBase(i int) *SkyConditionUpdate {
	scu.mutation.AddCloudBase(i)
	return scu
}

// ClearCloudBase clears the value of the "cloud_base" field.
func (scu *SkyConditionUpdate) ClearCloudBase() *SkyConditionUpdate {
	scu.mutation.ClearCloudBase()
	return scu
}

// SetMetarID sets the "metar" edge to the Metar entity by ID.
func (scu *SkyConditionUpdate) SetMetarID(id int) *SkyConditionUpdate {
	scu.mutation.SetMetarID(id)
	return scu
}

// SetMetar sets the "metar" edge to the Metar entity.
func (scu *SkyConditionUpdate) SetMetar(m *Metar) *SkyConditionUpdate {
	return scu.SetMetarID(m.ID)
}

// Mutation returns the SkyConditionMutation object of the builder.
func (scu *SkyConditionUpdate) Mutation() *SkyConditionMutation {
	return scu.mutation
}

// ClearMetar clears the "metar" edge to the Metar entity.
func (scu *SkyConditionUpdate) ClearMetar() *SkyConditionUpdate {
	scu.mutation.ClearMetar()
	return scu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *SkyConditionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(scu.hooks) == 0 {
		if err = scu.check(); err != nil {
			return 0, err
		}
		affected, err = scu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkyConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scu.check(); err != nil {
				return 0, err
			}
			scu.mutation = mutation
			affected, err = scu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scu.hooks) - 1; i >= 0; i-- {
			if scu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (scu *SkyConditionUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *SkyConditionUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *SkyConditionUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *SkyConditionUpdate) check() error {
	if v, ok := scu.mutation.SkyCover(); ok {
		if err := skycondition.SkyCoverValidator(v); err != nil {
			return &ValidationError{Name: "sky_cover", err: fmt.Errorf(`ent: validator failed for field "SkyCondition.sky_cover": %w`, err)}
		}
	}
	if _, ok := scu.mutation.MetarID(); scu.mutation.MetarCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SkyCondition.metar"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scu *SkyConditionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SkyConditionUpdate {
	scu.modifiers = append(scu.modifiers, modifiers...)
	return scu
}

func (scu *SkyConditionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skycondition.Table,
			Columns: skycondition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skycondition.FieldID,
			},
		},
	}
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.SkyCover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: skycondition.FieldSkyCover,
		})
	}
	if value, ok := scu.mutation.CloudBase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: skycondition.FieldCloudBase,
		})
	}
	if value, ok := scu.mutation.AddedCloudBase(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: skycondition.FieldCloudBase,
		})
	}
	if scu.mutation.CloudBaseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: skycondition.FieldCloudBase,
		})
	}
	if scu.mutation.MetarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skycondition.MetarTable,
			Columns: []string{skycondition.MetarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metar.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.MetarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skycondition.MetarTable,
			Columns: []string{skycondition.MetarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metar.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = scu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skycondition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SkyConditionUpdateOne is the builder for updating a single SkyCondition entity.
type SkyConditionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SkyConditionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetSkyCover sets the "sky_cover" field.
func (scuo *SkyConditionUpdateOne) SetSkyCover(sc skycondition.SkyCover) *SkyConditionUpdateOne {
	scuo.mutation.SetSkyCover(sc)
	return scuo
}

// SetCloudBase sets the "cloud_base" field.
func (scuo *SkyConditionUpdateOne) SetCloudBase(i int) *SkyConditionUpdateOne {
	scuo.mutation.ResetCloudBase()
	scuo.mutation.SetCloudBase(i)
	return scuo
}

// SetNillableCloudBase sets the "cloud_base" field if the given value is not nil.
func (scuo *SkyConditionUpdateOne) SetNillableCloudBase(i *int) *SkyConditionUpdateOne {
	if i != nil {
		scuo.SetCloudBase(*i)
	}
	return scuo
}

// AddCloudBase adds i to the "cloud_base" field.
func (scuo *SkyConditionUpdateOne) AddCloudBase(i int) *SkyConditionUpdateOne {
	scuo.mutation.AddCloudBase(i)
	return scuo
}

// ClearCloudBase clears the value of the "cloud_base" field.
func (scuo *SkyConditionUpdateOne) ClearCloudBase() *SkyConditionUpdateOne {
	scuo.mutation.ClearCloudBase()
	return scuo
}

// SetMetarID sets the "metar" edge to the Metar entity by ID.
func (scuo *SkyConditionUpdateOne) SetMetarID(id int) *SkyConditionUpdateOne {
	scuo.mutation.SetMetarID(id)
	return scuo
}

// SetMetar sets the "metar" edge to the Metar entity.
func (scuo *SkyConditionUpdateOne) SetMetar(m *Metar) *SkyConditionUpdateOne {
	return scuo.SetMetarID(m.ID)
}

// Mutation returns the SkyConditionMutation object of the builder.
func (scuo *SkyConditionUpdateOne) Mutation() *SkyConditionMutation {
	return scuo.mutation
}

// ClearMetar clears the "metar" edge to the Metar entity.
func (scuo *SkyConditionUpdateOne) ClearMetar() *SkyConditionUpdateOne {
	scuo.mutation.ClearMetar()
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *SkyConditionUpdateOne) Select(field string, fields ...string) *SkyConditionUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated SkyCondition entity.
func (scuo *SkyConditionUpdateOne) Save(ctx context.Context) (*SkyCondition, error) {
	var (
		err  error
		node *SkyCondition
	)
	if len(scuo.hooks) == 0 {
		if err = scuo.check(); err != nil {
			return nil, err
		}
		node, err = scuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkyConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scuo.check(); err != nil {
				return nil, err
			}
			scuo.mutation = mutation
			node, err = scuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(scuo.hooks) - 1; i >= 0; i-- {
			if scuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, scuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SkyCondition)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SkyConditionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *SkyConditionUpdateOne) SaveX(ctx context.Context) *SkyCondition {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *SkyConditionUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *SkyConditionUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *SkyConditionUpdateOne) check() error {
	if v, ok := scuo.mutation.SkyCover(); ok {
		if err := skycondition.SkyCoverValidator(v); err != nil {
			return &ValidationError{Name: "sky_cover", err: fmt.Errorf(`ent: validator failed for field "SkyCondition.sky_cover": %w`, err)}
		}
	}
	if _, ok := scuo.mutation.MetarID(); scuo.mutation.MetarCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SkyCondition.metar"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scuo *SkyConditionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SkyConditionUpdateOne {
	scuo.modifiers = append(scuo.modifiers, modifiers...)
	return scuo
}

func (scuo *SkyConditionUpdateOne) sqlSave(ctx context.Context) (_node *SkyCondition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skycondition.Table,
			Columns: skycondition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skycondition.FieldID,
			},
		},
	}
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SkyCondition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skycondition.FieldID)
		for _, f := range fields {
			if !skycondition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skycondition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.SkyCover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: skycondition.FieldSkyCover,
		})
	}
	if value, ok := scuo.mutation.CloudBase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: skycondition.FieldCloudBase,
		})
	}
	if value, ok := scuo.mutation.AddedCloudBase(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: skycondition.FieldCloudBase,
		})
	}
	if scuo.mutation.CloudBaseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: skycondition.FieldCloudBase,
		})
	}
	if scuo.mutation.MetarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skycondition.MetarTable,
			Columns: []string{skycondition.MetarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metar.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.MetarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skycondition.MetarTable,
			Columns: []string{skycondition.MetarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metar.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = scuo.modifiers
	_node = &SkyCondition{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skycondition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}