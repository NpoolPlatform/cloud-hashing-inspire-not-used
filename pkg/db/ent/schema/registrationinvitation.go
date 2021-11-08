package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// RegistrationInvitation holds the schema definition for the RegistrationInvitation entity.
type RegistrationInvitation struct {
	ent.Schema
}

// Fields of the RegistrationInvitation.
func (RegistrationInvitation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("inviter_id", uuid.UUID{}),
		field.UUID("invitee_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
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

// Edges of the RegistrationInvitation.
func (RegistrationInvitation) Edges() []ent.Edge {
	return nil
}

// Indexs of the RegistrationInvitation.
func (RegistrationInvitation) Indexs() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "inviter_id").
			Unique(),
		index.Fields("app_id", "invitee_id").
			Unique(),
	}
}
