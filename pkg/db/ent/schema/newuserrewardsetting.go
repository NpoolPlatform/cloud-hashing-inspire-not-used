package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// NewUserRewardSetting holds the schema definition for the NewUserRewardSetting entity.
type NewUserRewardSetting struct {
	ent.Schema
}

// Fields of the NewUserRewardSetting.
func (NewUserRewardSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}).
			Unique(),
		field.UUID("registration_coupon_id", uuid.UUID{}).
			Unique(),
		field.UUID("kyc_coupon_id", uuid.UUID{}).
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

// Edges of the NewUserRewardSetting.
func (NewUserRewardSetting) Edges() []ent.Edge {
	return nil
}
