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
	"github.com/google/uuid"
	"metar.gg/ent/forecast"
	"metar.gg/ent/predicate"
	"metar.gg/ent/skycondition"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
)

// TafUpdate is the builder for updating Taf entities.
type TafUpdate struct {
	config
	hooks     []Hook
	mutation  *TafMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TafUpdate builder.
func (tu *TafUpdate) Where(ps ...predicate.Taf) *TafUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetRawText sets the "raw_text" field.
func (tu *TafUpdate) SetRawText(s string) *TafUpdate {
	tu.mutation.SetRawText(s)
	return tu
}

// SetIssueTime sets the "issue_time" field.
func (tu *TafUpdate) SetIssueTime(t time.Time) *TafUpdate {
	tu.mutation.SetIssueTime(t)
	return tu
}

// SetImportTime sets the "import_time" field.
func (tu *TafUpdate) SetImportTime(t time.Time) *TafUpdate {
	tu.mutation.SetImportTime(t)
	return tu
}

// SetNillableImportTime sets the "import_time" field if the given value is not nil.
func (tu *TafUpdate) SetNillableImportTime(t *time.Time) *TafUpdate {
	if t != nil {
		tu.SetImportTime(*t)
	}
	return tu
}

// SetBulletinTime sets the "bulletin_time" field.
func (tu *TafUpdate) SetBulletinTime(t time.Time) *TafUpdate {
	tu.mutation.SetBulletinTime(t)
	return tu
}

// SetValidFromTime sets the "valid_from_time" field.
func (tu *TafUpdate) SetValidFromTime(t time.Time) *TafUpdate {
	tu.mutation.SetValidFromTime(t)
	return tu
}

// SetValidToTime sets the "valid_to_time" field.
func (tu *TafUpdate) SetValidToTime(t time.Time) *TafUpdate {
	tu.mutation.SetValidToTime(t)
	return tu
}

// SetRemarks sets the "remarks" field.
func (tu *TafUpdate) SetRemarks(s string) *TafUpdate {
	tu.mutation.SetRemarks(s)
	return tu
}

// SetHash sets the "hash" field.
func (tu *TafUpdate) SetHash(s string) *TafUpdate {
	tu.mutation.SetHash(s)
	return tu
}

// SetStationID sets the "station" edge to the WeatherStation entity by ID.
func (tu *TafUpdate) SetStationID(id uuid.UUID) *TafUpdate {
	tu.mutation.SetStationID(id)
	return tu
}

// SetStation sets the "station" edge to the WeatherStation entity.
func (tu *TafUpdate) SetStation(w *WeatherStation) *TafUpdate {
	return tu.SetStationID(w.ID)
}

// AddSkyConditionIDs adds the "sky_conditions" edge to the SkyCondition entity by IDs.
func (tu *TafUpdate) AddSkyConditionIDs(ids ...uuid.UUID) *TafUpdate {
	tu.mutation.AddSkyConditionIDs(ids...)
	return tu
}

// AddSkyConditions adds the "sky_conditions" edges to the SkyCondition entity.
func (tu *TafUpdate) AddSkyConditions(s ...*SkyCondition) *TafUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.AddSkyConditionIDs(ids...)
}

// AddForecastIDs adds the "forecast" edge to the Forecast entity by IDs.
func (tu *TafUpdate) AddForecastIDs(ids ...uuid.UUID) *TafUpdate {
	tu.mutation.AddForecastIDs(ids...)
	return tu
}

// AddForecast adds the "forecast" edges to the Forecast entity.
func (tu *TafUpdate) AddForecast(f ...*Forecast) *TafUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.AddForecastIDs(ids...)
}

// Mutation returns the TafMutation object of the builder.
func (tu *TafUpdate) Mutation() *TafMutation {
	return tu.mutation
}

// ClearStation clears the "station" edge to the WeatherStation entity.
func (tu *TafUpdate) ClearStation() *TafUpdate {
	tu.mutation.ClearStation()
	return tu
}

// ClearSkyConditions clears all "sky_conditions" edges to the SkyCondition entity.
func (tu *TafUpdate) ClearSkyConditions() *TafUpdate {
	tu.mutation.ClearSkyConditions()
	return tu
}

// RemoveSkyConditionIDs removes the "sky_conditions" edge to SkyCondition entities by IDs.
func (tu *TafUpdate) RemoveSkyConditionIDs(ids ...uuid.UUID) *TafUpdate {
	tu.mutation.RemoveSkyConditionIDs(ids...)
	return tu
}

// RemoveSkyConditions removes "sky_conditions" edges to SkyCondition entities.
func (tu *TafUpdate) RemoveSkyConditions(s ...*SkyCondition) *TafUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.RemoveSkyConditionIDs(ids...)
}

// ClearForecast clears all "forecast" edges to the Forecast entity.
func (tu *TafUpdate) ClearForecast() *TafUpdate {
	tu.mutation.ClearForecast()
	return tu
}

