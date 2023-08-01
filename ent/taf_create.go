// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"metar.gg/ent/forecast"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
)

// TafCreate is the builder for creating a Taf entity.
type TafCreate struct {
	config
	mutation *TafMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRawText sets the "raw_text" field.
func (tc *TafCreate) SetRawText(s string) *TafCreate {
	tc.mutation.SetRawText(s)
	return tc
}

// SetIssueTime sets the "issue_time" field.
func (tc *TafCreate) SetIssueTime(t time.Time) *TafCreate {
	tc.mutation.SetIssueTime(t)
	return tc
}

// SetImportTime sets the "import_time" field.
func (tc *TafCreate) SetImportTime(t time.Time) *TafCreate {
	tc.mutation.SetImportTime(t)
	return tc
}

// SetNillableImportTime sets the "import_time" field if the given value is not nil.
func (tc *TafCreate) SetNillableImportTime(t *time.Time) *TafCreate {
	if t != nil {
		tc.SetImportTime(*t)
	}
	return tc
}

// SetBulletinTime sets the "bulletin_time" field.
func (tc *TafCreate) SetBulletinTime(t time.Time) *TafCreate {
	tc.mutation.SetBulletinTime(t)
	return tc
}

// SetValidFromTime sets the "valid_from_time" field.
func (tc *TafCreate) SetValidFromTime(t time.Time) *TafCreate {
	tc.mutation.SetValidFromTime(t)
	return tc
}

// SetValidToTime sets the "valid_to_time" field.
func (tc *TafCreate) SetValidToTime(t time.Time) *TafCreate {
	tc.mutation.SetValidToTime(t)
	return tc
}

// SetRemarks sets the "remarks" field.
func (tc *TafCreate) SetRemarks(s string) *TafCreate {
	tc.mutation.SetRemarks(s)
	return tc
}

// SetHash sets the "hash" field.
func (tc *TafCreate) SetHash(s string) *TafCreate {
	tc.mutation.SetHash(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TafCreate) SetID(u uuid.UUID) *TafCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TafCreate) SetNillableID(u *uuid.UUID) *TafCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetStationID sets the "station" edge to the WeatherStation entity by ID.
func (tc *TafCreate) SetStationID(id uuid.UUID) *TafCreate {
	tc.mutation.SetStationID(id)
	return tc
}

// SetStation sets the "station" edge to the WeatherStation entity.
func (tc *TafCreate) SetStation(w *WeatherStation) *TafCreate {
	return tc.SetStationID(w.ID)
}

// AddForecastIDs adds the "forecast" edge to the Forecast entity by IDs.
func (tc *TafCreate) AddForecastIDs(ids ...uuid.UUID) *TafCreate {
	tc.mutation.AddForecastIDs(ids...)
	return tc
}

// AddForecast adds the "forecast" edges to the Forecast entity.
func (tc *TafCreate) AddForecast(f ...*Forecast) *TafCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tc.AddForecastIDs(ids...)
}

// Mutation returns the TafMutation object of the builder.
func (tc *TafCreate) Mutation() *TafMutation {
	return tc.mutation
}

