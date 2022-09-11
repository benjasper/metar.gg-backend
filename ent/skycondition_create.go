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
	"metar.gg/ent/skycondition"
)

// SkyConditionCreate is the builder for creating a SkyCondition entity.
type SkyConditionCreate struct {
	config
	mutation *SkyConditionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSkyCover sets the "sky_cover" field.
func (scc *SkyConditionCreate) SetSkyCover(sc skycondition.SkyCover) *SkyConditionCreate {
	scc.mutation.SetSkyCover(sc)
	return scc
}

// SetCloudBase sets the "cloud_base" field.
func (scc *SkyConditionCreate) SetCloudBase(i int) *SkyConditionCreate {
	scc.mutation.SetCloudBase(i)
	return scc
}

// SetNillableCloudBase sets the "cloud_base" field if the given value is not nil.
func (scc *SkyConditionCreate) SetNillableCloudBase(i *int) *SkyConditionCreate {
	if i != nil {
		scc.SetCloudBase(*i)
	}
	return scc
}

// SetCloudType sets the "cloud_type" field.
func (scc *SkyConditionCreate) SetCloudType(st skycondition.CloudType) *SkyConditionCreate {
	scc.mutation.SetCloudType(st)
	return scc
}

// SetNillableCloudType sets the "cloud_type" field if the given value is not nil.
func (scc *SkyConditionCreate) SetNillableCloudType(st *skycondition.CloudType) *SkyConditionCreate {
	if st != nil {
		scc.SetCloudType(*st)
	}
	return scc
}

// SetID sets the "id" field.
func (scc *SkyConditionCreate) SetID(u uuid.UUID) *SkyConditionCreate {
	scc.mutation.SetID(u)
	return scc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (scc *SkyConditionCreate) SetNillableID(u *uuid.UUID) *SkyConditionCreate {
	if u != nil {
		scc.SetID(*u)
	}
	return scc
}

// Mutation returns the SkyConditionMutation object of the builder.
func (scc *SkyConditionCreate) Mutation() *SkyConditionMutation {
	return scc.mutation
}

// Save creates the SkyCondition in the database.
func (scc *SkyConditionCreate) Save(ctx context.Context) (*SkyCondition, error) {
	var (
		err  error
		node *SkyCondition
	)
	scc.defaults()
	if len(scc.hooks) == 0 {
		if err = scc.check(); err != nil {
			return nil, err
		}
		node, err = scc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkyConditionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scc.check(); err != nil {
				return nil, err
			}
			scc.mutation = mutation
			if node, err = scc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(scc.hooks) - 1; i >= 0; i-- {
			if scc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, scc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (scc *SkyConditionCreate) SaveX(ctx context.Context) *SkyCondition {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SkyConditionCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SkyConditionCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *SkyConditionCreate) defaults() {
	if _, ok := scc.mutation.ID(); !ok {
		v := skycondition.DefaultID()
		scc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SkyConditionCreate) check() error {
	if _, ok := scc.mutation.SkyCover(); !ok {
		return &ValidationError{Name: "sky_cover", err: errors.New(`ent: missing required field "SkyCondition.sky_cover"`)}
	}
	if v, ok := scc.mutation.SkyCover(); ok {
		if err := skycondition.SkyCoverValidator(v); err != nil {
			return &ValidationError{Name: "sky_cover", err: fmt.Errorf(`ent: validator failed for field "SkyCondition.sky_cover": %w`, err)}
		}
	}
	if v, ok := scc.mutation.CloudType(); ok {
		if err := skycondition.CloudTypeValidator(v); err != nil {
			return &ValidationError{Name: "cloud_type", err: fmt.Errorf(`ent: validator failed for field "SkyCondition.cloud_type": %w`, err)}
		}
	}
	return nil
}

func (scc *SkyConditionCreate) sqlSave(ctx context.Context) (*SkyCondition, error) {
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
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
	return _node, nil
}

func (scc *SkyConditionCreate) createSpec() (*SkyCondition, *sqlgraph.CreateSpec) {
	var (
		_node = &SkyCondition{config: scc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: skycondition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: skycondition.FieldID,
			},
		}
	)
	_spec.OnConflict = scc.conflict
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := scc.mutation.SkyCover(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: skycondition.FieldSkyCover,
		})
		_node.SkyCover = value
	}
	if value, ok := scc.mutation.CloudBase(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: skycondition.FieldCloudBase,
		})
		_node.CloudBase = &value
	}
	if value, ok := scc.mutation.CloudType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: skycondition.FieldCloudType,
		})
		_node.CloudType = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SkyCondition.Create().
