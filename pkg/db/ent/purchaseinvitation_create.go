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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/purchaseinvitation"
	"github.com/google/uuid"
)

// PurchaseInvitationCreate is the builder for creating a PurchaseInvitation entity.
type PurchaseInvitationCreate struct {
	config
	mutation *PurchaseInvitationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (pic *PurchaseInvitationCreate) SetAppID(u uuid.UUID) *PurchaseInvitationCreate {
	pic.mutation.SetAppID(u)
	return pic
}

// SetOrderID sets the "order_id" field.
func (pic *PurchaseInvitationCreate) SetOrderID(u uuid.UUID) *PurchaseInvitationCreate {
	pic.mutation.SetOrderID(u)
	return pic
}

// SetInvitationCodeID sets the "invitation_code_id" field.
func (pic *PurchaseInvitationCreate) SetInvitationCodeID(u uuid.UUID) *PurchaseInvitationCreate {
	pic.mutation.SetInvitationCodeID(u)
	return pic
}

// SetCreateAt sets the "create_at" field.
func (pic *PurchaseInvitationCreate) SetCreateAt(u uint32) *PurchaseInvitationCreate {
	pic.mutation.SetCreateAt(u)
	return pic
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pic *PurchaseInvitationCreate) SetNillableCreateAt(u *uint32) *PurchaseInvitationCreate {
	if u != nil {
		pic.SetCreateAt(*u)
	}
	return pic
}

// SetUpdateAt sets the "update_at" field.
func (pic *PurchaseInvitationCreate) SetUpdateAt(u uint32) *PurchaseInvitationCreate {
	pic.mutation.SetUpdateAt(u)
	return pic
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (pic *PurchaseInvitationCreate) SetNillableUpdateAt(u *uint32) *PurchaseInvitationCreate {
	if u != nil {
		pic.SetUpdateAt(*u)
	}
	return pic
}

// SetDeleteAt sets the "delete_at" field.
func (pic *PurchaseInvitationCreate) SetDeleteAt(u uint32) *PurchaseInvitationCreate {
	pic.mutation.SetDeleteAt(u)
	return pic
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (pic *PurchaseInvitationCreate) SetNillableDeleteAt(u *uint32) *PurchaseInvitationCreate {
	if u != nil {
		pic.SetDeleteAt(*u)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *PurchaseInvitationCreate) SetID(u uuid.UUID) *PurchaseInvitationCreate {
	pic.mutation.SetID(u)
	return pic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pic *PurchaseInvitationCreate) SetNillableID(u *uuid.UUID) *PurchaseInvitationCreate {
	if u != nil {
		pic.SetID(*u)
	}
	return pic
}

// Mutation returns the PurchaseInvitationMutation object of the builder.
func (pic *PurchaseInvitationCreate) Mutation() *PurchaseInvitationMutation {
	return pic.mutation
}

// Save creates the PurchaseInvitation in the database.
func (pic *PurchaseInvitationCreate) Save(ctx context.Context) (*PurchaseInvitation, error) {
	var (
		err  error
		node *PurchaseInvitation
	)
	pic.defaults()
	if len(pic.hooks) == 0 {
		if err = pic.check(); err != nil {
			return nil, err
		}
		node, err = pic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PurchaseInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pic.check(); err != nil {
				return nil, err
			}
			pic.mutation = mutation
			if node, err = pic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pic.hooks) - 1; i >= 0; i-- {
			if pic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PurchaseInvitationCreate) SaveX(ctx context.Context) *PurchaseInvitation {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PurchaseInvitationCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PurchaseInvitationCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *PurchaseInvitationCreate) defaults() {
	if _, ok := pic.mutation.CreateAt(); !ok {
		v := purchaseinvitation.DefaultCreateAt()
		pic.mutation.SetCreateAt(v)
	}
	if _, ok := pic.mutation.UpdateAt(); !ok {
		v := purchaseinvitation.DefaultUpdateAt()
		pic.mutation.SetUpdateAt(v)
	}
	if _, ok := pic.mutation.DeleteAt(); !ok {
		v := purchaseinvitation.DefaultDeleteAt()
		pic.mutation.SetDeleteAt(v)
	}
	if _, ok := pic.mutation.ID(); !ok {
		v := purchaseinvitation.DefaultID()
		pic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PurchaseInvitationCreate) check() error {
	if _, ok := pic.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "PurchaseInvitation.app_id"`)}
	}
	if _, ok := pic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "PurchaseInvitation.order_id"`)}
	}
	if _, ok := pic.mutation.InvitationCodeID(); !ok {
		return &ValidationError{Name: "invitation_code_id", err: errors.New(`ent: missing required field "PurchaseInvitation.invitation_code_id"`)}
	}
	if _, ok := pic.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "PurchaseInvitation.create_at"`)}
	}
	if _, ok := pic.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "PurchaseInvitation.update_at"`)}
	}
	if _, ok := pic.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "PurchaseInvitation.delete_at"`)}
	}
	return nil
}

func (pic *PurchaseInvitationCreate) sqlSave(ctx context.Context) (*PurchaseInvitation, error) {
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
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

func (pic *PurchaseInvitationCreate) createSpec() (*PurchaseInvitation, *sqlgraph.CreateSpec) {
	var (
		_node = &PurchaseInvitation{config: pic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: purchaseinvitation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: purchaseinvitation.FieldID,
			},
		}
	)
	_spec.OnConflict = pic.conflict
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pic.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := pic.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := pic.mutation.InvitationCodeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldInvitationCodeID,
		})
		_node.InvitationCodeID = value
	}
	if value, ok := pic.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := pic.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := pic.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PurchaseInvitation.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PurchaseInvitationUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (pic *PurchaseInvitationCreate) OnConflict(opts ...sql.ConflictOption) *PurchaseInvitationUpsertOne {
	pic.conflict = opts
	return &PurchaseInvitationUpsertOne{
		create: pic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PurchaseInvitation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pic *PurchaseInvitationCreate) OnConflictColumns(columns ...string) *PurchaseInvitationUpsertOne {
	pic.conflict = append(pic.conflict, sql.ConflictColumns(columns...))
	return &PurchaseInvitationUpsertOne{
		create: pic,
	}
}

type (
	// PurchaseInvitationUpsertOne is the builder for "upsert"-ing
	//  one PurchaseInvitation node.
	PurchaseInvitationUpsertOne struct {
		create *PurchaseInvitationCreate
	}

	// PurchaseInvitationUpsert is the "OnConflict" setter.
	PurchaseInvitationUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *PurchaseInvitationUpsert) SetAppID(v uuid.UUID) *PurchaseInvitationUpsert {
	u.Set(purchaseinvitation.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsert) UpdateAppID() *PurchaseInvitationUpsert {
	u.SetExcluded(purchaseinvitation.FieldAppID)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *PurchaseInvitationUpsert) SetOrderID(v uuid.UUID) *PurchaseInvitationUpsert {
	u.Set(purchaseinvitation.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsert) UpdateOrderID() *PurchaseInvitationUpsert {
	u.SetExcluded(purchaseinvitation.FieldOrderID)
	return u
}

// SetInvitationCodeID sets the "invitation_code_id" field.
func (u *PurchaseInvitationUpsert) SetInvitationCodeID(v uuid.UUID) *PurchaseInvitationUpsert {
	u.Set(purchaseinvitation.FieldInvitationCodeID, v)
	return u
}

// UpdateInvitationCodeID sets the "invitation_code_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsert) UpdateInvitationCodeID() *PurchaseInvitationUpsert {
	u.SetExcluded(purchaseinvitation.FieldInvitationCodeID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *PurchaseInvitationUpsert) SetCreateAt(v uint32) *PurchaseInvitationUpsert {
	u.Set(purchaseinvitation.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsert) UpdateCreateAt() *PurchaseInvitationUpsert {
	u.SetExcluded(purchaseinvitation.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *PurchaseInvitationUpsert) AddCreateAt(v uint32) *PurchaseInvitationUpsert {
	u.Add(purchaseinvitation.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *PurchaseInvitationUpsert) SetUpdateAt(v uint32) *PurchaseInvitationUpsert {
	u.Set(purchaseinvitation.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsert) UpdateUpdateAt() *PurchaseInvitationUpsert {
	u.SetExcluded(purchaseinvitation.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PurchaseInvitationUpsert) AddUpdateAt(v uint32) *PurchaseInvitationUpsert {
	u.Add(purchaseinvitation.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *PurchaseInvitationUpsert) SetDeleteAt(v uint32) *PurchaseInvitationUpsert {
	u.Set(purchaseinvitation.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsert) UpdateDeleteAt() *PurchaseInvitationUpsert {
	u.SetExcluded(purchaseinvitation.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PurchaseInvitationUpsert) AddDeleteAt(v uint32) *PurchaseInvitationUpsert {
	u.Add(purchaseinvitation.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PurchaseInvitation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(purchaseinvitation.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PurchaseInvitationUpsertOne) UpdateNewValues() *PurchaseInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(purchaseinvitation.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.PurchaseInvitation.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PurchaseInvitationUpsertOne) Ignore() *PurchaseInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PurchaseInvitationUpsertOne) DoNothing() *PurchaseInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PurchaseInvitationCreate.OnConflict
// documentation for more info.
func (u *PurchaseInvitationUpsertOne) Update(set func(*PurchaseInvitationUpsert)) *PurchaseInvitationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PurchaseInvitationUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *PurchaseInvitationUpsertOne) SetAppID(v uuid.UUID) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertOne) UpdateAppID() *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateAppID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *PurchaseInvitationUpsertOne) SetOrderID(v uuid.UUID) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertOne) UpdateOrderID() *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateOrderID()
	})
}

// SetInvitationCodeID sets the "invitation_code_id" field.
func (u *PurchaseInvitationUpsertOne) SetInvitationCodeID(v uuid.UUID) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetInvitationCodeID(v)
	})
}

// UpdateInvitationCodeID sets the "invitation_code_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertOne) UpdateInvitationCodeID() *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateInvitationCodeID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PurchaseInvitationUpsertOne) SetCreateAt(v uint32) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *PurchaseInvitationUpsertOne) AddCreateAt(v uint32) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertOne) UpdateCreateAt() *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PurchaseInvitationUpsertOne) SetUpdateAt(v uint32) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PurchaseInvitationUpsertOne) AddUpdateAt(v uint32) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertOne) UpdateUpdateAt() *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PurchaseInvitationUpsertOne) SetDeleteAt(v uint32) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PurchaseInvitationUpsertOne) AddDeleteAt(v uint32) *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertOne) UpdateDeleteAt() *PurchaseInvitationUpsertOne {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PurchaseInvitationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PurchaseInvitationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PurchaseInvitationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PurchaseInvitationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PurchaseInvitationUpsertOne.ID is not supported by MySQL driver. Use PurchaseInvitationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PurchaseInvitationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PurchaseInvitationCreateBulk is the builder for creating many PurchaseInvitation entities in bulk.
type PurchaseInvitationCreateBulk struct {
	config
	builders []*PurchaseInvitationCreate
	conflict []sql.ConflictOption
}

// Save creates the PurchaseInvitation entities in the database.
func (picb *PurchaseInvitationCreateBulk) Save(ctx context.Context) ([]*PurchaseInvitation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PurchaseInvitation, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PurchaseInvitationMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = picb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PurchaseInvitationCreateBulk) SaveX(ctx context.Context) []*PurchaseInvitation {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PurchaseInvitationCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PurchaseInvitationCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PurchaseInvitation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PurchaseInvitationUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (picb *PurchaseInvitationCreateBulk) OnConflict(opts ...sql.ConflictOption) *PurchaseInvitationUpsertBulk {
	picb.conflict = opts
	return &PurchaseInvitationUpsertBulk{
		create: picb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PurchaseInvitation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (picb *PurchaseInvitationCreateBulk) OnConflictColumns(columns ...string) *PurchaseInvitationUpsertBulk {
	picb.conflict = append(picb.conflict, sql.ConflictColumns(columns...))
	return &PurchaseInvitationUpsertBulk{
		create: picb,
	}
}

// PurchaseInvitationUpsertBulk is the builder for "upsert"-ing
// a bulk of PurchaseInvitation nodes.
type PurchaseInvitationUpsertBulk struct {
	create *PurchaseInvitationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PurchaseInvitation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(purchaseinvitation.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PurchaseInvitationUpsertBulk) UpdateNewValues() *PurchaseInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(purchaseinvitation.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PurchaseInvitation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PurchaseInvitationUpsertBulk) Ignore() *PurchaseInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PurchaseInvitationUpsertBulk) DoNothing() *PurchaseInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PurchaseInvitationCreateBulk.OnConflict
// documentation for more info.
func (u *PurchaseInvitationUpsertBulk) Update(set func(*PurchaseInvitationUpsert)) *PurchaseInvitationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PurchaseInvitationUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *PurchaseInvitationUpsertBulk) SetAppID(v uuid.UUID) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertBulk) UpdateAppID() *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateAppID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *PurchaseInvitationUpsertBulk) SetOrderID(v uuid.UUID) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertBulk) UpdateOrderID() *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateOrderID()
	})
}

// SetInvitationCodeID sets the "invitation_code_id" field.
func (u *PurchaseInvitationUpsertBulk) SetInvitationCodeID(v uuid.UUID) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetInvitationCodeID(v)
	})
}

// UpdateInvitationCodeID sets the "invitation_code_id" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertBulk) UpdateInvitationCodeID() *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateInvitationCodeID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *PurchaseInvitationUpsertBulk) SetCreateAt(v uint32) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *PurchaseInvitationUpsertBulk) AddCreateAt(v uint32) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertBulk) UpdateCreateAt() *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *PurchaseInvitationUpsertBulk) SetUpdateAt(v uint32) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *PurchaseInvitationUpsertBulk) AddUpdateAt(v uint32) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertBulk) UpdateUpdateAt() *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *PurchaseInvitationUpsertBulk) SetDeleteAt(v uint32) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *PurchaseInvitationUpsertBulk) AddDeleteAt(v uint32) *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *PurchaseInvitationUpsertBulk) UpdateDeleteAt() *PurchaseInvitationUpsertBulk {
	return u.Update(func(s *PurchaseInvitationUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *PurchaseInvitationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PurchaseInvitationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PurchaseInvitationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PurchaseInvitationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