// Save creates the Taf in the database.
func (tc *TafCreate) Save(ctx context.Context) (*Taf, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TafCreate) SaveX(ctx context.Context) *Taf {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TafCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TafCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TafCreate) defaults() {
	if _, ok := tc.mutation.ImportTime(); !ok {
		v := taf.DefaultImportTime()
		tc.mutation.SetImportTime(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := taf.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TafCreate) check() error {
	if _, ok := tc.mutation.RawText(); !ok {
		return &ValidationError{Name: "raw_text", err: errors.New(`ent: missing required field "Taf.raw_text"`)}
	}
	if _, ok := tc.mutation.IssueTime(); !ok {
		return &ValidationError{Name: "issue_time", err: errors.New(`ent: missing required field "Taf.issue_time"`)}
	}
	if _, ok := tc.mutation.ImportTime(); !ok {
		return &ValidationError{Name: "import_time", err: errors.New(`ent: missing required field "Taf.import_time"`)}
	}
	if _, ok := tc.mutation.BulletinTime(); !ok {
		return &ValidationError{Name: "bulletin_time", err: errors.New(`ent: missing required field "Taf.bulletin_time"`)}
	}
	if _, ok := tc.mutation.ValidFromTime(); !ok {
		return &ValidationError{Name: "valid_from_time", err: errors.New(`ent: missing required field "Taf.valid_from_time"`)}
	}
	if _, ok := tc.mutation.ValidToTime(); !ok {
		return &ValidationError{Name: "valid_to_time", err: errors.New(`ent: missing required field "Taf.valid_to_time"`)}
	}
	if _, ok := tc.mutation.Remarks(); !ok {
		return &ValidationError{Name: "remarks", err: errors.New(`ent: missing required field "Taf.remarks"`)}
	}
	if _, ok := tc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Taf.hash"`)}
	}
	if _, ok := tc.mutation.StationID(); !ok {
		return &ValidationError{Name: "station", err: errors.New(`ent: missing required edge "Taf.station"`)}
	}
	return nil
}

func (tc *TafCreate) sqlSave(ctx context.Context) (*Taf, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TafCreate) createSpec() (*Taf, *sqlgraph.CreateSpec) {
	var (
		_node = &Taf{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(taf.Table, sqlgraph.NewFieldSpec(taf.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.RawText(); ok {
		_spec.SetField(taf.FieldRawText, field.TypeString, value)
		_node.RawText = value
	}
	if value, ok := tc.mutation.IssueTime(); ok {
		_spec.SetField(taf.FieldIssueTime, field.TypeTime, value)
		_node.IssueTime = value
	}
	if value, ok := tc.mutation.ImportTime(); ok {
		_spec.SetField(taf.FieldImportTime, field.TypeTime, value)
		_node.ImportTime = value
	}
	if value, ok := tc.mutation.BulletinTime(); ok {
		_spec.SetField(taf.FieldBulletinTime, field.TypeTime, value)
		_node.BulletinTime = value
	}
	if value, ok := tc.mutation.ValidFromTime(); ok {
		_spec.SetField(taf.FieldValidFromTime, field.TypeTime, value)
		_node.ValidFromTime = value
	}
	if value, ok := tc.mutation.ValidToTime(); ok {
		_spec.SetField(taf.FieldValidToTime, field.TypeTime, value)
		_node.ValidToTime = value
	}
	if value, ok := tc.mutation.Remarks(); ok {
		_spec.SetField(taf.FieldRemarks, field.TypeString, value)
		_node.Remarks = value
	}
	if value, ok := tc.mutation.Hash(); ok {
		_spec.SetField(taf.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if nodes := tc.mutation.StationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   taf.StationTable,
			Columns: []string{taf.StationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weatherstation.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.weather_station_tafs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ForecastIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   taf.ForecastTable,
			Columns: []string{taf.ForecastColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forecast.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Taf.Create().
//		SetRawText(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TafUpsert) {
//			SetRawText(v+v).
//		}).
//		Exec(ctx)
func (tc *TafCreate) OnConflict(opts ...sql.ConflictOption) *TafUpsertOne {
	tc.conflict = opts
	return &TafUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Taf.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TafCreate) OnConflictColumns(columns ...string) *TafUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TafUpsertOne{
		create: tc,
	}
}

type (
	// TafUpsertOne is the builder for "upsert"-ing
	//  one Taf node.
	TafUpsertOne struct {
		create *TafCreate
	}

	// TafUpsert is the "OnConflict" setter.
	TafUpsert struct {
		*sql.UpdateSet
	}
)

// SetRawText sets the "raw_text" field.
func (u *TafUpsert) SetRawText(v string) *TafUpsert {
	u.Set(taf.FieldRawText, v)
	return u
}

// UpdateRawText sets the "raw_text" field to the value that was provided on create.
func (u *TafUpsert) UpdateRawText() *TafUpsert {
	u.SetExcluded(taf.FieldRawText)
	return u
}

// SetIssueTime sets the "issue_time" field.
func (u *TafUpsert) SetIssueTime(v time.Time) *TafUpsert {
	u.Set(taf.FieldIssueTime, v)
	return u
}

// UpdateIssueTime sets the "issue_time" field to the value that was provided on create.
func (u *TafUpsert) UpdateIssueTime() *TafUpsert {
	u.SetExcluded(taf.FieldIssueTime)
	return u
}

// SetImportTime sets the "import_time" field.
func (u *TafUpsert) SetImportTime(v time.Time) *TafUpsert {
	u.Set(taf.FieldImportTime, v)
	return u
}

// UpdateImportTime sets the "import_time" field to the value that was provided on create.
func (u *TafUpsert) UpdateImportTime() *TafUpsert {
	u.SetExcluded(taf.FieldImportTime)
	return u
}

// SetBulletinTime sets the "bulletin_time" field.
func (u *TafUpsert) SetBulletinTime(v time.Time) *TafUpsert {
	u.Set(taf.FieldBulletinTime, v)
	return u
}

// UpdateBulletinTime sets the "bulletin_time" field to the value that was provided on create.
func (u *TafUpsert) UpdateBulletinTime() *TafUpsert {
	u.SetExcluded(taf.FieldBulletinTime)
	return u
}

// SetValidFromTime sets the "valid_from_time" field.
func (u *TafUpsert) SetValidFromTime(v time.Time) *TafUpsert {
	u.Set(taf.FieldValidFromTime, v)
	return u
}

// UpdateValidFromTime sets the "valid_from_time" field to the value that was provided on create.
func (u *TafUpsert) UpdateValidFromTime() *TafUpsert {
	u.SetExcluded(taf.FieldValidFromTime)
	return u
}

// SetValidToTime sets the "valid_to_time" field.
func (u *TafUpsert) SetValidToTime(v time.Time) *TafUpsert {
	u.Set(taf.FieldValidToTime, v)
	return u
}

// UpdateValidToTime sets the "valid_to_time" field to the value that was provided on create.
func (u *TafUpsert) UpdateValidToTime() *TafUpsert {
	u.SetExcluded(taf.FieldValidToTime)
	return u
}

// SetRemarks sets the "remarks" field.
func (u *TafUpsert) SetRemarks(v string) *TafUpsert {
	u.Set(taf.FieldRemarks, v)
	return u
}

// UpdateRemarks sets the "remarks" field to the value that was provided on create.
func (u *TafUpsert) UpdateRemarks() *TafUpsert {
	u.SetExcluded(taf.FieldRemarks)
	return u
}

// SetHash sets the "hash" field.
func (u *TafUpsert) SetHash(v string) *TafUpsert {
	u.Set(taf.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TafUpsert) UpdateHash() *TafUpsert {
	u.SetExcluded(taf.FieldHash)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Taf.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taf.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TafUpsertOne) UpdateNewValues() *TafUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(taf.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Taf.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TafUpsertOne) Ignore() *TafUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TafUpsertOne) DoNothing() *TafUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TafCreate.OnConflict
// documentation for more info.
func (u *TafUpsertOne) Update(set func(*TafUpsert)) *TafUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TafUpsert{UpdateSet: update})
	}))
	return u
}

// SetRawText sets the "raw_text" field.
func (u *TafUpsertOne) SetRawText(v string) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetRawText(v)
	})
}

// UpdateRawText sets the "raw_text" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateRawText() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateRawText()
	})
}

