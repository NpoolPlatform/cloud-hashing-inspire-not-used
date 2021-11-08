package schema

import "entgo.io/ent"

// PurchaseInvitation holds the schema definition for the PurchaseInvitation entity.
type PurchaseInvitation struct {
	ent.Schema
}

// Fields of the PurchaseInvitation.
func (PurchaseInvitation) Fields() []ent.Field {
	return nil
}

// Edges of the PurchaseInvitation.
func (PurchaseInvitation) Edges() []ent.Edge {
	return nil
}
