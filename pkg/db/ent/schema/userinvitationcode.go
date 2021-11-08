package schema

import "entgo.io/ent"

// UserInvitationCode holds the schema definition for the UserInvitationCode entity.
type UserInvitationCode struct {
	ent.Schema
}

// Fields of the UserInvitationCode.
func (UserInvitationCode) Fields() []ent.Field {
	return nil
}

// Edges of the UserInvitationCode.
func (UserInvitationCode) Edges() []ent.Edge {
	return nil
}
