// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/eventcoupon"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// EventCouponUpdate is the builder for updating EventCoupon entities.
type EventCouponUpdate struct {
	config
	hooks    []Hook
	mutation *EventCouponMutation
}

// Where appends a list predicates to the EventCouponUpdate builder.
func (ecu *EventCouponUpdate) Where(ps ...predicate.EventCoupon) *EventCouponUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetAppID sets the "app_id" field.
func (ecu *EventCouponUpdate) SetAppID(u uuid.UUID) *EventCouponUpdate {
	ecu.mutation.SetAppID(u)
	return ecu
}

// SetActivityID sets the "activity_id" field.
func (ecu *EventCouponUpdate) SetActivityID(u uuid.UUID) *EventCouponUpdate {
	ecu.mutation.SetActivityID(u)
	return ecu
}

// SetEvent sets the "event" field.
func (ecu *EventCouponUpdate) SetEvent(s string) *EventCouponUpdate {
	ecu.mutation.SetEvent(s)
	return ecu
}

// SetCouponID sets the "coupon_id" field.
func (ecu *EventCouponUpdate) SetCouponID(u uuid.UUID) *EventCouponUpdate {
	ecu.mutation.SetCouponID(u)
	return ecu
}

// SetCreateAt sets the "create_at" field.
func (ecu *EventCouponUpdate) SetCreateAt(u uint32) *EventCouponUpdate {
	ecu.mutation.ResetCreateAt()
	ecu.mutation.SetCreateAt(u)
	return ecu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableCreateAt(u *uint32) *EventCouponUpdate {
	if u != nil {
		ecu.SetCreateAt(*u)
	}
	return ecu
}

// AddCreateAt adds u to the "create_at" field.
func (ecu *EventCouponUpdate) AddCreateAt(u int32) *EventCouponUpdate {
	ecu.mutation.AddCreateAt(u)
	return ecu
}

// SetUpdateAt sets the "update_at" field.
func (ecu *EventCouponUpdate) SetUpdateAt(u uint32) *EventCouponUpdate {
	ecu.mutation.ResetUpdateAt()
	ecu.mutation.SetUpdateAt(u)
	return ecu
}

// AddUpdateAt adds u to the "update_at" field.
func (ecu *EventCouponUpdate) AddUpdateAt(u int32) *EventCouponUpdate {
	ecu.mutation.AddUpdateAt(u)
	return ecu
}

// SetDeleteAt sets the "delete_at" field.
func (ecu *EventCouponUpdate) SetDeleteAt(u uint32) *EventCouponUpdate {
	ecu.mutation.ResetDeleteAt()
	ecu.mutation.SetDeleteAt(u)
	return ecu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ecu *EventCouponUpdate) SetNillableDeleteAt(u *uint32) *EventCouponUpdate {
	if u != nil {
		ecu.SetDeleteAt(*u)
	}
	return ecu
}

// AddDeleteAt adds u to the "delete_at" field.
func (ecu *EventCouponUpdate) AddDeleteAt(u int32) *EventCouponUpdate {
	ecu.mutation.AddDeleteAt(u)
	return ecu
}

// Mutation returns the EventCouponMutation object of the builder.
func (ecu *EventCouponUpdate) Mutation() *EventCouponMutation {
	return ecu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EventCouponUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ecu.defaults()
	if len(ecu.hooks) == 0 {
		affected, err = ecu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ecu.mutation = mutation
			affected, err = ecu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ecu.hooks) - 1; i >= 0; i-- {
			if ecu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ecu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ecu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EventCouponUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EventCouponUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EventCouponUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecu *EventCouponUpdate) defaults() {
	if _, ok := ecu.mutation.UpdateAt(); !ok {
		v := eventcoupon.UpdateDefaultUpdateAt()
		ecu.mutation.SetUpdateAt(v)
	}
}

func (ecu *EventCouponUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventcoupon.Table,
			Columns: eventcoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: eventcoupon.FieldID,
			},
		},
	}
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldAppID,
		})
	}
	if value, ok := ecu.mutation.ActivityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldActivityID,
		})
	}
	if value, ok := ecu.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventcoupon.FieldEvent,
		})
	}
	if value, ok := ecu.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldCouponID,
		})
	}
	if value, ok := ecu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreateAt,
		})
	}
	if value, ok := ecu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreateAt,
		})
	}
	if value, ok := ecu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdateAt,
		})
	}
	if value, ok := ecu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdateAt,
		})
	}
	if value, ok := ecu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeleteAt,
		})
	}
	if value, ok := ecu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventcoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventCouponUpdateOne is the builder for updating a single EventCoupon entity.
type EventCouponUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventCouponMutation
}

// SetAppID sets the "app_id" field.
func (ecuo *EventCouponUpdateOne) SetAppID(u uuid.UUID) *EventCouponUpdateOne {
	ecuo.mutation.SetAppID(u)
	return ecuo
}

