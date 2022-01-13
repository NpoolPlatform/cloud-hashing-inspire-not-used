package newuserrewardsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"             //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/new-user-reward-setting" //nolint

	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetNewUserRewardSettingDetailRequest) (*npool.GetNewUserRewardSettingDetailResponse, error) {
	info, err := newuserrewardsetting.Get(ctx, &npool.GetNewUserRewardSettingRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get new user reward setting: %v", err)
	}

	registrationCouponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
		ID: info.Info.RegistrationCouponID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get new user reward setting registration coupon pool: %v", err)
	}

	kycCouponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
		ID: info.Info.KycCouponID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get new user reward setting kyc coupon pool: %v", err)
	}

	return &npool.GetNewUserRewardSettingDetailResponse{
		Info: &npool.NewUserRewardSettingDetail{
			ID:                         info.Info.ID,
			AppID:                      info.Info.AppID,
			AutoGenerateInvitationCode: info.Info.AutoGenerateInvitationCode,
			RegistrationCoupon:         registrationCouponPool.Info,
			KycCoupon:                  kycCouponPool.Info,
		},
	}, nil
}
