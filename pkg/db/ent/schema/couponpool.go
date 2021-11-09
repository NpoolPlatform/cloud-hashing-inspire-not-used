package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// CouponPool holds the schema definition for the CouponPool entity.
type CouponPool struct {
	ent.Schema
}

// Fields of the CouponPool.
func (CouponPool) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Uint64("denomination"),
		field.Int32("circulation"),
		field.UUID("release_by_user_id", uuid.UUID{}),
		field.Uint32("start"),
		field.Int32("duration_days"),
		field.UUID("app_id", uuid.UUID{}),
		field.String("message").
			MaxLen(512),
		field.String("name").
			MaxLen(64).
			Unique(),
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

// Edges of the CouponPool.
func (CouponPool) Edges() []ent.Edge {
	return nil
}
