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
		field.Int("registration_reward_threshold"),
		field.Int("registration_reward_amount"),
		field.Int("kyc_reward_threshold"),
		field.Int("kyc_reward_amount"),
		field.Int("purchase_reward_percent"),
		field.Int("purchase_reward_chain_levels"),
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
