package appuserpurchaseamountsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appuserpurchaseamountsetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppUserPurchaseAmountSetting(info *npool.AppUserPurchaseAmountSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToAppUserPurchaseAmountSetting(row *ent.AppUserPurchaseAmountSetting) *npool.AppUserPurchaseAmountSetting {
	return &npool.AppUserPurchaseAmountSetting{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		Amount:     price.DBPriceToVisualPrice(row.Amount),
		Percent:    row.Percent,
		Title:      row.Title,
		BadgeLarge: row.BadgeLarge,
		BadgeSmall: row.BadgeSmall,
		Start:      row.Start,
		End:        row.End,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserPurchaseAmountSettingRequest) (*npool.CreateAppUserPurchaseAmountSettingResponse, error) {
	if err := validateAppUserPurchaseAmountSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserPurchaseAmountSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(0).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app user purchase amount setting: %v", err)
	}

	return &npool.CreateAppUserPurchaseAmountSettingResponse{
		Info: dbRowToAppUserPurchaseAmountSetting(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserPurchaseAmountSettingRequest) (*npool.UpdateAppUserPurchaseAmountSettingResponse, error) {
	if err := validateAppUserPurchaseAmountSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserPurchaseAmountSetting.
		UpdateOneID(id).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		SetEnd(in.GetInfo().GetEnd()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app user purchase amount setting: %v", err)
	}

	return &npool.UpdateAppUserPurchaseAmountSettingResponse{
		Info: dbRowToAppUserPurchaseAmountSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingRequest) (*npool.GetAppUserPurchaseAmountSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserPurchaseAmountSetting.
		Query().
		Where(
			appuserpurchaseamountsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user purchase amount setting: %v", err)
	}

	var setting *npool.AppUserPurchaseAmountSetting
	for _, info := range infos {
		setting = dbRowToAppUserPurchaseAmountSetting(info)
		break
	}

	return &npool.GetAppUserPurchaseAmountSettingResponse{
		Info: setting,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingsByAppRequest) (*npool.GetAppUserPurchaseAmountSettingsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserPurchaseAmountSetting.
		Query().
		Where(
			appuserpurchaseamountsetting.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user purchase amount setting: %v", err)
	}

	settings := []*npool.AppUserPurchaseAmountSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppUserPurchaseAmountSetting(info))
	}

	return &npool.GetAppUserPurchaseAmountSettingsByAppResponse{
		Infos: settings,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingsByAppUserRequest) (*npool.GetAppUserPurchaseAmountSettingsByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserPurchaseAmountSetting.
		Query().
		Where(
			appuserpurchaseamountsetting.And(
				appuserpurchaseamountsetting.AppID(appID),
				appuserpurchaseamountsetting.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user purchase amount setting: %v", err)
	}

	settings := []*npool.AppUserPurchaseAmountSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppUserPurchaseAmountSetting(info))
	}

	return &npool.GetAppUserPurchaseAmountSettingsByAppUserResponse{
		Infos: settings,
	}, nil
}
