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
	"metar.gg/ent/airport"
	"metar.gg/ent/region"
)

// RegionCreate is the builder for creating a Region entity.
type RegionCreate struct {
	config
	mutation *RegionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetImportID sets the "import_id" field.
func (rc *RegionCreate) SetImportID(i int) *RegionCreate {
	rc.mutation.SetImportID(i)
	return rc
}

// SetHash sets the "hash" field.
func (rc *RegionCreate) SetHash(s string) *RegionCreate {
	rc.mutation.SetHash(s)
	return rc
}

// SetImportFlag sets the "import_flag" field.
func (rc *RegionCreate) SetImportFlag(b bool) *RegionCreate {
	rc.mutation.SetImportFlag(b)
	return rc
}

// SetNillableImportFlag sets the "import_flag" field if the given value is not nil.
func (rc *RegionCreate) SetNillableImportFlag(b *bool) *RegionCreate {
	if b != nil {
		rc.SetImportFlag(*b)
	}
	return rc
}

// SetLastUpdated sets the "last_updated" field.
func (rc *RegionCreate) SetLastUpdated(t time.Time) *RegionCreate {
	rc.mutation.SetLastUpdated(t)
	return rc
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (rc *RegionCreate) SetNillableLastUpdated(t *time.Time) *RegionCreate {
	if t != nil {
		rc.SetLastUpdated(*t)
	}
	return rc
}

// SetCode sets the "code" field.
func (rc *RegionCreate) SetCode(s string) *RegionCreate {
	rc.mutation.SetCode(s)
	return rc
}

// SetLocalCode sets the "local_code" field.
func (rc *RegionCreate) SetLocalCode(s string) *RegionCreate {
	rc.mutation.SetLocalCode(s)
	return rc
}

// SetName sets the "name" field.
func (rc *RegionCreate) SetName(s string) *RegionCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetWikipediaLink sets the "wikipedia_link" field.
func (rc *RegionCreate) SetWikipediaLink(s string) *RegionCreate {
	rc.mutation.SetWikipediaLink(s)
	return rc
}

// SetKeywords sets the "keywords" field.
func (rc *RegionCreate) SetKeywords(s []string) *RegionCreate {
	rc.mutation.SetKeywords(s)
	return rc
}

// SetID sets the "id" field.
func (rc *RegionCreate) SetID(u uuid.UUID) *RegionCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RegionCreate) SetNillableID(u *uuid.UUID) *RegionCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// AddAirportIDs adds the "airports" edge to the Airport entity by IDs.
func (rc *RegionCreate) AddAirportIDs(ids ...uuid.UUID) *RegionCreate {
	rc.mutation.AddAirportIDs(ids...)
	return rc
}

// AddAirports adds the "airports" edges to the Airport entity.
func (rc *RegionCreate) AddAirports(a ...*Airport) *RegionCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAirportIDs(ids...)
}

// Mutation returns the RegionMutation object of the builder.
func (rc *RegionCreate) Mutation() *RegionMutation {
	return rc.mutation
}

// Save creates the Region in the database.
func (rc *RegionCreate) Save(ctx context.Context) (*Region, error) {
	var (
		err  error
		node *Region
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Region)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RegionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RegionCreate) SaveX(ctx context.Context) *Region {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RegionCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RegionCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RegionCreate) defaults() {
	if _, ok := rc.mutation.ImportFlag(); !ok {
		v := region.DefaultImportFlag
		rc.mutation.SetImportFlag(v)
	}
	if _, ok := rc.mutation.LastUpdated(); !ok {
		v := region.DefaultLastUpdated()
		rc.mutation.SetLastUpdated(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := region.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RegionCreate) check() error {
	if _, ok := rc.mutation.ImportID(); !ok {
		return &ValidationError{Name: "import_id", err: errors.New(`ent: missing required field "Region.import_id"`)}
	}
	if _, ok := rc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Region.hash"`)}
	}
	if _, ok := rc.mutation.ImportFlag(); !ok {
		return &ValidationError{Name: "import_flag", err: errors.New(`ent: missing required field "Region.import_flag"`)}
	}
	if _, ok := rc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "last_updated", err: errors.New(`ent: missing required field "Region.last_updated"`)}
	}
	if _, ok := rc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Region.code"`)}
	}
	if _, ok := rc.mutation.LocalCode(); !ok {
		return &ValidationError{Name: "local_code", err: errors.New(`ent: missing required field "Region.local_code"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Region.name"`)}
	}
	if _, ok := rc.mutation.WikipediaLink(); !ok {
		return &ValidationError{Name: "wikipedia_link", err: errors.New(`ent: missing required field "Region.wikipedia_link"`)}
	}
	if _, ok := rc.mutation.Keywords(); !ok {
		return &ValidationError{Name: "keywords", err: errors.New(`ent: missing required field "Region.keywords"`)}
	}
	return nil
}

