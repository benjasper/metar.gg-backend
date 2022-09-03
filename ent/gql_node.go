// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
	"metar.gg/ent/airport"
	"metar.gg/ent/frequency"
	"metar.gg/ent/runway"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (a *Airport) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Airport",
		Fields: make([]*Field, 17),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(a.Identifier); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "identifier",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Type); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "airport.Type",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Latitude); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "float64",
		Name:  "latitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Longitude); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "float64",
		Name:  "longitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Elevation); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "elevation",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Continent); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "continent",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Country); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "country",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Region); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "region",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Municipality); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "municipality",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.ScheduledService); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "bool",
		Name:  "scheduled_service",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.GpsCode); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "gps_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.IataCode); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "string",
		Name:  "iata_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.LocalCode); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "string",
		Name:  "local_code",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Website); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "string",
		Name:  "website",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Wikipedia); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "string",
		Name:  "wikipedia",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Keywords); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "[]string",
		Name:  "keywords",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Runway",
		Name: "runways",
	}
	err = a.QueryRunways().
		Select(runway.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (f *Frequency) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     f.ID,
		Type:   "Frequency",
		Fields: make([]*Field, 0),
		Edges:  make([]*Edge, 0),
	}
	return node, nil
}

func (r *Runway) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Runway",
		Fields: make([]*Field, 18),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(r.AirportIdentifier); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "airport_identifier",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Length); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "int",
		Name:  "length",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Width); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "width",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Surface); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "surface",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Lighted); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "bool",
		Name:  "lighted",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Closed); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "bool",
		Name:  "closed",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayIdentifier); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "low_runway_identifier",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayLatitude); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "float64",
		Name:  "low_runway_latitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayLongitude); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "float64",
		Name:  "low_runway_longitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayElevation); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "int",
		Name:  "low_runway_elevation",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayHeading); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "int",
		Name:  "low_runway_heading",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.LowRunwayDisplaced); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "int",
		Name:  "low_runway_displaced",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayIdentifier); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "string",
		Name:  "high_runway_identifier",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayLatitude); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "float64",
		Name:  "high_runway_latitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayLongitude); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "float64",
		Name:  "high_runway_longitude",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayElevation); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "int",
		Name:  "high_runway_elevation",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayHeading); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "int",
		Name:  "high_runway_heading",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.HighRunwayDisplaced); err != nil {
		return nil, err
	}
	node.Fields[17] = &Field{
		Type:  "int",
		Name:  "high_runway_displaced",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Airport",
		Name: "airport",
	}
	err = r.QueryAirport().
		Select(airport.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case airport.Table:
		query := c.Airport.Query().
			Where(airport.ID(id))
		query, err := query.CollectFields(ctx, "Airport")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case frequency.Table:
		query := c.Frequency.Query().
			Where(frequency.ID(id))
		query, err := query.CollectFields(ctx, "Frequency")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case runway.Table:
		query := c.Runway.Query().
			Where(runway.ID(id))
		query, err := query.CollectFields(ctx, "Runway")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case airport.Table:
		query := c.Airport.Query().
			Where(airport.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Airport")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case frequency.Table:
		query := c.Frequency.Query().
			Where(frequency.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Frequency")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case runway.Table:
		query := c.Runway.Query().
			Where(runway.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Runway")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
