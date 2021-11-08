package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// UserKpiSetting holds the schema definition for the UserKpiSetting entity.
type UserKpiSetting struct {
	ent.Schema
}

// Fields of the UserKpiSetting.
func (UserKpiSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Uint64("amount"),
		field.Int32("percent"),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
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

// Edges of the UserKpiSetting.
func (UserKpiSetting) Edges() []ent.Edge {
	return nil
}

// Indexs of the UserKpiSetting.
func (UserKpiSetting) Indexs() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "good_id", "user_id").
			Unique(),
	}
}
