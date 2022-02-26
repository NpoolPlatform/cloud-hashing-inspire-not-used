// Code generated by entc, DO NOT EDIT.

package appuserpurchaseamountsetting

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// BadgeLarge applies equality check predicate on the "badge_large" field. It's identical to BadgeLargeEQ.
func BadgeLarge(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBadgeLarge), v))
	})
}

// BadgeSmall applies equality check predicate on the "badge_small" field. It's identical to BadgeSmallEQ.
func BadgeSmall(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBadgeSmall), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v uint64) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Percent applies equality check predicate on the "percent" field. It's identical to PercentEQ.
func Percent(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPercent), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// BadgeLargeEQ applies the EQ predicate on the "badge_large" field.
func BadgeLargeEQ(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeNEQ applies the NEQ predicate on the "badge_large" field.
func BadgeLargeNEQ(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeIn applies the In predicate on the "badge_large" field.
func BadgeLargeIn(vs ...string) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBadgeLarge), v...))
	})
}

// BadgeLargeNotIn applies the NotIn predicate on the "badge_large" field.
func BadgeLargeNotIn(vs ...string) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBadgeLarge), v...))
	})
}

// BadgeLargeGT applies the GT predicate on the "badge_large" field.
func BadgeLargeGT(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeGTE applies the GTE predicate on the "badge_large" field.
func BadgeLargeGTE(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeLT applies the LT predicate on the "badge_large" field.
func BadgeLargeLT(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeLTE applies the LTE predicate on the "badge_large" field.
func BadgeLargeLTE(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeContains applies the Contains predicate on the "badge_large" field.
func BadgeLargeContains(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeHasPrefix applies the HasPrefix predicate on the "badge_large" field.
func BadgeLargeHasPrefix(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeHasSuffix applies the HasSuffix predicate on the "badge_large" field.
func BadgeLargeHasSuffix(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeEqualFold applies the EqualFold predicate on the "badge_large" field.
func BadgeLargeEqualFold(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBadgeLarge), v))
	})
}

// BadgeLargeContainsFold applies the ContainsFold predicate on the "badge_large" field.
func BadgeLargeContainsFold(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBadgeLarge), v))
	})
}

// BadgeSmallEQ applies the EQ predicate on the "badge_small" field.
func BadgeSmallEQ(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallNEQ applies the NEQ predicate on the "badge_small" field.
func BadgeSmallNEQ(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallIn applies the In predicate on the "badge_small" field.
func BadgeSmallIn(vs ...string) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBadgeSmall), v...))
	})
}

// BadgeSmallNotIn applies the NotIn predicate on the "badge_small" field.
func BadgeSmallNotIn(vs ...string) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBadgeSmall), v...))
	})
}

// BadgeSmallGT applies the GT predicate on the "badge_small" field.
func BadgeSmallGT(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallGTE applies the GTE predicate on the "badge_small" field.
func BadgeSmallGTE(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallLT applies the LT predicate on the "badge_small" field.
func BadgeSmallLT(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallLTE applies the LTE predicate on the "badge_small" field.
func BadgeSmallLTE(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallContains applies the Contains predicate on the "badge_small" field.
func BadgeSmallContains(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallHasPrefix applies the HasPrefix predicate on the "badge_small" field.
func BadgeSmallHasPrefix(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallHasSuffix applies the HasSuffix predicate on the "badge_small" field.
func BadgeSmallHasSuffix(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallEqualFold applies the EqualFold predicate on the "badge_small" field.
func BadgeSmallEqualFold(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBadgeSmall), v))
	})
}

// BadgeSmallContainsFold applies the ContainsFold predicate on the "badge_small" field.
func BadgeSmallContainsFold(v string) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBadgeSmall), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v uint64) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v uint64) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...uint64) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...uint64) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v uint64) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v uint64) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v uint64) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v uint64) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// PercentEQ applies the EQ predicate on the "percent" field.
func PercentEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPercent), v))
	})
}

// PercentNEQ applies the NEQ predicate on the "percent" field.
func PercentNEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPercent), v))
	})
}

// PercentIn applies the In predicate on the "percent" field.
func PercentIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPercent), v...))
	})
}

// PercentNotIn applies the NotIn predicate on the "percent" field.
func PercentNotIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPercent), v...))
	})
}

// PercentGT applies the GT predicate on the "percent" field.
func PercentGT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPercent), v))
	})
}

// PercentGTE applies the GTE predicate on the "percent" field.
func PercentGTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPercent), v))
	})
}

// PercentLT applies the LT predicate on the "percent" field.
func PercentLT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPercent), v))
	})
}

// PercentLTE applies the LTE predicate on the "percent" field.
func PercentLTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPercent), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.AppUserPurchaseAmountSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppUserPurchaseAmountSetting) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppUserPurchaseAmountSetting) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppUserPurchaseAmountSetting) predicate.AppUserPurchaseAmountSetting {
	return predicate.AppUserPurchaseAmountSetting(func(s *sql.Selector) {
		p(s.Not())
	})
}
