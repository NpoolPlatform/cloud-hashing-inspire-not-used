package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// AgencySetting holds the schema definition for the AgencySetting entity.
type AgencySetting struct {
	ent.Schema
}

// Fields of the AgencySetting.
func (AgencySetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}).
			Unique(),
		field.Int32("registration_reward_threshold"),
		field.UUID("registration_coupon_id", uuid.UUID{}),
		field.Int32("kyc_reward_threshold"),
		field.UUID("kyc_coupon_id", uuid.UUID{}),
		field.Int32("total_purchase_reward_percent"),
		field.Int32("purchase_reward_chain_levels"),
		field.Int32("level_purchase_reward_percent"),
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

// Edges of the AgencySetting.
func (AgencySetting) Edges() []ent.Edge {
	return nil
}
