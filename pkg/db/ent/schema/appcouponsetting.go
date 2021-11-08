package schema

import "entgo.io/ent"

// AppCouponSetting holds the schema definition for the AppCouponSetting entity.
type AppCouponSetting struct {
	ent.Schema
}

// Fields of the AppCouponSetting.
func (AppCouponSetting) Fields() []ent.Field {
	return nil
}

// Edges of the AppCouponSetting.
func (AppCouponSetting) Edges() []ent.Edge {
	return nil
}
