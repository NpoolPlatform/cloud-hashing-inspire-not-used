package schema

import "entgo.io/ent"

// AgencySetting holds the schema definition for the AgencySetting entity.
type AgencySetting struct {
	ent.Schema
}

// Fields of the AgencySetting.
func (AgencySetting) Fields() []ent.Field {
	return nil
}

// Edges of the AgencySetting.
func (AgencySetting) Edges() []ent.Edge {
	return nil
}
