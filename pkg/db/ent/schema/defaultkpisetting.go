package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// DefaultKpiSetting holds the schema definition for the DefaultKpiSetting entity.
type DefaultKpiSetting struct {
	ent.Schema
}

// Fields of the DefaultKpiSetting.
func (DefaultKpiSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Uint64("amount"),
		field.Int32("percent"),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
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

// Edges of the DefaultKpiSetting.
func (DefaultKpiSetting) Edges() []ent.Edge {
	return nil
}

// Indexs of the DefaultKpiSetting.
func (DefaultKpiSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "good_id").
			Unique(),
	}
}
