package appcouponsetting

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcouponsetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppCouponSetting(info *npool.AppCouponSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToAppCouponSetting(row *ent.AppCouponSetting) *npool.AppCouponSetting {
	return &npool.AppCouponSetting{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		DominationLimit: price.DBPriceToVisualPrice(row.DominationLimit),
		TotalLimit:      row.TotalLimit,
	}
}

func Create(ctx context.Context, in *npool.CreateAppCouponSettingRequest) (*npool.CreateAppCouponSettingResponse, error) {
	if err := validateAppCouponSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppCouponSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetDominationLimit(price.VisualPriceToDBPrice(in.GetInfo().GetDominationLimit())).
		SetTotalLimit(in.GetInfo().GetTotalLimit()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app coupon setting: %v", err)
	}

	return &npool.CreateAppCouponSettingResponse{
		Info: dbRowToAppCouponSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppCouponSettingRequest) (*npool.GetAppCouponSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppCouponSetting.
		Query().
		Where(
			appcouponsetting.And(
				appcouponsetting.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app coupon setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty app coupon setting")
	}

	return &npool.GetAppCouponSettingResponse{
		Info: dbRowToAppCouponSetting(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppCouponSettingByAppRequest) (*npool.GetAppCouponSettingByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppCouponSetting.
		Query().
		Where(
			appcouponsetting.And(
				appcouponsetting.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app coupon setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty app coupon setting")
	}

	return &npool.GetAppCouponSettingByAppResponse{
		Info: dbRowToAppCouponSetting(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppCouponSettingRequest) (*npool.UpdateAppCouponSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	if err := validateAppCouponSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppCouponSetting.
		UpdateOneID(id).
		SetDominationLimit(price.VisualPriceToDBPrice(in.GetInfo().GetDominationLimit())).
		SetTotalLimit(in.GetInfo().GetTotalLimit()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app coupon setting: %v", err)
	}

	return &npool.UpdateAppCouponSettingResponse{
		Info: dbRowToAppCouponSetting(info),
	}, nil
}
