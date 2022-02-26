// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appuserinvitationsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserInvitationSettingUpdate is the builder for updating AppUserInvitationSetting entities.
type AppUserInvitationSettingUpdate struct {
	config
	hooks    []Hook
	mutation *AppUserInvitationSettingMutation
}

// Where appends a list predicates to the AppUserInvitationSettingUpdate builder.
func (auisu *AppUserInvitationSettingUpdate) Where(ps ...predicate.AppUserInvitationSetting) *AppUserInvitationSettingUpdate {
	auisu.mutation.Where(ps...)
	return auisu
}

// SetAppID sets the "app_id" field.
func (auisu *AppUserInvitationSettingUpdate) SetAppID(u uuid.UUID) *AppUserInvitationSettingUpdate {
	auisu.mutation.SetAppID(u)
	return auisu
}

// SetUserID sets the "user_id" field.
func (auisu *AppUserInvitationSettingUpdate) SetUserID(u uuid.UUID) *AppUserInvitationSettingUpdate {
	auisu.mutation.SetUserID(u)
	return auisu
}

// SetCount sets the "count" field.
func (auisu *AppUserInvitationSettingUpdate) SetCount(u uint32) *AppUserInvitationSettingUpdate {
	auisu.mutation.ResetCount()
	auisu.mutation.SetCount(u)
	return auisu
}

// AddCount adds u to the "count" field.
func (auisu *AppUserInvitationSettingUpdate) AddCount(u int32) *AppUserInvitationSettingUpdate {
	auisu.mutation.AddCount(u)
	return auisu
}

// SetDiscount sets the "discount" field.
func (auisu *AppUserInvitationSettingUpdate) SetDiscount(u uint32) *AppUserInvitationSettingUpdate {
	auisu.mutation.ResetDiscount()
	auisu.mutation.SetDiscount(u)
	return auisu
}

// AddDiscount adds u to the "discount" field.
func (auisu *AppUserInvitationSettingUpdate) AddDiscount(u int32) *AppUserInvitationSettingUpdate {
	auisu.mutation.AddDiscount(u)
	return auisu
}

// SetCreateAt sets the "create_at" field.
func (auisu *AppUserInvitationSettingUpdate) SetCreateAt(u uint32) *AppUserInvitationSettingUpdate {
	auisu.mutation.ResetCreateAt()
	auisu.mutation.SetCreateAt(u)
	return auisu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auisu *AppUserInvitationSettingUpdate) SetNillableCreateAt(u *uint32) *AppUserInvitationSettingUpdate {
	if u != nil {
		auisu.SetCreateAt(*u)
	}
	return auisu
}

// AddCreateAt adds u to the "create_at" field.
func (auisu *AppUserInvitationSettingUpdate) AddCreateAt(u int32) *AppUserInvitationSettingUpdate {
	auisu.mutation.AddCreateAt(u)
	return auisu
}

// SetUpdateAt sets the "update_at" field.
func (auisu *AppUserInvitationSettingUpdate) SetUpdateAt(u uint32) *AppUserInvitationSettingUpdate {
	auisu.mutation.ResetUpdateAt()
	auisu.mutation.SetUpdateAt(u)
	return auisu
}

// AddUpdateAt adds u to the "update_at" field.
func (auisu *AppUserInvitationSettingUpdate) AddUpdateAt(u int32) *AppUserInvitationSettingUpdate {
	auisu.mutation.AddUpdateAt(u)
	return auisu
}

// SetDeleteAt sets the "delete_at" field.
func (auisu *AppUserInvitationSettingUpdate) SetDeleteAt(u uint32) *AppUserInvitationSettingUpdate {
	auisu.mutation.ResetDeleteAt()
	auisu.mutation.SetDeleteAt(u)
	return auisu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auisu *AppUserInvitationSettingUpdate) SetNillableDeleteAt(u *uint32) *AppUserInvitationSettingUpdate {
	if u != nil {
		auisu.SetDeleteAt(*u)
	}
	return auisu
}

// AddDeleteAt adds u to the "delete_at" field.
func (auisu *AppUserInvitationSettingUpdate) AddDeleteAt(u int32) *AppUserInvitationSettingUpdate {
	auisu.mutation.AddDeleteAt(u)
	return auisu
}

// SetTitle sets the "title" field.
func (auisu *AppUserInvitationSettingUpdate) SetTitle(s string) *AppUserInvitationSettingUpdate {
	auisu.mutation.SetTitle(s)
	return auisu
}

// SetBadgeLarge sets the "badge_large" field.
func (auisu *AppUserInvitationSettingUpdate) SetBadgeLarge(s string) *AppUserInvitationSettingUpdate {
	auisu.mutation.SetBadgeLarge(s)
	return auisu
}