func (rc *RegionCreate) sqlSave(ctx context.Context) (*Region, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
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

func (rc *RegionCreate) createSpec() (*Region, *sqlgraph.CreateSpec) {
	var (
		_node = &Region{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: region.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: region.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.ImportID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: region.FieldImportID,
		})
		_node.ImportID = value
	}
	if value, ok := rc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: region.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := rc.mutation.ImportFlag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: region.FieldImportFlag,
		})
		_node.ImportFlag = value
	}
	if value, ok := rc.mutation.LastUpdated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: region.FieldLastUpdated,
		})
		_node.LastUpdated = value
	}
	if value, ok := rc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: region.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := rc.mutation.LocalCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: region.FieldLocalCode,
		})
		_node.LocalCode = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: region.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.WikipediaLink(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: region.FieldWikipediaLink,
		})
		_node.WikipediaLink = value
	}
	if value, ok := rc.mutation.Keywords(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: region.FieldKeywords,
		})
		_node.Keywords = value
	}
	if nodes := rc.mutation.AirportsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   region.AirportsTable,
			Columns: []string{region.AirportsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: airport.FieldID,
				},
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
//	client.Region.Create().
//		SetImportID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RegionUpsert) {
//			SetImportID(v+v).
//		}).
//		Exec(ctx)
func (rc *RegionCreate) OnConflict(opts ...sql.ConflictOption) *RegionUpsertOne {
	rc.conflict = opts
	return &RegionUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Region.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RegionCreate) OnConflictColumns(columns ...string) *RegionUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RegionUpsertOne{
		create: rc,
	}
}

type (
	// RegionUpsertOne is the builder for "upsert"-ing
	//  one Region node.
	RegionUpsertOne struct {
		create *RegionCreate
	}

	// RegionUpsert is the "OnConflict" setter.
	RegionUpsert struct {
		*sql.UpdateSet
	}
)

// SetImportID sets the "import_id" field.
func (u *RegionUpsert) SetImportID(v int) *RegionUpsert {
	u.Set(region.FieldImportID, v)
	return u
}

// UpdateImportID sets the "import_id" field to the value that was provided on create.
func (u *RegionUpsert) UpdateImportID() *RegionUpsert {
	u.SetExcluded(region.FieldImportID)
	return u
}

// AddImportID adds v to the "import_id" field.
func (u *RegionUpsert) AddImportID(v int) *RegionUpsert {
	u.Add(region.FieldImportID, v)
	return u
}

