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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/commissioncoinsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CommissionCoinSettingQuery is the builder for querying CommissionCoinSetting entities.
type CommissionCoinSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CommissionCoinSetting
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommissionCoinSettingQuery builder.
func (ccsq *CommissionCoinSettingQuery) Where(ps ...predicate.CommissionCoinSetting) *CommissionCoinSettingQuery {
	ccsq.predicates = append(ccsq.predicates, ps...)
	return ccsq
}

// Limit adds a limit step to the query.
func (ccsq *CommissionCoinSettingQuery) Limit(limit int) *CommissionCoinSettingQuery {
	ccsq.limit = &limit
	return ccsq
}

// Offset adds an offset step to the query.
func (ccsq *CommissionCoinSettingQuery) Offset(offset int) *CommissionCoinSettingQuery {
	ccsq.offset = &offset
	return ccsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccsq *CommissionCoinSettingQuery) Unique(unique bool) *CommissionCoinSettingQuery {
	ccsq.unique = &unique
	return ccsq
}

// Order adds an order step to the query.
func (ccsq *CommissionCoinSettingQuery) Order(o ...OrderFunc) *CommissionCoinSettingQuery {
	ccsq.order = append(ccsq.order, o...)
	return ccsq
}

// First returns the first CommissionCoinSetting entity from the query.
// Returns a *NotFoundError when no CommissionCoinSetting was found.
func (ccsq *CommissionCoinSettingQuery) First(ctx context.Context) (*CommissionCoinSetting, error) {
	nodes, err := ccsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{commissioncoinsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) FirstX(ctx context.Context) *CommissionCoinSetting {
	node, err := ccsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommissionCoinSetting ID from the query.
// Returns a *NotFoundError when no CommissionCoinSetting ID was found.
func (ccsq *CommissionCoinSettingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ccsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{commissioncoinsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ccsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommissionCoinSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CommissionCoinSetting entity is not found.
// Returns a *NotFoundError when no CommissionCoinSetting entities are found.
func (ccsq *CommissionCoinSettingQuery) Only(ctx context.Context) (*CommissionCoinSetting, error) {
	nodes, err := ccsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{commissioncoinsetting.Label}
	default:
		return nil, &NotSingularError{commissioncoinsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) OnlyX(ctx context.Context) *CommissionCoinSetting {
	node, err := ccsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommissionCoinSetting ID in the query.
// Returns a *NotSingularError when exactly one CommissionCoinSetting ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ccsq *CommissionCoinSettingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ccsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = &NotSingularError{commissioncoinsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ccsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommissionCoinSettings.
func (ccsq *CommissionCoinSettingQuery) All(ctx context.Context) ([]*CommissionCoinSetting, error) {
	if err := ccsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ccsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) AllX(ctx context.Context) []*CommissionCoinSetting {
	nodes, err := ccsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommissionCoinSetting IDs.
func (ccsq *CommissionCoinSettingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ccsq.Select(commissioncoinsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ccsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccsq *CommissionCoinSettingQuery) Count(ctx context.Context) (int, error) {
	if err := ccsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ccsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) CountX(ctx context.Context) int {
	count, err := ccsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccsq *CommissionCoinSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := ccsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ccsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ccsq *CommissionCoinSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := ccsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommissionCoinSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccsq *CommissionCoinSettingQuery) Clone() *CommissionCoinSettingQuery {
	if ccsq == nil {
		return nil
	}
	return &CommissionCoinSettingQuery{
		config:     ccsq.config,
		limit:      ccsq.limit,
		offset:     ccsq.offset,
		order:      append([]OrderFunc{}, ccsq.order...),
		predicates: append([]predicate.CommissionCoinSetting{}, ccsq.predicates...),
		// clone intermediate query.
		sql:  ccsq.sql.Clone(),
		path: ccsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CommissionCoinSetting.Query().
//		GroupBy(commissioncoinsetting.FieldCoinTypeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ccsq *CommissionCoinSettingQuery) GroupBy(field string, fields ...string) *CommissionCoinSettingGroupBy {
	group := &CommissionCoinSettingGroupBy{config: ccsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ccsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ccsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
//	}
//
//	client.CommissionCoinSetting.Query().
//		Select(commissioncoinsetting.FieldCoinTypeID).
//		Scan(ctx, &v)
//
func (ccsq *CommissionCoinSettingQuery) Select(fields ...string) *CommissionCoinSettingSelect {
	ccsq.fields = append(ccsq.fields, fields...)
	return &CommissionCoinSettingSelect{CommissionCoinSettingQuery: ccsq}
}

func (ccsq *CommissionCoinSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ccsq.fields {
		if !commissioncoinsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ccsq.path != nil {
		prev, err := ccsq.path(ctx)
		if err != nil {
			return err
		}
		ccsq.sql = prev
	}
	return nil
}

func (ccsq *CommissionCoinSettingQuery) sqlAll(ctx context.Context) ([]*CommissionCoinSetting, error) {
	var (
		nodes = []*CommissionCoinSetting{}
		_spec = ccsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CommissionCoinSetting{config: ccsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ccsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ccsq *CommissionCoinSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccsq.querySpec()
	_spec.Node.Columns = ccsq.fields
	if len(ccsq.fields) > 0 {
		_spec.Unique = ccsq.unique != nil && *ccsq.unique
	}
	return sqlgraph.CountNodes(ctx, ccsq.driver, _spec)
}

func (ccsq *CommissionCoinSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ccsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ccsq *CommissionCoinSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commissioncoinsetting.Table,
			Columns: commissioncoinsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: commissioncoinsetting.FieldID,
			},
		},
		From:   ccsq.sql,
		Unique: true,
	}
	if unique := ccsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ccsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commissioncoinsetting.FieldID)
		for i := range fields {
			if fields[i] != commissioncoinsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccsq *CommissionCoinSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccsq.driver.Dialect())
	t1 := builder.Table(commissioncoinsetting.Table)
	columns := ccsq.fields
	if len(columns) == 0 {
		columns = commissioncoinsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccsq.sql != nil {
		selector = ccsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccsq.unique != nil && *ccsq.unique {
		selector.Distinct()
	}
	for _, p := range ccsq.predicates {
		p(selector)
	}
	for _, p := range ccsq.order {
		p(selector)
	}
	if offset := ccsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommissionCoinSettingGroupBy is the group-by builder for CommissionCoinSetting entities.
type CommissionCoinSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccsgb *CommissionCoinSettingGroupBy) Aggregate(fns ...AggregateFunc) *CommissionCoinSettingGroupBy {
	ccsgb.fns = append(ccsgb.fns, fns...)
	return ccsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ccsgb *CommissionCoinSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ccsgb.path(ctx)
	if err != nil {
		return err
	}
	ccsgb.sql = query
	return ccsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ccsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ccsgb.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ccsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := ccsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ccsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) StringX(ctx context.Context) string {
	v, err := ccsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ccsgb.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ccsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := ccsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ccsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) IntX(ctx context.Context) int {
	v, err := ccsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ccsgb.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ccsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ccsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ccsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ccsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ccsgb.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ccsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ccsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccsgb *CommissionCoinSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ccsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ccsgb *CommissionCoinSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := ccsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ccsgb *CommissionCoinSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ccsgb.fields {
		if !commissioncoinsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ccsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ccsgb *CommissionCoinSettingGroupBy) sqlQuery() *sql.Selector {
	selector := ccsgb.sql.Select()
	aggregation := make([]string, 0, len(ccsgb.fns))
	for _, fn := range ccsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ccsgb.fields)+len(ccsgb.fns))
		for _, f := range ccsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ccsgb.fields...)...)
}

// CommissionCoinSettingSelect is the builder for selecting fields of CommissionCoinSetting entities.
type CommissionCoinSettingSelect struct {
	*CommissionCoinSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ccss *CommissionCoinSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ccss.prepareQuery(ctx); err != nil {
		return err
	}
	ccss.sql = ccss.CommissionCoinSettingQuery.sqlQuery(ctx)
	return ccss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ccss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ccss.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ccss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) StringsX(ctx context.Context) []string {
	v, err := ccss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ccss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) StringX(ctx context.Context) string {
	v, err := ccss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ccss.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ccss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) IntsX(ctx context.Context) []int {
	v, err := ccss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ccss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) IntX(ctx context.Context) int {
	v, err := ccss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ccss.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ccss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ccss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ccss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := ccss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ccss.fields) > 1 {
		return nil, errors.New("ent: CommissionCoinSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ccss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := ccss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ccss *CommissionCoinSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ccss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{commissioncoinsetting.Label}
	default:
		err = fmt.Errorf("ent: CommissionCoinSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ccss *CommissionCoinSettingSelect) BoolX(ctx context.Context) bool {
	v, err := ccss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ccss *CommissionCoinSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ccss.sql.Query()
	if err := ccss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
