package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppInvitationSetting holds the schema definition for the AppInvitationSetting entity.
type AppInvitationSetting struct {
	ent.Schema
}

// Fields of the AppInvitationSetting.
func (AppInvitationSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("app_id", uuid.UUID{}),
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

// Edges of the AppInvitationSetting.
func (AppInvitationSetting) Edges() []ent.Edge {
	return nil
}

// Indexes of the AppInvitationSetting.
func (AppInvitationSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "count").Unique(),
	}
}
