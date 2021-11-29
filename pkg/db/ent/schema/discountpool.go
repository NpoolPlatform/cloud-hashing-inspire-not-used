package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// DiscountPool holds the schema definition for the DiscountPool entity.
type DiscountPool struct {
	ent.Schema
}

// Fields of the DiscountPool.
func (DiscountPool) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.Uint64("value"),
		field.UUID("release_by_user_id", uuid.UUID{}),
		field.Uint32("start"),
		field.Int32("duration_days"),
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

// Edges of the DiscountPool.
func (DiscountPool) Edges() []ent.Edge {
	return nil
}
