// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CouponPoolQuery is the builder for querying CouponPool entities.
type CouponPoolQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CouponPool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CouponPoolQuery builder.
func (cpq *CouponPoolQuery) Where(ps ...predicate.CouponPool) *CouponPoolQuery {
	cpq.predicates = append(cpq.predicates, ps...)
	return cpq
}

// Limit adds a limit step to the query.
func (cpq *CouponPoolQuery) Limit(limit int) *CouponPoolQuery {
	cpq.limit = &limit
	return cpq
}

// Offset adds an offset step to the query.
func (cpq *CouponPoolQuery) Offset(offset int) *CouponPoolQuery {
	cpq.offset = &offset
	return cpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cpq *CouponPoolQuery) Unique(unique bool) *CouponPoolQuery {
	cpq.unique = &unique
	return cpq
}

// Order adds an order step to the query.
func (cpq *CouponPoolQuery) Order(o ...OrderFunc) *CouponPoolQuery {
	cpq.order = append(cpq.order, o...)
	return cpq
}

// First returns the first CouponPool entity from the query.
// Returns a *NotFoundError when no CouponPool was found.
func (cpq *CouponPoolQuery) First(ctx context.Context) (*CouponPool, error) {
	nodes, err := cpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{couponpool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpq *CouponPoolQuery) FirstX(ctx context.Context) *CouponPool {
	node, err := cpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CouponPool ID from the query.
// Returns a *NotFoundError when no CouponPool ID was found.
func (cpq *CouponPoolQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{couponpool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cpq *CouponPoolQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CouponPool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CouponPool entity is not found.
// Returns a *NotFoundError when no CouponPool entities are found.
func (cpq *CouponPoolQuery) Only(ctx context.Context) (*CouponPool, error) {
	nodes, err := cpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{couponpool.Label}
	default:
		return nil, &NotSingularError{couponpool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpq *CouponPoolQuery) OnlyX(ctx context.Context) *CouponPool {
	node, err := cpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CouponPool ID in the query.
// Returns a *NotSingularError when exactly one CouponPool ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cpq *CouponPoolQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = &NotSingularError{couponpool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cpq *CouponPoolQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CouponPools.
func (cpq *CouponPoolQuery) All(ctx context.Context) ([]*CouponPool, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpq *CouponPoolQuery) AllX(ctx context.Context) []*CouponPool {
	nodes, err := cpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CouponPool IDs.
func (cpq *CouponPoolQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cpq.Select(couponpool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpq *CouponPoolQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpq *CouponPoolQuery) Count(ctx context.Context) (int, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpq *CouponPoolQuery) CountX(ctx context.Context) int {
	count, err := cpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpq *CouponPoolQuery) Exist(ctx context.Context) (bool, error) {
	if err := cpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpq *CouponPoolQuery) ExistX(ctx context.Context) bool {
	exist, err := cpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CouponPoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpq *CouponPoolQuery) Clone() *CouponPoolQuery {
	if cpq == nil {
		return nil
	}
	return &CouponPoolQuery{
		config:     cpq.config,
		limit:      cpq.limit,
		offset:     cpq.offset,
		order:      append([]OrderFunc{}, cpq.order...),
		predicates: append([]predicate.CouponPool{}, cpq.predicates...),
		// clone intermediate query.
		sql:  cpq.sql.Clone(),
		path: cpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Denomination uint64 `json:"denomination,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CouponPool.Query().
//		GroupBy(couponpool.FieldDenomination).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cpq *CouponPoolQuery) GroupBy(field string, fields ...string) *CouponPoolGroupBy {
	group := &CouponPoolGroupBy{config: cpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Denomination uint64 `json:"denomination,omitempty"`
//	}
//
//	client.CouponPool.Query().
//		Select(couponpool.FieldDenomination).
//		Scan(ctx, &v)
//
func (cpq *CouponPoolQuery) Select(fields ...string) *CouponPoolSelect {
	cpq.fields = append(cpq.fields, fields...)
	return &CouponPoolSelect{CouponPoolQuery: cpq}
}

func (cpq *CouponPoolQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cpq.fields {
		if !couponpool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cpq.path != nil {
		prev, err := cpq.path(ctx)
		if err != nil {
			return err
		}
		cpq.sql = prev
	}
	return nil
}

func (cpq *CouponPoolQuery) sqlAll(ctx context.Context) ([]*CouponPool, error) {
	var (
		nodes = []*CouponPool{}
		_spec = cpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CouponPool{config: cpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, cpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cpq *CouponPoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cpq.querySpec()
	return sqlgraph.CountNodes(ctx, cpq.driver, _spec)
}

func (cpq *CouponPoolQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cpq *CouponPoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponpool.Table,
			Columns: couponpool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponpool.FieldID,
			},
		},
		From:   cpq.sql,
		Unique: true,
	}
	if unique := cpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponpool.FieldID)
		for i := range fields {
			if fields[i] != couponpool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cpq *CouponPoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cpq.driver.Dialect())
	t1 := builder.Table(couponpool.Table)
	columns := cpq.fields
	if len(columns) == 0 {
		columns = couponpool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cpq.sql != nil {
		selector = cpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range cpq.predicates {
		p(selector)
	}
	for _, p := range cpq.order {
		p(selector)
	}
	if offset := cpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CouponPoolGroupBy is the group-by builder for CouponPool entities.
type CouponPoolGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpgb *CouponPoolGroupBy) Aggregate(fns ...AggregateFunc) *CouponPoolGroupBy {
	cpgb.fns = append(cpgb.fns, fns...)
	return cpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cpgb *CouponPoolGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cpgb.path(ctx)
	if err != nil {
		return err
	}
	cpgb.sql = query
	return cpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CouponPoolGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) StringsX(ctx context.Context) []string {
	v, err := cpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) StringX(ctx context.Context) string {
	v, err := cpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CouponPoolGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) IntsX(ctx context.Context) []int {
	v, err := cpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) IntX(ctx context.Context) int {
	v, err := cpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CouponPoolGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CouponPoolGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cpgb *CouponPoolGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cpgb *CouponPoolGroupBy) BoolX(ctx context.Context) bool {
	v, err := cpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cpgb *CouponPoolGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cpgb.fields {
		if !couponpool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpgb *CouponPoolGroupBy) sqlQuery() *sql.Selector {
	selector := cpgb.sql.Select()
	aggregation := make([]string, 0, len(cpgb.fns))
	for _, fn := range cpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cpgb.fields)+len(cpgb.fns))
		for _, f := range cpgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cpgb.fields...)...)
}

// CouponPoolSelect is the builder for selecting fields of CouponPool entities.
type CouponPoolSelect struct {
	*CouponPoolQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cps *CouponPoolSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cps.prepareQuery(ctx); err != nil {
		return err
	}
	cps.sql = cps.CouponPoolQuery.sqlQuery(ctx)
	return cps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cps *CouponPoolSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CouponPoolSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cps *CouponPoolSelect) StringsX(ctx context.Context) []string {
	v, err := cps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cps *CouponPoolSelect) StringX(ctx context.Context) string {
	v, err := cps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CouponPoolSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cps *CouponPoolSelect) IntsX(ctx context.Context) []int {
	v, err := cps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cps *CouponPoolSelect) IntX(ctx context.Context) int {
	v, err := cps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CouponPoolSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cps *CouponPoolSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cps *CouponPoolSelect) Float64X(ctx context.Context) float64 {
	v, err := cps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CouponPoolSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cps *CouponPoolSelect) BoolsX(ctx context.Context) []bool {
	v, err := cps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cps *CouponPoolSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{couponpool.Label}
	default:
		err = fmt.Errorf("ent: CouponPoolSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cps *CouponPoolSelect) BoolX(ctx context.Context) bool {
	v, err := cps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cps *CouponPoolSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cps.sql.Query()
	if err := cps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
