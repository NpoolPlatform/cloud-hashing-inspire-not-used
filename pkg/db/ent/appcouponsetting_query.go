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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcouponsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppCouponSettingQuery is the builder for querying AppCouponSetting entities.
type AppCouponSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppCouponSetting
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppCouponSettingQuery builder.
func (acsq *AppCouponSettingQuery) Where(ps ...predicate.AppCouponSetting) *AppCouponSettingQuery {
	acsq.predicates = append(acsq.predicates, ps...)
	return acsq
}

// Limit adds a limit step to the query.
func (acsq *AppCouponSettingQuery) Limit(limit int) *AppCouponSettingQuery {
	acsq.limit = &limit
	return acsq
}

// Offset adds an offset step to the query.
func (acsq *AppCouponSettingQuery) Offset(offset int) *AppCouponSettingQuery {
	acsq.offset = &offset
	return acsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acsq *AppCouponSettingQuery) Unique(unique bool) *AppCouponSettingQuery {
	acsq.unique = &unique
	return acsq
}

// Order adds an order step to the query.
func (acsq *AppCouponSettingQuery) Order(o ...OrderFunc) *AppCouponSettingQuery {
	acsq.order = append(acsq.order, o...)
	return acsq
}

// First returns the first AppCouponSetting entity from the query.
// Returns a *NotFoundError when no AppCouponSetting was found.
func (acsq *AppCouponSettingQuery) First(ctx context.Context) (*AppCouponSetting, error) {
	nodes, err := acsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appcouponsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) FirstX(ctx context.Context) *AppCouponSetting {
	node, err := acsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppCouponSetting ID from the query.
// Returns a *NotFoundError when no AppCouponSetting ID was found.
func (acsq *AppCouponSettingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appcouponsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppCouponSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppCouponSetting entity is not found.
// Returns a *NotFoundError when no AppCouponSetting entities are found.
func (acsq *AppCouponSettingQuery) Only(ctx context.Context) (*AppCouponSetting, error) {
	nodes, err := acsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appcouponsetting.Label}
	default:
		return nil, &NotSingularError{appcouponsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) OnlyX(ctx context.Context) *AppCouponSetting {
	node, err := acsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppCouponSetting ID in the query.
// Returns a *NotSingularError when exactly one AppCouponSetting ID is not found.
// Returns a *NotFoundError when no entities are found.
func (acsq *AppCouponSettingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = &NotSingularError{appcouponsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppCouponSettings.
func (acsq *AppCouponSettingQuery) All(ctx context.Context) ([]*AppCouponSetting, error) {
	if err := acsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return acsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) AllX(ctx context.Context) []*AppCouponSetting {
	nodes, err := acsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppCouponSetting IDs.
func (acsq *AppCouponSettingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := acsq.Select(appcouponsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acsq *AppCouponSettingQuery) Count(ctx context.Context) (int, error) {
	if err := acsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return acsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) CountX(ctx context.Context) int {
	count, err := acsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acsq *AppCouponSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := acsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return acsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (acsq *AppCouponSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := acsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppCouponSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acsq *AppCouponSettingQuery) Clone() *AppCouponSettingQuery {
	if acsq == nil {
		return nil
	}
	return &AppCouponSettingQuery{
		config:     acsq.config,
		limit:      acsq.limit,
		offset:     acsq.offset,
		order:      append([]OrderFunc{}, acsq.order...),
		predicates: append([]predicate.AppCouponSetting{}, acsq.predicates...),
		// clone intermediate query.
		sql:  acsq.sql.Clone(),
		path: acsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppCouponSetting.Query().
//		GroupBy(appcouponsetting.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (acsq *AppCouponSettingQuery) GroupBy(field string, fields ...string) *AppCouponSettingGroupBy {
	group := &AppCouponSettingGroupBy{config: acsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := acsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return acsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.AppCouponSetting.Query().
//		Select(appcouponsetting.FieldAppID).
//		Scan(ctx, &v)
//
func (acsq *AppCouponSettingQuery) Select(fields ...string) *AppCouponSettingSelect {
	acsq.fields = append(acsq.fields, fields...)
	return &AppCouponSettingSelect{AppCouponSettingQuery: acsq}
}

func (acsq *AppCouponSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range acsq.fields {
		if !appcouponsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acsq.path != nil {
		prev, err := acsq.path(ctx)
		if err != nil {
			return err
		}
		acsq.sql = prev
	}
	return nil
}

func (acsq *AppCouponSettingQuery) sqlAll(ctx context.Context) ([]*AppCouponSetting, error) {
	var (
		nodes = []*AppCouponSetting{}
		_spec = acsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppCouponSetting{config: acsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, acsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (acsq *AppCouponSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acsq.querySpec()
	return sqlgraph.CountNodes(ctx, acsq.driver, _spec)
}

func (acsq *AppCouponSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := acsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (acsq *AppCouponSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcouponsetting.Table,
			Columns: appcouponsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcouponsetting.FieldID,
			},
		},
		From:   acsq.sql,
		Unique: true,
	}
	if unique := acsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := acsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcouponsetting.FieldID)
		for i := range fields {
			if fields[i] != appcouponsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acsq *AppCouponSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acsq.driver.Dialect())
	t1 := builder.Table(appcouponsetting.Table)
	columns := acsq.fields
	if len(columns) == 0 {
		columns = appcouponsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acsq.sql != nil {
		selector = acsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range acsq.predicates {
		p(selector)
	}
	for _, p := range acsq.order {
		p(selector)
	}
	if offset := acsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppCouponSettingGroupBy is the group-by builder for AppCouponSetting entities.
type AppCouponSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acsgb *AppCouponSettingGroupBy) Aggregate(fns ...AggregateFunc) *AppCouponSettingGroupBy {
	acsgb.fns = append(acsgb.fns, fns...)
	return acsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (acsgb *AppCouponSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := acsgb.path(ctx)
	if err != nil {
		return err
	}
	acsgb.sql = query
	return acsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := acsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := acsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) StringX(ctx context.Context) string {
	v, err := acsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := acsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) IntX(ctx context.Context) int {
	v, err := acsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := acsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := acsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(acsgb.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := acsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := acsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acsgb *AppCouponSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acsgb *AppCouponSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := acsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acsgb *AppCouponSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range acsgb.fields {
		if !appcouponsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := acsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (acsgb *AppCouponSettingGroupBy) sqlQuery() *sql.Selector {
	selector := acsgb.sql.Select()
	aggregation := make([]string, 0, len(acsgb.fns))
	for _, fn := range acsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(acsgb.fields)+len(acsgb.fns))
		for _, f := range acsgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(acsgb.fields...)...)
}

// AppCouponSettingSelect is the builder for selecting fields of AppCouponSetting entities.
type AppCouponSettingSelect struct {
	*AppCouponSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (acss *AppCouponSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := acss.prepareQuery(ctx); err != nil {
		return err
	}
	acss.sql = acss.AppCouponSettingQuery.sqlQuery(ctx)
	return acss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acss *AppCouponSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := acss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acss *AppCouponSettingSelect) StringsX(ctx context.Context) []string {
	v, err := acss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acss *AppCouponSettingSelect) StringX(ctx context.Context) string {
	v, err := acss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acss *AppCouponSettingSelect) IntsX(ctx context.Context) []int {
	v, err := acss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acss *AppCouponSettingSelect) IntX(ctx context.Context) int {
	v, err := acss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acss *AppCouponSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := acss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acss *AppCouponSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := acss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(acss.fields) > 1 {
		return nil, errors.New("ent: AppCouponSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := acss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acss *AppCouponSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := acss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (acss *AppCouponSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcouponsetting.Label}
	default:
		err = fmt.Errorf("ent: AppCouponSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acss *AppCouponSettingSelect) BoolX(ctx context.Context) bool {
	v, err := acss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acss *AppCouponSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := acss.sql.Query()
	if err := acss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}