//		SetSkyCover(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SkyConditionUpsert) {
//			SetSkyCover(v+v).
//		}).
//		Exec(ctx)
func (scc *SkyConditionCreate) OnConflict(opts ...sql.ConflictOption) *SkyConditionUpsertOne {
	scc.conflict = opts
	return &SkyConditionUpsertOne{
		create: scc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SkyCondition.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scc *SkyConditionCreate) OnConflictColumns(columns ...string) *SkyConditionUpsertOne {
	scc.conflict = append(scc.conflict, sql.ConflictColumns(columns...))
	return &SkyConditionUpsertOne{
		create: scc,
	}
}

type (
	// SkyConditionUpsertOne is the builder for "upsert"-ing
	//  one SkyCondition node.
	SkyConditionUpsertOne struct {
		create *SkyConditionCreate
	}

	// SkyConditionUpsert is the "OnConflict" setter.
	SkyConditionUpsert struct {
		*sql.UpdateSet
	}
)

// SetSkyCover sets the "sky_cover" field.
func (u *SkyConditionUpsert) SetSkyCover(v skycondition.SkyCover) *SkyConditionUpsert {
	u.Set(skycondition.FieldSkyCover, v)
	return u
}

// UpdateSkyCover sets the "sky_cover" field to the value that was provided on create.
func (u *SkyConditionUpsert) UpdateSkyCover() *SkyConditionUpsert {
	u.SetExcluded(skycondition.FieldSkyCover)
	return u
}

// SetCloudBase sets the "cloud_base" field.
func (u *SkyConditionUpsert) SetCloudBase(v int) *SkyConditionUpsert {
	u.Set(skycondition.FieldCloudBase, v)
	return u
}

// UpdateCloudBase sets the "cloud_base" field to the value that was provided on create.
func (u *SkyConditionUpsert) UpdateCloudBase() *SkyConditionUpsert {
	u.SetExcluded(skycondition.FieldCloudBase)
	return u
}

// AddCloudBase adds v to the "cloud_base" field.
func (u *SkyConditionUpsert) AddCloudBase(v int) *SkyConditionUpsert {
	u.Add(skycondition.FieldCloudBase, v)
	return u
}

// ClearCloudBase clears the value of the "cloud_base" field.
func (u *SkyConditionUpsert) ClearCloudBase() *SkyConditionUpsert {
	u.SetNull(skycondition.FieldCloudBase)
	return u
}

// SetCloudType sets the "cloud_type" field.
func (u *SkyConditionUpsert) SetCloudType(v skycondition.CloudType) *SkyConditionUpsert {
	u.Set(skycondition.FieldCloudType, v)
	return u
}

// UpdateCloudType sets the "cloud_type" field to the value that was provided on create.
func (u *SkyConditionUpsert) UpdateCloudType() *SkyConditionUpsert {
	u.SetExcluded(skycondition.FieldCloudType)
	return u
}

// ClearCloudType clears the value of the "cloud_type" field.
func (u *SkyConditionUpsert) ClearCloudType() *SkyConditionUpsert {
	u.SetNull(skycondition.FieldCloudType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SkyCondition.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(skycondition.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SkyConditionUpsertOne) UpdateNewValues() *SkyConditionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(skycondition.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SkyCondition.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SkyConditionUpsertOne) Ignore() *SkyConditionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SkyConditionUpsertOne) DoNothing() *SkyConditionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SkyConditionCreate.OnConflict
// documentation for more info.
func (u *SkyConditionUpsertOne) Update(set func(*SkyConditionUpsert)) *SkyConditionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SkyConditionUpsert{UpdateSet: update})
	}))
	return u
}

// SetSkyCover sets the "sky_cover" field.
func (u *SkyConditionUpsertOne) SetSkyCover(v skycondition.SkyCover) *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.SetSkyCover(v)
	})
}

// UpdateSkyCover sets the "sky_cover" field to the value that was provided on create.
func (u *SkyConditionUpsertOne) UpdateSkyCover() *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.UpdateSkyCover()
	})
}

// SetCloudBase sets the "cloud_base" field.
func (u *SkyConditionUpsertOne) SetCloudBase(v int) *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.SetCloudBase(v)
	})
}

// AddCloudBase adds v to the "cloud_base" field.
func (u *SkyConditionUpsertOne) AddCloudBase(v int) *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.AddCloudBase(v)
	})
}

// UpdateCloudBase sets the "cloud_base" field to the value that was provided on create.
func (u *SkyConditionUpsertOne) UpdateCloudBase() *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.UpdateCloudBase()
	})
}

// ClearCloudBase clears the value of the "cloud_base" field.
func (u *SkyConditionUpsertOne) ClearCloudBase() *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.ClearCloudBase()
	})
}

