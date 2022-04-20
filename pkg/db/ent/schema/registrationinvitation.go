package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// RegistrationInvitation holds the schema definition for the RegistrationInvitation entity.
type RegistrationInvitation struct {
	ent.Schema
}

func (RegistrationInvitation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
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
	}
}

// Edges of the RegistrationInvitation.
func (RegistrationInvitation) Edges() []ent.Edge {
	return nil
}

// Indexes of the RegistrationInvitation.
func (RegistrationInvitation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "invitee_id").
			Unique(),
	}
}
