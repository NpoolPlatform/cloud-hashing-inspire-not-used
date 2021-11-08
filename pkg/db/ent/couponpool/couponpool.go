// Code generated by entc, DO NOT EDIT.

package couponpool

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the couponpool type in the database.
	Label = "coupon_pool"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDenomination holds the string denoting the denomination field in the database.
	FieldDenomination = "denomination"
	// FieldCiculation holds the string denoting the ciculation field in the database.
	FieldCiculation = "ciculation"
	// FieldUsed holds the string denoting the used field in the database.
	FieldUsed = "used"
	// FieldReleaseByUserID holds the string denoting the release_by_user_id field in the database.
	FieldReleaseByUserID = "release_by_user_id"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldDurationDays holds the string denoting the duration_days field in the database.
	FieldDurationDays = "duration_days"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the couponpool in the database.
	Table = "coupon_pools"
)

// Columns holds all SQL columns for couponpool fields.
var Columns = []string{
	FieldID,
	FieldDenomination,
	FieldCiculation,
	FieldUsed,
	FieldReleaseByUserID,
	FieldStart,
	FieldDurationDays,
	FieldAppID,
	FieldMessage,
	FieldName,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUsed holds the default value on creation for the "used" field.
	DefaultUsed int
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
