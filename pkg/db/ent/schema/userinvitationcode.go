package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// UserInvitationCode holds the schema definition for the UserInvitationCode entity.
type UserInvitationCode struct {
	ent.Schema
}

// Fields of the UserInvitationCode.
func (UserInvitationCode) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("app_id", uuid.UUID{}),
		field.String("invitation_code").
			Unique(),
		field.Bool("confirmed").Default(false),
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

// Edges of the UserInvitationCode.
func (UserInvitationCode) Edges() []ent.Edge {
	return nil
}

// Indexes of the UserInvitationCode.
func (UserInvitationCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id").
			Unique(),
	}
}
