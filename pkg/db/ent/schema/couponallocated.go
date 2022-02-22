package schema

import (
	"time"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// CouponAllocated holds the schema definition for the CouponAllocated entity.
type CouponAllocated struct {
	ent.Schema
}

// Fields of the CouponAllocated.
func (CouponAllocated) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Enum("type").
			Values(constant.CouponTypeDiscount, constant.CouponTypeCoupon),
		field.UUID("coupon_id", uuid.UUID{}),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the CouponAllocated.
func (CouponAllocated) Edges() []ent.Edge {
	return nil
}
