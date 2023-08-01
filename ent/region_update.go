// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"metar.gg/ent/airport"
	"metar.gg/ent/predicate"
	"metar.gg/ent/region"
)

// RegionUpdate is the builder for updating Region entities.
type RegionUpdate struct {
	config
	hooks     []Hook
	mutation  *RegionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RegionUpdate builder.
func (ru *RegionUpdate) Where(ps ...predicate.Region) *RegionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetImportID sets the "import_id" field.
func (ru *RegionUpdate) SetImportID(i int) *RegionUpdate {
	ru.mutation.ResetImportID()
	ru.mutation.SetImportID(i)
	return ru
}

// AddImportID adds i to the "import_id" field.
func (ru *RegionUpdate) AddImportID(i int) *RegionUpdate {
	ru.mutation.AddImportID(i)
	return ru
}

// SetHash sets the "hash" field.
func (ru *RegionUpdate) SetHash(s string) *RegionUpdate {
	ru.mutation.SetHash(s)
	return ru
}

// SetImportFlag sets the "import_flag" field.
func (ru *RegionUpdate) SetImportFlag(b bool) *RegionUpdate {
	ru.mutation.SetImportFlag(b)
	return ru
}

// SetNillableImportFlag sets the "import_flag" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableImportFlag(b *bool) *RegionUpdate {
	if b != nil {
		ru.SetImportFlag(*b)
	}
	return ru
}

// SetLastUpdated sets the "last_updated" field.
func (ru *RegionUpdate) SetLastUpdated(t time.Time) *RegionUpdate {
	ru.mutation.SetLastUpdated(t)
	return ru
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (ru *RegionUpdate) SetNillableLastUpdated(t *time.Time) *RegionUpdate {
	if t != nil {
		ru.SetLastUpdated(*t)
	}
	return ru
}

// SetCode sets the "code" field.
func (ru *RegionUpdate) SetCode(s string) *RegionUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetLocalCode sets the "local_code" field.
func (ru *RegionUpdate) SetLocalCode(s string) *RegionUpdate {
	ru.mutation.SetLocalCode(s)
	return ru
}

// SetName sets the "name" field.
func (ru *RegionUpdate) SetName(s string) *RegionUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetWikipediaLink sets the "wikipedia_link" field.
func (ru *RegionUpdate) SetWikipediaLink(s string) *RegionUpdate {
	ru.mutation.SetWikipediaLink(s)
	return ru
}

// SetKeywords sets the "keywords" field.
func (ru *RegionUpdate) SetKeywords(s []string) *RegionUpdate {
	ru.mutation.SetKeywords(s)
	return ru
}

// AppendKeywords appends s to the "keywords" field.
func (ru *RegionUpdate) AppendKeywords(s []string) *RegionUpdate {
	ru.mutation.AppendKeywords(s)
	return ru
}

// AddAirportIDs adds the "airports" edge to the Airport entity by IDs.
func (ru *RegionUpdate) AddAirportIDs(ids ...uuid.UUID) *RegionUpdate {
	ru.mutation.AddAirportIDs(ids...)
	return ru
}

// AddAirports adds the "airports" edges to the Airport entity.
func (ru *RegionUpdate) AddAirports(a ...*Airport) *RegionUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAirportIDs(ids...)
}

// Mutation returns the RegionMutation object of the builder.
func (ru *RegionUpdate) Mutation() *RegionMutation {
	return ru.mutation
}

// ClearAirports clears all "airports" edges to the Airport entity.
func (ru *RegionUpdate) ClearAirports() *RegionUpdate {
	ru.mutation.ClearAirports()
	return ru
}

// RemoveAirportIDs removes the "airports" edge to Airport entities by IDs.
func (ru *RegionUpdate) RemoveAirportIDs(ids ...uuid.UUID) *RegionUpdate {
	ru.mutation.RemoveAirportIDs(ids...)
	return ru
}

