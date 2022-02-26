package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// CommissionCoinSetting holds the schema definition for the CommissionCoinSetting entity.
type CommissionCoinSetting struct {
	ent.Schema
}

// Fields of the CommissionCoinSetting.
func (CommissionCoinSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("coin_type_id", uuid.UUID{}).Unique(),
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

// Edges of the CommissionCoinSetting.
func (CommissionCoinSetting) Edges() []ent.Edge {
	return nil
}
