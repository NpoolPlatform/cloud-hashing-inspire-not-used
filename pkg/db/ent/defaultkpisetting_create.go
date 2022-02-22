// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/defaultkpisetting"
	"github.com/google/uuid"
)

// DefaultKpiSettingCreate is the builder for creating a DefaultKpiSetting entity.
type DefaultKpiSettingCreate struct {
	config
	mutation *DefaultKpiSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAmount sets the "amount" field.
func (dksc *DefaultKpiSettingCreate) SetAmount(u uint64) *DefaultKpiSettingCreate {
	dksc.mutation.SetAmount(u)
	return dksc
}

// SetPercent sets the "percent" field.
func (dksc *DefaultKpiSettingCreate) SetPercent(i int32) *DefaultKpiSettingCreate {
	dksc.mutation.SetPercent(i)
	return dksc
}

// SetAppID sets the "app_id" field.
func (dksc *DefaultKpiSettingCreate) SetAppID(u uuid.UUID) *DefaultKpiSettingCreate {
	dksc.mutation.SetAppID(u)
	return dksc
}

// SetGoodID sets the "good_id" field.
func (dksc *DefaultKpiSettingCreate) SetGoodID(u uuid.UUID) *DefaultKpiSettingCreate {
	dksc.mutation.SetGoodID(u)
	return dksc
}

// SetCreateAt sets the "create_at" field.
func (dksc *DefaultKpiSettingCreate) SetCreateAt(u uint32) *DefaultKpiSettingCreate {
	dksc.mutation.SetCreateAt(u)
	return dksc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (dksc *DefaultKpiSettingCreate) SetNillableCreateAt(u *uint32) *DefaultKpiSettingCreate {
	if u != nil {
		dksc.SetCreateAt(*u)
	}
	return dksc
}

// SetUpdateAt sets the "update_at" field.
func (dksc *DefaultKpiSettingCreate) SetUpdateAt(u uint32) *DefaultKpiSettingCreate {
	dksc.mutation.SetUpdateAt(u)
	return dksc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (dksc *DefaultKpiSettingCreate) SetNillableUpdateAt(u *uint32) *DefaultKpiSettingCreate {
	if u != nil {
		dksc.SetUpdateAt(*u)
	}
	return dksc
}

// SetDeleteAt sets the "delete_at" field.
func (dksc *DefaultKpiSettingCreate) SetDeleteAt(u uint32) *DefaultKpiSettingCreate {
	dksc.mutation.SetDeleteAt(u)
	return dksc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (dksc *DefaultKpiSettingCreate) SetNillableDeleteAt(u *uint32) *DefaultKpiSettingCreate {
	if u != nil {
		dksc.SetDeleteAt(*u)
	}
	return dksc
}

// SetID sets the "id" field.
func (dksc *DefaultKpiSettingCreate) SetID(u uuid.UUID) *DefaultKpiSettingCreate {
	dksc.mutation.SetID(u)
	return dksc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dksc *DefaultKpiSettingCreate) SetNillableID(u *uuid.UUID) *DefaultKpiSettingCreate {
	if u != nil {
		dksc.SetID(*u)
	}
	return dksc
}

// Mutation returns the DefaultKpiSettingMutation object of the builder.
func (dksc *DefaultKpiSettingCreate) Mutation() *DefaultKpiSettingMutation {
	return dksc.mutation
}

// Save creates the DefaultKpiSetting in the database.
func (dksc *DefaultKpiSettingCreate) Save(ctx context.Context) (*DefaultKpiSetting, error) {
	var (
		err  error
		node *DefaultKpiSetting
	)
	dksc.defaults()
	if len(dksc.hooks) == 0 {
		if err = dksc.check(); err != nil {
			return nil, err
		}
		node, err = dksc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DefaultKpiSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dksc.check(); err != nil {
				return nil, err
			}
			dksc.mutation = mutation
			if node, err = dksc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dksc.hooks) - 1; i >= 0; i-- {
			if dksc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dksc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dksc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dksc *DefaultKpiSettingCreate) SaveX(ctx context.Context) *DefaultKpiSetting {
	v, err := dksc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dksc *DefaultKpiSettingCreate) Exec(ctx context.Context) error {
	_, err := dksc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dksc *DefaultKpiSettingCreate) ExecX(ctx context.Context) {
	if err := dksc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dksc *DefaultKpiSettingCreate) defaults() {
	if _, ok := dksc.mutation.CreateAt(); !ok {
		v := defaultkpisetting.DefaultCreateAt()
		dksc.mutation.SetCreateAt(v)
	}
	if _, ok := dksc.mutation.UpdateAt(); !ok {
		v := defaultkpisetting.DefaultUpdateAt()
		dksc.mutation.SetUpdateAt(v)
	}
	if _, ok := dksc.mutation.DeleteAt(); !ok {
		v := defaultkpisetting.DefaultDeleteAt()
		dksc.mutation.SetDeleteAt(v)
	}
	if _, ok := dksc.mutation.ID(); !ok {
		v := defaultkpisetting.DefaultID()
		dksc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dksc *DefaultKpiSettingCreate) check() error {
	if _, ok := dksc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "DefaultKpiSetting.amount"`)}
	}
	if _, ok := dksc.mutation.Percent(); !ok {
		return &ValidationError{Name: "percent", err: errors.New(`ent: missing required field "DefaultKpiSetting.percent"`)}
	}
	if _, ok := dksc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "DefaultKpiSetting.app_id"`)}
	}
	if _, ok := dksc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "DefaultKpiSetting.good_id"`)}
	}
	if _, ok := dksc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "DefaultKpiSetting.create_at"`)}
	}
	if _, ok := dksc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "DefaultKpiSetting.update_at"`)}
	}
	if _, ok := dksc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "DefaultKpiSetting.delete_at"`)}
	}
	return nil
}

