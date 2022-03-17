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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userkpisetting"
	"github.com/google/uuid"
)

// UserKpiSettingQuery is the builder for querying UserKpiSetting entities.
type UserKpiSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserKpiSetting
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserKpiSettingQuery builder.
func (uksq *UserKpiSettingQuery) Where(ps ...predicate.UserKpiSetting) *UserKpiSettingQuery {
	uksq.predicates = append(uksq.predicates, ps...)
	return uksq
}

// Limit adds a limit step to the query.
func (uksq *UserKpiSettingQuery) Limit(limit int) *UserKpiSettingQuery {
	uksq.limit = &limit
	return uksq
}

// Offset adds an offset step to the query.
func (uksq *UserKpiSettingQuery) Offset(offset int) *UserKpiSettingQuery {
	uksq.offset = &offset
	return uksq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uksq *UserKpiSettingQuery) Unique(unique bool) *UserKpiSettingQuery {
	uksq.unique = &unique
	return uksq
}

// Order adds an order step to the query.
func (uksq *UserKpiSettingQuery) Order(o ...OrderFunc) *UserKpiSettingQuery {
	uksq.order = append(uksq.order, o...)
	return uksq
}

// First returns the first UserKpiSetting entity from the query.
// Returns a *NotFoundError when no UserKpiSetting was found.
func (uksq *UserKpiSettingQuery) First(ctx context.Context) (*UserKpiSetting, error) {
	nodes, err := uksq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userkpisetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) FirstX(ctx context.Context) *UserKpiSetting {
	node, err := uksq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserKpiSetting ID from the query.
// Returns a *NotFoundError when no UserKpiSetting ID was found.
func (uksq *UserKpiSettingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uksq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userkpisetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uksq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserKpiSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserKpiSetting entity is found.
// Returns a *NotFoundError when no UserKpiSetting entities are found.
func (uksq *UserKpiSettingQuery) Only(ctx context.Context) (*UserKpiSetting, error) {
	nodes, err := uksq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userkpisetting.Label}
	default:
		return nil, &NotSingularError{userkpisetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) OnlyX(ctx context.Context) *UserKpiSetting {
	node, err := uksq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserKpiSetting ID in the query.
// Returns a *NotSingularError when more than one UserKpiSetting ID is found.
// Returns a *NotFoundError when no entities are found.
func (uksq *UserKpiSettingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uksq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = &NotSingularError{userkpisetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uksq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserKpiSettings.
func (uksq *UserKpiSettingQuery) All(ctx context.Context) ([]*UserKpiSetting, error) {
	if err := uksq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uksq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) AllX(ctx context.Context) []*UserKpiSetting {
	nodes, err := uksq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserKpiSetting IDs.
func (uksq *UserKpiSettingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := uksq.Select(userkpisetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uksq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uksq *UserKpiSettingQuery) Count(ctx context.Context) (int, error) {
	if err := uksq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uksq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) CountX(ctx context.Context) int {
	count, err := uksq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uksq *UserKpiSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := uksq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uksq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uksq *UserKpiSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := uksq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserKpiSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uksq *UserKpiSettingQuery) Clone() *UserKpiSettingQuery {
	if uksq == nil {
		return nil
	}
	return &UserKpiSettingQuery{
		config:     uksq.config,
		limit:      uksq.limit,
		offset:     uksq.offset,
		order:      append([]OrderFunc{}, uksq.order...),
		predicates: append([]predicate.UserKpiSetting{}, uksq.predicates...),
		// clone intermediate query.
		sql:    uksq.sql.Clone(),
		path:   uksq.path,
		unique: uksq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Amount uint64 `json:"amount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserKpiSetting.Query().
//		GroupBy(userkpisetting.FieldAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uksq *UserKpiSettingQuery) GroupBy(field string, fields ...string) *UserKpiSettingGroupBy {
	group := &UserKpiSettingGroupBy{config: uksq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uksq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Amount uint64 `json:"amount,omitempty"`
//	}
//
//	client.UserKpiSetting.Query().
//		Select(userkpisetting.FieldAmount).
//		Scan(ctx, &v)
//
func (uksq *UserKpiSettingQuery) Select(fields ...string) *UserKpiSettingSelect {
	uksq.fields = append(uksq.fields, fields...)
	return &UserKpiSettingSelect{UserKpiSettingQuery: uksq}
}

func (uksq *UserKpiSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uksq.fields {
		if !userkpisetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uksq.path != nil {
		prev, err := uksq.path(ctx)
		if err != nil {
			return err
		}
		uksq.sql = prev
	}
	return nil
}

func (uksq *UserKpiSettingQuery) sqlAll(ctx context.Context) ([]*UserKpiSetting, error) {
	var (
		nodes = []*UserKpiSetting{}
		_spec = uksq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserKpiSetting{config: uksq.config}
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
	if err := sqlgraph.QueryNodes(ctx, uksq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uksq *UserKpiSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uksq.querySpec()
	_spec.Node.Columns = uksq.fields
	if len(uksq.fields) > 0 {
		_spec.Unique = uksq.unique != nil && *uksq.unique
	}
	return sqlgraph.CountNodes(ctx, uksq.driver, _spec)
}

func (uksq *UserKpiSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uksq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uksq *UserKpiSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userkpisetting.Table,
			Columns: userkpisetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userkpisetting.FieldID,
			},
		},
		From:   uksq.sql,
		Unique: true,
	}
	if unique := uksq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uksq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userkpisetting.FieldID)
		for i := range fields {
			if fields[i] != userkpisetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uksq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uksq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uksq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uksq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uksq *UserKpiSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uksq.driver.Dialect())
	t1 := builder.Table(userkpisetting.Table)
	columns := uksq.fields
	if len(columns) == 0 {
		columns = userkpisetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uksq.sql != nil {
		selector = uksq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uksq.unique != nil && *uksq.unique {
		selector.Distinct()
	}
	for _, p := range uksq.predicates {
		p(selector)
	}
	for _, p := range uksq.order {
		p(selector)
	}
	if offset := uksq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uksq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserKpiSettingGroupBy is the group-by builder for UserKpiSetting entities.
type UserKpiSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uksgb *UserKpiSettingGroupBy) Aggregate(fns ...AggregateFunc) *UserKpiSettingGroupBy {
	uksgb.fns = append(uksgb.fns, fns...)
	return uksgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uksgb *UserKpiSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uksgb.path(ctx)
	if err != nil {
		return err
	}
	uksgb.sql = query
	return uksgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uksgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uksgb.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := uksgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uksgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) StringX(ctx context.Context) string {
	v, err := uksgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uksgb.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := uksgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uksgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) IntX(ctx context.Context) int {
	v, err := uksgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uksgb.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uksgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uksgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uksgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uksgb.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uksgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uksgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uksgb *UserKpiSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uksgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uksgb *UserKpiSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := uksgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uksgb *UserKpiSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uksgb.fields {
		if !userkpisetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uksgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uksgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uksgb *UserKpiSettingGroupBy) sqlQuery() *sql.Selector {
	selector := uksgb.sql.Select()
	aggregation := make([]string, 0, len(uksgb.fns))
	for _, fn := range uksgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uksgb.fields)+len(uksgb.fns))
		for _, f := range uksgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uksgb.fields...)...)
}

// UserKpiSettingSelect is the builder for selecting fields of UserKpiSetting entities.
type UserKpiSettingSelect struct {
	*UserKpiSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ukss *UserKpiSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ukss.prepareQuery(ctx); err != nil {
		return err
	}
	ukss.sql = ukss.UserKpiSettingQuery.sqlQuery(ctx)
	return ukss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ukss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ukss.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ukss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) StringsX(ctx context.Context) []string {
	v, err := ukss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ukss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) StringX(ctx context.Context) string {
	v, err := ukss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ukss.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ukss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) IntsX(ctx context.Context) []int {
	v, err := ukss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ukss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) IntX(ctx context.Context) int {
	v, err := ukss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ukss.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ukss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ukss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ukss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := ukss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ukss.fields) > 1 {
		return nil, errors.New("ent: UserKpiSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ukss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := ukss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ukss *UserKpiSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ukss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userkpisetting.Label}
	default:
		err = fmt.Errorf("ent: UserKpiSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ukss *UserKpiSettingSelect) BoolX(ctx context.Context) bool {
	v, err := ukss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ukss *UserKpiSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ukss.sql.Query()
	if err := ukss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
