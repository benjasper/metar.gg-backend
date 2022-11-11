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
	"metar.gg/ent/predicate"
	"metar.gg/ent/turbulencecondition"
)

// TurbulenceConditionQuery is the builder for querying TurbulenceCondition entities.
type TurbulenceConditionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TurbulenceCondition
	withFKs    bool
	loadTotal  []func(context.Context, []*TurbulenceCondition) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TurbulenceConditionQuery builder.
func (tcq *TurbulenceConditionQuery) Where(ps ...predicate.TurbulenceCondition) *TurbulenceConditionQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit adds a limit step to the query.
func (tcq *TurbulenceConditionQuery) Limit(limit int) *TurbulenceConditionQuery {
	tcq.limit = &limit
	return tcq
}

// Offset adds an offset step to the query.
func (tcq *TurbulenceConditionQuery) Offset(offset int) *TurbulenceConditionQuery {
	tcq.offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TurbulenceConditionQuery) Unique(unique bool) *TurbulenceConditionQuery {
	tcq.unique = &unique
	return tcq
}

// Order adds an order step to the query.
func (tcq *TurbulenceConditionQuery) Order(o ...OrderFunc) *TurbulenceConditionQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// First returns the first TurbulenceCondition entity from the query.
// Returns a *NotFoundError when no TurbulenceCondition was found.
func (tcq *TurbulenceConditionQuery) First(ctx context.Context) (*TurbulenceCondition, error) {
	nodes, err := tcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{turbulencecondition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) FirstX(ctx context.Context) *TurbulenceCondition {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TurbulenceCondition ID from the query.
// Returns a *NotFoundError when no TurbulenceCondition ID was found.
func (tcq *TurbulenceConditionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{turbulencecondition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TurbulenceCondition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TurbulenceCondition entity is found.
// Returns a *NotFoundError when no TurbulenceCondition entities are found.
func (tcq *TurbulenceConditionQuery) Only(ctx context.Context) (*TurbulenceCondition, error) {
	nodes, err := tcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{turbulencecondition.Label}
	default:
		return nil, &NotSingularError{turbulencecondition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) OnlyX(ctx context.Context) *TurbulenceCondition {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TurbulenceCondition ID in the query.
// Returns a *NotSingularError when more than one TurbulenceCondition ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TurbulenceConditionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{turbulencecondition.Label}
	default:
		err = &NotSingularError{turbulencecondition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TurbulenceConditions.
func (tcq *TurbulenceConditionQuery) All(ctx context.Context) ([]*TurbulenceCondition, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) AllX(ctx context.Context) []*TurbulenceCondition {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TurbulenceCondition IDs.
func (tcq *TurbulenceConditionQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := tcq.Select(turbulencecondition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TurbulenceConditionQuery) Count(ctx context.Context) (int, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TurbulenceConditionQuery) Exist(ctx context.Context) (bool, error) {
	if err := tcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TurbulenceConditionQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TurbulenceConditionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TurbulenceConditionQuery) Clone() *TurbulenceConditionQuery {
	if tcq == nil {
		return nil
	}
	return &TurbulenceConditionQuery{
		config:     tcq.config,
		limit:      tcq.limit,
		offset:     tcq.offset,
		order:      append([]OrderFunc{}, tcq.order...),
		predicates: append([]predicate.TurbulenceCondition{}, tcq.predicates...),
		// clone intermediate query.
		sql:    tcq.sql.Clone(),
		path:   tcq.path,
		unique: tcq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Intensity string `json:"intensity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TurbulenceCondition.Query().
//		GroupBy(turbulencecondition.FieldIntensity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tcq *TurbulenceConditionQuery) GroupBy(field string, fields ...string) *TurbulenceConditionGroupBy {
	grbuild := &TurbulenceConditionGroupBy{config: tcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tcq.sqlQuery(ctx), nil
	}
	grbuild.label = turbulencecondition.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Intensity string `json:"intensity,omitempty"`
//	}
//
//	client.TurbulenceCondition.Query().
//		Select(turbulencecondition.FieldIntensity).
//		Scan(ctx, &v)
func (tcq *TurbulenceConditionQuery) Select(fields ...string) *TurbulenceConditionSelect {
	tcq.fields = append(tcq.fields, fields...)
	selbuild := &TurbulenceConditionSelect{TurbulenceConditionQuery: tcq}
	selbuild.label = turbulencecondition.Label
	selbuild.flds, selbuild.scan = &tcq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a TurbulenceConditionSelect configured with the given aggregations.
func (tcq *TurbulenceConditionQuery) Aggregate(fns ...AggregateFunc) *TurbulenceConditionSelect {
	return tcq.Select().Aggregate(fns...)
}

func (tcq *TurbulenceConditionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tcq.fields {
		if !turbulencecondition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *TurbulenceConditionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TurbulenceCondition, error) {
	var (
		nodes   = []*TurbulenceCondition{}
		withFKs = tcq.withFKs
		_spec   = tcq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, turbulencecondition.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TurbulenceCondition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TurbulenceCondition{config: tcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tcq.modifiers) > 0 {
		_spec.Modifiers = tcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range tcq.loadTotal {
		if err := tcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tcq *TurbulenceConditionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	if len(tcq.modifiers) > 0 {
		_spec.Modifiers = tcq.modifiers
	}
	_spec.Node.Columns = tcq.fields
	if len(tcq.fields) > 0 {
		_spec.Unique = tcq.unique != nil && *tcq.unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TurbulenceConditionQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := tcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (tcq *TurbulenceConditionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   turbulencecondition.Table,
			Columns: turbulencecondition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: turbulencecondition.FieldID,
			},
		},
		From:   tcq.sql,
		Unique: true,
	}
	if unique := tcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, turbulencecondition.FieldID)
		for i := range fields {
			if fields[i] != turbulencecondition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *TurbulenceConditionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(turbulencecondition.Table)
	columns := tcq.fields
	if len(columns) == 0 {
		columns = turbulencecondition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcq.unique != nil && *tcq.unique {
		selector.Distinct()
	}
	for _, m := range tcq.modifiers {
		m(selector)
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tcq *TurbulenceConditionQuery) Modify(modifiers ...func(s *sql.Selector)) *TurbulenceConditionSelect {
	tcq.modifiers = append(tcq.modifiers, modifiers...)
	return tcq.Select()
}

// TurbulenceConditionGroupBy is the group-by builder for TurbulenceCondition entities.
type TurbulenceConditionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TurbulenceConditionGroupBy) Aggregate(fns ...AggregateFunc) *TurbulenceConditionGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tcgb *TurbulenceConditionGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tcgb.path(ctx)
	if err != nil {
		return err
	}
	tcgb.sql = query
	return tcgb.sqlScan(ctx, v)
}

func (tcgb *TurbulenceConditionGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tcgb.fields {
		if !turbulencecondition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tcgb *TurbulenceConditionGroupBy) sqlQuery() *sql.Selector {
	selector := tcgb.sql.Select()
	aggregation := make([]string, 0, len(tcgb.fns))
	for _, fn := range tcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tcgb.fields)+len(tcgb.fns))
		for _, f := range tcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tcgb.fields...)...)
}

// TurbulenceConditionSelect is the builder for selecting fields of TurbulenceCondition entities.
type TurbulenceConditionSelect struct {
	*TurbulenceConditionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tcs *TurbulenceConditionSelect) Aggregate(fns ...AggregateFunc) *TurbulenceConditionSelect {
	tcs.fns = append(tcs.fns, fns...)
	return tcs
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TurbulenceConditionSelect) Scan(ctx context.Context, v any) error {
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	tcs.sql = tcs.TurbulenceConditionQuery.sqlQuery(ctx)
	return tcs.sqlScan(ctx, v)
}

func (tcs *TurbulenceConditionSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(tcs.fns))
	for _, fn := range tcs.fns {
		aggregation = append(aggregation, fn(tcs.sql))
	}
	switch n := len(*tcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		tcs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		tcs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := tcs.sql.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tcs *TurbulenceConditionSelect) Modify(modifiers ...func(s *sql.Selector)) *TurbulenceConditionSelect {
	tcs.modifiers = append(tcs.modifiers, modifiers...)
	return tcs
}