// SetCloudType sets the "cloud_type" field.
func (u *SkyConditionUpsertOne) SetCloudType(v skycondition.CloudType) *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.SetCloudType(v)
	})
}

// UpdateCloudType sets the "cloud_type" field to the value that was provided on create.
func (u *SkyConditionUpsertOne) UpdateCloudType() *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.UpdateCloudType()
	})
}

// ClearCloudType clears the value of the "cloud_type" field.
func (u *SkyConditionUpsertOne) ClearCloudType() *SkyConditionUpsertOne {
	return u.Update(func(s *SkyConditionUpsert) {
		s.ClearCloudType()
	})
}

// Exec executes the query.
func (u *SkyConditionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SkyConditionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SkyConditionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SkyConditionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SkyConditionUpsertOne.ID is not supported by MySQL driver. Use SkyConditionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SkyConditionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SkyConditionCreateBulk is the builder for creating many SkyCondition entities in bulk.
type SkyConditionCreateBulk struct {
	config
	builders []*SkyConditionCreate
	conflict []sql.ConflictOption
}

// Save creates the SkyCondition entities in the database.
func (sccb *SkyConditionCreateBulk) Save(ctx context.Context) ([]*SkyCondition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SkyCondition, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkyConditionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SkyConditionCreateBulk) SaveX(ctx context.Context) []*SkyCondition {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SkyConditionCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SkyConditionCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SkyCondition.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SkyConditionUpsert) {
//			SetSkyCover(v+v).
//		}).
//		Exec(ctx)
func (sccb *SkyConditionCreateBulk) OnConflict(opts ...sql.ConflictOption) *SkyConditionUpsertBulk {
	sccb.conflict = opts
	return &SkyConditionUpsertBulk{
		create: sccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SkyCondition.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sccb *SkyConditionCreateBulk) OnConflictColumns(columns ...string) *SkyConditionUpsertBulk {
	sccb.conflict = append(sccb.conflict, sql.ConflictColumns(columns...))
	return &SkyConditionUpsertBulk{
		create: sccb,
	}
}

// SkyConditionUpsertBulk is the builder for "upsert"-ing
// a bulk of SkyCondition nodes.
type SkyConditionUpsertBulk struct {
	create *SkyConditionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SkyCondition.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(skycondition.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SkyConditionUpsertBulk) UpdateNewValues() *SkyConditionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(skycondition.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SkyCondition.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SkyConditionUpsertBulk) Ignore() *SkyConditionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SkyConditionUpsertBulk) DoNothing() *SkyConditionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SkyConditionCreateBulk.OnConflict
// documentation for more info.
func (u *SkyConditionUpsertBulk) Update(set func(*SkyConditionUpsert)) *SkyConditionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SkyConditionUpsert{UpdateSet: update})
	}))
	return u
}

// SetSkyCover sets the "sky_cover" field.
func (u *SkyConditionUpsertBulk) SetSkyCover(v skycondition.SkyCover) *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.SetSkyCover(v)
	})
}

// UpdateSkyCover sets the "sky_cover" field to the value that was provided on create.
func (u *SkyConditionUpsertBulk) UpdateSkyCover() *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.UpdateSkyCover()
	})
}

// SetCloudBase sets the "cloud_base" field.
func (u *SkyConditionUpsertBulk) SetCloudBase(v int) *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.SetCloudBase(v)
	})
}

// AddCloudBase adds v to the "cloud_base" field.
func (u *SkyConditionUpsertBulk) AddCloudBase(v int) *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.AddCloudBase(v)
	})
}

// UpdateCloudBase sets the "cloud_base" field to the value that was provided on create.
func (u *SkyConditionUpsertBulk) UpdateCloudBase() *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.UpdateCloudBase()
	})
}

// ClearCloudBase clears the value of the "cloud_base" field.
func (u *SkyConditionUpsertBulk) ClearCloudBase() *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.ClearCloudBase()
	})
}

// SetCloudType sets the "cloud_type" field.
func (u *SkyConditionUpsertBulk) SetCloudType(v skycondition.CloudType) *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.SetCloudType(v)
	})
}

// UpdateCloudType sets the "cloud_type" field to the value that was provided on create.
func (u *SkyConditionUpsertBulk) UpdateCloudType() *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.UpdateCloudType()
	})
}

// ClearCloudType clears the value of the "cloud_type" field.
func (u *SkyConditionUpsertBulk) ClearCloudType() *SkyConditionUpsertBulk {
	return u.Update(func(s *SkyConditionUpsert) {
		s.ClearCloudType()
	})
}

// Exec executes the query.
func (u *SkyConditionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SkyConditionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SkyConditionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SkyConditionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
