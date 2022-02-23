package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// EventCoupon holds the schema definition for the EventCoupon entity.
type EventCoupon struct {
	ent.Schema
}

// Fields of the EventCoupon.
func (EventCoupon) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("activity_id", uuid.UUID{}),
		field.String("type"),
		field.UUID("coupon_id", uuid.UUID{}),
		field.String("event").
			MaxLen(32),
		field.Uint32("count"),
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

// Edges of the EventCoupon.
func (EventCoupon) Edges() []ent.Edge {
	return nil
}

// Indexes of the EventCoupon.
func (EventCoupon) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "activity_id", "type", "coupon_id", "event").
			Unique(),
	}
}
