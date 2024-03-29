// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
)

// WeatherStationCreate is the builder for creating a WeatherStation entity.
type WeatherStationCreate struct {
	config
	mutation *WeatherStationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStationID sets the "station_id" field.
func (wsc *WeatherStationCreate) SetStationID(s string) *WeatherStationCreate {
	wsc.mutation.SetStationID(s)
	return wsc
}

// SetLatitude sets the "latitude" field.
func (wsc *WeatherStationCreate) SetLatitude(f float64) *WeatherStationCreate {
	wsc.mutation.SetLatitude(f)
	return wsc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (wsc *WeatherStationCreate) SetNillableLatitude(f *float64) *WeatherStationCreate {
	if f != nil {
		wsc.SetLatitude(*f)
	}
	return wsc
}

// SetLongitude sets the "longitude" field.
func (wsc *WeatherStationCreate) SetLongitude(f float64) *WeatherStationCreate {
	wsc.mutation.SetLongitude(f)
	return wsc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (wsc *WeatherStationCreate) SetNillableLongitude(f *float64) *WeatherStationCreate {
	if f != nil {
		wsc.SetLongitude(*f)
	}
	return wsc
}

// SetElevation sets the "elevation" field.
func (wsc *WeatherStationCreate) SetElevation(f float64) *WeatherStationCreate {
	wsc.mutation.SetElevation(f)
	return wsc
}

// SetNillableElevation sets the "elevation" field if the given value is not nil.
func (wsc *WeatherStationCreate) SetNillableElevation(f *float64) *WeatherStationCreate {
	if f != nil {
		wsc.SetElevation(*f)
	}
	return wsc
}

// SetHash sets the "hash" field.
func (wsc *WeatherStationCreate) SetHash(s string) *WeatherStationCreate {
	wsc.mutation.SetHash(s)
	return wsc
}

// SetID sets the "id" field.
func (wsc *WeatherStationCreate) SetID(u uuid.UUID) *WeatherStationCreate {
	wsc.mutation.SetID(u)
	return wsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wsc *WeatherStationCreate) SetNillableID(u *uuid.UUID) *WeatherStationCreate {
	if u != nil {
		wsc.SetID(*u)
	}
	return wsc
}

// SetAirportID sets the "airport" edge to the Airport entity by ID.
func (wsc *WeatherStationCreate) SetAirportID(id uuid.UUID) *WeatherStationCreate {
	wsc.mutation.SetAirportID(id)
	return wsc
}

// SetNillableAirportID sets the "airport" edge to the Airport entity by ID if the given value is not nil.
func (wsc *WeatherStationCreate) SetNillableAirportID(id *uuid.UUID) *WeatherStationCreate {
	if id != nil {
		wsc = wsc.SetAirportID(*id)
	}
	return wsc
}

// SetAirport sets the "airport" edge to the Airport entity.
func (wsc *WeatherStationCreate) SetAirport(a *Airport) *WeatherStationCreate {
	return wsc.SetAirportID(a.ID)
}

// AddMetarIDs adds the "metars" edge to the Metar entity by IDs.
func (wsc *WeatherStationCreate) AddMetarIDs(ids ...uuid.UUID) *WeatherStationCreate {
	wsc.mutation.AddMetarIDs(ids...)
	return wsc
}

// AddMetars adds the "metars" edges to the Metar entity.
func (wsc *WeatherStationCreate) AddMetars(m ...*Metar) *WeatherStationCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return wsc.AddMetarIDs(ids...)
}

// AddTafIDs adds the "tafs" edge to the Taf entity by IDs.
func (wsc *WeatherStationCreate) AddTafIDs(ids ...uuid.UUID) *WeatherStationCreate {
	wsc.mutation.AddTafIDs(ids...)
	return wsc
}

// AddTafs adds the "tafs" edges to the Taf entity.
func (wsc *WeatherStationCreate) AddTafs(t ...*Taf) *WeatherStationCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return wsc.AddTafIDs(ids...)
}

// Mutation returns the WeatherStationMutation object of the builder.
func (wsc *WeatherStationCreate) Mutation() *WeatherStationMutation {
	return wsc.mutation
}

// Save creates the WeatherStation in the database.
func (wsc *WeatherStationCreate) Save(ctx context.Context) (*WeatherStation, error) {
	wsc.defaults()
	return withHooks(ctx, wsc.sqlSave, wsc.mutation, wsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wsc *WeatherStationCreate) SaveX(ctx context.Context) *WeatherStation {
	v, err := wsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wsc *WeatherStationCreate) Exec(ctx context.Context) error {
	_, err := wsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsc *WeatherStationCreate) ExecX(ctx context.Context) {
	if err := wsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wsc *WeatherStationCreate) defaults() {
	if _, ok := wsc.mutation.ID(); !ok {
		v := weatherstation.DefaultID()
		wsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wsc *WeatherStationCreate) check() error {
	if _, ok := wsc.mutation.StationID(); !ok {
		return &ValidationError{Name: "station_id", err: errors.New(`ent: missing required field "WeatherStation.station_id"`)}
	}
	if _, ok := wsc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "WeatherStation.hash"`)}
	}
	return nil
}

func (wsc *WeatherStationCreate) sqlSave(ctx context.Context) (*WeatherStation, error) {
	if err := wsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wsc.driver, _spec); err != nil {
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
	wsc.mutation.id = &_node.ID
	wsc.mutation.done = true
	return _node, nil
}

func (wsc *WeatherStationCreate) createSpec() (*WeatherStation, *sqlgraph.CreateSpec) {
	var (
		_node = &WeatherStation{config: wsc.config}
		_spec = sqlgraph.NewCreateSpec(weatherstation.Table, sqlgraph.NewFieldSpec(weatherstation.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = wsc.conflict
	if id, ok := wsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wsc.mutation.StationID(); ok {
		_spec.SetField(weatherstation.FieldStationID, field.TypeString, value)
		_node.StationID = value
	}
	if value, ok := wsc.mutation.Latitude(); ok {
		_spec.SetField(weatherstation.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = &value
	}
	if value, ok := wsc.mutation.Longitude(); ok {
		_spec.SetField(weatherstation.FieldLongitude, field.TypeFloat64, value)
		_node.Longitude = &value
	}
	if value, ok := wsc.mutation.Elevation(); ok {
		_spec.SetField(weatherstation.FieldElevation, field.TypeFloat64, value)
		_node.Elevation = &value
	}
	if value, ok := wsc.mutation.Hash(); ok {
		_spec.SetField(weatherstation.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if nodes := wsc.mutation.AirportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   weatherstation.AirportTable,
			Columns: []string{weatherstation.AirportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.airport_station = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wsc.mutation.MetarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weatherstation.MetarsTable,
			Columns: []string{weatherstation.MetarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metar.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wsc.mutation.TafsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weatherstation.TafsTable,
			Columns: []string{weatherstation.TafsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taf.FieldID, field.TypeUUID),
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
//	client.WeatherStation.Create().
//		SetStationID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WeatherStationUpsert) {
//			SetStationID(v+v).
//		}).
//		Exec(ctx)
func (wsc *WeatherStationCreate) OnConflict(opts ...sql.ConflictOption) *WeatherStationUpsertOne {
	wsc.conflict = opts
	return &WeatherStationUpsertOne{
		create: wsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WeatherStation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wsc *WeatherStationCreate) OnConflictColumns(columns ...string) *WeatherStationUpsertOne {
	wsc.conflict = append(wsc.conflict, sql.ConflictColumns(columns...))
	return &WeatherStationUpsertOne{
		create: wsc,
	}
}

type (
	// WeatherStationUpsertOne is the builder for "upsert"-ing
	//  one WeatherStation node.
	WeatherStationUpsertOne struct {
		create *WeatherStationCreate
	}

	// WeatherStationUpsert is the "OnConflict" setter.
	WeatherStationUpsert struct {
		*sql.UpdateSet
	}
)

// SetLatitude sets the "latitude" field.
func (u *WeatherStationUpsert) SetLatitude(v float64) *WeatherStationUpsert {
	u.Set(weatherstation.FieldLatitude, v)
	return u
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *WeatherStationUpsert) UpdateLatitude() *WeatherStationUpsert {
	u.SetExcluded(weatherstation.FieldLatitude)
	return u
}

// AddLatitude adds v to the "latitude" field.
func (u *WeatherStationUpsert) AddLatitude(v float64) *WeatherStationUpsert {
	u.Add(weatherstation.FieldLatitude, v)
	return u
}

// ClearLatitude clears the value of the "latitude" field.
func (u *WeatherStationUpsert) ClearLatitude() *WeatherStationUpsert {
	u.SetNull(weatherstation.FieldLatitude)
	return u
}

// SetLongitude sets the "longitude" field.
func (u *WeatherStationUpsert) SetLongitude(v float64) *WeatherStationUpsert {
	u.Set(weatherstation.FieldLongitude, v)
	return u
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *WeatherStationUpsert) UpdateLongitude() *WeatherStationUpsert {
	u.SetExcluded(weatherstation.FieldLongitude)
	return u
}

// AddLongitude adds v to the "longitude" field.
func (u *WeatherStationUpsert) AddLongitude(v float64) *WeatherStationUpsert {
	u.Add(weatherstation.FieldLongitude, v)
	return u
}

// ClearLongitude clears the value of the "longitude" field.
func (u *WeatherStationUpsert) ClearLongitude() *WeatherStationUpsert {
	u.SetNull(weatherstation.FieldLongitude)
	return u
}

// SetElevation sets the "elevation" field.
func (u *WeatherStationUpsert) SetElevation(v float64) *WeatherStationUpsert {
	u.Set(weatherstation.FieldElevation, v)
	return u
}

// UpdateElevation sets the "elevation" field to the value that was provided on create.
func (u *WeatherStationUpsert) UpdateElevation() *WeatherStationUpsert {
	u.SetExcluded(weatherstation.FieldElevation)
	return u
}

// AddElevation adds v to the "elevation" field.
func (u *WeatherStationUpsert) AddElevation(v float64) *WeatherStationUpsert {
	u.Add(weatherstation.FieldElevation, v)
	return u
}

// ClearElevation clears the value of the "elevation" field.
func (u *WeatherStationUpsert) ClearElevation() *WeatherStationUpsert {
	u.SetNull(weatherstation.FieldElevation)
	return u
}

// SetHash sets the "hash" field.
func (u *WeatherStationUpsert) SetHash(v string) *WeatherStationUpsert {
	u.Set(weatherstation.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *WeatherStationUpsert) UpdateHash() *WeatherStationUpsert {
	u.SetExcluded(weatherstation.FieldHash)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.WeatherStation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(weatherstation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *WeatherStationUpsertOne) UpdateNewValues() *WeatherStationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(weatherstation.FieldID)
		}
		if _, exists := u.create.mutation.StationID(); exists {
			s.SetIgnore(weatherstation.FieldStationID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WeatherStation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *WeatherStationUpsertOne) Ignore() *WeatherStationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WeatherStationUpsertOne) DoNothing() *WeatherStationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WeatherStationCreate.OnConflict
// documentation for more info.
func (u *WeatherStationUpsertOne) Update(set func(*WeatherStationUpsert)) *WeatherStationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WeatherStationUpsert{UpdateSet: update})
	}))
	return u
}

// SetLatitude sets the "latitude" field.
func (u *WeatherStationUpsertOne) SetLatitude(v float64) *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetLatitude(v)
	})
}

// AddLatitude adds v to the "latitude" field.
func (u *WeatherStationUpsertOne) AddLatitude(v float64) *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.AddLatitude(v)
	})
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *WeatherStationUpsertOne) UpdateLatitude() *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateLatitude()
	})
}

// ClearLatitude clears the value of the "latitude" field.
func (u *WeatherStationUpsertOne) ClearLatitude() *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.ClearLatitude()
	})
}

// SetLongitude sets the "longitude" field.
func (u *WeatherStationUpsertOne) SetLongitude(v float64) *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetLongitude(v)
	})
}

// AddLongitude adds v to the "longitude" field.
func (u *WeatherStationUpsertOne) AddLongitude(v float64) *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.AddLongitude(v)
	})
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *WeatherStationUpsertOne) UpdateLongitude() *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateLongitude()
	})
}

// ClearLongitude clears the value of the "longitude" field.
func (u *WeatherStationUpsertOne) ClearLongitude() *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.ClearLongitude()
	})
}

// SetElevation sets the "elevation" field.
func (u *WeatherStationUpsertOne) SetElevation(v float64) *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetElevation(v)
	})
}

// AddElevation adds v to the "elevation" field.
func (u *WeatherStationUpsertOne) AddElevation(v float64) *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.AddElevation(v)
	})
}

// UpdateElevation sets the "elevation" field to the value that was provided on create.
func (u *WeatherStationUpsertOne) UpdateElevation() *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateElevation()
	})
}

// ClearElevation clears the value of the "elevation" field.
func (u *WeatherStationUpsertOne) ClearElevation() *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.ClearElevation()
	})
}

// SetHash sets the "hash" field.
func (u *WeatherStationUpsertOne) SetHash(v string) *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *WeatherStationUpsertOne) UpdateHash() *WeatherStationUpsertOne {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateHash()
	})
}

// Exec executes the query.
func (u *WeatherStationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WeatherStationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WeatherStationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *WeatherStationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: WeatherStationUpsertOne.ID is not supported by MySQL driver. Use WeatherStationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *WeatherStationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// WeatherStationCreateBulk is the builder for creating many WeatherStation entities in bulk.
type WeatherStationCreateBulk struct {
	config
	builders []*WeatherStationCreate
	conflict []sql.ConflictOption
}

// Save creates the WeatherStation entities in the database.
func (wscb *WeatherStationCreateBulk) Save(ctx context.Context) ([]*WeatherStation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wscb.builders))
	nodes := make([]*WeatherStation, len(wscb.builders))
	mutators := make([]Mutator, len(wscb.builders))
	for i := range wscb.builders {
		func(i int, root context.Context) {
			builder := wscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WeatherStationMutation)
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
					_, err = mutators[i+1].Mutate(root, wscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = wscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wscb *WeatherStationCreateBulk) SaveX(ctx context.Context) []*WeatherStation {
	v, err := wscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wscb *WeatherStationCreateBulk) Exec(ctx context.Context) error {
	_, err := wscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wscb *WeatherStationCreateBulk) ExecX(ctx context.Context) {
	if err := wscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.WeatherStation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WeatherStationUpsert) {
//			SetStationID(v+v).
//		}).
//		Exec(ctx)
func (wscb *WeatherStationCreateBulk) OnConflict(opts ...sql.ConflictOption) *WeatherStationUpsertBulk {
	wscb.conflict = opts
	return &WeatherStationUpsertBulk{
		create: wscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WeatherStation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wscb *WeatherStationCreateBulk) OnConflictColumns(columns ...string) *WeatherStationUpsertBulk {
	wscb.conflict = append(wscb.conflict, sql.ConflictColumns(columns...))
	return &WeatherStationUpsertBulk{
		create: wscb,
	}
}

// WeatherStationUpsertBulk is the builder for "upsert"-ing
// a bulk of WeatherStation nodes.
type WeatherStationUpsertBulk struct {
	create *WeatherStationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.WeatherStation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(weatherstation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *WeatherStationUpsertBulk) UpdateNewValues() *WeatherStationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(weatherstation.FieldID)
			}
			if _, exists := b.mutation.StationID(); exists {
				s.SetIgnore(weatherstation.FieldStationID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WeatherStation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *WeatherStationUpsertBulk) Ignore() *WeatherStationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WeatherStationUpsertBulk) DoNothing() *WeatherStationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WeatherStationCreateBulk.OnConflict
// documentation for more info.
func (u *WeatherStationUpsertBulk) Update(set func(*WeatherStationUpsert)) *WeatherStationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WeatherStationUpsert{UpdateSet: update})
	}))
	return u
}

// SetLatitude sets the "latitude" field.
func (u *WeatherStationUpsertBulk) SetLatitude(v float64) *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetLatitude(v)
	})
}

// AddLatitude adds v to the "latitude" field.
func (u *WeatherStationUpsertBulk) AddLatitude(v float64) *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.AddLatitude(v)
	})
}

// UpdateLatitude sets the "latitude" field to the value that was provided on create.
func (u *WeatherStationUpsertBulk) UpdateLatitude() *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateLatitude()
	})
}

// ClearLatitude clears the value of the "latitude" field.
func (u *WeatherStationUpsertBulk) ClearLatitude() *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.ClearLatitude()
	})
}

// SetLongitude sets the "longitude" field.
func (u *WeatherStationUpsertBulk) SetLongitude(v float64) *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetLongitude(v)
	})
}

// AddLongitude adds v to the "longitude" field.
func (u *WeatherStationUpsertBulk) AddLongitude(v float64) *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.AddLongitude(v)
	})
}

// UpdateLongitude sets the "longitude" field to the value that was provided on create.
func (u *WeatherStationUpsertBulk) UpdateLongitude() *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateLongitude()
	})
}

// ClearLongitude clears the value of the "longitude" field.
func (u *WeatherStationUpsertBulk) ClearLongitude() *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.ClearLongitude()
	})
}

// SetElevation sets the "elevation" field.
func (u *WeatherStationUpsertBulk) SetElevation(v float64) *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetElevation(v)
	})
}

// AddElevation adds v to the "elevation" field.
func (u *WeatherStationUpsertBulk) AddElevation(v float64) *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.AddElevation(v)
	})
}

// UpdateElevation sets the "elevation" field to the value that was provided on create.
func (u *WeatherStationUpsertBulk) UpdateElevation() *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateElevation()
	})
}

// ClearElevation clears the value of the "elevation" field.
func (u *WeatherStationUpsertBulk) ClearElevation() *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.ClearElevation()
	})
}

// SetHash sets the "hash" field.
func (u *WeatherStationUpsertBulk) SetHash(v string) *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *WeatherStationUpsertBulk) UpdateHash() *WeatherStationUpsertBulk {
	return u.Update(func(s *WeatherStationUpsert) {
		s.UpdateHash()
	})
}

// Exec executes the query.
func (u *WeatherStationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the WeatherStationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WeatherStationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WeatherStationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
