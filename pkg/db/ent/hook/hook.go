// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
)

// The ActivityFunc type is an adapter to allow the use of ordinary
// function as Activity mutator.
type ActivityFunc func(context.Context, *ent.ActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ActivityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ActivityMutation", m)
	}
	return f(ctx, mv)
}

// The AppCommissionSettingFunc type is an adapter to allow the use of ordinary
// function as AppCommissionSetting mutator.
type AppCommissionSettingFunc func(context.Context, *ent.AppCommissionSettingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppCommissionSettingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppCommissionSettingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppCommissionSettingMutation", m)
	}
	return f(ctx, mv)
}

// The AppCouponSettingFunc type is an adapter to allow the use of ordinary
// function as AppCouponSetting mutator.
type AppCouponSettingFunc func(context.Context, *ent.AppCouponSettingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppCouponSettingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppCouponSettingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppCouponSettingMutation", m)
	}
	return f(ctx, mv)
}

// The AppInvitationSettingFunc type is an adapter to allow the use of ordinary
// function as AppInvitationSetting mutator.
type AppInvitationSettingFunc func(context.Context, *ent.AppInvitationSettingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppInvitationSettingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppInvitationSettingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppInvitationSettingMutation", m)
	}
	return f(ctx, mv)
}

// The AppPurchaseAmountSettingFunc type is an adapter to allow the use of ordinary
// function as AppPurchaseAmountSetting mutator.
type AppPurchaseAmountSettingFunc func(context.Context, *ent.AppPurchaseAmountSettingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppPurchaseAmountSettingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppPurchaseAmountSettingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppPurchaseAmountSettingMutation", m)
	}
	return f(ctx, mv)
}

// The CommissionCoinSettingFunc type is an adapter to allow the use of ordinary
// function as CommissionCoinSetting mutator.
type CommissionCoinSettingFunc func(context.Context, *ent.CommissionCoinSettingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommissionCoinSettingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CommissionCoinSettingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommissionCoinSettingMutation", m)
	}
	return f(ctx, mv)
}

// The CouponAllocatedFunc type is an adapter to allow the use of ordinary
// function as CouponAllocated mutator.
type CouponAllocatedFunc func(context.Context, *ent.CouponAllocatedMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CouponAllocatedFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CouponAllocatedMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CouponAllocatedMutation", m)
	}
	return f(ctx, mv)
}

// The CouponPoolFunc type is an adapter to allow the use of ordinary
// function as CouponPool mutator.
type CouponPoolFunc func(context.Context, *ent.CouponPoolMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CouponPoolFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CouponPoolMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CouponPoolMutation", m)
	}
	return f(ctx, mv)
}

// The DefaultKpiSettingFunc type is an adapter to allow the use of ordinary
// function as DefaultKpiSetting mutator.
type DefaultKpiSettingFunc func(context.Context, *ent.DefaultKpiSettingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DefaultKpiSettingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DefaultKpiSettingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DefaultKpiSettingMutation", m)
	}
	return f(ctx, mv)
}

// The DiscountPoolFunc type is an adapter to allow the use of ordinary
// function as DiscountPool mutator.
type DiscountPoolFunc func(context.Context, *ent.DiscountPoolMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DiscountPoolFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DiscountPoolMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DiscountPoolMutation", m)
	}
	return f(ctx, mv)
}

// The EventCouponFunc type is an adapter to allow the use of ordinary
// function as EventCoupon mutator.
type EventCouponFunc func(context.Context, *ent.EventCouponMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EventCouponFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EventCouponMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EventCouponMutation", m)
	}
	return f(ctx, mv)
}

// The RegistrationInvitationFunc type is an adapter to allow the use of ordinary
// function as RegistrationInvitation mutator.
type RegistrationInvitationFunc func(context.Context, *ent.RegistrationInvitationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegistrationInvitationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RegistrationInvitationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RegistrationInvitationMutation", m)
	}
	return f(ctx, mv)
}

// The UserInvitationCodeFunc type is an adapter to allow the use of ordinary
// function as UserInvitationCode mutator.
type UserInvitationCodeFunc func(context.Context, *ent.UserInvitationCodeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserInvitationCodeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserInvitationCodeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserInvitationCodeMutation", m)
	}
	return f(ctx, mv)
}

// The UserKpiSettingFunc type is an adapter to allow the use of ordinary
// function as UserKpiSetting mutator.
type UserKpiSettingFunc func(context.Context, *ent.UserKpiSettingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserKpiSettingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserKpiSettingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserKpiSettingMutation", m)
	}
	return f(ctx, mv)
}

// The UserSpecialReductionFunc type is an adapter to allow the use of ordinary
// function as UserSpecialReduction mutator.
type UserSpecialReductionFunc func(context.Context, *ent.UserSpecialReductionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserSpecialReductionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserSpecialReductionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserSpecialReductionMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
