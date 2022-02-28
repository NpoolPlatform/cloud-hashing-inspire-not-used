// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActivitiesColumns holds the columns for the "activities" table.
	ActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "created_by", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "start", Type: field.TypeUint32},
		{Name: "end", Type: field.TypeUint32},
		{Name: "system_activity", Type: field.TypeBool},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// ActivitiesTable holds the schema information for the "activities" table.
	ActivitiesTable = &schema.Table{
		Name:       "activities",
		Columns:    ActivitiesColumns,
		PrimaryKey: []*schema.Column{ActivitiesColumns[0]},
	}
	// AppCommissionSettingsColumns holds the columns for the "app_commission_settings" table.
	AppCommissionSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "level", Type: field.TypeUint32},
		{Name: "invitation_discount", Type: field.TypeBool},
		{Name: "unique_setting", Type: field.TypeBool},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppCommissionSettingsTable holds the schema information for the "app_commission_settings" table.
	AppCommissionSettingsTable = &schema.Table{
		Name:       "app_commission_settings",
		Columns:    AppCommissionSettingsColumns,
		PrimaryKey: []*schema.Column{AppCommissionSettingsColumns[0]},
	}
	// AppCouponSettingsColumns holds the columns for the "app_coupon_settings" table.
	AppCouponSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "domination_limit", Type: field.TypeUint64},
		{Name: "total_limit", Type: field.TypeInt32},
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
	// AppInvitationSettingsColumns holds the columns for the "app_invitation_settings" table.
	AppInvitationSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "count", Type: field.TypeUint32},
		{Name: "discount", Type: field.TypeUint32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
		{Name: "title", Type: field.TypeString},
		{Name: "badge_large", Type: field.TypeString},
		{Name: "badge_small", Type: field.TypeString},
	}
	// AppInvitationSettingsTable holds the schema information for the "app_invitation_settings" table.
	AppInvitationSettingsTable = &schema.Table{
		Name:       "app_invitation_settings",
		Columns:    AppInvitationSettingsColumns,
		PrimaryKey: []*schema.Column{AppInvitationSettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appinvitationsetting_app_id_count",
				Unique:  true,
				Columns: []*schema.Column{AppInvitationSettingsColumns[1], AppInvitationSettingsColumns[2]},
			},
		},
	}
	// AppPurchaseAmountSettingsColumns holds the columns for the "app_purchase_amount_settings" table.
	AppPurchaseAmountSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "percent", Type: field.TypeUint32},
		{Name: "start", Type: field.TypeUint32},
		{Name: "end", Type: field.TypeUint32},
		{Name: "badge_large", Type: field.TypeString},
		{Name: "badge_small", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppPurchaseAmountSettingsTable holds the schema information for the "app_purchase_amount_settings" table.
	AppPurchaseAmountSettingsTable = &schema.Table{
		Name:       "app_purchase_amount_settings",
		Columns:    AppPurchaseAmountSettingsColumns,
		PrimaryKey: []*schema.Column{AppPurchaseAmountSettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "apppurchaseamountsetting_app_id_amount",
				Unique:  true,
				Columns: []*schema.Column{AppPurchaseAmountSettingsColumns[1], AppPurchaseAmountSettingsColumns[3]},
			},
		},
	}
	// AppUserInvitationSettingsColumns holds the columns for the "app_user_invitation_settings" table.
	AppUserInvitationSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "count", Type: field.TypeUint32},
		{Name: "discount", Type: field.TypeUint32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
		{Name: "title", Type: field.TypeString},
		{Name: "badge_large", Type: field.TypeString},
		{Name: "badge_small", Type: field.TypeString},
	}
	// AppUserInvitationSettingsTable holds the schema information for the "app_user_invitation_settings" table.
	AppUserInvitationSettingsTable = &schema.Table{
		Name:       "app_user_invitation_settings",
		Columns:    AppUserInvitationSettingsColumns,
		PrimaryKey: []*schema.Column{AppUserInvitationSettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appuserinvitationsetting_app_id_user_id_count",
				Unique:  true,
				Columns: []*schema.Column{AppUserInvitationSettingsColumns[1], AppUserInvitationSettingsColumns[2], AppUserInvitationSettingsColumns[3]},
			},
		},
	}
	// AppUserPurchaseAmountSettingsColumns holds the columns for the "app_user_purchase_amount_settings" table.
	AppUserPurchaseAmountSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "percent", Type: field.TypeUint32},
		{Name: "start", Type: field.TypeUint32},
		{Name: "end", Type: field.TypeUint32},
		{Name: "badge_large", Type: field.TypeString},
		{Name: "badge_small", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppUserPurchaseAmountSettingsTable holds the schema information for the "app_user_purchase_amount_settings" table.
	AppUserPurchaseAmountSettingsTable = &schema.Table{
		Name:       "app_user_purchase_amount_settings",
		Columns:    AppUserPurchaseAmountSettingsColumns,
		PrimaryKey: []*schema.Column{AppUserPurchaseAmountSettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appuserpurchaseamountsetting_app_id_user_id_amount",
				Unique:  true,
				Columns: []*schema.Column{AppUserPurchaseAmountSettingsColumns[1], AppUserPurchaseAmountSettingsColumns[2], AppUserPurchaseAmountSettingsColumns[4]},
			},
		},
	}
	// CommissionCoinSettingsColumns holds the columns for the "commission_coin_settings" table.
	CommissionCoinSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Unique: true},
		{Name: "using", Type: field.TypeBool, Default: false},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CommissionCoinSettingsTable holds the schema information for the "commission_coin_settings" table.
	CommissionCoinSettingsTable = &schema.Table{
		Name:       "commission_coin_settings",
		Columns:    CommissionCoinSettingsColumns,
		PrimaryKey: []*schema.Column{CommissionCoinSettingsColumns[0]},
	}
	// CouponAllocatedsColumns holds the columns for the "coupon_allocateds" table.
	CouponAllocatedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString},
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
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "denomination", Type: field.TypeUint64},
		{Name: "circulation", Type: field.TypeInt32},
		{Name: "release_by_user_id", Type: field.TypeUUID},
		{Name: "start", Type: field.TypeUint32},
		{Name: "duration_days", Type: field.TypeInt32},
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
	// DefaultKpiSettingsColumns holds the columns for the "default_kpi_settings" table.
	DefaultKpiSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "percent", Type: field.TypeInt32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// DefaultKpiSettingsTable holds the schema information for the "default_kpi_settings" table.
	DefaultKpiSettingsTable = &schema.Table{
		Name:       "default_kpi_settings",
		Columns:    DefaultKpiSettingsColumns,
		PrimaryKey: []*schema.Column{DefaultKpiSettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "defaultkpisetting_app_id_good_id",
				Unique:  true,
				Columns: []*schema.Column{DefaultKpiSettingsColumns[3], DefaultKpiSettingsColumns[4]},
			},
		},
	}
	// DiscountPoolsColumns holds the columns for the "discount_pools" table.
	DiscountPoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "discount", Type: field.TypeUint32},
		{Name: "release_by_user_id", Type: field.TypeUUID},
		{Name: "start", Type: field.TypeUint32},
		{Name: "duration_days", Type: field.TypeInt32},
		{Name: "message", Type: field.TypeString, Size: 512},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// DiscountPoolsTable holds the schema information for the "discount_pools" table.
	DiscountPoolsTable = &schema.Table{
		Name:       "discount_pools",
		Columns:    DiscountPoolsColumns,
		PrimaryKey: []*schema.Column{DiscountPoolsColumns[0]},
	}
	// EventCouponsColumns holds the columns for the "event_coupons" table.
	EventCouponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "activity_id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString},
		{Name: "coupon_id", Type: field.TypeUUID},
		{Name: "event", Type: field.TypeString},
		{Name: "count", Type: field.TypeUint32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// EventCouponsTable holds the schema information for the "event_coupons" table.
	EventCouponsTable = &schema.Table{
		Name:       "event_coupons",
		Columns:    EventCouponsColumns,
		PrimaryKey: []*schema.Column{EventCouponsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "eventcoupon_app_id_activity_id_type_coupon_id_event",
				Unique:  true,
				Columns: []*schema.Column{EventCouponsColumns[1], EventCouponsColumns[2], EventCouponsColumns[3], EventCouponsColumns[4], EventCouponsColumns[5]},
			},
		},
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
		Indexes: []*schema.Index{
			{
				Name:    "registrationinvitation_app_id_invitee_id",
				Unique:  true,
				Columns: []*schema.Column{RegistrationInvitationsColumns[6], RegistrationInvitationsColumns[5]},
			},
		},
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
		Indexes: []*schema.Index{
			{
				Name:    "userinvitationcode_app_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{UserInvitationCodesColumns[2], UserInvitationCodesColumns[1]},
			},
		},
	}
	// UserKpiSettingsColumns holds the columns for the "user_kpi_settings" table.
	UserKpiSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "percent", Type: field.TypeInt32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserKpiSettingsTable holds the schema information for the "user_kpi_settings" table.
	UserKpiSettingsTable = &schema.Table{
		Name:       "user_kpi_settings",
		Columns:    UserKpiSettingsColumns,
		PrimaryKey: []*schema.Column{UserKpiSettingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userkpisetting_app_id_good_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{UserKpiSettingsColumns[3], UserKpiSettingsColumns[4], UserKpiSettingsColumns[5]},
			},
		},
	}
	// UserSpecialReductionsColumns holds the columns for the "user_special_reductions" table.
	UserSpecialReductionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "release_by_user_id", Type: field.TypeUUID},
		{Name: "start", Type: field.TypeUint32},
		{Name: "duration_days", Type: field.TypeInt32},
		{Name: "message", Type: field.TypeString, Size: 512},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserSpecialReductionsTable holds the schema information for the "user_special_reductions" table.
	UserSpecialReductionsTable = &schema.Table{
		Name:       "user_special_reductions",
		Columns:    UserSpecialReductionsColumns,
		PrimaryKey: []*schema.Column{UserSpecialReductionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActivitiesTable,
		AppCommissionSettingsTable,
		AppCouponSettingsTable,
		AppInvitationSettingsTable,
		AppPurchaseAmountSettingsTable,
		AppUserInvitationSettingsTable,
		AppUserPurchaseAmountSettingsTable,
		CommissionCoinSettingsTable,
		CouponAllocatedsTable,
		CouponPoolsTable,
		DefaultKpiSettingsTable,
		DiscountPoolsTable,
		EventCouponsTable,
		RegistrationInvitationsTable,
		UserInvitationCodesTable,
		UserKpiSettingsTable,
		UserSpecialReductionsTable,
	}
)

func init() {
}