// SetBadgeSmall sets the "badge_small" field.
func (auisu *AppUserInvitationSettingUpdate) SetBadgeSmall(s string) *AppUserInvitationSettingUpdate {
	auisu.mutation.SetBadgeSmall(s)
	return auisu
}

// Mutation returns the AppUserInvitationSettingMutation object of the builder.
func (auisu *AppUserInvitationSettingUpdate) Mutation() *AppUserInvitationSettingMutation {
	return auisu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auisu *AppUserInvitationSettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	auisu.defaults()
	if len(auisu.hooks) == 0 {
		affected, err = auisu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserInvitationSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auisu.mutation = mutation
			affected, err = auisu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(auisu.hooks) - 1; i >= 0; i-- {
			if auisu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auisu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auisu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (auisu *AppUserInvitationSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := auisu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auisu *AppUserInvitationSettingUpdate) Exec(ctx context.Context) error {
	_, err := auisu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auisu *AppUserInvitationSettingUpdate) ExecX(ctx context.Context) {
	if err := auisu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auisu *AppUserInvitationSettingUpdate) defaults() {
	if _, ok := auisu.mutation.UpdateAt(); !ok {
		v := appuserinvitationsetting.UpdateDefaultUpdateAt()
		auisu.mutation.SetUpdateAt(v)
	}
}

func (auisu *AppUserInvitationSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserinvitationsetting.Table,
			Columns: appuserinvitationsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserinvitationsetting.FieldID,
			},
		},
	}
	if ps := auisu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auisu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserinvitationsetting.FieldAppID,
		})
	}
	if value, ok := auisu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserinvitationsetting.FieldUserID,
		})
	}
	if value, ok := auisu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCount,
		})
	}
	if value, ok := auisu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCount,
		})
	}
	if value, ok := auisu.mutation.Discount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDiscount,
		})
	}
	if value, ok := auisu.mutation.AddedDiscount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDiscount,
		})
	}
	if value, ok := auisu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCreateAt,
		})
	}
	if value, ok := auisu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCreateAt,
		})
	}
	if value, ok := auisu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldUpdateAt,
		})
	}
	if value, ok := auisu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldUpdateAt,
		})
	}
	if value, ok := auisu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDeleteAt,
		})
	}
	if value, ok := auisu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDeleteAt,
		})
	}
	if value, ok := auisu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserinvitationsetting.FieldTitle,
		})
	}
	if value, ok := auisu.mutation.BadgeLarge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserinvitationsetting.FieldBadgeLarge,
		})
	}
	if value, ok := auisu.mutation.BadgeSmall(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserinvitationsetting.FieldBadgeSmall,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, auisu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserinvitationsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserInvitationSettingUpdateOne is the builder for updating a single AppUserInvitationSetting entity.
type AppUserInvitationSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppUserInvitationSettingMutation
}

// SetAppID sets the "app_id" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetAppID(u uuid.UUID) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.SetAppID(u)
	return auisuo
}

// SetUserID sets the "user_id" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetUserID(u uuid.UUID) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.SetUserID(u)
	return auisuo
}

// SetCount sets the "count" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetCount(u uint32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.ResetCount()
	auisuo.mutation.SetCount(u)
	return auisuo
}

// AddCount adds u to the "count" field.
func (auisuo *AppUserInvitationSettingUpdateOne) AddCount(u int32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.AddCount(u)
	return auisuo
}

// SetDiscount sets the "discount" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetDiscount(u uint32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.ResetDiscount()
	auisuo.mutation.SetDiscount(u)
	return auisuo
}

// AddDiscount adds u to the "discount" field.
func (auisuo *AppUserInvitationSettingUpdateOne) AddDiscount(u int32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.AddDiscount(u)
	return auisuo
}

// SetCreateAt sets the "create_at" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetCreateAt(u uint32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.ResetCreateAt()
	auisuo.mutation.SetCreateAt(u)
	return auisuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auisuo *AppUserInvitationSettingUpdateOne) SetNillableCreateAt(u *uint32) *AppUserInvitationSettingUpdateOne {
	if u != nil {
		auisuo.SetCreateAt(*u)
	}
	return auisuo
}

// AddCreateAt adds u to the "create_at" field.
func (auisuo *AppUserInvitationSettingUpdateOne) AddCreateAt(u int32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.AddCreateAt(u)
	return auisuo
}

// SetUpdateAt sets the "update_at" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetUpdateAt(u uint32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.ResetUpdateAt()
	auisuo.mutation.SetUpdateAt(u)
	return auisuo
}