// SetHash sets the "hash" field.
func (u *RegionUpsert) SetHash(v string) *RegionUpsert {
	u.Set(region.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *RegionUpsert) UpdateHash() *RegionUpsert {
	u.SetExcluded(region.FieldHash)
	return u
}

// SetImportFlag sets the "import_flag" field.
func (u *RegionUpsert) SetImportFlag(v bool) *RegionUpsert {
	u.Set(region.FieldImportFlag, v)
	return u
}

// UpdateImportFlag sets the "import_flag" field to the value that was provided on create.
func (u *RegionUpsert) UpdateImportFlag() *RegionUpsert {
	u.SetExcluded(region.FieldImportFlag)
	return u
}

// SetLastUpdated sets the "last_updated" field.
func (u *RegionUpsert) SetLastUpdated(v time.Time) *RegionUpsert {
	u.Set(region.FieldLastUpdated, v)
	return u
}

// UpdateLastUpdated sets the "last_updated" field to the value that was provided on create.
func (u *RegionUpsert) UpdateLastUpdated() *RegionUpsert {
	u.SetExcluded(region.FieldLastUpdated)
	return u
}

// SetCode sets the "code" field.
func (u *RegionUpsert) SetCode(v string) *RegionUpsert {
	u.Set(region.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *RegionUpsert) UpdateCode() *RegionUpsert {
	u.SetExcluded(region.FieldCode)
	return u
}

// SetLocalCode sets the "local_code" field.
func (u *RegionUpsert) SetLocalCode(v string) *RegionUpsert {
	u.Set(region.FieldLocalCode, v)
	return u
}

// UpdateLocalCode sets the "local_code" field to the value that was provided on create.
func (u *RegionUpsert) UpdateLocalCode() *RegionUpsert {
	u.SetExcluded(region.FieldLocalCode)
	return u
}

// SetName sets the "name" field.
func (u *RegionUpsert) SetName(v string) *RegionUpsert {
	u.Set(region.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RegionUpsert) UpdateName() *RegionUpsert {
	u.SetExcluded(region.FieldName)
	return u
}

// SetWikipediaLink sets the "wikipedia_link" field.
func (u *RegionUpsert) SetWikipediaLink(v string) *RegionUpsert {
	u.Set(region.FieldWikipediaLink, v)
	return u
}

// UpdateWikipediaLink sets the "wikipedia_link" field to the value that was provided on create.
func (u *RegionUpsert) UpdateWikipediaLink() *RegionUpsert {
	u.SetExcluded(region.FieldWikipediaLink)
	return u
}

// SetKeywords sets the "keywords" field.
func (u *RegionUpsert) SetKeywords(v []string) *RegionUpsert {
	u.Set(region.FieldKeywords, v)
	return u
}

// UpdateKeywords sets the "keywords" field to the value that was provided on create.
func (u *RegionUpsert) UpdateKeywords() *RegionUpsert {
	u.SetExcluded(region.FieldKeywords)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Region.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(region.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RegionUpsertOne) UpdateNewValues() *RegionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(region.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Region.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RegionUpsertOne) Ignore() *RegionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RegionUpsertOne) DoNothing() *RegionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RegionCreate.OnConflict
// documentation for more info.
func (u *RegionUpsertOne) Update(set func(*RegionUpsert)) *RegionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RegionUpsert{UpdateSet: update})
	}))
	return u
}

// SetImportID sets the "import_id" field.
func (u *RegionUpsertOne) SetImportID(v int) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetImportID(v)
	})
}

// AddImportID adds v to the "import_id" field.
func (u *RegionUpsertOne) AddImportID(v int) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.AddImportID(v)
	})
}

// UpdateImportID sets the "import_id" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateImportID() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateImportID()
	})
}

// SetHash sets the "hash" field.
func (u *RegionUpsertOne) SetHash(v string) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateHash() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateHash()
	})
}

// SetImportFlag sets the "import_flag" field.
func (u *RegionUpsertOne) SetImportFlag(v bool) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetImportFlag(v)
	})
}

// UpdateImportFlag sets the "import_flag" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateImportFlag() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateImportFlag()
	})
}

// SetLastUpdated sets the "last_updated" field.
func (u *RegionUpsertOne) SetLastUpdated(v time.Time) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetLastUpdated(v)
	})
}

// UpdateLastUpdated sets the "last_updated" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateLastUpdated() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateLastUpdated()
	})
}

// SetCode sets the "code" field.
func (u *RegionUpsertOne) SetCode(v string) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateCode() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateCode()
	})
}

// SetLocalCode sets the "local_code" field.
func (u *RegionUpsertOne) SetLocalCode(v string) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetLocalCode(v)
	})
}

// UpdateLocalCode sets the "local_code" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateLocalCode() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateLocalCode()
	})
}

// SetName sets the "name" field.
func (u *RegionUpsertOne) SetName(v string) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateName() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateName()
	})
}

// SetWikipediaLink sets the "wikipedia_link" field.
func (u *RegionUpsertOne) SetWikipediaLink(v string) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetWikipediaLink(v)
	})
}

// UpdateWikipediaLink sets the "wikipedia_link" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateWikipediaLink() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateWikipediaLink()
	})
}

// SetKeywords sets the "keywords" field.
func (u *RegionUpsertOne) SetKeywords(v []string) *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.SetKeywords(v)
	})
}

