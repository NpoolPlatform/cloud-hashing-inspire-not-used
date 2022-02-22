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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/purchaseinvitation"
	"github.com/google/uuid"
)

// PurchaseInvitationQuery is the builder for querying PurchaseInvitation entities.
type PurchaseInvitationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PurchaseInvitation
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PurchaseInvitationQuery builder.
func (piq *PurchaseInvitationQuery) Where(ps ...predicate.PurchaseInvitation) *PurchaseInvitationQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit adds a limit step to the query.
func (piq *PurchaseInvitationQuery) Limit(limit int) *PurchaseInvitationQuery {
	piq.limit = &limit
	return piq
}

// Offset adds an offset step to the query.
func (piq *PurchaseInvitationQuery) Offset(offset int) *PurchaseInvitationQuery {
	piq.offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *PurchaseInvitationQuery) Unique(unique bool) *PurchaseInvitationQuery {
	piq.unique = &unique
	return piq
}

// Order adds an order step to the query.
func (piq *PurchaseInvitationQuery) Order(o ...OrderFunc) *PurchaseInvitationQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// First returns the first PurchaseInvitation entity from the query.
// Returns a *NotFoundError when no PurchaseInvitation was found.
func (piq *PurchaseInvitationQuery) First(ctx context.Context) (*PurchaseInvitation, error) {
	nodes, err := piq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{purchaseinvitation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) FirstX(ctx context.Context) *PurchaseInvitation {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PurchaseInvitation ID from the query.
// Returns a *NotFoundError when no PurchaseInvitation ID was found.
func (piq *PurchaseInvitationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = piq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{purchaseinvitation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PurchaseInvitation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PurchaseInvitation entity is not found.
// Returns a *NotFoundError when no PurchaseInvitation entities are found.
func (piq *PurchaseInvitationQuery) Only(ctx context.Context) (*PurchaseInvitation, error) {
	nodes, err := piq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{purchaseinvitation.Label}
	default:
		return nil, &NotSingularError{purchaseinvitation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) OnlyX(ctx context.Context) *PurchaseInvitation {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PurchaseInvitation ID in the query.
// Returns a *NotSingularError when exactly one PurchaseInvitation ID is not found.
// Returns a *NotFoundError when no entities are found.
func (piq *PurchaseInvitationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = piq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = &NotSingularError{purchaseinvitation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PurchaseInvitations.
func (piq *PurchaseInvitationQuery) All(ctx context.Context) ([]*PurchaseInvitation, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return piq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) AllX(ctx context.Context) []*PurchaseInvitation {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PurchaseInvitation IDs.
func (piq *PurchaseInvitationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := piq.Select(purchaseinvitation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *PurchaseInvitationQuery) Count(ctx context.Context) (int, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return piq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *PurchaseInvitationQuery) Exist(ctx context.Context) (bool, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return piq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *PurchaseInvitationQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PurchaseInvitationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *PurchaseInvitationQuery) Clone() *PurchaseInvitationQuery {
	if piq == nil {
		return nil
	}
	return &PurchaseInvitationQuery{
		config:     piq.config,
		limit:      piq.limit,
		offset:     piq.offset,
		order:      append([]OrderFunc{}, piq.order...),
		predicates: append([]predicate.PurchaseInvitation{}, piq.predicates...),
		// clone intermediate query.
		sql:  piq.sql.Clone(),
		path: piq.path,
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
//	client.PurchaseInvitation.Query().
//		GroupBy(purchaseinvitation.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (piq *PurchaseInvitationQuery) GroupBy(field string, fields ...string) *PurchaseInvitationGroupBy {
	group := &PurchaseInvitationGroupBy{config: piq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return piq.sqlQuery(ctx), nil
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
//	client.PurchaseInvitation.Query().
//		Select(purchaseinvitation.FieldAppID).
//		Scan(ctx, &v)
//
func (piq *PurchaseInvitationQuery) Select(fields ...string) *PurchaseInvitationSelect {
	piq.fields = append(piq.fields, fields...)
	return &PurchaseInvitationSelect{PurchaseInvitationQuery: piq}
}

func (piq *PurchaseInvitationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range piq.fields {
		if !purchaseinvitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *PurchaseInvitationQuery) sqlAll(ctx context.Context) ([]*PurchaseInvitation, error) {
	var (
		nodes = []*PurchaseInvitation{}
		_spec = piq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PurchaseInvitation{config: piq.config}
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
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (piq *PurchaseInvitationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	_spec.Node.Columns = piq.fields
	if len(piq.fields) > 0 {
		_spec.Unique = piq.unique != nil && *piq.unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *PurchaseInvitationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := piq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (piq *PurchaseInvitationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   purchaseinvitation.Table,
			Columns: purchaseinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: purchaseinvitation.FieldID,
			},
		},
		From:   piq.sql,
		Unique: true,
	}
	if unique := piq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := piq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, purchaseinvitation.FieldID)
		for i := range fields {
			if fields[i] != purchaseinvitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *PurchaseInvitationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(purchaseinvitation.Table)
	columns := piq.fields
	if len(columns) == 0 {
		columns = purchaseinvitation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.unique != nil && *piq.unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PurchaseInvitationGroupBy is the group-by builder for PurchaseInvitation entities.
type PurchaseInvitationGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *PurchaseInvitationGroupBy) Aggregate(fns ...AggregateFunc) *PurchaseInvitationGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the group-by query and scans the result into the given value.
func (pigb *PurchaseInvitationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pigb.path(ctx)
	if err != nil {
		return err
	}
	pigb.sql = query
	return pigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) StringsX(ctx context.Context) []string {
	v, err := pigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) StringX(ctx context.Context) string {
	v, err := pigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) IntsX(ctx context.Context) []int {
	v, err := pigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) IntX(ctx context.Context) int {
	v, err := pigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pigb.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pigb *PurchaseInvitationGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pigb *PurchaseInvitationGroupBy) BoolX(ctx context.Context) bool {
	v, err := pigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pigb *PurchaseInvitationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pigb.fields {
		if !purchaseinvitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pigb *PurchaseInvitationGroupBy) sqlQuery() *sql.Selector {
	selector := pigb.sql.Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pigb.fields)+len(pigb.fns))
		for _, f := range pigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pigb.fields...)...)
}

// PurchaseInvitationSelect is the builder for selecting fields of PurchaseInvitation entities.
type PurchaseInvitationSelect struct {
	*PurchaseInvitationQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pis *PurchaseInvitationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	pis.sql = pis.PurchaseInvitationQuery.sqlQuery(ctx)
	return pis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) StringsX(ctx context.Context) []string {
	v, err := pis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) StringX(ctx context.Context) string {
	v, err := pis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) IntsX(ctx context.Context) []int {
	v, err := pis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) IntX(ctx context.Context) int {
	v, err := pis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) Float64X(ctx context.Context) float64 {
	v, err := pis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pis.fields) > 1 {
		return nil, errors.New("ent: PurchaseInvitationSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) BoolsX(ctx context.Context) []bool {
	v, err := pis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pis *PurchaseInvitationSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{purchaseinvitation.Label}
	default:
		err = fmt.Errorf("ent: PurchaseInvitationSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pis *PurchaseInvitationSelect) BoolX(ctx context.Context) bool {
	v, err := pis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pis *PurchaseInvitationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pis.sql.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