func (dksc *DefaultKpiSettingCreate) sqlSave(ctx context.Context) (*DefaultKpiSetting, error) {
	_node, _spec := dksc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dksc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (dksc *DefaultKpiSettingCreate) createSpec() (*DefaultKpiSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &DefaultKpiSetting{config: dksc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: defaultkpisetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: defaultkpisetting.FieldID,
			},
		}
	)
	_spec.OnConflict = dksc.conflict
	if id, ok := dksc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dksc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: defaultkpisetting.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := dksc.mutation.Percent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: defaultkpisetting.FieldPercent,
		})
		_node.Percent = value
	}
	if value, ok := dksc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: defaultkpisetting.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := dksc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: defaultkpisetting.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := dksc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: defaultkpisetting.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := dksc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: defaultkpisetting.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := dksc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: defaultkpisetting.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DefaultKpiSetting.Create().
//		SetAmount(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DefaultKpiSettingUpsert) {
//			SetAmount(v+v).
//		}).
//		Exec(ctx)
//
func (dksc *DefaultKpiSettingCreate) OnConflict(opts ...sql.ConflictOption) *DefaultKpiSettingUpsertOne {
	dksc.conflict = opts
	return &DefaultKpiSettingUpsertOne{
		create: dksc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DefaultKpiSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dksc *DefaultKpiSettingCreate) OnConflictColumns(columns ...string) *DefaultKpiSettingUpsertOne {
	dksc.conflict = append(dksc.conflict, sql.ConflictColumns(columns...))
	return &DefaultKpiSettingUpsertOne{
		create: dksc,
	}
}

type (
	// DefaultKpiSettingUpsertOne is the builder for "upsert"-ing
	//  one DefaultKpiSetting node.
	DefaultKpiSettingUpsertOne struct {
		create *DefaultKpiSettingCreate
	}

	// DefaultKpiSettingUpsert is the "OnConflict" setter.
	DefaultKpiSettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetAmount sets the "amount" field.
func (u *DefaultKpiSettingUpsert) SetAmount(v uint64) *DefaultKpiSettingUpsert {
	u.Set(defaultkpisetting.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsert) UpdateAmount() *DefaultKpiSettingUpsert {
	u.SetExcluded(defaultkpisetting.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *DefaultKpiSettingUpsert) AddAmount(v uint64) *DefaultKpiSettingUpsert {
	u.Add(defaultkpisetting.FieldAmount, v)
	return u
}

// SetPercent sets the "percent" field.
func (u *DefaultKpiSettingUpsert) SetPercent(v int32) *DefaultKpiSettingUpsert {
	u.Set(defaultkpisetting.FieldPercent, v)
	return u
}

// UpdatePercent sets the "percent" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsert) UpdatePercent() *DefaultKpiSettingUpsert {
	u.SetExcluded(defaultkpisetting.FieldPercent)
	return u
}

// AddPercent adds v to the "percent" field.
func (u *DefaultKpiSettingUpsert) AddPercent(v int32) *DefaultKpiSettingUpsert {
	u.Add(defaultkpisetting.FieldPercent, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *DefaultKpiSettingUpsert) SetAppID(v uuid.UUID) *DefaultKpiSettingUpsert {
	u.Set(defaultkpisetting.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsert) UpdateAppID() *DefaultKpiSettingUpsert {
	u.SetExcluded(defaultkpisetting.FieldAppID)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *DefaultKpiSettingUpsert) SetGoodID(v uuid.UUID) *DefaultKpiSettingUpsert {
	u.Set(defaultkpisetting.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsert) UpdateGoodID() *DefaultKpiSettingUpsert {
	u.SetExcluded(defaultkpisetting.FieldGoodID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *DefaultKpiSettingUpsert) SetCreateAt(v uint32) *DefaultKpiSettingUpsert {
	u.Set(defaultkpisetting.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsert) UpdateCreateAt() *DefaultKpiSettingUpsert {
	u.SetExcluded(defaultkpisetting.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *DefaultKpiSettingUpsert) AddCreateAt(v uint32) *DefaultKpiSettingUpsert {
	u.Add(defaultkpisetting.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *DefaultKpiSettingUpsert) SetUpdateAt(v uint32) *DefaultKpiSettingUpsert {
	u.Set(defaultkpisetting.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsert) UpdateUpdateAt() *DefaultKpiSettingUpsert {
	u.SetExcluded(defaultkpisetting.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *DefaultKpiSettingUpsert) AddUpdateAt(v uint32) *DefaultKpiSettingUpsert {
	u.Add(defaultkpisetting.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *DefaultKpiSettingUpsert) SetDeleteAt(v uint32) *DefaultKpiSettingUpsert {
	u.Set(defaultkpisetting.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsert) UpdateDeleteAt() *DefaultKpiSettingUpsert {
	u.SetExcluded(defaultkpisetting.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *DefaultKpiSettingUpsert) AddDeleteAt(v uint32) *DefaultKpiSettingUpsert {
	u.Add(defaultkpisetting.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DefaultKpiSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(defaultkpisetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DefaultKpiSettingUpsertOne) UpdateNewValues() *DefaultKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(defaultkpisetting.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.DefaultKpiSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DefaultKpiSettingUpsertOne) Ignore() *DefaultKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DefaultKpiSettingUpsertOne) DoNothing() *DefaultKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DefaultKpiSettingCreate.OnConflict
// documentation for more info.
func (u *DefaultKpiSettingUpsertOne) Update(set func(*DefaultKpiSettingUpsert)) *DefaultKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DefaultKpiSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAmount sets the "amount" field.
func (u *DefaultKpiSettingUpsertOne) SetAmount(v uint64) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *DefaultKpiSettingUpsertOne) AddAmount(v uint64) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertOne) UpdateAmount() *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateAmount()
	})
}

// SetPercent sets the "percent" field.
func (u *DefaultKpiSettingUpsertOne) SetPercent(v int32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetPercent(v)
	})
}

// AddPercent adds v to the "percent" field.
func (u *DefaultKpiSettingUpsertOne) AddPercent(v int32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddPercent(v)
	})
}

// UpdatePercent sets the "percent" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertOne) UpdatePercent() *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdatePercent()
	})
}

// SetAppID sets the "app_id" field.
func (u *DefaultKpiSettingUpsertOne) SetAppID(v uuid.UUID) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertOne) UpdateAppID() *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *DefaultKpiSettingUpsertOne) SetGoodID(v uuid.UUID) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertOne) UpdateGoodID() *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateGoodID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *DefaultKpiSettingUpsertOne) SetCreateAt(v uint32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *DefaultKpiSettingUpsertOne) AddCreateAt(v uint32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertOne) UpdateCreateAt() *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *DefaultKpiSettingUpsertOne) SetUpdateAt(v uint32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *DefaultKpiSettingUpsertOne) AddUpdateAt(v uint32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertOne) UpdateUpdateAt() *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *DefaultKpiSettingUpsertOne) SetDeleteAt(v uint32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *DefaultKpiSettingUpsertOne) AddDeleteAt(v uint32) *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertOne) UpdateDeleteAt() *DefaultKpiSettingUpsertOne {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *DefaultKpiSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DefaultKpiSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DefaultKpiSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DefaultKpiSettingUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DefaultKpiSettingUpsertOne.ID is not supported by MySQL driver. Use DefaultKpiSettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DefaultKpiSettingUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DefaultKpiSettingCreateBulk is the builder for creating many DefaultKpiSetting entities in bulk.
type DefaultKpiSettingCreateBulk struct {
	config
	builders []*DefaultKpiSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the DefaultKpiSetting entities in the database.
func (dkscb *DefaultKpiSettingCreateBulk) Save(ctx context.Context) ([]*DefaultKpiSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dkscb.builders))
	nodes := make([]*DefaultKpiSetting, len(dkscb.builders))
	mutators := make([]Mutator, len(dkscb.builders))
	for i := range dkscb.builders {
		func(i int, root context.Context) {
			builder := dkscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DefaultKpiSettingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dkscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dkscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dkscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dkscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dkscb *DefaultKpiSettingCreateBulk) SaveX(ctx context.Context) []*DefaultKpiSetting {
	v, err := dkscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dkscb *DefaultKpiSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := dkscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dkscb *DefaultKpiSettingCreateBulk) ExecX(ctx context.Context) {
	if err := dkscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DefaultKpiSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DefaultKpiSettingUpsert) {
//			SetAmount(v+v).
//		}).
//		Exec(ctx)
//
func (dkscb *DefaultKpiSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *DefaultKpiSettingUpsertBulk {
	dkscb.conflict = opts
	return &DefaultKpiSettingUpsertBulk{
		create: dkscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DefaultKpiSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dkscb *DefaultKpiSettingCreateBulk) OnConflictColumns(columns ...string) *DefaultKpiSettingUpsertBulk {
	dkscb.conflict = append(dkscb.conflict, sql.ConflictColumns(columns...))
	return &DefaultKpiSettingUpsertBulk{
		create: dkscb,
	}
}

// DefaultKpiSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of DefaultKpiSetting nodes.
type DefaultKpiSettingUpsertBulk struct {
	create *DefaultKpiSettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DefaultKpiSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(defaultkpisetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DefaultKpiSettingUpsertBulk) UpdateNewValues() *DefaultKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(defaultkpisetting.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DefaultKpiSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DefaultKpiSettingUpsertBulk) Ignore() *DefaultKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DefaultKpiSettingUpsertBulk) DoNothing() *DefaultKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DefaultKpiSettingCreateBulk.OnConflict
// documentation for more info.
func (u *DefaultKpiSettingUpsertBulk) Update(set func(*DefaultKpiSettingUpsert)) *DefaultKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DefaultKpiSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAmount sets the "amount" field.
func (u *DefaultKpiSettingUpsertBulk) SetAmount(v uint64) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *DefaultKpiSettingUpsertBulk) AddAmount(v uint64) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertBulk) UpdateAmount() *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateAmount()
	})
}