// UpdateKeywords sets the "keywords" field to the value that was provided on create.
func (u *RegionUpsertOne) UpdateKeywords() *RegionUpsertOne {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateKeywords()
	})
}

// Exec executes the query.
func (u *RegionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RegionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RegionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RegionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: RegionUpsertOne.ID is not supported by MySQL driver. Use RegionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RegionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RegionCreateBulk is the builder for creating many Region entities in bulk.
type RegionCreateBulk struct {
	config
	builders []*RegionCreate
	conflict []sql.ConflictOption
}

// Save creates the Region entities in the database.
func (rcb *RegionCreateBulk) Save(ctx context.Context) ([]*Region, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Region, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RegionMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RegionCreateBulk) SaveX(ctx context.Context) []*Region {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RegionCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RegionCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Region.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RegionUpsert) {
//			SetImportID(v+v).
//		}).
//		Exec(ctx)
func (rcb *RegionCreateBulk) OnConflict(opts ...sql.ConflictOption) *RegionUpsertBulk {
	rcb.conflict = opts
	return &RegionUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Region.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RegionCreateBulk) OnConflictColumns(columns ...string) *RegionUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RegionUpsertBulk{
		create: rcb,
	}
}

// RegionUpsertBulk is the builder for "upsert"-ing
// a bulk of Region nodes.
type RegionUpsertBulk struct {
	create *RegionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Region.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(region.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RegionUpsertBulk) UpdateNewValues() *RegionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(region.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Region.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RegionUpsertBulk) Ignore() *RegionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RegionUpsertBulk) DoNothing() *RegionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RegionCreateBulk.OnConflict
// documentation for more info.
func (u *RegionUpsertBulk) Update(set func(*RegionUpsert)) *RegionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RegionUpsert{UpdateSet: update})
	}))
	return u
}

// SetImportID sets the "import_id" field.
func (u *RegionUpsertBulk) SetImportID(v int) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetImportID(v)
	})
}

// AddImportID adds v to the "import_id" field.
func (u *RegionUpsertBulk) AddImportID(v int) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.AddImportID(v)
	})
}

// UpdateImportID sets the "import_id" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateImportID() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateImportID()
	})
}

// SetHash sets the "hash" field.
func (u *RegionUpsertBulk) SetHash(v string) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateHash() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateHash()
	})
}

// SetImportFlag sets the "import_flag" field.
func (u *RegionUpsertBulk) SetImportFlag(v bool) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetImportFlag(v)
	})
}

// UpdateImportFlag sets the "import_flag" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateImportFlag() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateImportFlag()
	})
}

// SetLastUpdated sets the "last_updated" field.
func (u *RegionUpsertBulk) SetLastUpdated(v time.Time) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetLastUpdated(v)
	})
}

// UpdateLastUpdated sets the "last_updated" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateLastUpdated() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateLastUpdated()
	})
}

// SetCode sets the "code" field.
func (u *RegionUpsertBulk) SetCode(v string) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateCode() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateCode()
	})
}

// SetLocalCode sets the "local_code" field.
func (u *RegionUpsertBulk) SetLocalCode(v string) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetLocalCode(v)
	})
}

// UpdateLocalCode sets the "local_code" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateLocalCode() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateLocalCode()
	})
}

// SetName sets the "name" field.
func (u *RegionUpsertBulk) SetName(v string) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateName() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateName()
	})
}

// SetWikipediaLink sets the "wikipedia_link" field.
func (u *RegionUpsertBulk) SetWikipediaLink(v string) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetWikipediaLink(v)
	})
}

// UpdateWikipediaLink sets the "wikipedia_link" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateWikipediaLink() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateWikipediaLink()
	})
}

// SetKeywords sets the "keywords" field.
func (u *RegionUpsertBulk) SetKeywords(v []string) *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.SetKeywords(v)
	})
}

// UpdateKeywords sets the "keywords" field to the value that was provided on create.
func (u *RegionUpsertBulk) UpdateKeywords() *RegionUpsertBulk {
	return u.Update(func(s *RegionUpsert) {
		s.UpdateKeywords()
	})
}

// Exec executes the query.
func (u *RegionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RegionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RegionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RegionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
