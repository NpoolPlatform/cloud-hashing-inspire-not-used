package schema

import "entgo.io/ent"

// CouponPool holds the schema definition for the CouponPool entity.
type CouponPool struct {
	ent.Schema
}

// Fields of the CouponPool.
func (CouponPool) Fields() []ent.Field {
	return nil
}

// Edges of the CouponPool.
func (CouponPool) Edges() []ent.Edge {
	return nil
}
