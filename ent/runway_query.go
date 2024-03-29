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
	ctx         *QueryContext
	order       []runway.OrderOption
	inters      []Interceptor
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

// Limit the number of records to be returned by this query.
func (rq *RunwayQuery) Limit(limit int) *RunwayQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RunwayQuery) Offset(offset int) *RunwayQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RunwayQuery) Unique(unique bool) *RunwayQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RunwayQuery) Order(o ...runway.OrderOption) *RunwayQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryAirport chains the current query on the "airport" edge.
func (rq *RunwayQuery) QueryAirport() *AirportQuery {
	query := (&AirportClient{config: rq.config}).Query()
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
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
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
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
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
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
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
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
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
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Runway, *RunwayQuery]()
	return withInterceptors[[]*Runway](ctx, rq, qr, rq.inters)
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
func (rq *RunwayQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(runway.FieldID).Scan(ctx, &ids); err != nil {
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
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RunwayQuery](), rq.inters)
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
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
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
		ctx:         rq.ctx.Clone(),
		order:       append([]runway.OrderOption{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Runway{}, rq.predicates...),
		withAirport: rq.withAirport.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithAirport tells the query-builder to eager-load the nodes that are connected to
// the "airport" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RunwayQuery) WithAirport(opts ...func(*AirportQuery)) *RunwayQuery {
	query := (&AirportClient{config: rq.config}).Query()
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
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RunwayGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = runway.Label
	grbuild.scan = grbuild.Scan
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
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RunwaySelect{RunwayQuery: rq}
	sbuild.label = runway.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RunwaySelect configured with the given aggregations.
func (rq *RunwayQuery) Aggregate(fns ...AggregateFunc) *RunwaySelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RunwayQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
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
	if len(ids) == 0 {
		return nil
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
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RunwayQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(runway.Table, runway.Columns, sqlgraph.NewFieldSpec(runway.FieldID, field.TypeUUID))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
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
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
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
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = runway.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
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
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
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
	selector
	build *RunwayQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RunwayGroupBy) Aggregate(fns ...AggregateFunc) *RunwayGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RunwayGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RunwayQuery, *RunwayGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RunwayGroupBy) sqlScan(ctx context.Context, root *RunwayQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RunwaySelect is the builder for selecting fields of Runway entities.
type RunwaySelect struct {
	*RunwayQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RunwaySelect) Aggregate(fns ...AggregateFunc) *RunwaySelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RunwaySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RunwayQuery, *RunwaySelect](ctx, rs.RunwayQuery, rs, rs.inters, v)
}

func (rs *RunwaySelect) sqlScan(ctx context.Context, root *RunwayQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
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