// SetPercent sets the "percent" field.
func (u *DefaultKpiSettingUpsertBulk) SetPercent(v int32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetPercent(v)
	})
}

// AddPercent adds v to the "percent" field.
func (u *DefaultKpiSettingUpsertBulk) AddPercent(v int32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddPercent(v)
	})
}

// UpdatePercent sets the "percent" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertBulk) UpdatePercent() *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdatePercent()
	})
}

// SetAppID sets the "app_id" field.
func (u *DefaultKpiSettingUpsertBulk) SetAppID(v uuid.UUID) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertBulk) UpdateAppID() *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *DefaultKpiSettingUpsertBulk) SetGoodID(v uuid.UUID) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertBulk) UpdateGoodID() *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateGoodID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *DefaultKpiSettingUpsertBulk) SetCreateAt(v uint32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *DefaultKpiSettingUpsertBulk) AddCreateAt(v uint32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertBulk) UpdateCreateAt() *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *DefaultKpiSettingUpsertBulk) SetUpdateAt(v uint32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *DefaultKpiSettingUpsertBulk) AddUpdateAt(v uint32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertBulk) UpdateUpdateAt() *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *DefaultKpiSettingUpsertBulk) SetDeleteAt(v uint32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *DefaultKpiSettingUpsertBulk) AddDeleteAt(v uint32) *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *DefaultKpiSettingUpsertBulk) UpdateDeleteAt() *DefaultKpiSettingUpsertBulk {
	return u.Update(func(s *DefaultKpiSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *DefaultKpiSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DefaultKpiSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DefaultKpiSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DefaultKpiSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