// SetActivityID sets the "activity_id" field.
func (ecuo *EventCouponUpdateOne) SetActivityID(u uuid.UUID) *EventCouponUpdateOne {
	ecuo.mutation.SetActivityID(u)
	return ecuo
}

// SetEvent sets the "event" field.
func (ecuo *EventCouponUpdateOne) SetEvent(s string) *EventCouponUpdateOne {
	ecuo.mutation.SetEvent(s)
	return ecuo
}

// SetCouponID sets the "coupon_id" field.
func (ecuo *EventCouponUpdateOne) SetCouponID(u uuid.UUID) *EventCouponUpdateOne {
	ecuo.mutation.SetCouponID(u)
	return ecuo
}

// SetCreateAt sets the "create_at" field.
func (ecuo *EventCouponUpdateOne) SetCreateAt(u uint32) *EventCouponUpdateOne {
	ecuo.mutation.ResetCreateAt()
	ecuo.mutation.SetCreateAt(u)
	return ecuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableCreateAt(u *uint32) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetCreateAt(*u)
	}
	return ecuo
}

// AddCreateAt adds u to the "create_at" field.
func (ecuo *EventCouponUpdateOne) AddCreateAt(u int32) *EventCouponUpdateOne {
	ecuo.mutation.AddCreateAt(u)
	return ecuo
}

// SetUpdateAt sets the "update_at" field.
func (ecuo *EventCouponUpdateOne) SetUpdateAt(u uint32) *EventCouponUpdateOne {
	ecuo.mutation.ResetUpdateAt()
	ecuo.mutation.SetUpdateAt(u)
	return ecuo
}

// AddUpdateAt adds u to the "update_at" field.
func (ecuo *EventCouponUpdateOne) AddUpdateAt(u int32) *EventCouponUpdateOne {
	ecuo.mutation.AddUpdateAt(u)
	return ecuo
}

// SetDeleteAt sets the "delete_at" field.
func (ecuo *EventCouponUpdateOne) SetDeleteAt(u uint32) *EventCouponUpdateOne {
	ecuo.mutation.ResetDeleteAt()
	ecuo.mutation.SetDeleteAt(u)
	return ecuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ecuo *EventCouponUpdateOne) SetNillableDeleteAt(u *uint32) *EventCouponUpdateOne {
	if u != nil {
		ecuo.SetDeleteAt(*u)
	}
	return ecuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (ecuo *EventCouponUpdateOne) AddDeleteAt(u int32) *EventCouponUpdateOne {
	ecuo.mutation.AddDeleteAt(u)
	return ecuo
}

// Mutation returns the EventCouponMutation object of the builder.
func (ecuo *EventCouponUpdateOne) Mutation() *EventCouponMutation {
	return ecuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EventCouponUpdateOne) Select(field string, fields ...string) *EventCouponUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EventCoupon entity.
func (ecuo *EventCouponUpdateOne) Save(ctx context.Context) (*EventCoupon, error) {
	var (
		err  error
		node *EventCoupon
	)
	ecuo.defaults()
	if len(ecuo.hooks) == 0 {
		node, err = ecuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ecuo.mutation = mutation
			node, err = ecuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ecuo.hooks) - 1; i >= 0; i-- {
			if ecuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ecuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ecuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EventCouponUpdateOne) SaveX(ctx context.Context) *EventCoupon {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EventCouponUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EventCouponUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecuo *EventCouponUpdateOne) defaults() {
	if _, ok := ecuo.mutation.UpdateAt(); !ok {
		v := eventcoupon.UpdateDefaultUpdateAt()
		ecuo.mutation.SetUpdateAt(v)
	}
}

func (ecuo *EventCouponUpdateOne) sqlSave(ctx context.Context) (_node *EventCoupon, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventcoupon.Table,
			Columns: eventcoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: eventcoupon.FieldID,
			},
		},
	}
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EventCoupon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventcoupon.FieldID)
		for _, f := range fields {
			if !eventcoupon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eventcoupon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldAppID,
		})
	}
	if value, ok := ecuo.mutation.ActivityID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldActivityID,
		})
	}
	if value, ok := ecuo.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventcoupon.FieldEvent,
		})
	}
	if value, ok := ecuo.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: eventcoupon.FieldCouponID,
		})
	}
	if value, ok := ecuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreateAt,
		})
	}
	if value, ok := ecuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldCreateAt,
		})
	}
	if value, ok := ecuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdateAt,
		})
	}
	if value, ok := ecuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldUpdateAt,
		})
	}
	if value, ok := ecuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeleteAt,
		})
	}
	if value, ok := ecuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: eventcoupon.FieldDeleteAt,
		})
	}
	_node = &EventCoupon{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventcoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
