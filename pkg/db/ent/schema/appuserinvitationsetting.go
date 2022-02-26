package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppUserInvitationSetting holds the schema definition for the AppUserInvitationSetting entity.
type AppUserInvitationSetting struct {
	ent.Schema
}

// Fields of the AppUserInvitationSetting.
func (AppUserInvitationSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Uint32("count"),
		field.Uint32("discount"),
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
		field.String("title"),
		field.String("badge_large"),
		field.String("badge_small"),
	}
}

// Edges of the AppUserInvitationSetting.
func (AppUserInvitationSetting) Edges() []ent.Edge {
	return nil
}

// Indexes of the AppUserInvitationSetting.
func (AppUserInvitationSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id", "count").Unique(),
	}
}
