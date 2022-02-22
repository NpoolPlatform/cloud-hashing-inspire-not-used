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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"
	"github.com/google/uuid"
)

// RegistrationInvitationQuery is the builder for querying RegistrationInvitation entities.
type RegistrationInvitationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RegistrationInvitation
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RegistrationInvitationQuery builder.
func (riq *RegistrationInvitationQuery) Where(ps ...predicate.RegistrationInvitation) *RegistrationInvitationQuery {
	riq.predicates = append(riq.predicates, ps...)
	return riq
}

// Limit adds a limit step to the query.
func (riq *RegistrationInvitationQuery) Limit(limit int) *RegistrationInvitationQuery {
	riq.limit = &limit
	return riq
}

// Offset adds an offset step to the query.
func (riq *RegistrationInvitationQuery) Offset(offset int) *RegistrationInvitationQuery {
	riq.offset = &offset
	return riq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (riq *RegistrationInvitationQuery) Unique(unique bool) *RegistrationInvitationQuery {
	riq.unique = &unique
	return riq
}

// Order adds an order step to the query.
func (riq *RegistrationInvitationQuery) Order(o ...OrderFunc) *RegistrationInvitationQuery {
	riq.order = append(riq.order, o...)
	return riq
}

// First returns the first RegistrationInvitation entity from the query.
// Returns a *NotFoundError when no RegistrationInvitation was found.
func (riq *RegistrationInvitationQuery) First(ctx context.Context) (*RegistrationInvitation, error) {
	nodes, err := riq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{registrationinvitation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) FirstX(ctx context.Context) *RegistrationInvitation {
	node, err := riq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RegistrationInvitation ID from the query.
// Returns a *NotFoundError when no RegistrationInvitation ID was found.
func (riq *RegistrationInvitationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = riq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{registrationinvitation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := riq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RegistrationInvitation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one RegistrationInvitation entity is not found.
// Returns a *NotFoundError when no RegistrationInvitation entities are found.
func (riq *RegistrationInvitationQuery) Only(ctx context.Context) (*RegistrationInvitation, error) {
	nodes, err := riq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{registrationinvitation.Label}
	default:
		return nil, &NotSingularError{registrationinvitation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) OnlyX(ctx context.Context) *RegistrationInvitation {
	node, err := riq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RegistrationInvitation ID in the query.
// Returns a *NotSingularError when exactly one RegistrationInvitation ID is not found.
// Returns a *NotFoundError when no entities are found.
func (riq *RegistrationInvitationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = riq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = &NotSingularError{registrationinvitation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := riq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RegistrationInvitations.
func (riq *RegistrationInvitationQuery) All(ctx context.Context) ([]*RegistrationInvitation, error) {
	if err := riq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return riq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) AllX(ctx context.Context) []*RegistrationInvitation {
	nodes, err := riq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RegistrationInvitation IDs.
func (riq *RegistrationInvitationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := riq.Select(registrationinvitation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := riq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (riq *RegistrationInvitationQuery) Count(ctx context.Context) (int, error) {
	if err := riq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return riq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) CountX(ctx context.Context) int {
	count, err := riq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (riq *RegistrationInvitationQuery) Exist(ctx context.Context) (bool, error) {
	if err := riq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return riq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (riq *RegistrationInvitationQuery) ExistX(ctx context.Context) bool {
	exist, err := riq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RegistrationInvitationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (riq *RegistrationInvitationQuery) Clone() *RegistrationInvitationQuery {
	if riq == nil {
		return nil
	}
	return &RegistrationInvitationQuery{
		config:     riq.config,
		limit:      riq.limit,
		offset:     riq.offset,
		order:      append([]OrderFunc{}, riq.order...),
		predicates: append([]predicate.RegistrationInvitation{}, riq.predicates...),
		// clone intermediate query.
		sql:  riq.sql.Clone(),
		path: riq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateAt uint32 `json:"create_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RegistrationInvitation.Query().
//		GroupBy(registrationinvitation.FieldCreateAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (riq *RegistrationInvitationQuery) GroupBy(field string, fields ...string) *RegistrationInvitationGroupBy {
	group := &RegistrationInvitationGroupBy{config: riq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := riq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return riq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateAt uint32 `json:"create_at,omitempty"`
//	}
//
//	client.RegistrationInvitation.Query().
//		Select(registrationinvitation.FieldCreateAt).
//		Scan(ctx, &v)
//
func (riq *RegistrationInvitationQuery) Select(fields ...string) *RegistrationInvitationSelect {
	riq.fields = append(riq.fields, fields...)
	return &RegistrationInvitationSelect{RegistrationInvitationQuery: riq}
}

func (riq *RegistrationInvitationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range riq.fields {
		if !registrationinvitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if riq.path != nil {
		prev, err := riq.path(ctx)
		if err != nil {
			return err
		}
		riq.sql = prev
	}
	return nil
}

func (riq *RegistrationInvitationQuery) sqlAll(ctx context.Context) ([]*RegistrationInvitation, error) {
	var (
		nodes = []*RegistrationInvitation{}
		_spec = riq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RegistrationInvitation{config: riq.config}
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
	if err := sqlgraph.QueryNodes(ctx, riq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (riq *RegistrationInvitationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := riq.querySpec()
	_spec.Node.Columns = riq.fields
	if len(riq.fields) > 0 {
		_spec.Unique = riq.unique != nil && *riq.unique
	}
	return sqlgraph.CountNodes(ctx, riq.driver, _spec)
}

func (riq *RegistrationInvitationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := riq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (riq *RegistrationInvitationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registrationinvitation.Table,
			Columns: registrationinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: registrationinvitation.FieldID,
			},
		},
		From:   riq.sql,
		Unique: true,
	}
	if unique := riq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := riq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, registrationinvitation.FieldID)
		for i := range fields {
			if fields[i] != registrationinvitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := riq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := riq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := riq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := riq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (riq *RegistrationInvitationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(riq.driver.Dialect())
	t1 := builder.Table(registrationinvitation.Table)
	columns := riq.fields
	if len(columns) == 0 {
		columns = registrationinvitation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if riq.sql != nil {
		selector = riq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if riq.unique != nil && *riq.unique {
		selector.Distinct()
	}
	for _, p := range riq.predicates {
		p(selector)
	}
	for _, p := range riq.order {
		p(selector)
	}
	if offset := riq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := riq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RegistrationInvitationGroupBy is the group-by builder for RegistrationInvitation entities.
type RegistrationInvitationGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rigb *RegistrationInvitationGroupBy) Aggregate(fns ...AggregateFunc) *RegistrationInvitationGroupBy {
	rigb.fns = append(rigb.fns, fns...)
	return rigb
}

// Scan applies the group-by query and scans the result into the given value.
func (rigb *RegistrationInvitationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rigb.path(ctx)
	if err != nil {
		return err
	}
	rigb.sql = query
	return rigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) StringsX(ctx context.Context) []string {
	v, err := rigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) StringX(ctx context.Context) string {
	v, err := rigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) IntsX(ctx context.Context) []int {
	v, err := rigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) IntX(ctx context.Context) int {
	v, err := rigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rigb.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rigb *RegistrationInvitationGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rigb *RegistrationInvitationGroupBy) BoolX(ctx context.Context) bool {
	v, err := rigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rigb *RegistrationInvitationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rigb.fields {
		if !registrationinvitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rigb *RegistrationInvitationGroupBy) sqlQuery() *sql.Selector {
	selector := rigb.sql.Select()
	aggregation := make([]string, 0, len(rigb.fns))
	for _, fn := range rigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rigb.fields)+len(rigb.fns))
		for _, f := range rigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rigb.fields...)...)
}

// RegistrationInvitationSelect is the builder for selecting fields of RegistrationInvitation entities.
type RegistrationInvitationSelect struct {
	*RegistrationInvitationQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ris *RegistrationInvitationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ris.prepareQuery(ctx); err != nil {
		return err
	}
	ris.sql = ris.RegistrationInvitationQuery.sqlQuery(ctx)
	return ris.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ris.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) StringsX(ctx context.Context) []string {
	v, err := ris.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ris.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) StringX(ctx context.Context) string {
	v, err := ris.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) IntsX(ctx context.Context) []int {
	v, err := ris.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ris.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) IntX(ctx context.Context) int {
	v, err := ris.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ris.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ris.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) Float64X(ctx context.Context) float64 {
	v, err := ris.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ris.fields) > 1 {
		return nil, errors.New("ent: RegistrationInvitationSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ris.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) BoolsX(ctx context.Context) []bool {
	v, err := ris.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ris *RegistrationInvitationSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ris.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registrationinvitation.Label}
	default:
		err = fmt.Errorf("ent: RegistrationInvitationSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ris *RegistrationInvitationSelect) BoolX(ctx context.Context) bool {
	v, err := ris.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ris *RegistrationInvitationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ris.sql.Query()
	if err := ris.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