// RemoveForecastIDs removes the "forecast" edge to Forecast entities by IDs.
func (tu *TafUpdate) RemoveForecastIDs(ids ...uuid.UUID) *TafUpdate {
	tu.mutation.RemoveForecastIDs(ids...)
	return tu
}

// RemoveForecast removes "forecast" edges to Forecast entities.
func (tu *TafUpdate) RemoveForecast(f ...*Forecast) *TafUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.RemoveForecastIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TafUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TafMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TafUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TafUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TafUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TafUpdate) check() error {
	if _, ok := tu.mutation.StationID(); tu.mutation.StationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Taf.station"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TafUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TafUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TafUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taf.Table,
			Columns: taf.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: taf.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.RawText(); ok {
		_spec.SetField(taf.FieldRawText, field.TypeString, value)
	}
	if value, ok := tu.mutation.IssueTime(); ok {
		_spec.SetField(taf.FieldIssueTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.ImportTime(); ok {
		_spec.SetField(taf.FieldImportTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.BulletinTime(); ok {
		_spec.SetField(taf.FieldBulletinTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.ValidFromTime(); ok {
		_spec.SetField(taf.FieldValidFromTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.ValidToTime(); ok {
		_spec.SetField(taf.FieldValidToTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Remarks(); ok {
		_spec.SetField(taf.FieldRemarks, field.TypeString, value)
	}
	if value, ok := tu.mutation.Hash(); ok {
		_spec.SetField(taf.FieldHash, field.TypeString, value)
	}
	if tu.mutation.StationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taf.StationTable,
			Columns: []string{taf.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: weatherstation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.StationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taf.StationTable,
			Columns: []string{taf.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: weatherstation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.SkyConditionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.SkyConditionsTable,
			Columns: []string{taf.SkyConditionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skycondition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedSkyConditionsIDs(); len(nodes) > 0 && !tu.mutation.SkyConditionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.SkyConditionsTable,
			Columns: []string{taf.SkyConditionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skycondition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SkyConditionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.SkyConditionsTable,
			Columns: []string{taf.SkyConditionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skycondition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ForecastCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.ForecastTable,
			Columns: []string{taf.ForecastColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: forecast.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedForecastIDs(); len(nodes) > 0 && !tu.mutation.ForecastCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.ForecastTable,
			Columns: []string{taf.ForecastColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: forecast.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ForecastIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.ForecastTable,
			Columns: []string{taf.ForecastColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: forecast.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TafUpdateOne is the builder for updating a single Taf entity.
type TafUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TafMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetRawText sets the "raw_text" field.
func (tuo *TafUpdateOne) SetRawText(s string) *TafUpdateOne {
	tuo.mutation.SetRawText(s)
	return tuo
}

// SetIssueTime sets the "issue_time" field.
func (tuo *TafUpdateOne) SetIssueTime(t time.Time) *TafUpdateOne {
	tuo.mutation.SetIssueTime(t)
	return tuo
}

// SetImportTime sets the "import_time" field.
func (tuo *TafUpdateOne) SetImportTime(t time.Time) *TafUpdateOne {
	tuo.mutation.SetImportTime(t)
	return tuo
}

// SetNillableImportTime sets the "import_time" field if the given value is not nil.
func (tuo *TafUpdateOne) SetNillableImportTime(t *time.Time) *TafUpdateOne {
	if t != nil {
		tuo.SetImportTime(*t)
	}
	return tuo
}

// SetBulletinTime sets the "bulletin_time" field.
func (tuo *TafUpdateOne) SetBulletinTime(t time.Time) *TafUpdateOne {
	tuo.mutation.SetBulletinTime(t)
	return tuo
}

// SetValidFromTime sets the "valid_from_time" field.
func (tuo *TafUpdateOne) SetValidFromTime(t time.Time) *TafUpdateOne {
	tuo.mutation.SetValidFromTime(t)
	return tuo
}

// SetValidToTime sets the "valid_to_time" field.
func (tuo *TafUpdateOne) SetValidToTime(t time.Time) *TafUpdateOne {
	tuo.mutation.SetValidToTime(t)
	return tuo
}

// SetRemarks sets the "remarks" field.
func (tuo *TafUpdateOne) SetRemarks(s string) *TafUpdateOne {
	tuo.mutation.SetRemarks(s)
	return tuo
}

// SetHash sets the "hash" field.
func (tuo *TafUpdateOne) SetHash(s string) *TafUpdateOne {
	tuo.mutation.SetHash(s)
	return tuo
}

// SetStationID sets the "station" edge to the WeatherStation entity by ID.
func (tuo *TafUpdateOne) SetStationID(id uuid.UUID) *TafUpdateOne {
	tuo.mutation.SetStationID(id)
	return tuo
}

// SetStation sets the "station" edge to the WeatherStation entity.
func (tuo *TafUpdateOne) SetStation(w *WeatherStation) *TafUpdateOne {
	return tuo.SetStationID(w.ID)
}

// AddSkyConditionIDs adds the "sky_conditions" edge to the SkyCondition entity by IDs.
func (tuo *TafUpdateOne) AddSkyConditionIDs(ids ...uuid.UUID) *TafUpdateOne {
	tuo.mutation.AddSkyConditionIDs(ids...)
	return tuo
}

// AddSkyConditions adds the "sky_conditions" edges to the SkyCondition entity.
func (tuo *TafUpdateOne) AddSkyConditions(s ...*SkyCondition) *TafUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.AddSkyConditionIDs(ids...)
}

// AddForecastIDs adds the "forecast" edge to the Forecast entity by IDs.
func (tuo *TafUpdateOne) AddForecastIDs(ids ...uuid.UUID) *TafUpdateOne {
	tuo.mutation.AddForecastIDs(ids...)
	return tuo
}

// AddForecast adds the "forecast" edges to the Forecast entity.
func (tuo *TafUpdateOne) AddForecast(f ...*Forecast) *TafUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.AddForecastIDs(ids...)
}

// Mutation returns the TafMutation object of the builder.
func (tuo *TafUpdateOne) Mutation() *TafMutation {
	return tuo.mutation
}

// ClearStation clears the "station" edge to the WeatherStation entity.
func (tuo *TafUpdateOne) ClearStation() *TafUpdateOne {
	tuo.mutation.ClearStation()
	return tuo
}

// ClearSkyConditions clears all "sky_conditions" edges to the SkyCondition entity.
func (tuo *TafUpdateOne) ClearSkyConditions() *TafUpdateOne {
	tuo.mutation.ClearSkyConditions()
	return tuo
}

// RemoveSkyConditionIDs removes the "sky_conditions" edge to SkyCondition entities by IDs.
func (tuo *TafUpdateOne) RemoveSkyConditionIDs(ids ...uuid.UUID) *TafUpdateOne {
	tuo.mutation.RemoveSkyConditionIDs(ids...)
	return tuo
}

// RemoveSkyConditions removes "sky_conditions" edges to SkyCondition entities.
func (tuo *TafUpdateOne) RemoveSkyConditions(s ...*SkyCondition) *TafUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.RemoveSkyConditionIDs(ids...)
}

// ClearForecast clears all "forecast" edges to the Forecast entity.
func (tuo *TafUpdateOne) ClearForecast() *TafUpdateOne {
	tuo.mutation.ClearForecast()
	return tuo
}

// RemoveForecastIDs removes the "forecast" edge to Forecast entities by IDs.
func (tuo *TafUpdateOne) RemoveForecastIDs(ids ...uuid.UUID) *TafUpdateOne {
	tuo.mutation.RemoveForecastIDs(ids...)
	return tuo
}

// RemoveForecast removes "forecast" edges to Forecast entities.
func (tuo *TafUpdateOne) RemoveForecast(f ...*Forecast) *TafUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.RemoveForecastIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TafUpdateOne) Select(field string, fields ...string) *TafUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Taf entity.
func (tuo *TafUpdateOne) Save(ctx context.Context) (*Taf, error) {
	var (
		err  error
		node *Taf
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TafMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Taf)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TafMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TafUpdateOne) SaveX(ctx context.Context) *Taf {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TafUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TafUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TafUpdateOne) check() error {
	if _, ok := tuo.mutation.StationID(); tuo.mutation.StationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Taf.station"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TafUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TafUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TafUpdateOne) sqlSave(ctx context.Context) (_node *Taf, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   taf.Table,
			Columns: taf.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: taf.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Taf.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taf.FieldID)
		for _, f := range fields {
			if !taf.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taf.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.RawText(); ok {
		_spec.SetField(taf.FieldRawText, field.TypeString, value)
	}
	if value, ok := tuo.mutation.IssueTime(); ok {
		_spec.SetField(taf.FieldIssueTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.ImportTime(); ok {
		_spec.SetField(taf.FieldImportTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.BulletinTime(); ok {
		_spec.SetField(taf.FieldBulletinTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.ValidFromTime(); ok {
		_spec.SetField(taf.FieldValidFromTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.ValidToTime(); ok {
		_spec.SetField(taf.FieldValidToTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Remarks(); ok {
		_spec.SetField(taf.FieldRemarks, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Hash(); ok {
		_spec.SetField(taf.FieldHash, field.TypeString, value)
	}
	if tuo.mutation.StationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taf.StationTable,
			Columns: []string{taf.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: weatherstation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.StationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taf.StationTable,
			Columns: []string{taf.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: weatherstation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.SkyConditionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.SkyConditionsTable,
			Columns: []string{taf.SkyConditionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skycondition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedSkyConditionsIDs(); len(nodes) > 0 && !tuo.mutation.SkyConditionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.SkyConditionsTable,
			Columns: []string{taf.SkyConditionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skycondition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SkyConditionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.SkyConditionsTable,
			Columns: []string{taf.SkyConditionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skycondition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ForecastCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.ForecastTable,
			Columns: []string{taf.ForecastColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: forecast.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedForecastIDs(); len(nodes) > 0 && !tuo.mutation.ForecastCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.ForecastTable,
			Columns: []string{taf.ForecastColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: forecast.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ForecastIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.ForecastTable,
			Columns: []string{taf.ForecastColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: forecast.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Taf{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
