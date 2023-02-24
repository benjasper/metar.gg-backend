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
	"metar.gg/ent/country"
	"metar.gg/ent/frequency"
	"metar.gg/ent/predicate"
	"metar.gg/ent/region"
	"metar.gg/ent/runway"
	"metar.gg/ent/weatherstation"
)

// AirportQuery is the builder for querying Airport entities.
type AirportQuery struct {
	config
	ctx                  *QueryContext
	order                []OrderFunc
	inters               []Interceptor
	predicates           []predicate.Airport
	withRegion           *RegionQuery
	withCountry          *CountryQuery
	withRunways          *RunwayQuery
	withFrequencies      *FrequencyQuery
	withStation          *WeatherStationQuery
	withFKs              bool
	loadTotal            []func(context.Context, []*Airport) error
	modifiers            []func(*sql.Selector)
	withNamedRunways     map[string]*RunwayQuery
	withNamedFrequencies map[string]*FrequencyQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AirportQuery builder.
func (aq *AirportQuery) Where(ps ...predicate.Airport) *AirportQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AirportQuery) Limit(limit int) *AirportQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AirportQuery) Offset(offset int) *AirportQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AirportQuery) Unique(unique bool) *AirportQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AirportQuery) Order(o ...OrderFunc) *AirportQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryRegion chains the current query on the "region" edge.
func (aq *AirportQuery) QueryRegion() *RegionQuery {
	query := (&RegionClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(region.Table, region.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, airport.RegionTable, airport.RegionColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCountry chains the current query on the "country" edge.
func (aq *AirportQuery) QueryCountry() *CountryQuery {
	query := (&CountryClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, airport.CountryTable, airport.CountryColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRunways chains the current query on the "runways" edge.
func (aq *AirportQuery) QueryRunways() *RunwayQuery {
	query := (&RunwayClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(runway.Table, runway.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, airport.RunwaysTable, airport.RunwaysColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFrequencies chains the current query on the "frequencies" edge.
func (aq *AirportQuery) QueryFrequencies() *FrequencyQuery {
	query := (&FrequencyClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(frequency.Table, frequency.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, airport.FrequenciesTable, airport.FrequenciesColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStation chains the current query on the "station" edge.
func (aq *AirportQuery) QueryStation() *WeatherStationQuery {
	query := (&WeatherStationClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(airport.Table, airport.FieldID, selector),
			sqlgraph.To(weatherstation.Table, weatherstation.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, airport.StationTable, airport.StationColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Airport entity from the query.
// Returns a *NotFoundError when no Airport was found.
func (aq *AirportQuery) First(ctx context.Context) (*Airport, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{airport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AirportQuery) FirstX(ctx context.Context) *Airport {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Airport ID from the query.
// Returns a *NotFoundError when no Airport ID was found.
func (aq *AirportQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{airport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AirportQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Airport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Airport entity is found.
// Returns a *NotFoundError when no Airport entities are found.
func (aq *AirportQuery) Only(ctx context.Context) (*Airport, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{airport.Label}
	default:
		return nil, &NotSingularError{airport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AirportQuery) OnlyX(ctx context.Context) *Airport {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Airport ID in the query.
// Returns a *NotSingularError when more than one Airport ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AirportQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{airport.Label}
	default:
		err = &NotSingularError{airport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AirportQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Airports.
func (aq *AirportQuery) All(ctx context.Context) ([]*Airport, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Airport, *AirportQuery]()
	return withInterceptors[[]*Airport](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AirportQuery) AllX(ctx context.Context) []*Airport {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Airport IDs.
func (aq *AirportQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(airport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AirportQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AirportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AirportQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AirportQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AirportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AirportQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AirportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AirportQuery) Clone() *AirportQuery {
	if aq == nil {
		return nil
	}
	return &AirportQuery{
		config:          aq.config,
		ctx:             aq.ctx.Clone(),
		order:           append([]OrderFunc{}, aq.order...),
		inters:          append([]Interceptor{}, aq.inters...),
		predicates:      append([]predicate.Airport{}, aq.predicates...),
		withRegion:      aq.withRegion.Clone(),
		withCountry:     aq.withCountry.Clone(),
		withRunways:     aq.withRunways.Clone(),
		withFrequencies: aq.withFrequencies.Clone(),
		withStation:     aq.withStation.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithRegion tells the query-builder to eager-load the nodes that are connected to
// the "region" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithRegion(opts ...func(*RegionQuery)) *AirportQuery {
	query := (&RegionClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withRegion = query
	return aq
}

// WithCountry tells the query-builder to eager-load the nodes that are connected to
// the "country" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithCountry(opts ...func(*CountryQuery)) *AirportQuery {
	query := (&CountryClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withCountry = query
	return aq
}

// WithRunways tells the query-builder to eager-load the nodes that are connected to
// the "runways" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithRunways(opts ...func(*RunwayQuery)) *AirportQuery {
	query := (&RunwayClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withRunways = query
	return aq
}

// WithFrequencies tells the query-builder to eager-load the nodes that are connected to
// the "frequencies" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithFrequencies(opts ...func(*FrequencyQuery)) *AirportQuery {
	query := (&FrequencyClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withFrequencies = query
	return aq
}

// WithStation tells the query-builder to eager-load the nodes that are connected to
// the "station" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithStation(opts ...func(*WeatherStationQuery)) *AirportQuery {
	query := (&WeatherStationClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withStation = query
	return aq
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
//	client.Airport.Query().
//		GroupBy(airport.FieldImportID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AirportQuery) GroupBy(field string, fields ...string) *AirportGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AirportGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = airport.Label
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
//	client.Airport.Query().
//		Select(airport.FieldImportID).
//		Scan(ctx, &v)
func (aq *AirportQuery) Select(fields ...string) *AirportSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AirportSelect{AirportQuery: aq}
	sbuild.label = airport.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AirportSelect configured with the given aggregations.
func (aq *AirportQuery) Aggregate(fns ...AggregateFunc) *AirportSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AirportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !airport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AirportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Airport, error) {
	var (
		nodes       = []*Airport{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [5]bool{
			aq.withRegion != nil,
			aq.withCountry != nil,
			aq.withRunways != nil,
			aq.withFrequencies != nil,
			aq.withStation != nil,
		}
	)
	if aq.withRegion != nil || aq.withCountry != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, airport.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Airport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Airport{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withRegion; query != nil {
		if err := aq.loadRegion(ctx, query, nodes, nil,
			func(n *Airport, e *Region) { n.Edges.Region = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withCountry; query != nil {
		if err := aq.loadCountry(ctx, query, nodes, nil,
			func(n *Airport, e *Country) { n.Edges.Country = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withRunways; query != nil {
		if err := aq.loadRunways(ctx, query, nodes,
			func(n *Airport) { n.Edges.Runways = []*Runway{} },
			func(n *Airport, e *Runway) { n.Edges.Runways = append(n.Edges.Runways, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withFrequencies; query != nil {
		if err := aq.loadFrequencies(ctx, query, nodes,
			func(n *Airport) { n.Edges.Frequencies = []*Frequency{} },
			func(n *Airport, e *Frequency) { n.Edges.Frequencies = append(n.Edges.Frequencies, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withStation; query != nil {
		if err := aq.loadStation(ctx, query, nodes, nil,
			func(n *Airport, e *WeatherStation) { n.Edges.Station = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedRunways {
		if err := aq.loadRunways(ctx, query, nodes,
			func(n *Airport) { n.appendNamedRunways(name) },
			func(n *Airport, e *Runway) { n.appendNamedRunways(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedFrequencies {
		if err := aq.loadFrequencies(ctx, query, nodes,
			func(n *Airport) { n.appendNamedFrequencies(name) },
			func(n *Airport, e *Frequency) { n.appendNamedFrequencies(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AirportQuery) loadRegion(ctx context.Context, query *RegionQuery, nodes []*Airport, init func(*Airport), assign func(*Airport, *Region)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Airport)
	for i := range nodes {
		if nodes[i].region_airports == nil {
			continue
		}
		fk := *nodes[i].region_airports
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(region.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "region_airports" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AirportQuery) loadCountry(ctx context.Context, query *CountryQuery, nodes []*Airport, init func(*Airport), assign func(*Airport, *Country)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Airport)
	for i := range nodes {
		if nodes[i].country_airports == nil {
			continue
		}
		fk := *nodes[i].country_airports
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(country.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "country_airports" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AirportQuery) loadRunways(ctx context.Context, query *RunwayQuery, nodes []*Airport, init func(*Airport), assign func(*Airport, *Runway)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Airport)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Runway(func(s *sql.Selector) {
		s.Where(sql.InValues(airport.RunwaysColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.airport_runways
		if fk == nil {
			return fmt.Errorf(`foreign-key "airport_runways" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "airport_runways" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AirportQuery) loadFrequencies(ctx context.Context, query *FrequencyQuery, nodes []*Airport, init func(*Airport), assign func(*Airport, *Frequency)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Airport)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Frequency(func(s *sql.Selector) {
		s.Where(sql.InValues(airport.FrequenciesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.airport_frequencies
		if fk == nil {
			return fmt.Errorf(`foreign-key "airport_frequencies" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "airport_frequencies" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AirportQuery) loadStation(ctx context.Context, query *WeatherStationQuery, nodes []*Airport, init func(*Airport), assign func(*Airport, *WeatherStation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Airport)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.WeatherStation(func(s *sql.Selector) {
		s.Where(sql.InValues(airport.StationColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.airport_station
		if fk == nil {
			return fmt.Errorf(`foreign-key "airport_station" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "airport_station" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AirportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AirportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeUUID))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airport.FieldID)
		for i := range fields {
			if fields[i] != airport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AirportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(airport.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = airport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aq *AirportQuery) Modify(modifiers ...func(s *sql.Selector)) *AirportSelect {
	aq.modifiers = append(aq.modifiers, modifiers...)
	return aq.Select()
}

// WithNamedRunways tells the query-builder to eager-load the nodes that are connected to the "runways"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithNamedRunways(name string, opts ...func(*RunwayQuery)) *AirportQuery {
	query := (&RunwayClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedRunways == nil {
		aq.withNamedRunways = make(map[string]*RunwayQuery)
	}
	aq.withNamedRunways[name] = query
	return aq
}

// WithNamedFrequencies tells the query-builder to eager-load the nodes that are connected to the "frequencies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AirportQuery) WithNamedFrequencies(name string, opts ...func(*FrequencyQuery)) *AirportQuery {
	query := (&FrequencyClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedFrequencies == nil {
		aq.withNamedFrequencies = make(map[string]*FrequencyQuery)
	}
	aq.withNamedFrequencies[name] = query
	return aq
}

// AirportGroupBy is the group-by builder for Airport entities.
type AirportGroupBy struct {
	selector
	build *AirportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AirportGroupBy) Aggregate(fns ...AggregateFunc) *AirportGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AirportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AirportQuery, *AirportGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AirportGroupBy) sqlScan(ctx context.Context, root *AirportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AirportSelect is the builder for selecting fields of Airport entities.
type AirportSelect struct {
	*AirportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AirportSelect) Aggregate(fns ...AggregateFunc) *AirportSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AirportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AirportQuery, *AirportSelect](ctx, as.AirportQuery, as, as.inters, v)
}

func (as *AirportSelect) sqlScan(ctx context.Context, root *AirportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (as *AirportSelect) Modify(modifiers ...func(s *sql.Selector)) *AirportSelect {
	as.modifiers = append(as.modifiers, modifiers...)
	return as
}
