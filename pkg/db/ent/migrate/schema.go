// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgencySettingsColumns holds the columns for the "agency_settings" table.
	AgencySettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "registration_reward_threshold", Type: field.TypeInt32},
		{Name: "registration_coupon_id", Type: field.TypeUUID},
		{Name: "kyc_reward_threshold", Type: field.TypeInt32},
		{Name: "kyc_coupon_id", Type: field.TypeUUID},
		{Name: "purchase_reward_percent", Type: field.TypeInt32},
		{Name: "purchase_reward_chain_levels", Type: field.TypeInt32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AgencySettingsTable holds the schema information for the "agency_settings" table.
	AgencySettingsTable = &schema.Table{
		Name:       "agency_settings",
		Columns:    AgencySettingsColumns,
		PrimaryKey: []*schema.Column{AgencySettingsColumns[0]},
	}
	// AppCouponSettingsColumns holds the columns for the "app_coupon_settings" table.
	AppCouponSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "domination_limit", Type: field.TypeInt},
		{Name: "total_limit", Type: field.TypeInt},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppCouponSettingsTable holds the schema information for the "app_coupon_settings" table.
	AppCouponSettingsTable = &schema.Table{
		Name:       "app_coupon_settings",
		Columns:    AppCouponSettingsColumns,
		PrimaryKey: []*schema.Column{AppCouponSettingsColumns[0]},
	}
	// CouponAllocatedsColumns holds the columns for the "coupon_allocateds" table.
	CouponAllocatedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "used", Type: field.TypeBool, Default: false},
		{Name: "coupon_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CouponAllocatedsTable holds the schema information for the "coupon_allocateds" table.
	CouponAllocatedsTable = &schema.Table{
		Name:       "coupon_allocateds",
		Columns:    CouponAllocatedsColumns,
		PrimaryKey: []*schema.Column{CouponAllocatedsColumns[0]},
	}
	// CouponPoolsColumns holds the columns for the "coupon_pools" table.
	CouponPoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "denomination", Type: field.TypeUint64},
		{Name: "ciculation", Type: field.TypeInt},
		{Name: "used", Type: field.TypeInt, Default: 0},
		{Name: "release_by_user_id", Type: field.TypeUUID},
		{Name: "start", Type: field.TypeUint32},
		{Name: "duration_days", Type: field.TypeInt},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Size: 512},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CouponPoolsTable holds the schema information for the "coupon_pools" table.
	CouponPoolsTable = &schema.Table{
		Name:       "coupon_pools",
		Columns:    CouponPoolsColumns,
		PrimaryKey: []*schema.Column{CouponPoolsColumns[0]},
	}
	// NewUserRewardSettingsColumns holds the columns for the "new_user_reward_settings" table.
	NewUserRewardSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "registration_coupon_id", Type: field.TypeUUID, Unique: true},
		{Name: "kyc_coupon_id", Type: field.TypeUUID, Unique: true},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// NewUserRewardSettingsTable holds the schema information for the "new_user_reward_settings" table.
	NewUserRewardSettingsTable = &schema.Table{
		Name:       "new_user_reward_settings",
		Columns:    NewUserRewardSettingsColumns,
		PrimaryKey: []*schema.Column{NewUserRewardSettingsColumns[0]},
	}
	// PurchaseInvitationsColumns holds the columns for the "purchase_invitations" table.
	PurchaseInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "invitation_code_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// PurchaseInvitationsTable holds the schema information for the "purchase_invitations" table.
	PurchaseInvitationsTable = &schema.Table{
		Name:       "purchase_invitations",
		Columns:    PurchaseInvitationsColumns,
		PrimaryKey: []*schema.Column{PurchaseInvitationsColumns[0]},
	}
	// RegistrationInvitationsColumns holds the columns for the "registration_invitations" table.
	RegistrationInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
		{Name: "inviter_id", Type: field.TypeUUID},
		{Name: "invitee_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
	}
	// RegistrationInvitationsTable holds the schema information for the "registration_invitations" table.
	RegistrationInvitationsTable = &schema.Table{
		Name:       "registration_invitations",
		Columns:    RegistrationInvitationsColumns,
		PrimaryKey: []*schema.Column{RegistrationInvitationsColumns[0]},
	}
	// UserInvitationCodesColumns holds the columns for the "user_invitation_codes" table.
	UserInvitationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "invitation_code", Type: field.TypeString, Unique: true},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserInvitationCodesTable holds the schema information for the "user_invitation_codes" table.
	UserInvitationCodesTable = &schema.Table{
		Name:       "user_invitation_codes",
		Columns:    UserInvitationCodesColumns,
		PrimaryKey: []*schema.Column{UserInvitationCodesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgencySettingsTable,
		AppCouponSettingsTable,
		CouponAllocatedsTable,
		CouponPoolsTable,
		NewUserRewardSettingsTable,
		PurchaseInvitationsTable,
		RegistrationInvitationsTable,
		UserInvitationCodesTable,
	}
)

func init() {
}