// SetIssueTime sets the "issue_time" field.
func (u *TafUpsertOne) SetIssueTime(v time.Time) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetIssueTime(v)
	})
}

// UpdateIssueTime sets the "issue_time" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateIssueTime() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateIssueTime()
	})
}

// SetImportTime sets the "import_time" field.
func (u *TafUpsertOne) SetImportTime(v time.Time) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetImportTime(v)
	})
}

// UpdateImportTime sets the "import_time" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateImportTime() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateImportTime()
	})
}

// SetBulletinTime sets the "bulletin_time" field.
func (u *TafUpsertOne) SetBulletinTime(v time.Time) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetBulletinTime(v)
	})
}

// UpdateBulletinTime sets the "bulletin_time" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateBulletinTime() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateBulletinTime()
	})
}

// SetValidFromTime sets the "valid_from_time" field.
func (u *TafUpsertOne) SetValidFromTime(v time.Time) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetValidFromTime(v)
	})
}

// UpdateValidFromTime sets the "valid_from_time" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateValidFromTime() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateValidFromTime()
	})
}

// SetValidToTime sets the "valid_to_time" field.
func (u *TafUpsertOne) SetValidToTime(v time.Time) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetValidToTime(v)
	})
}

// UpdateValidToTime sets the "valid_to_time" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateValidToTime() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateValidToTime()
	})
}

