package agencysetting

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/agency-setting" //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"    //nolint

	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetAgencySettingDetailRequest) (*npool.GetAgencySettingDetailResponse, error) {
	info, err := agencysetting.Get(ctx, &npool.GetAgencySettingRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get agency setting: %v", err)
	}

	registrationCouponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
		ID: info.Info.RegistrationCouponID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get agency setting registration coupon pool: %v", err)
	}

	kycCouponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
		ID: info.Info.KycCouponID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get agency setting kyc coupon pool: %v", err)
	}

	return &npool.GetAgencySettingDetailResponse{
		Info: &npool.AgencySettingDetail{
			ID:                          info.Info.ID,
			AppID:                       info.Info.AppID,
			GoodID:                      info.Info.GoodID,
			RegistrationRewardThreshold: info.Info.RegistrationRewardThreshold,
			KycRewardThreshold:          info.Info.KycRewardThreshold,
			RegistrationCoupon:          registrationCouponPool.Info,
			KycCoupon:                   kycCouponPool.Info,
			TotalPurchaseRewardPercent:  info.Info.TotalPurchaseRewardPercent,
			PurchaseRewardChainLevels:   info.Info.PurchaseRewardChainLevels,
			LevelPurchaseRewardPercent:  info.Info.LevelPurchaseRewardPercent,
		},
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAgencySettingDetailByAppRequest) (*npool.GetAgencySettingDetailByAppResponse, error) {
	info, err := agencysetting.GetByApp(ctx, &npool.GetAgencySettingByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get agency setting: %v", err)
	}

	registrationCouponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
		ID: info.Info.RegistrationCouponID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get agency setting registration coupon pool: %v", err)
	}

	kycCouponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
		ID: info.Info.KycCouponID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get agency setting kyc coupon pool: %v", err)
	}

	return &npool.GetAgencySettingDetailByAppResponse{
		Info: &npool.AgencySettingDetail{
			ID:                          info.Info.ID,
			AppID:                       info.Info.AppID,
			GoodID:                      info.Info.GoodID,
			RegistrationRewardThreshold: info.Info.RegistrationRewardThreshold,
			KycRewardThreshold:          info.Info.KycRewardThreshold,
			RegistrationCoupon:          registrationCouponPool.Info,
			KycCoupon:                   kycCouponPool.Info,
			TotalPurchaseRewardPercent:  info.Info.TotalPurchaseRewardPercent,
			PurchaseRewardChainLevels:   info.Info.PurchaseRewardChainLevels,
			LevelPurchaseRewardPercent:  info.Info.LevelPurchaseRewardPercent,
		},
	}, nil
}
