// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
	"metar.gg/ent/skycondition"
)

// SkyConditionQuery is the builder for querying SkyCondition entities.
type SkyConditionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SkyCondition
	withMetar  *MetarQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*SkyCondition) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SkyConditionQuery builder.
func (scq *SkyConditionQuery) Where(ps ...predicate.SkyCondition) *SkyConditionQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit adds a limit step to the query.
func (scq *SkyConditionQuery) Limit(limit int) *SkyConditionQuery {
	scq.limit = &limit
	return scq
}

// Offset adds an offset step to the query.
func (scq *SkyConditionQuery) Offset(offset int) *SkyConditionQuery {
	scq.offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *SkyConditionQuery) Unique(unique bool) *SkyConditionQuery {
	scq.unique = &unique
	return scq
}

// Order adds an order step to the query.
func (scq *SkyConditionQuery) Order(o ...OrderFunc) *SkyConditionQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryMetar chains the current query on the "metar" edge.
func (scq *SkyConditionQuery) QueryMetar() *MetarQuery {
	query := &MetarQuery{config: scq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(skycondition.Table, skycondition.FieldID, selector),
			sqlgraph.To(metar.Table, metar.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, skycondition.MetarTable, skycondition.MetarColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SkyCondition entity from the query.
// Returns a *NotFoundError when no SkyCondition was found.
func (scq *SkyConditionQuery) First(ctx context.Context) (*SkyCondition, error) {
	nodes, err := scq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{skycondition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *SkyConditionQuery) FirstX(ctx context.Context) *SkyCondition {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SkyCondition ID from the query.
// Returns a *NotFoundError when no SkyCondition ID was found.
func (scq *SkyConditionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{skycondition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *SkyConditionQuery) FirstIDX(ctx context.Context) int {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SkyCondition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SkyCondition entity is found.
// Returns a *NotFoundError when no SkyCondition entities are found.
func (scq *SkyConditionQuery) Only(ctx context.Context) (*SkyCondition, error) {
	nodes, err := scq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{skycondition.Label}
	default:
		return nil, &NotSingularError{skycondition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *SkyConditionQuery) OnlyX(ctx context.Context) *SkyCondition {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SkyCondition ID in the query.
// Returns a *NotSingularError when more than one SkyCondition ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *SkyConditionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{skycondition.Label}
	default:
		err = &NotSingularError{skycondition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *SkyConditionQuery) OnlyIDX(ctx context.Context) int {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SkyConditions.
func (scq *SkyConditionQuery) All(ctx context.Context) ([]*SkyCondition, error) {
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return scq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (scq *SkyConditionQuery) AllX(ctx context.Context) []*SkyCondition {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SkyCondition IDs.
func (scq *SkyConditionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := scq.Select(skycondition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *SkyConditionQuery) IDsX(ctx context.Context) []int {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *SkyConditionQuery) Count(ctx context.Context) (int, error) {
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return scq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (scq *SkyConditionQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *SkyConditionQuery) Exist(ctx context.Context) (bool, error) {
	if err := scq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return scq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *SkyConditionQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SkyConditionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *SkyConditionQuery) Clone() *SkyConditionQuery {
	if scq == nil {
		return nil
	}
	return &SkyConditionQuery{
		config:     scq.config,
		limit:      scq.limit,
		offset:     scq.offset,
		order:      append([]OrderFunc{}, scq.order...),
		predicates: append([]predicate.SkyCondition{}, scq.predicates...),
		withMetar:  scq.withMetar.Clone(),
		// clone intermediate query.
		sql:    scq.sql.Clone(),
		path:   scq.path,
		unique: scq.unique,
	}
}

// WithMetar tells the query-builder to eager-load the nodes that are connected to
// the "metar" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *SkyConditionQuery) WithMetar(opts ...func(*MetarQuery)) *SkyConditionQuery {
	query := &MetarQuery{config: scq.config}
	for _, opt := range opts {
		opt(query)
	}
	scq.withMetar = query
	return scq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SkyCover skycondition.SkyCover `json:"sky_cover,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SkyCondition.Query().
//		GroupBy(skycondition.FieldSkyCover).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *SkyConditionQuery) GroupBy(field string, fields ...string) *SkyConditionGroupBy {
	grbuild := &SkyConditionGroupBy{config: scq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return scq.sqlQuery(ctx), nil
	}
	grbuild.label = skycondition.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SkyCover skycondition.SkyCover `json:"sky_cover,omitempty"`
//	}
//
//	client.SkyCondition.Query().
//		Select(skycondition.FieldSkyCover).
//		Scan(ctx, &v)
func (scq *SkyConditionQuery) Select(fields ...string) *SkyConditionSelect {
	scq.fields = append(scq.fields, fields...)
	selbuild := &SkyConditionSelect{SkyConditionQuery: scq}
	selbuild.label = skycondition.Label
	selbuild.flds, selbuild.scan = &scq.fields, selbuild.Scan
	return selbuild
}

func (scq *SkyConditionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range scq.fields {
		if !skycondition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *SkyConditionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SkyCondition, error) {
	var (
		nodes       = []*SkyCondition{}
		withFKs     = scq.withFKs
		_spec       = scq.querySpec()
		loadedTypes = [1]bool{
			scq.withMetar != nil,
		}
	)
	if scq.withMetar != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, skycondition.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SkyCondition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SkyCondition{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withMetar; query != nil {
		if err := scq.loadMetar(ctx, query, nodes, nil,
			func(n *SkyCondition, e *Metar) { n.Edges.Metar = e }); err != nil {
			return nil, err
		}
	}
	for i := range scq.loadTotal {
		if err := scq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *SkyConditionQuery) loadMetar(ctx context.Context, query *MetarQuery, nodes []*SkyCondition, init func(*SkyCondition), assign func(*SkyCondition, *Metar)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SkyCondition)
	for i := range nodes {
		if nodes[i].metar_sky_conditions == nil {
			continue
		}
		fk := *nodes[i].metar_sky_conditions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(metar.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "metar_sky_conditions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (scq *SkyConditionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	_spec.Node.Columns = scq.fields
	if len(scq.fields) > 0 {
		_spec.Unique = scq.unique != nil && *scq.unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *SkyConditionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := scq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (scq *SkyConditionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skycondition.Table,
			Columns: skycondition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skycondition.FieldID,
			},
		},
		From:   scq.sql,
		Unique: true,
	}
	if unique := scq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := scq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skycondition.FieldID)
		for i := range fields {
			if fields[i] != skycondition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *SkyConditionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(skycondition.Table)
	columns := scq.fields
	if len(columns) == 0 {
		columns = skycondition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.unique != nil && *scq.unique {
		selector.Distinct()
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SkyConditionGroupBy is the group-by builder for SkyCondition entities.
type SkyConditionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *SkyConditionGroupBy) Aggregate(fns ...AggregateFunc) *SkyConditionGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the group-by query and scans the result into the given value.
func (scgb *SkyConditionGroupBy) Scan(ctx context.Context, v any) error {
	query, err := scgb.path(ctx)
	if err != nil {
		return err
	}
	scgb.sql = query
	return scgb.sqlScan(ctx, v)
}

func (scgb *SkyConditionGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range scgb.fields {
		if !skycondition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := scgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (scgb *SkyConditionGroupBy) sqlQuery() *sql.Selector {
	selector := scgb.sql.Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(scgb.fields)+len(scgb.fns))
		for _, f := range scgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(scgb.fields...)...)
}

// SkyConditionSelect is the builder for selecting fields of SkyCondition entities.
type SkyConditionSelect struct {
	*SkyConditionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (scs *SkyConditionSelect) Scan(ctx context.Context, v any) error {
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	scs.sql = scs.SkyConditionQuery.sqlQuery(ctx)
	return scs.sqlScan(ctx, v)
}

func (scs *SkyConditionSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := scs.sql.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
