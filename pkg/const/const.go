package constant

type DTMAction struct {
	Action string
	Revert string
}

const (
	CouponTypeCoupon   = "coupon"
	CouponTypeDiscount = "discount"

	EventSharing                 = "sharing"
	EventInvitationRegisteration = "invitation-registeration"
	EventInvitationPurchase      = "invitation-purchase"
	EventRegister                = "register"
	EventKYCAuthenticate         = "kyc-authenticate"
	EventTotalAmount             = "total-amount"
	EventSingleAmount            = "single-amount"

	CommissionByAmount = "commission-by-amount"

	CreateRegistrationInvitation       = "CreateRegistrationInvitation"
	CreateRegistrationInvitationRevert = "CreateRegistrationInvitationRevert"
)

var DTMEntry = map[string]*DTMAction{
	CreateRegistrationInvitation: &DTMAction{ //nolint
		Action: CreateRegistrationInvitation,
		Revert: CreateRegistrationInvitationRevert,
	},
}
