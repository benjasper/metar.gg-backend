// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"metar.gg/ent/airport"
	"metar.gg/ent/predicate"
	"metar.gg/ent/runway"
)

// RunwayQuery is the builder for querying Runway entities.
type RunwayQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.Runway
	withAirport *AirportQuery
	withFKs     bool
	loadTotal   []func(context.Context, []*Runway) error
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RunwayQuery builder.
func (rq *RunwayQuery) Where(ps ...predicate.Runway) *RunwayQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RunwayQuery) Limit(limit int) *RunwayQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RunwayQuery) Offset(offset int) *RunwayQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RunwayQuery) Unique(unique bool) *RunwayQuery {
	rq.unique = &unique
	return rq
}

// Order adds an order step to the query.
func (rq *RunwayQuery) Order(o ...OrderFunc) *RunwayQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryAirport chains the current query on the "airport" edge.
func (rq *RunwayQuery) QueryAirport() *AirportQuery {
	query := &AirportQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(runway.Table, runway.FieldID, selector),
			sqlgraph.To(airport.Table, airport.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, runway.AirportTable, runway.AirportColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Runway entity from the query.
// Returns a *NotFoundError when no Runway was found.
func (rq *RunwayQuery) First(ctx context.Context) (*Runway, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{runway.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RunwayQuery) FirstX(ctx context.Context) *Runway {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Runway ID from the query.
// Returns a *NotFoundError when no Runway ID was found.
func (rq *RunwayQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{runway.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RunwayQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Runway entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Runway entity is found.
// Returns a *NotFoundError when no Runway entities are found.
func (rq *RunwayQuery) Only(ctx context.Context) (*Runway, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{runway.Label}
	default:
		return nil, &NotSingularError{runway.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RunwayQuery) OnlyX(ctx context.Context) *Runway {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Runway ID in the query.
// Returns a *NotSingularError when more than one Runway ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RunwayQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{runway.Label}
	default:
		err = &NotSingularError{runway.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RunwayQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Runways.
func (rq *RunwayQuery) All(ctx context.Context) ([]*Runway, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RunwayQuery) AllX(ctx context.Context) []*Runway {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Runway IDs.
func (rq *RunwayQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rq.Select(runway.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RunwayQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RunwayQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RunwayQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RunwayQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RunwayQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RunwayQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RunwayQuery) Clone() *RunwayQuery {
	if rq == nil {
		return nil
	}
	return &RunwayQuery{
		config:      rq.config,
		limit:       rq.limit,
		offset:      rq.offset,
		order:       append([]OrderFunc{}, rq.order...),
		predicates:  append([]predicate.Runway{}, rq.predicates...),
		withAirport: rq.withAirport.Clone(),
		// clone intermediate query.
		sql:    rq.sql.Clone(),
		path:   rq.path,
		unique: rq.unique,
	}
}

// WithAirport tells the query-builder to eager-load the nodes that are connected to
// the "airport" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RunwayQuery) WithAirport(opts ...func(*AirportQuery)) *RunwayQuery {
	query := &AirportQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withAirport = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ImportID int `json:"import_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Runway.Query().
//		GroupBy(runway.FieldImportID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RunwayQuery) GroupBy(field string, fields ...string) *RunwayGroupBy {
	grbuild := &RunwayGroupBy{config: rq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(ctx), nil
	}
	grbuild.label = runway.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ImportID int `json:"import_id,omitempty"`
//	}
//
//	client.Runway.Query().
//		Select(runway.FieldImportID).
//		Scan(ctx, &v)
func (rq *RunwayQuery) Select(fields ...string) *RunwaySelect {
	rq.fields = append(rq.fields, fields...)
	selbuild := &RunwaySelect{RunwayQuery: rq}
	selbuild.label = runway.Label
	selbuild.flds, selbuild.scan = &rq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a RunwaySelect configured with the given aggregations.
func (rq *RunwayQuery) Aggregate(fns ...AggregateFunc) *RunwaySelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RunwayQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rq.fields {
		if !runway.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RunwayQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Runway, error) {
	var (
		nodes       = []*Runway{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [1]bool{
			rq.withAirport != nil,
		}
	)
	if rq.withAirport != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, runway.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Runway).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Runway{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withAirport; query != nil {
		if err := rq.loadAirport(ctx, query, nodes, nil,
			func(n *Runway, e *Airport) { n.Edges.Airport = e }); err != nil {
			return nil, err
		}
	}
	for i := range rq.loadTotal {
		if err := rq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RunwayQuery) loadAirport(ctx context.Context, query *AirportQuery, nodes []*Runway, init func(*Runway), assign func(*Runway, *Airport)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Runway)
	for i := range nodes {
		if nodes[i].airport_runways == nil {
			continue
		}
		fk := *nodes[i].airport_runways
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(airport.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "airport_runways" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RunwayQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	if len(rq.modifiers) > 0 {
		_spec.Modifiers = rq.modifiers
	}
	_spec.Node.Columns = rq.fields
	if len(rq.fields) > 0 {
		_spec.Unique = rq.unique != nil && *rq.unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RunwayQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (rq *RunwayQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   runway.Table,
			Columns: runway.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: runway.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, runway.FieldID)
		for i := range fields {
			if fields[i] != runway.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RunwayQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(runway.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = runway.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.unique != nil && *rq.unique {
		selector.Distinct()
	}
	for _, m := range rq.modifiers {
		m(selector)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rq *RunwayQuery) Modify(modifiers ...func(s *sql.Selector)) *RunwaySelect {
	rq.modifiers = append(rq.modifiers, modifiers...)
	return rq.Select()
}

// RunwayGroupBy is the group-by builder for Runway entities.
type RunwayGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RunwayGroupBy) Aggregate(fns ...AggregateFunc) *RunwayGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rgb *RunwayGroupBy) Scan(ctx context.Context, v any) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

func (rgb *RunwayGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range rgb.fields {
		if !runway.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RunwayGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql.Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
		for _, f := range rgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rgb.fields...)...)
}

// RunwaySelect is the builder for selecting fields of Runway entities.
type RunwaySelect struct {
	*RunwayQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RunwaySelect) Aggregate(fns ...AggregateFunc) *RunwaySelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RunwaySelect) Scan(ctx context.Context, v any) error {
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	rs.sql = rs.RunwayQuery.sqlQuery(ctx)
	return rs.sqlScan(ctx, v)
}

func (rs *RunwaySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(rs.sql))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		rs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		rs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := rs.sql.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rs *RunwaySelect) Modify(modifiers ...func(s *sql.Selector)) *RunwaySelect {
	rs.modifiers = append(rs.modifiers, modifiers...)
	return rs
}