// RemoveAirports removes "airports" edges to Airport entities.
func (ru *RegionUpdate) RemoveAirports(a ...*Airport) *RegionUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAirportIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RegionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RegionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RegionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RegionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RegionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RegionUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RegionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(region.Table, region.Columns, sqlgraph.NewFieldSpec(region.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.ImportID(); ok {
		_spec.SetField(region.FieldImportID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedImportID(); ok {
		_spec.AddField(region.FieldImportID, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Hash(); ok {
		_spec.SetField(region.FieldHash, field.TypeString, value)
	}
	if value, ok := ru.mutation.ImportFlag(); ok {
		_spec.SetField(region.FieldImportFlag, field.TypeBool, value)
	}
	if value, ok := ru.mutation.LastUpdated(); ok {
		_spec.SetField(region.FieldLastUpdated, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(region.FieldCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.LocalCode(); ok {
		_spec.SetField(region.FieldLocalCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(region.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.WikipediaLink(); ok {
		_spec.SetField(region.FieldWikipediaLink, field.TypeString, value)
	}
	if value, ok := ru.mutation.Keywords(); ok {
		_spec.SetField(region.FieldKeywords, field.TypeJSON, value)
	}
	if value, ok := ru.mutation.AppendedKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, region.FieldKeywords, value)
		})
	}
	if ru.mutation.AirportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.AirportsTable,
			Columns: []string{region.AirportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAirportsIDs(); len(nodes) > 0 && !ru.mutation.AirportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.AirportsTable,
			Columns: []string{region.AirportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AirportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.AirportsTable,
			Columns: []string{region.AirportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{region.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RegionUpdateOne is the builder for updating a single Region entity.
type RegionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RegionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetImportID sets the "import_id" field.
func (ruo *RegionUpdateOne) SetImportID(i int) *RegionUpdateOne {
	ruo.mutation.ResetImportID()
	ruo.mutation.SetImportID(i)
	return ruo
}

// AddImportID adds i to the "import_id" field.
func (ruo *RegionUpdateOne) AddImportID(i int) *RegionUpdateOne {
	ruo.mutation.AddImportID(i)
	return ruo
}

// SetHash sets the "hash" field.
func (ruo *RegionUpdateOne) SetHash(s string) *RegionUpdateOne {
	ruo.mutation.SetHash(s)
	return ruo
}

// SetImportFlag sets the "import_flag" field.
func (ruo *RegionUpdateOne) SetImportFlag(b bool) *RegionUpdateOne {
	ruo.mutation.SetImportFlag(b)
	return ruo
}

// SetNillableImportFlag sets the "import_flag" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableImportFlag(b *bool) *RegionUpdateOne {
	if b != nil {
		ruo.SetImportFlag(*b)
	}
	return ruo
}

// SetLastUpdated sets the "last_updated" field.
func (ruo *RegionUpdateOne) SetLastUpdated(t time.Time) *RegionUpdateOne {
	ruo.mutation.SetLastUpdated(t)
	return ruo
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (ruo *RegionUpdateOne) SetNillableLastUpdated(t *time.Time) *RegionUpdateOne {
	if t != nil {
		ruo.SetLastUpdated(*t)
	}
	return ruo
}

// SetCode sets the "code" field.
func (ruo *RegionUpdateOne) SetCode(s string) *RegionUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetLocalCode sets the "local_code" field.
func (ruo *RegionUpdateOne) SetLocalCode(s string) *RegionUpdateOne {
	ruo.mutation.SetLocalCode(s)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RegionUpdateOne) SetName(s string) *RegionUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetWikipediaLink sets the "wikipedia_link" field.
func (ruo *RegionUpdateOne) SetWikipediaLink(s string) *RegionUpdateOne {
	ruo.mutation.SetWikipediaLink(s)
	return ruo
}

// SetKeywords sets the "keywords" field.
func (ruo *RegionUpdateOne) SetKeywords(s []string) *RegionUpdateOne {
	ruo.mutation.SetKeywords(s)
	return ruo
}

// AppendKeywords appends s to the "keywords" field.
func (ruo *RegionUpdateOne) AppendKeywords(s []string) *RegionUpdateOne {
	ruo.mutation.AppendKeywords(s)
	return ruo
}

// AddAirportIDs adds the "airports" edge to the Airport entity by IDs.
func (ruo *RegionUpdateOne) AddAirportIDs(ids ...uuid.UUID) *RegionUpdateOne {
	ruo.mutation.AddAirportIDs(ids...)
	return ruo
}

// AddAirports adds the "airports" edges to the Airport entity.
func (ruo *RegionUpdateOne) AddAirports(a ...*Airport) *RegionUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAirportIDs(ids...)
}

// Mutation returns the RegionMutation object of the builder.
func (ruo *RegionUpdateOne) Mutation() *RegionMutation {
	return ruo.mutation
}

// ClearAirports clears all "airports" edges to the Airport entity.
func (ruo *RegionUpdateOne) ClearAirports() *RegionUpdateOne {
	ruo.mutation.ClearAirports()
	return ruo
}

// RemoveAirportIDs removes the "airports" edge to Airport entities by IDs.
func (ruo *RegionUpdateOne) RemoveAirportIDs(ids ...uuid.UUID) *RegionUpdateOne {
	ruo.mutation.RemoveAirportIDs(ids...)
	return ruo
}

// RemoveAirports removes "airports" edges to Airport entities.
func (ruo *RegionUpdateOne) RemoveAirports(a ...*Airport) *RegionUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAirportIDs(ids...)
}

// Where appends a list predicates to the RegionUpdate builder.
func (ruo *RegionUpdateOne) Where(ps ...predicate.Region) *RegionUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RegionUpdateOne) Select(field string, fields ...string) *RegionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Region entity.
func (ruo *RegionUpdateOne) Save(ctx context.Context) (*Region, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RegionUpdateOne) SaveX(ctx context.Context) *Region {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RegionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RegionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RegionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RegionUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RegionUpdateOne) sqlSave(ctx context.Context) (_node *Region, err error) {
	_spec := sqlgraph.NewUpdateSpec(region.Table, region.Columns, sqlgraph.NewFieldSpec(region.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Region.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, region.FieldID)
		for _, f := range fields {
			if !region.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != region.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.ImportID(); ok {
		_spec.SetField(region.FieldImportID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedImportID(); ok {
		_spec.AddField(region.FieldImportID, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Hash(); ok {
		_spec.SetField(region.FieldHash, field.TypeString, value)
	}
	if value, ok := ruo.mutation.ImportFlag(); ok {
		_spec.SetField(region.FieldImportFlag, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.LastUpdated(); ok {
		_spec.SetField(region.FieldLastUpdated, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(region.FieldCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.LocalCode(); ok {
		_spec.SetField(region.FieldLocalCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(region.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.WikipediaLink(); ok {
		_spec.SetField(region.FieldWikipediaLink, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Keywords(); ok {
		_spec.SetField(region.FieldKeywords, field.TypeJSON, value)
	}
	if value, ok := ruo.mutation.AppendedKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, region.FieldKeywords, value)
		})
	}
	if ruo.mutation.AirportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.AirportsTable,
			Columns: []string{region.AirportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAirportsIDs(); len(nodes) > 0 && !ruo.mutation.AirportsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.AirportsTable,
			Columns: []string{region.AirportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AirportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.AirportsTable,
			Columns: []string{region.AirportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Region{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{region.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