// SetRemarks sets the "remarks" field.
func (u *TafUpsertOne) SetRemarks(v string) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetRemarks(v)
	})
}

// UpdateRemarks sets the "remarks" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateRemarks() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateRemarks()
	})
}

// SetHash sets the "hash" field.
func (u *TafUpsertOne) SetHash(v string) *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TafUpsertOne) UpdateHash() *TafUpsertOne {
	return u.Update(func(s *TafUpsert) {
		s.UpdateHash()
	})
}

// Exec executes the query.
func (u *TafUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TafCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TafUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TafUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TafUpsertOne.ID is not supported by MySQL driver. Use TafUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TafUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TafCreateBulk is the builder for creating many Taf entities in bulk.
type TafCreateBulk struct {
	config
	builders []*TafCreate
	conflict []sql.ConflictOption
}

// Save creates the Taf entities in the database.
func (tcb *TafCreateBulk) Save(ctx context.Context) ([]*Taf, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Taf, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TafMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TafCreateBulk) SaveX(ctx context.Context) []*Taf {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TafCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TafCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Taf.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TafUpsert) {
//			SetRawText(v+v).
//		}).
//		Exec(ctx)
func (tcb *TafCreateBulk) OnConflict(opts ...sql.ConflictOption) *TafUpsertBulk {
	tcb.conflict = opts
	return &TafUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Taf.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TafCreateBulk) OnConflictColumns(columns ...string) *TafUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TafUpsertBulk{
		create: tcb,
	}
}

// TafUpsertBulk is the builder for "upsert"-ing
// a bulk of Taf nodes.
type TafUpsertBulk struct {
	create *TafCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Taf.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(taf.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TafUpsertBulk) UpdateNewValues() *TafUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(taf.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Taf.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TafUpsertBulk) Ignore() *TafUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TafUpsertBulk) DoNothing() *TafUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TafCreateBulk.OnConflict
// documentation for more info.
func (u *TafUpsertBulk) Update(set func(*TafUpsert)) *TafUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TafUpsert{UpdateSet: update})
	}))
	return u
}

// SetRawText sets the "raw_text" field.
func (u *TafUpsertBulk) SetRawText(v string) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetRawText(v)
	})
}

// UpdateRawText sets the "raw_text" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateRawText() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateRawText()
	})
}

// SetIssueTime sets the "issue_time" field.
func (u *TafUpsertBulk) SetIssueTime(v time.Time) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetIssueTime(v)
	})
}

// UpdateIssueTime sets the "issue_time" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateIssueTime() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateIssueTime()
	})
}

// SetImportTime sets the "import_time" field.
func (u *TafUpsertBulk) SetImportTime(v time.Time) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetImportTime(v)
	})
}

// UpdateImportTime sets the "import_time" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateImportTime() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateImportTime()
	})
}

// SetBulletinTime sets the "bulletin_time" field.
func (u *TafUpsertBulk) SetBulletinTime(v time.Time) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetBulletinTime(v)
	})
}

// UpdateBulletinTime sets the "bulletin_time" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateBulletinTime() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateBulletinTime()
	})
}

// SetValidFromTime sets the "valid_from_time" field.
func (u *TafUpsertBulk) SetValidFromTime(v time.Time) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetValidFromTime(v)
	})
}

// UpdateValidFromTime sets the "valid_from_time" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateValidFromTime() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateValidFromTime()
	})
}

// SetValidToTime sets the "valid_to_time" field.
func (u *TafUpsertBulk) SetValidToTime(v time.Time) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetValidToTime(v)
	})
}

// UpdateValidToTime sets the "valid_to_time" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateValidToTime() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateValidToTime()
	})
}

// SetRemarks sets the "remarks" field.
func (u *TafUpsertBulk) SetRemarks(v string) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetRemarks(v)
	})
}

// UpdateRemarks sets the "remarks" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateRemarks() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateRemarks()
	})
}

// SetHash sets the "hash" field.
func (u *TafUpsertBulk) SetHash(v string) *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TafUpsertBulk) UpdateHash() *TafUpsertBulk {
	return u.Update(func(s *TafUpsert) {
		s.UpdateHash()
	})
}

// Exec executes the query.
func (u *TafUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TafCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TafCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TafUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