// AddUpdateAt adds u to the "update_at" field.
func (auisuo *AppUserInvitationSettingUpdateOne) AddUpdateAt(u int32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.AddUpdateAt(u)
	return auisuo
}

// SetDeleteAt sets the "delete_at" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetDeleteAt(u uint32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.ResetDeleteAt()
	auisuo.mutation.SetDeleteAt(u)
	return auisuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auisuo *AppUserInvitationSettingUpdateOne) SetNillableDeleteAt(u *uint32) *AppUserInvitationSettingUpdateOne {
	if u != nil {
		auisuo.SetDeleteAt(*u)
	}
	return auisuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (auisuo *AppUserInvitationSettingUpdateOne) AddDeleteAt(u int32) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.AddDeleteAt(u)
	return auisuo
}

// SetTitle sets the "title" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetTitle(s string) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.SetTitle(s)
	return auisuo
}

// SetBadgeLarge sets the "badge_large" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetBadgeLarge(s string) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.SetBadgeLarge(s)
	return auisuo
}

// SetBadgeSmall sets the "badge_small" field.
func (auisuo *AppUserInvitationSettingUpdateOne) SetBadgeSmall(s string) *AppUserInvitationSettingUpdateOne {
	auisuo.mutation.SetBadgeSmall(s)
	return auisuo
}

// Mutation returns the AppUserInvitationSettingMutation object of the builder.
func (auisuo *AppUserInvitationSettingUpdateOne) Mutation() *AppUserInvitationSettingMutation {
	return auisuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auisuo *AppUserInvitationSettingUpdateOne) Select(field string, fields ...string) *AppUserInvitationSettingUpdateOne {
	auisuo.fields = append([]string{field}, fields...)
	return auisuo
}

// Save executes the query and returns the updated AppUserInvitationSetting entity.
func (auisuo *AppUserInvitationSettingUpdateOne) Save(ctx context.Context) (*AppUserInvitationSetting, error) {
	var (
		err  error
		node *AppUserInvitationSetting
	)
	auisuo.defaults()
	if len(auisuo.hooks) == 0 {
		node, err = auisuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserInvitationSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auisuo.mutation = mutation
			node, err = auisuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auisuo.hooks) - 1; i >= 0; i-- {
			if auisuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auisuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auisuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auisuo *AppUserInvitationSettingUpdateOne) SaveX(ctx context.Context) *AppUserInvitationSetting {
	node, err := auisuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auisuo *AppUserInvitationSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := auisuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auisuo *AppUserInvitationSettingUpdateOne) ExecX(ctx context.Context) {
	if err := auisuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auisuo *AppUserInvitationSettingUpdateOne) defaults() {
	if _, ok := auisuo.mutation.UpdateAt(); !ok {
		v := appuserinvitationsetting.UpdateDefaultUpdateAt()
		auisuo.mutation.SetUpdateAt(v)
	}
}

func (auisuo *AppUserInvitationSettingUpdateOne) sqlSave(ctx context.Context) (_node *AppUserInvitationSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserinvitationsetting.Table,
			Columns: appuserinvitationsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserinvitationsetting.FieldID,
			},
		},
	}
	id, ok := auisuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppUserInvitationSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auisuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuserinvitationsetting.FieldID)
		for _, f := range fields {
			if !appuserinvitationsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appuserinvitationsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auisuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auisuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserinvitationsetting.FieldAppID,
		})
	}
	if value, ok := auisuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserinvitationsetting.FieldUserID,
		})
	}
	if value, ok := auisuo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCount,
		})
	}
	if value, ok := auisuo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCount,
		})
	}
	if value, ok := auisuo.mutation.Discount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDiscount,
		})
	}
	if value, ok := auisuo.mutation.AddedDiscount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDiscount,
		})
	}
	if value, ok := auisuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCreateAt,
		})
	}
	if value, ok := auisuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldCreateAt,
		})
	}
	if value, ok := auisuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldUpdateAt,
		})
	}
	if value, ok := auisuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldUpdateAt,
		})
	}
	if value, ok := auisuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDeleteAt,
		})
	}
	if value, ok := auisuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserinvitationsetting.FieldDeleteAt,
		})
	}
	if value, ok := auisuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserinvitationsetting.FieldTitle,
		})
	}
	if value, ok := auisuo.mutation.BadgeLarge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserinvitationsetting.FieldBadgeLarge,
		})
	}
	if value, ok := auisuo.mutation.BadgeSmall(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserinvitationsetting.FieldBadgeSmall,
		})
	}
	_node = &AppUserInvitationSetting{config: auisuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auisuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserinvitationsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
