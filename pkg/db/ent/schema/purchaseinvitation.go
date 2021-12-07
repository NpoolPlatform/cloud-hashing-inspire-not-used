package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// PurchaseInvitation holds the schema definition for the PurchaseInvitation entity.
type PurchaseInvitation struct {
	ent.Schema
}

// Fields of the PurchaseInvitation.
func (PurchaseInvitation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("order_id", uuid.UUID{}),
		field.UUID("invitation_code_id", uuid.UUID{}),
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

// Edges of the PurchaseInvitation.
func (PurchaseInvitation) Edges() []ent.Edge {
	return nil
}

// Indexs of the PurchaseInvitation.
func (PurchaseInvitation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "order_id").
			Unique(),
	}
}
