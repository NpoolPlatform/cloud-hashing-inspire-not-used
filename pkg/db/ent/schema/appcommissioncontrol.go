package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// AppCommissionSetting holds the schema definition for the AppCommissionSetting entity.
type AppCommissionSetting struct {
	ent.Schema
}

// Fields of the AppCommissionSetting.
func (AppCommissionSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("app_id", uuid.UUID{}).Unique(),
		field.String("type"),
		field.Uint32("level"),
		field.Bool("invitation_discount"),
		field.Bool("unique_setting"),
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

// Edges of the AppCommissionSetting.
func (AppCommissionSetting) Edges() []ent.Edge {
	return nil
}
