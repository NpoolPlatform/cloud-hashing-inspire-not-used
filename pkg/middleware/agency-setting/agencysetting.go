package agencysetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/agency-setting" //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"    //nolint

	"golang.org/x/xerrors"
)

func constructDetail(setting *npool.AgencySetting, regCouponPool, kycCouponPool *npool.CouponPool) *npool.AgencySettingDetail {
	return &npool.AgencySettingDetail{
		Setting:            setting,
		RegistrationCoupon: regCouponPool,
		KycCoupon:          kycCouponPool,
	}
}

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
		Info: constructDetail(info.Info, registrationCouponPool.Info, kycCouponPool.Info),
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
		Info: constructDetail(info.Info, registrationCouponPool.Info, kycCouponPool.Info),
	}, nil
}
