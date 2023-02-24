// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"metar.gg/ent/forecast"
	"metar.gg/ent/predicate"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
)

// TafQuery is the builder for querying Taf entities.
type TafQuery struct {
	config
	ctx               *QueryContext
	order             []OrderFunc
	inters            []Interceptor
	predicates        []predicate.Taf
	withStation       *WeatherStationQuery
	withForecast      *ForecastQuery
	withFKs           bool
	loadTotal         []func(context.Context, []*Taf) error
	modifiers         []func(*sql.Selector)
	withNamedForecast map[string]*ForecastQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TafQuery builder.
func (tq *TafQuery) Where(ps ...predicate.Taf) *TafQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TafQuery) Limit(limit int) *TafQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TafQuery) Offset(offset int) *TafQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TafQuery) Unique(unique bool) *TafQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TafQuery) Order(o ...OrderFunc) *TafQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryStation chains the current query on the "station" edge.
func (tq *TafQuery) QueryStation() *WeatherStationQuery {
	query := (&WeatherStationClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taf.Table, taf.FieldID, selector),
			sqlgraph.To(weatherstation.Table, weatherstation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, taf.StationTable, taf.StationColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryForecast chains the current query on the "forecast" edge.
func (tq *TafQuery) QueryForecast() *ForecastQuery {
	query := (&ForecastClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taf.Table, taf.FieldID, selector),
			sqlgraph.To(forecast.Table, forecast.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, taf.ForecastTable, taf.ForecastColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Taf entity from the query.
// Returns a *NotFoundError when no Taf was found.
func (tq *TafQuery) First(ctx context.Context) (*Taf, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{taf.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TafQuery) FirstX(ctx context.Context) *Taf {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Taf ID from the query.
// Returns a *NotFoundError when no Taf ID was found.
func (tq *TafQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{taf.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TafQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Taf entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Taf entity is found.
// Returns a *NotFoundError when no Taf entities are found.
func (tq *TafQuery) Only(ctx context.Context) (*Taf, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{taf.Label}
	default:
		return nil, &NotSingularError{taf.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TafQuery) OnlyX(ctx context.Context) *Taf {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Taf ID in the query.
// Returns a *NotSingularError when more than one Taf ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TafQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{taf.Label}
	default:
		err = &NotSingularError{taf.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TafQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tafs.
func (tq *TafQuery) All(ctx context.Context) ([]*Taf, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Taf, *TafQuery]()
	return withInterceptors[[]*Taf](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TafQuery) AllX(ctx context.Context) []*Taf {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Taf IDs.
func (tq *TafQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(taf.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TafQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TafQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TafQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TafQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TafQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TafQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TafQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TafQuery) Clone() *TafQuery {
	if tq == nil {
		return nil
	}
	return &TafQuery{
		config:       tq.config,
		ctx:          tq.ctx.Clone(),
		order:        append([]OrderFunc{}, tq.order...),
		inters:       append([]Interceptor{}, tq.inters...),
		predicates:   append([]predicate.Taf{}, tq.predicates...),
		withStation:  tq.withStation.Clone(),
		withForecast: tq.withForecast.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithStation tells the query-builder to eager-load the nodes that are connected to
// the "station" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TafQuery) WithStation(opts ...func(*WeatherStationQuery)) *TafQuery {
	query := (&WeatherStationClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withStation = query
	return tq
}

// WithForecast tells the query-builder to eager-load the nodes that are connected to
// the "forecast" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TafQuery) WithForecast(opts ...func(*ForecastQuery)) *TafQuery {
	query := (&ForecastClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withForecast = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RawText string `json:"raw_text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Taf.Query().
//		GroupBy(taf.FieldRawText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TafQuery) GroupBy(field string, fields ...string) *TafGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TafGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = taf.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RawText string `json:"raw_text,omitempty"`
//	}
//
//	client.Taf.Query().
//		Select(taf.FieldRawText).
//		Scan(ctx, &v)
func (tq *TafQuery) Select(fields ...string) *TafSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TafSelect{TafQuery: tq}
	sbuild.label = taf.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TafSelect configured with the given aggregations.
func (tq *TafQuery) Aggregate(fns ...AggregateFunc) *TafSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TafQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !taf.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TafQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Taf, error) {
	var (
		nodes       = []*Taf{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withStation != nil,
			tq.withForecast != nil,
		}
	)
	if tq.withStation != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, taf.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Taf).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Taf{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withStation; query != nil {
		if err := tq.loadStation(ctx, query, nodes, nil,
			func(n *Taf, e *WeatherStation) { n.Edges.Station = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withForecast; query != nil {
		if err := tq.loadForecast(ctx, query, nodes,
			func(n *Taf) { n.Edges.Forecast = []*Forecast{} },
			func(n *Taf, e *Forecast) { n.Edges.Forecast = append(n.Edges.Forecast, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedForecast {
		if err := tq.loadForecast(ctx, query, nodes,
			func(n *Taf) { n.appendNamedForecast(name) },
			func(n *Taf, e *Forecast) { n.appendNamedForecast(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TafQuery) loadStation(ctx context.Context, query *WeatherStationQuery, nodes []*Taf, init func(*Taf), assign func(*Taf, *WeatherStation)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Taf)
	for i := range nodes {
		if nodes[i].weather_station_tafs == nil {
			continue
		}
		fk := *nodes[i].weather_station_tafs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(weatherstation.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "weather_station_tafs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TafQuery) loadForecast(ctx context.Context, query *ForecastQuery, nodes []*Taf, init func(*Taf), assign func(*Taf, *Forecast)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Taf)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Forecast(func(s *sql.Selector) {
		s.Where(sql.InValues(taf.ForecastColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.taf_forecast
		if fk == nil {
			return fmt.Errorf(`foreign-key "taf_forecast" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "taf_forecast" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TafQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TafQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(taf.Table, taf.Columns, sqlgraph.NewFieldSpec(taf.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taf.FieldID)
		for i := range fields {
			if fields[i] != taf.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TafQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(taf.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = taf.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tq *TafQuery) Modify(modifiers ...func(s *sql.Selector)) *TafSelect {
	tq.modifiers = append(tq.modifiers, modifiers...)
	return tq.Select()
}

// WithNamedForecast tells the query-builder to eager-load the nodes that are connected to the "forecast"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TafQuery) WithNamedForecast(name string, opts ...func(*ForecastQuery)) *TafQuery {
	query := (&ForecastClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedForecast == nil {
		tq.withNamedForecast = make(map[string]*ForecastQuery)
	}
	tq.withNamedForecast[name] = query
	return tq
}

// TafGroupBy is the group-by builder for Taf entities.
type TafGroupBy struct {
	selector
	build *TafQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TafGroupBy) Aggregate(fns ...AggregateFunc) *TafGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TafGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TafQuery, *TafGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TafGroupBy) sqlScan(ctx context.Context, root *TafQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TafSelect is the builder for selecting fields of Taf entities.
type TafSelect struct {
	*TafQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TafSelect) Aggregate(fns ...AggregateFunc) *TafSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TafSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TafQuery, *TafSelect](ctx, ts.TafQuery, ts, ts.inters, v)
}

func (ts *TafSelect) sqlScan(ctx context.Context, root *TafQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ts *TafSelect) Modify(modifiers ...func(s *sql.Selector)) *TafSelect {
	ts.modifiers = append(ts.modifiers, modifiers...)
	return ts
}
