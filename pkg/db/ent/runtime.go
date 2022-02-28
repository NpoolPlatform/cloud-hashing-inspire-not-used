// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/activity"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcommissionsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcouponsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appinvitationsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/apppurchaseamountsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appuserinvitationsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appuserpurchaseamountsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/commissioncoinsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/defaultkpisetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/discountpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/eventcoupon"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userkpisetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userspecialreduction"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	activityFields := schema.Activity{}.Fields()
	_ = activityFields
	// activityDescCreateAt is the schema descriptor for create_at field.
	activityDescCreateAt := activityFields[7].Descriptor()
	// activity.DefaultCreateAt holds the default value on creation for the create_at field.
	activity.DefaultCreateAt = activityDescCreateAt.Default.(func() uint32)
	// activityDescUpdateAt is the schema descriptor for update_at field.
	activityDescUpdateAt := activityFields[8].Descriptor()
	// activity.DefaultUpdateAt holds the default value on creation for the update_at field.
	activity.DefaultUpdateAt = activityDescUpdateAt.Default.(func() uint32)
	// activity.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	activity.UpdateDefaultUpdateAt = activityDescUpdateAt.UpdateDefault.(func() uint32)
	// activityDescDeleteAt is the schema descriptor for delete_at field.
	activityDescDeleteAt := activityFields[9].Descriptor()
	// activity.DefaultDeleteAt holds the default value on creation for the delete_at field.
	activity.DefaultDeleteAt = activityDescDeleteAt.Default.(func() uint32)
	// activityDescID is the schema descriptor for id field.
	activityDescID := activityFields[0].Descriptor()
	// activity.DefaultID holds the default value on creation for the id field.
	activity.DefaultID = activityDescID.Default.(func() uuid.UUID)
	appcommissionsettingFields := schema.AppCommissionSetting{}.Fields()
	_ = appcommissionsettingFields
	// appcommissionsettingDescCreateAt is the schema descriptor for create_at field.
	appcommissionsettingDescCreateAt := appcommissionsettingFields[6].Descriptor()
	// appcommissionsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	appcommissionsetting.DefaultCreateAt = appcommissionsettingDescCreateAt.Default.(func() uint32)
	// appcommissionsettingDescUpdateAt is the schema descriptor for update_at field.
	appcommissionsettingDescUpdateAt := appcommissionsettingFields[7].Descriptor()
	// appcommissionsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	appcommissionsetting.DefaultUpdateAt = appcommissionsettingDescUpdateAt.Default.(func() uint32)
	// appcommissionsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appcommissionsetting.UpdateDefaultUpdateAt = appcommissionsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// appcommissionsettingDescDeleteAt is the schema descriptor for delete_at field.
	appcommissionsettingDescDeleteAt := appcommissionsettingFields[8].Descriptor()
	// appcommissionsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appcommissionsetting.DefaultDeleteAt = appcommissionsettingDescDeleteAt.Default.(func() uint32)
	// appcommissionsettingDescID is the schema descriptor for id field.
	appcommissionsettingDescID := appcommissionsettingFields[0].Descriptor()
	// appcommissionsetting.DefaultID holds the default value on creation for the id field.
	appcommissionsetting.DefaultID = appcommissionsettingDescID.Default.(func() uuid.UUID)
	appcouponsettingFields := schema.AppCouponSetting{}.Fields()
	_ = appcouponsettingFields
	// appcouponsettingDescCreateAt is the schema descriptor for create_at field.
	appcouponsettingDescCreateAt := appcouponsettingFields[4].Descriptor()
	// appcouponsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	appcouponsetting.DefaultCreateAt = appcouponsettingDescCreateAt.Default.(func() uint32)
	// appcouponsettingDescUpdateAt is the schema descriptor for update_at field.
	appcouponsettingDescUpdateAt := appcouponsettingFields[5].Descriptor()
	// appcouponsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	appcouponsetting.DefaultUpdateAt = appcouponsettingDescUpdateAt.Default.(func() uint32)
	// appcouponsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appcouponsetting.UpdateDefaultUpdateAt = appcouponsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// appcouponsettingDescDeleteAt is the schema descriptor for delete_at field.
	appcouponsettingDescDeleteAt := appcouponsettingFields[6].Descriptor()
	// appcouponsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appcouponsetting.DefaultDeleteAt = appcouponsettingDescDeleteAt.Default.(func() uint32)
	// appcouponsettingDescID is the schema descriptor for id field.
	appcouponsettingDescID := appcouponsettingFields[0].Descriptor()
	// appcouponsetting.DefaultID holds the default value on creation for the id field.
	appcouponsetting.DefaultID = appcouponsettingDescID.Default.(func() uuid.UUID)
	appinvitationsettingFields := schema.AppInvitationSetting{}.Fields()
	_ = appinvitationsettingFields
	// appinvitationsettingDescCreateAt is the schema descriptor for create_at field.
	appinvitationsettingDescCreateAt := appinvitationsettingFields[4].Descriptor()
	// appinvitationsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	appinvitationsetting.DefaultCreateAt = appinvitationsettingDescCreateAt.Default.(func() uint32)
	// appinvitationsettingDescUpdateAt is the schema descriptor for update_at field.
	appinvitationsettingDescUpdateAt := appinvitationsettingFields[5].Descriptor()
	// appinvitationsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	appinvitationsetting.DefaultUpdateAt = appinvitationsettingDescUpdateAt.Default.(func() uint32)
	// appinvitationsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appinvitationsetting.UpdateDefaultUpdateAt = appinvitationsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// appinvitationsettingDescDeleteAt is the schema descriptor for delete_at field.
	appinvitationsettingDescDeleteAt := appinvitationsettingFields[6].Descriptor()
	// appinvitationsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appinvitationsetting.DefaultDeleteAt = appinvitationsettingDescDeleteAt.Default.(func() uint32)
	// appinvitationsettingDescID is the schema descriptor for id field.
	appinvitationsettingDescID := appinvitationsettingFields[0].Descriptor()
	// appinvitationsetting.DefaultID holds the default value on creation for the id field.
	appinvitationsetting.DefaultID = appinvitationsettingDescID.Default.(func() uuid.UUID)
	apppurchaseamountsettingFields := schema.AppPurchaseAmountSetting{}.Fields()
	_ = apppurchaseamountsettingFields
	// apppurchaseamountsettingDescCreateAt is the schema descriptor for create_at field.
	apppurchaseamountsettingDescCreateAt := apppurchaseamountsettingFields[9].Descriptor()
	// apppurchaseamountsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	apppurchaseamountsetting.DefaultCreateAt = apppurchaseamountsettingDescCreateAt.Default.(func() uint32)
	// apppurchaseamountsettingDescUpdateAt is the schema descriptor for update_at field.
	apppurchaseamountsettingDescUpdateAt := apppurchaseamountsettingFields[10].Descriptor()
	// apppurchaseamountsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	apppurchaseamountsetting.DefaultUpdateAt = apppurchaseamountsettingDescUpdateAt.Default.(func() uint32)
	// apppurchaseamountsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	apppurchaseamountsetting.UpdateDefaultUpdateAt = apppurchaseamountsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// apppurchaseamountsettingDescDeleteAt is the schema descriptor for delete_at field.
	apppurchaseamountsettingDescDeleteAt := apppurchaseamountsettingFields[11].Descriptor()
	// apppurchaseamountsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	apppurchaseamountsetting.DefaultDeleteAt = apppurchaseamountsettingDescDeleteAt.Default.(func() uint32)
	// apppurchaseamountsettingDescID is the schema descriptor for id field.
	apppurchaseamountsettingDescID := apppurchaseamountsettingFields[0].Descriptor()
	// apppurchaseamountsetting.DefaultID holds the default value on creation for the id field.
	apppurchaseamountsetting.DefaultID = apppurchaseamountsettingDescID.Default.(func() uuid.UUID)
	appuserinvitationsettingFields := schema.AppUserInvitationSetting{}.Fields()
	_ = appuserinvitationsettingFields
	// appuserinvitationsettingDescCreateAt is the schema descriptor for create_at field.
	appuserinvitationsettingDescCreateAt := appuserinvitationsettingFields[5].Descriptor()
	// appuserinvitationsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	appuserinvitationsetting.DefaultCreateAt = appuserinvitationsettingDescCreateAt.Default.(func() uint32)
	// appuserinvitationsettingDescUpdateAt is the schema descriptor for update_at field.
	appuserinvitationsettingDescUpdateAt := appuserinvitationsettingFields[6].Descriptor()
	// appuserinvitationsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	appuserinvitationsetting.DefaultUpdateAt = appuserinvitationsettingDescUpdateAt.Default.(func() uint32)
	// appuserinvitationsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appuserinvitationsetting.UpdateDefaultUpdateAt = appuserinvitationsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// appuserinvitationsettingDescDeleteAt is the schema descriptor for delete_at field.
	appuserinvitationsettingDescDeleteAt := appuserinvitationsettingFields[7].Descriptor()
	// appuserinvitationsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appuserinvitationsetting.DefaultDeleteAt = appuserinvitationsettingDescDeleteAt.Default.(func() uint32)
	// appuserinvitationsettingDescID is the schema descriptor for id field.
	appuserinvitationsettingDescID := appuserinvitationsettingFields[0].Descriptor()
	// appuserinvitationsetting.DefaultID holds the default value on creation for the id field.
	appuserinvitationsetting.DefaultID = appuserinvitationsettingDescID.Default.(func() uuid.UUID)
	appuserpurchaseamountsettingFields := schema.AppUserPurchaseAmountSetting{}.Fields()
	_ = appuserpurchaseamountsettingFields
	// appuserpurchaseamountsettingDescCreateAt is the schema descriptor for create_at field.
	appuserpurchaseamountsettingDescCreateAt := appuserpurchaseamountsettingFields[10].Descriptor()
	// appuserpurchaseamountsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	appuserpurchaseamountsetting.DefaultCreateAt = appuserpurchaseamountsettingDescCreateAt.Default.(func() uint32)
	// appuserpurchaseamountsettingDescUpdateAt is the schema descriptor for update_at field.
	appuserpurchaseamountsettingDescUpdateAt := appuserpurchaseamountsettingFields[11].Descriptor()
	// appuserpurchaseamountsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	appuserpurchaseamountsetting.DefaultUpdateAt = appuserpurchaseamountsettingDescUpdateAt.Default.(func() uint32)
	// appuserpurchaseamountsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appuserpurchaseamountsetting.UpdateDefaultUpdateAt = appuserpurchaseamountsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// appuserpurchaseamountsettingDescDeleteAt is the schema descriptor for delete_at field.
	appuserpurchaseamountsettingDescDeleteAt := appuserpurchaseamountsettingFields[12].Descriptor()
	// appuserpurchaseamountsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appuserpurchaseamountsetting.DefaultDeleteAt = appuserpurchaseamountsettingDescDeleteAt.Default.(func() uint32)
	// appuserpurchaseamountsettingDescID is the schema descriptor for id field.
	appuserpurchaseamountsettingDescID := appuserpurchaseamountsettingFields[0].Descriptor()
	// appuserpurchaseamountsetting.DefaultID holds the default value on creation for the id field.
	appuserpurchaseamountsetting.DefaultID = appuserpurchaseamountsettingDescID.Default.(func() uuid.UUID)
	commissioncoinsettingFields := schema.CommissionCoinSetting{}.Fields()
	_ = commissioncoinsettingFields
	// commissioncoinsettingDescUsing is the schema descriptor for using field.
	commissioncoinsettingDescUsing := commissioncoinsettingFields[2].Descriptor()
	// commissioncoinsetting.DefaultUsing holds the default value on creation for the using field.
	commissioncoinsetting.DefaultUsing = commissioncoinsettingDescUsing.Default.(bool)
	// commissioncoinsettingDescCreateAt is the schema descriptor for create_at field.
	commissioncoinsettingDescCreateAt := commissioncoinsettingFields[3].Descriptor()
	// commissioncoinsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	commissioncoinsetting.DefaultCreateAt = commissioncoinsettingDescCreateAt.Default.(func() uint32)
	// commissioncoinsettingDescUpdateAt is the schema descriptor for update_at field.
	commissioncoinsettingDescUpdateAt := commissioncoinsettingFields[4].Descriptor()
	// commissioncoinsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	commissioncoinsetting.DefaultUpdateAt = commissioncoinsettingDescUpdateAt.Default.(func() uint32)
	// commissioncoinsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	commissioncoinsetting.UpdateDefaultUpdateAt = commissioncoinsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// commissioncoinsettingDescDeleteAt is the schema descriptor for delete_at field.
	commissioncoinsettingDescDeleteAt := commissioncoinsettingFields[5].Descriptor()
	// commissioncoinsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	commissioncoinsetting.DefaultDeleteAt = commissioncoinsettingDescDeleteAt.Default.(func() uint32)
	// commissioncoinsettingDescID is the schema descriptor for id field.
	commissioncoinsettingDescID := commissioncoinsettingFields[0].Descriptor()
	// commissioncoinsetting.DefaultID holds the default value on creation for the id field.
	commissioncoinsetting.DefaultID = commissioncoinsettingDescID.Default.(func() uuid.UUID)
	couponallocatedFields := schema.CouponAllocated{}.Fields()
	_ = couponallocatedFields
	// couponallocatedDescCreateAt is the schema descriptor for create_at field.
	couponallocatedDescCreateAt := couponallocatedFields[5].Descriptor()
	// couponallocated.DefaultCreateAt holds the default value on creation for the create_at field.
	couponallocated.DefaultCreateAt = couponallocatedDescCreateAt.Default.(func() uint32)
	// couponallocatedDescUpdateAt is the schema descriptor for update_at field.
	couponallocatedDescUpdateAt := couponallocatedFields[6].Descriptor()
	// couponallocated.DefaultUpdateAt holds the default value on creation for the update_at field.
	couponallocated.DefaultUpdateAt = couponallocatedDescUpdateAt.Default.(func() uint32)
	// couponallocated.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	couponallocated.UpdateDefaultUpdateAt = couponallocatedDescUpdateAt.UpdateDefault.(func() uint32)
	// couponallocatedDescDeleteAt is the schema descriptor for delete_at field.
	couponallocatedDescDeleteAt := couponallocatedFields[7].Descriptor()
	// couponallocated.DefaultDeleteAt holds the default value on creation for the delete_at field.
	couponallocated.DefaultDeleteAt = couponallocatedDescDeleteAt.Default.(func() uint32)
	// couponallocatedDescID is the schema descriptor for id field.
	couponallocatedDescID := couponallocatedFields[0].Descriptor()
	// couponallocated.DefaultID holds the default value on creation for the id field.
	couponallocated.DefaultID = couponallocatedDescID.Default.(func() uuid.UUID)
	couponpoolFields := schema.CouponPool{}.Fields()
	_ = couponpoolFields
	// couponpoolDescMessage is the schema descriptor for message field.
	couponpoolDescMessage := couponpoolFields[7].Descriptor()
	// couponpool.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	couponpool.MessageValidator = couponpoolDescMessage.Validators[0].(func(string) error)
	// couponpoolDescName is the schema descriptor for name field.
	couponpoolDescName := couponpoolFields[8].Descriptor()
	// couponpool.NameValidator is a validator for the "name" field. It is called by the builders before save.
	couponpool.NameValidator = couponpoolDescName.Validators[0].(func(string) error)
	// couponpoolDescCreateAt is the schema descriptor for create_at field.
	couponpoolDescCreateAt := couponpoolFields[9].Descriptor()
	// couponpool.DefaultCreateAt holds the default value on creation for the create_at field.
	couponpool.DefaultCreateAt = couponpoolDescCreateAt.Default.(func() uint32)
	// couponpoolDescUpdateAt is the schema descriptor for update_at field.
	couponpoolDescUpdateAt := couponpoolFields[10].Descriptor()
	// couponpool.DefaultUpdateAt holds the default value on creation for the update_at field.
	couponpool.DefaultUpdateAt = couponpoolDescUpdateAt.Default.(func() uint32)
	// couponpool.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	couponpool.UpdateDefaultUpdateAt = couponpoolDescUpdateAt.UpdateDefault.(func() uint32)
	// couponpoolDescDeleteAt is the schema descriptor for delete_at field.
	couponpoolDescDeleteAt := couponpoolFields[11].Descriptor()
	// couponpool.DefaultDeleteAt holds the default value on creation for the delete_at field.
	couponpool.DefaultDeleteAt = couponpoolDescDeleteAt.Default.(func() uint32)
	// couponpoolDescID is the schema descriptor for id field.
	couponpoolDescID := couponpoolFields[0].Descriptor()
	// couponpool.DefaultID holds the default value on creation for the id field.
	couponpool.DefaultID = couponpoolDescID.Default.(func() uuid.UUID)
	defaultkpisettingFields := schema.DefaultKpiSetting{}.Fields()
	_ = defaultkpisettingFields
	// defaultkpisettingDescCreateAt is the schema descriptor for create_at field.
	defaultkpisettingDescCreateAt := defaultkpisettingFields[5].Descriptor()
	// defaultkpisetting.DefaultCreateAt holds the default value on creation for the create_at field.
	defaultkpisetting.DefaultCreateAt = defaultkpisettingDescCreateAt.Default.(func() uint32)
	// defaultkpisettingDescUpdateAt is the schema descriptor for update_at field.
	defaultkpisettingDescUpdateAt := defaultkpisettingFields[6].Descriptor()
	// defaultkpisetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	defaultkpisetting.DefaultUpdateAt = defaultkpisettingDescUpdateAt.Default.(func() uint32)
	// defaultkpisetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	defaultkpisetting.UpdateDefaultUpdateAt = defaultkpisettingDescUpdateAt.UpdateDefault.(func() uint32)
	// defaultkpisettingDescDeleteAt is the schema descriptor for delete_at field.
	defaultkpisettingDescDeleteAt := defaultkpisettingFields[7].Descriptor()
	// defaultkpisetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	defaultkpisetting.DefaultDeleteAt = defaultkpisettingDescDeleteAt.Default.(func() uint32)
	// defaultkpisettingDescID is the schema descriptor for id field.
	defaultkpisettingDescID := defaultkpisettingFields[0].Descriptor()
	// defaultkpisetting.DefaultID holds the default value on creation for the id field.
	defaultkpisetting.DefaultID = defaultkpisettingDescID.Default.(func() uuid.UUID)
	discountpoolFields := schema.DiscountPool{}.Fields()
	_ = discountpoolFields
	// discountpoolDescMessage is the schema descriptor for message field.
	discountpoolDescMessage := discountpoolFields[6].Descriptor()
	// discountpool.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	discountpool.MessageValidator = discountpoolDescMessage.Validators[0].(func(string) error)
	// discountpoolDescName is the schema descriptor for name field.
	discountpoolDescName := discountpoolFields[7].Descriptor()
	// discountpool.NameValidator is a validator for the "name" field. It is called by the builders before save.
	discountpool.NameValidator = discountpoolDescName.Validators[0].(func(string) error)
	// discountpoolDescCreateAt is the schema descriptor for create_at field.
	discountpoolDescCreateAt := discountpoolFields[8].Descriptor()
	// discountpool.DefaultCreateAt holds the default value on creation for the create_at field.
	discountpool.DefaultCreateAt = discountpoolDescCreateAt.Default.(func() uint32)
	// discountpoolDescUpdateAt is the schema descriptor for update_at field.
	discountpoolDescUpdateAt := discountpoolFields[9].Descriptor()
	// discountpool.DefaultUpdateAt holds the default value on creation for the update_at field.
	discountpool.DefaultUpdateAt = discountpoolDescUpdateAt.Default.(func() uint32)
	// discountpool.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	discountpool.UpdateDefaultUpdateAt = discountpoolDescUpdateAt.UpdateDefault.(func() uint32)
	// discountpoolDescDeleteAt is the schema descriptor for delete_at field.
	discountpoolDescDeleteAt := discountpoolFields[10].Descriptor()
	// discountpool.DefaultDeleteAt holds the default value on creation for the delete_at field.
	discountpool.DefaultDeleteAt = discountpoolDescDeleteAt.Default.(func() uint32)
	// discountpoolDescID is the schema descriptor for id field.
	discountpoolDescID := discountpoolFields[0].Descriptor()
	// discountpool.DefaultID holds the default value on creation for the id field.
	discountpool.DefaultID = discountpoolDescID.Default.(func() uuid.UUID)
	eventcouponFields := schema.EventCoupon{}.Fields()
	_ = eventcouponFields
	// eventcouponDescCreateAt is the schema descriptor for create_at field.
	eventcouponDescCreateAt := eventcouponFields[7].Descriptor()
	// eventcoupon.DefaultCreateAt holds the default value on creation for the create_at field.
	eventcoupon.DefaultCreateAt = eventcouponDescCreateAt.Default.(func() uint32)
	// eventcouponDescUpdateAt is the schema descriptor for update_at field.
	eventcouponDescUpdateAt := eventcouponFields[8].Descriptor()
	// eventcoupon.DefaultUpdateAt holds the default value on creation for the update_at field.
	eventcoupon.DefaultUpdateAt = eventcouponDescUpdateAt.Default.(func() uint32)
	// eventcoupon.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	eventcoupon.UpdateDefaultUpdateAt = eventcouponDescUpdateAt.UpdateDefault.(func() uint32)
	// eventcouponDescDeleteAt is the schema descriptor for delete_at field.
	eventcouponDescDeleteAt := eventcouponFields[9].Descriptor()
	// eventcoupon.DefaultDeleteAt holds the default value on creation for the delete_at field.
	eventcoupon.DefaultDeleteAt = eventcouponDescDeleteAt.Default.(func() uint32)
	// eventcouponDescID is the schema descriptor for id field.
	eventcouponDescID := eventcouponFields[0].Descriptor()
	// eventcoupon.DefaultID holds the default value on creation for the id field.
	eventcoupon.DefaultID = eventcouponDescID.Default.(func() uuid.UUID)
	registrationinvitationFields := schema.RegistrationInvitation{}.Fields()
	_ = registrationinvitationFields
	// registrationinvitationDescCreateAt is the schema descriptor for create_at field.
	registrationinvitationDescCreateAt := registrationinvitationFields[1].Descriptor()
	// registrationinvitation.DefaultCreateAt holds the default value on creation for the create_at field.
	registrationinvitation.DefaultCreateAt = registrationinvitationDescCreateAt.Default.(func() uint32)
	// registrationinvitationDescUpdateAt is the schema descriptor for update_at field.
	registrationinvitationDescUpdateAt := registrationinvitationFields[2].Descriptor()
	// registrationinvitation.DefaultUpdateAt holds the default value on creation for the update_at field.
	registrationinvitation.DefaultUpdateAt = registrationinvitationDescUpdateAt.Default.(func() uint32)
	// registrationinvitation.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	registrationinvitation.UpdateDefaultUpdateAt = registrationinvitationDescUpdateAt.UpdateDefault.(func() uint32)
	// registrationinvitationDescDeleteAt is the schema descriptor for delete_at field.
	registrationinvitationDescDeleteAt := registrationinvitationFields[3].Descriptor()
	// registrationinvitation.DefaultDeleteAt holds the default value on creation for the delete_at field.
	registrationinvitation.DefaultDeleteAt = registrationinvitationDescDeleteAt.Default.(func() uint32)
	// registrationinvitationDescID is the schema descriptor for id field.
	registrationinvitationDescID := registrationinvitationFields[0].Descriptor()
	// registrationinvitation.DefaultID holds the default value on creation for the id field.
	registrationinvitation.DefaultID = registrationinvitationDescID.Default.(func() uuid.UUID)
	userinvitationcodeFields := schema.UserInvitationCode{}.Fields()
	_ = userinvitationcodeFields
	// userinvitationcodeDescCreateAt is the schema descriptor for create_at field.
	userinvitationcodeDescCreateAt := userinvitationcodeFields[4].Descriptor()
	// userinvitationcode.DefaultCreateAt holds the default value on creation for the create_at field.
	userinvitationcode.DefaultCreateAt = userinvitationcodeDescCreateAt.Default.(func() uint32)
	// userinvitationcodeDescUpdateAt is the schema descriptor for update_at field.
	userinvitationcodeDescUpdateAt := userinvitationcodeFields[5].Descriptor()
	// userinvitationcode.DefaultUpdateAt holds the default value on creation for the update_at field.
	userinvitationcode.DefaultUpdateAt = userinvitationcodeDescUpdateAt.Default.(func() uint32)
	// userinvitationcode.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userinvitationcode.UpdateDefaultUpdateAt = userinvitationcodeDescUpdateAt.UpdateDefault.(func() uint32)
	// userinvitationcodeDescDeleteAt is the schema descriptor for delete_at field.
	userinvitationcodeDescDeleteAt := userinvitationcodeFields[6].Descriptor()
	// userinvitationcode.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userinvitationcode.DefaultDeleteAt = userinvitationcodeDescDeleteAt.Default.(func() uint32)
	// userinvitationcodeDescID is the schema descriptor for id field.
	userinvitationcodeDescID := userinvitationcodeFields[0].Descriptor()
	// userinvitationcode.DefaultID holds the default value on creation for the id field.
	userinvitationcode.DefaultID = userinvitationcodeDescID.Default.(func() uuid.UUID)
	userkpisettingFields := schema.UserKpiSetting{}.Fields()
	_ = userkpisettingFields
	// userkpisettingDescCreateAt is the schema descriptor for create_at field.
	userkpisettingDescCreateAt := userkpisettingFields[6].Descriptor()
	// userkpisetting.DefaultCreateAt holds the default value on creation for the create_at field.
	userkpisetting.DefaultCreateAt = userkpisettingDescCreateAt.Default.(func() uint32)
	// userkpisettingDescUpdateAt is the schema descriptor for update_at field.
	userkpisettingDescUpdateAt := userkpisettingFields[7].Descriptor()
	// userkpisetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	userkpisetting.DefaultUpdateAt = userkpisettingDescUpdateAt.Default.(func() uint32)
	// userkpisetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userkpisetting.UpdateDefaultUpdateAt = userkpisettingDescUpdateAt.UpdateDefault.(func() uint32)
	// userkpisettingDescDeleteAt is the schema descriptor for delete_at field.
	userkpisettingDescDeleteAt := userkpisettingFields[8].Descriptor()
	// userkpisetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userkpisetting.DefaultDeleteAt = userkpisettingDescDeleteAt.Default.(func() uint32)
	// userkpisettingDescID is the schema descriptor for id field.
	userkpisettingDescID := userkpisettingFields[0].Descriptor()
	// userkpisetting.DefaultID holds the default value on creation for the id field.
	userkpisetting.DefaultID = userkpisettingDescID.Default.(func() uuid.UUID)
	userspecialreductionFields := schema.UserSpecialReduction{}.Fields()
	_ = userspecialreductionFields
	// userspecialreductionDescMessage is the schema descriptor for message field.
	userspecialreductionDescMessage := userspecialreductionFields[7].Descriptor()
	// userspecialreduction.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	userspecialreduction.MessageValidator = userspecialreductionDescMessage.Validators[0].(func(string) error)
	// userspecialreductionDescCreateAt is the schema descriptor for create_at field.
	userspecialreductionDescCreateAt := userspecialreductionFields[8].Descriptor()
	// userspecialreduction.DefaultCreateAt holds the default value on creation for the create_at field.
	userspecialreduction.DefaultCreateAt = userspecialreductionDescCreateAt.Default.(func() uint32)
	// userspecialreductionDescUpdateAt is the schema descriptor for update_at field.
	userspecialreductionDescUpdateAt := userspecialreductionFields[9].Descriptor()
	// userspecialreduction.DefaultUpdateAt holds the default value on creation for the update_at field.
	userspecialreduction.DefaultUpdateAt = userspecialreductionDescUpdateAt.Default.(func() uint32)
	// userspecialreduction.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userspecialreduction.UpdateDefaultUpdateAt = userspecialreductionDescUpdateAt.UpdateDefault.(func() uint32)
	// userspecialreductionDescDeleteAt is the schema descriptor for delete_at field.
	userspecialreductionDescDeleteAt := userspecialreductionFields[10].Descriptor()
	// userspecialreduction.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userspecialreduction.DefaultDeleteAt = userspecialreductionDescDeleteAt.Default.(func() uint32)
	// userspecialreductionDescID is the schema descriptor for id field.
	userspecialreductionDescID := userspecialreductionFields[0].Descriptor()
	// userspecialreduction.DefaultID holds the default value on creation for the id field.
	userspecialreduction.DefaultID = userspecialreductionDescID.Default.(func() uuid.UUID)
}
