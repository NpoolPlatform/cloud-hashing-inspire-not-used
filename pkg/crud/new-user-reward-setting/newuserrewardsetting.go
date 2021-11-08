package newuserrewardsetting

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateNewUserRewardSetting(info *npool.NewUserRewardSetting) error {
	if _, err := uuid.Parse(info.GetRegistrationCouponID()); err != nil {
		return xerrors.Errorf("invalid registration coupon id: %v", err)
	}
	if _, err := uuid.Parse(info.GetKycCouponID()); err != nil {
		return xerrors.Errorf("invalid kyc coupon id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToNewUserRewardSetting(row *ent.NewUserRewardSetting) *npool.NewUserRewardSetting {
	return &npool.NewUserRewardSetting{
		AppID:                row.AppID.String(),
		RegistrationCouponID: row.RegistrationCouponID.String(),
		KycCouponID:          row.KycCouponID.String(),
		ID:                   row.ID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateNewUserRewardSettingRequest) (*npool.CreateNewUserRewardSettingResponse, error) {
	if err := validateNewUserRewardSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		NewUserRewardSetting.
		Create().
		SetRegistrationCouponID(uuid.MustParse(in.GetInfo().GetRegistrationCouponID())).
		SetKycCouponID(uuid.MustParse(in.GetInfo().GetKycCouponID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create new user reward setting: %v", err)
	}

	return &npool.CreateNewUserRewardSettingResponse{
		Info: dbRowToNewUserRewardSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetNewUserRewardSettingRequest) (*npool.GetNewUserRewardSettingResponse, error) {
	return nil, nil
}

func GetByApp(ctx context.Context, in *npool.GetNewUserRewardSettingByAppRequest) (*npool.GetNewUserRewardSettingByAppResponse, error) {
	return nil, nil
}

func Update(ctx context.Context, in *npool.UpdateNewUserRewardSettingRequest) (*npool.UpdateNewUserRewardSettingResponse, error) {
	return nil, nil
}
