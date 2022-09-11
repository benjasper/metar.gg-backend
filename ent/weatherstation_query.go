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
	"metar.gg/ent/airport"
	"metar.gg/ent/metar"
	"metar.gg/ent/predicate"
	"metar.gg/ent/taf"
	"metar.gg/ent/weatherstation"
)

// WeatherStationQuery is the builder for querying WeatherStation entities.
type WeatherStationQuery struct {
	config
	limit           *int
	offset          *int
	unique          *bool
	order           []OrderFunc
	fields          []string
	predicates      []predicate.WeatherStation
	withAirport     *AirportQuery
	withMetars      *MetarQuery
	withTafs        *TafQuery
	withFKs         bool
	loadTotal       []func(context.Context, []*WeatherStation) error
	modifiers       []func(*sql.Selector)
	withNamedMetars map[string]*MetarQuery
	withNamedTafs   map[string]*TafQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WeatherStationQuery builder.
func (wsq *WeatherStationQuery) Where(ps ...predicate.WeatherStation) *WeatherStationQuery {
	wsq.predicates = append(wsq.predicates, ps...)
	return wsq
}

// Limit adds a limit step to the query.
func (wsq *WeatherStationQuery) Limit(limit int) *WeatherStationQuery {
	wsq.limit = &limit
	return wsq
}

// Offset adds an offset step to the query.
func (wsq *WeatherStationQuery) Offset(offset int) *WeatherStationQuery {
	wsq.offset = &offset
	return wsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wsq *WeatherStationQuery) Unique(unique bool) *WeatherStationQuery {
	wsq.unique = &unique
	return wsq
}

// Order adds an order step to the query.
func (wsq *WeatherStationQuery) Order(o ...OrderFunc) *WeatherStationQuery {
	wsq.order = append(wsq.order, o...)
	return wsq
}

// QueryAirport chains the current query on the "airport" edge.
func (wsq *WeatherStationQuery) QueryAirport() *AirportQuery {
	query := &AirportQuery{config: wsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(weatherstation.Table, weatherstation.FieldID, selector),
			sqlgraph.To(airport.Table, airport.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, weatherstation.AirportTable, weatherstation.AirportColumn),
		)
		fromU = sqlgraph.SetNeighbors(wsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMetars chains the current query on the "metars" edge.
func (wsq *WeatherStationQuery) QueryMetars() *MetarQuery {
	query := &MetarQuery{config: wsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(weatherstation.Table, weatherstation.FieldID, selector),
			sqlgraph.To(metar.Table, metar.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, weatherstation.MetarsTable, weatherstation.MetarsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTafs chains the current query on the "tafs" edge.
func (wsq *WeatherStationQuery) QueryTafs() *TafQuery {
	query := &TafQuery{config: wsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(weatherstation.Table, weatherstation.FieldID, selector),
			sqlgraph.To(taf.Table, taf.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, weatherstation.TafsTable, weatherstation.TafsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WeatherStation entity from the query.
// Returns a *NotFoundError when no WeatherStation was found.
func (wsq *WeatherStationQuery) First(ctx context.Context) (*WeatherStation, error) {
	nodes, err := wsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{weatherstation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wsq *WeatherStationQuery) FirstX(ctx context.Context) *WeatherStation {
	node, err := wsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WeatherStation ID from the query.
// Returns a *NotFoundError when no WeatherStation ID was found.
func (wsq *WeatherStationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{weatherstation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wsq *WeatherStationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WeatherStation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WeatherStation entity is found.
// Returns a *NotFoundError when no WeatherStation entities are found.
func (wsq *WeatherStationQuery) Only(ctx context.Context) (*WeatherStation, error) {
	nodes, err := wsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{weatherstation.Label}
	default:
		return nil, &NotSingularError{weatherstation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wsq *WeatherStationQuery) OnlyX(ctx context.Context) *WeatherStation {
	node, err := wsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WeatherStation ID in the query.
// Returns a *NotSingularError when more than one WeatherStation ID is found.
// Returns a *NotFoundError when no entities are found.
func (wsq *WeatherStationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{weatherstation.Label}
	default:
		err = &NotSingularError{weatherstation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wsq *WeatherStationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WeatherStations.
func (wsq *WeatherStationQuery) All(ctx context.Context) ([]*WeatherStation, error) {
	if err := wsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wsq *WeatherStationQuery) AllX(ctx context.Context) []*WeatherStation {
	nodes, err := wsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WeatherStation IDs.
func (wsq *WeatherStationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := wsq.Select(weatherstation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wsq *WeatherStationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wsq *WeatherStationQuery) Count(ctx context.Context) (int, error) {
	if err := wsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wsq *WeatherStationQuery) CountX(ctx context.Context) int {
	count, err := wsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wsq *WeatherStationQuery) Exist(ctx context.Context) (bool, error) {
	if err := wsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wsq *WeatherStationQuery) ExistX(ctx context.Context) bool {
	exist, err := wsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WeatherStationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wsq *WeatherStationQuery) Clone() *WeatherStationQuery {
	if wsq == nil {
		return nil
	}
	return &WeatherStationQuery{
		config:      wsq.config,
		limit:       wsq.limit,
		offset:      wsq.offset,
		order:       append([]OrderFunc{}, wsq.order...),
		predicates:  append([]predicate.WeatherStation{}, wsq.predicates...),
		withAirport: wsq.withAirport.Clone(),
		withMetars:  wsq.withMetars.Clone(),
		withTafs:    wsq.withTafs.Clone(),
		// clone intermediate query.
		sql:    wsq.sql.Clone(),
		path:   wsq.path,
		unique: wsq.unique,
	}
}

// WithAirport tells the query-builder to eager-load the nodes that are connected to
// the "airport" edge. The optional arguments are used to configure the query builder of the edge.
func (wsq *WeatherStationQuery) WithAirport(opts ...func(*AirportQuery)) *WeatherStationQuery {
	query := &AirportQuery{config: wsq.config}
	for _, opt := range opts {
		opt(query)
	}
	wsq.withAirport = query
	return wsq
}

// WithMetars tells the query-builder to eager-load the nodes that are connected to
// the "metars" edge. The optional arguments are used to configure the query builder of the edge.
func (wsq *WeatherStationQuery) WithMetars(opts ...func(*MetarQuery)) *WeatherStationQuery {
	query := &MetarQuery{config: wsq.config}
	for _, opt := range opts {
		opt(query)
	}
	wsq.withMetars = query
	return wsq
}

// WithTafs tells the query-builder to eager-load the nodes that are connected to
// the "tafs" edge. The optional arguments are used to configure the query builder of the edge.
func (wsq *WeatherStationQuery) WithTafs(opts ...func(*TafQuery)) *WeatherStationQuery {
	query := &TafQuery{config: wsq.config}
	for _, opt := range opts {
		opt(query)
	}
	wsq.withTafs = query
	return wsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StationID string `json:"station_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WeatherStation.Query().
//		GroupBy(weatherstation.FieldStationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wsq *WeatherStationQuery) GroupBy(field string, fields ...string) *WeatherStationGroupBy {
	grbuild := &WeatherStationGroupBy{config: wsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wsq.sqlQuery(ctx), nil
	}
	grbuild.label = weatherstation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StationID string `json:"station_id,omitempty"`
//	}
//
//	client.WeatherStation.Query().
//		Select(weatherstation.FieldStationID).
//		Scan(ctx, &v)
func (wsq *WeatherStationQuery) Select(fields ...string) *WeatherStationSelect {
	wsq.fields = append(wsq.fields, fields...)
	selbuild := &WeatherStationSelect{WeatherStationQuery: wsq}
	selbuild.label = weatherstation.Label
	selbuild.flds, selbuild.scan = &wsq.fields, selbuild.Scan
	return selbuild
}

func (wsq *WeatherStationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wsq.fields {
		if !weatherstation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wsq.path != nil {
		prev, err := wsq.path(ctx)
		if err != nil {
			return err
		}
		wsq.sql = prev
	}
	return nil
}

func (wsq *WeatherStationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WeatherStation, error) {
	var (
		nodes       = []*WeatherStation{}
		withFKs     = wsq.withFKs
		_spec       = wsq.querySpec()
		loadedTypes = [3]bool{
			wsq.withAirport != nil,
			wsq.withMetars != nil,
			wsq.withTafs != nil,
		}
	)
	if wsq.withAirport != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, weatherstation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WeatherStation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WeatherStation{config: wsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wsq.modifiers) > 0 {
		_spec.Modifiers = wsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wsq.withAirport; query != nil {
		if err := wsq.loadAirport(ctx, query, nodes, nil,
			func(n *WeatherStation, e *Airport) { n.Edges.Airport = e }); err != nil {
			return nil, err
		}
	}
	if query := wsq.withMetars; query != nil {
		if err := wsq.loadMetars(ctx, query, nodes,
			func(n *WeatherStation) { n.Edges.Metars = []*Metar{} },
			func(n *WeatherStation, e *Metar) { n.Edges.Metars = append(n.Edges.Metars, e) }); err != nil {
			return nil, err
		}
	}
	if query := wsq.withTafs; query != nil {
		if err := wsq.loadTafs(ctx, query, nodes,
			func(n *WeatherStation) { n.Edges.Tafs = []*Taf{} },
			func(n *WeatherStation, e *Taf) { n.Edges.Tafs = append(n.Edges.Tafs, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wsq.withNamedMetars {
		if err := wsq.loadMetars(ctx, query, nodes,
			func(n *WeatherStation) { n.appendNamedMetars(name) },
			func(n *WeatherStation, e *Metar) { n.appendNamedMetars(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wsq.withNamedTafs {
		if err := wsq.loadTafs(ctx, query, nodes,
			func(n *WeatherStation) { n.appendNamedTafs(name) },
			func(n *WeatherStation, e *Taf) { n.appendNamedTafs(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range wsq.loadTotal {
		if err := wsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wsq *WeatherStationQuery) loadAirport(ctx context.Context, query *AirportQuery, nodes []*WeatherStation, init func(*WeatherStation), assign func(*WeatherStation, *Airport)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WeatherStation)
	for i := range nodes {
		if nodes[i].airport_station == nil {
			continue
		}
		fk := *nodes[i].airport_station
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
			return fmt.Errorf(`unexpected foreign-key "airport_station" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wsq *WeatherStationQuery) loadMetars(ctx context.Context, query *MetarQuery, nodes []*WeatherStation, init func(*WeatherStation), assign func(*WeatherStation, *Metar)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*WeatherStation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Metar(func(s *sql.Selector) {
		s.Where(sql.InValues(weatherstation.MetarsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.weather_station_metars
		if fk == nil {
			return fmt.Errorf(`foreign-key "weather_station_metars" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "weather_station_metars" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wsq *WeatherStationQuery) loadTafs(ctx context.Context, query *TafQuery, nodes []*WeatherStation, init func(*WeatherStation), assign func(*WeatherStation, *Taf)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*WeatherStation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Taf(func(s *sql.Selector) {
		s.Where(sql.InValues(weatherstation.TafsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.weather_station_tafs
		if fk == nil {
			return fmt.Errorf(`foreign-key "weather_station_tafs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "weather_station_tafs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wsq *WeatherStationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wsq.querySpec()
	if len(wsq.modifiers) > 0 {
		_spec.Modifiers = wsq.modifiers
	}
	_spec.Node.Columns = wsq.fields
	if len(wsq.fields) > 0 {
		_spec.Unique = wsq.unique != nil && *wsq.unique
	}
	return sqlgraph.CountNodes(ctx, wsq.driver, _spec)
}

func (wsq *WeatherStationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wsq *WeatherStationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   weatherstation.Table,
			Columns: weatherstation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: weatherstation.FieldID,
			},
		},
		From:   wsq.sql,
		Unique: true,
	}
	if unique := wsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, weatherstation.FieldID)
		for i := range fields {
			if fields[i] != weatherstation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wsq *WeatherStationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wsq.driver.Dialect())
	t1 := builder.Table(weatherstation.Table)
	columns := wsq.fields
	if len(columns) == 0 {
		columns = weatherstation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wsq.sql != nil {
		selector = wsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wsq.unique != nil && *wsq.unique {
		selector.Distinct()
	}
	for _, m := range wsq.modifiers {
		m(selector)
	}
	for _, p := range wsq.predicates {
		p(selector)
	}
	for _, p := range wsq.order {
		p(selector)
	}
	if offset := wsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wsq *WeatherStationQuery) Modify(modifiers ...func(s *sql.Selector)) *WeatherStationSelect {
	wsq.modifiers = append(wsq.modifiers, modifiers...)
	return wsq.Select()
}

// WithNamedMetars tells the query-builder to eager-load the nodes that are connected to the "metars"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wsq *WeatherStationQuery) WithNamedMetars(name string, opts ...func(*MetarQuery)) *WeatherStationQuery {
	query := &MetarQuery{config: wsq.config}
	for _, opt := range opts {
		opt(query)
	}
	if wsq.withNamedMetars == nil {
		wsq.withNamedMetars = make(map[string]*MetarQuery)
	}
	wsq.withNamedMetars[name] = query
	return wsq
}

// WithNamedTafs tells the query-builder to eager-load the nodes that are connected to the "tafs"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wsq *WeatherStationQuery) WithNamedTafs(name string, opts ...func(*TafQuery)) *WeatherStationQuery {
	query := &TafQuery{config: wsq.config}
	for _, opt := range opts {
		opt(query)
	}
	if wsq.withNamedTafs == nil {
		wsq.withNamedTafs = make(map[string]*TafQuery)
	}
	wsq.withNamedTafs[name] = query
	return wsq
}

// WeatherStationGroupBy is the group-by builder for WeatherStation entities.
type WeatherStationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wsgb *WeatherStationGroupBy) Aggregate(fns ...AggregateFunc) *WeatherStationGroupBy {
	wsgb.fns = append(wsgb.fns, fns...)
	return wsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wsgb *WeatherStationGroupBy) Scan(ctx context.Context, v any) error {
	query, err := wsgb.path(ctx)
	if err != nil {
		return err
	}
	wsgb.sql = query
	return wsgb.sqlScan(ctx, v)
}

func (wsgb *WeatherStationGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range wsgb.fields {
		if !weatherstation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wsgb *WeatherStationGroupBy) sqlQuery() *sql.Selector {
	selector := wsgb.sql.Select()
	aggregation := make([]string, 0, len(wsgb.fns))
	for _, fn := range wsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wsgb.fields)+len(wsgb.fns))
		for _, f := range wsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wsgb.fields...)...)
}

// WeatherStationSelect is the builder for selecting fields of WeatherStation entities.
type WeatherStationSelect struct {
	*WeatherStationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wss *WeatherStationSelect) Scan(ctx context.Context, v any) error {
	if err := wss.prepareQuery(ctx); err != nil {
		return err
	}
	wss.sql = wss.WeatherStationQuery.sqlQuery(ctx)
	return wss.sqlScan(ctx, v)
}

func (wss *WeatherStationSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := wss.sql.Query()
	if err := wss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wss *WeatherStationSelect) Modify(modifiers ...func(s *sql.Selector)) *WeatherStationSelect {
	wss.modifiers = append(wss.modifiers, modifiers...)
	return wss
}
