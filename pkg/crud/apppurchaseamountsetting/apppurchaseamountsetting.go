package apppurchaseamountsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/apppurchaseamountsetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppPurchaseAmountSetting(info *npool.AppPurchaseAmountSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToAppPurchaseAmountSetting(row *ent.AppPurchaseAmountSetting) *npool.AppPurchaseAmountSetting {
	return &npool.AppPurchaseAmountSetting{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		GoodID:     row.GoodID.String(),
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

func Create(ctx context.Context, in *npool.CreateAppPurchaseAmountSettingRequest) (*npool.CreateAppPurchaseAmountSettingResponse, error) {
	if err := validateAppPurchaseAmountSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	userID, err := uuid.Parse(in.GetInfo().GetUserID())
	if err != nil {
		userID = uuid.UUID{}
	}

	goodID, err := uuid.Parse(in.GetInfo().GetGoodID())
	if err != nil {
		goodID = uuid.UUID{}
	}

	info, err := cli.
		AppPurchaseAmountSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(userID).
		SetGoodID(goodID).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(0).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app purchase amount setting: %v", err)
	}

	return &npool.CreateAppPurchaseAmountSettingResponse{
		Info: dbRowToAppPurchaseAmountSetting(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppPurchaseAmountSettingRequest) (*npool.UpdateAppPurchaseAmountSettingResponse, error) {
	if err := validateAppPurchaseAmountSetting(in.GetInfo()); err != nil {
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
		AppPurchaseAmountSetting.
		UpdateOneID(id).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		SetEnd(in.GetInfo().GetEnd()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app purchase amount setting: %v", err)
	}

	return &npool.UpdateAppPurchaseAmountSettingResponse{
		Info: dbRowToAppPurchaseAmountSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppPurchaseAmountSettingRequest) (*npool.GetAppPurchaseAmountSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppPurchaseAmountSetting.
		Query().
		Where(
			apppurchaseamountsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app purchase amount setting: %v", err)
	}

	var setting *npool.AppPurchaseAmountSetting
	for _, info := range infos {
		setting = dbRowToAppPurchaseAmountSetting(info)
		break
	}

	return &npool.GetAppPurchaseAmountSettingResponse{
		Info: setting,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppPurchaseAmountSettingsByAppRequest) (*npool.GetAppPurchaseAmountSettingsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppPurchaseAmountSetting.
		Query().
		Where(
			apppurchaseamountsetting.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app purchase amount setting: %v", err)
	}

	settings := []*npool.AppPurchaseAmountSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppPurchaseAmountSetting(info))
	}

	return &npool.GetAppPurchaseAmountSettingsByAppResponse{
		Infos: settings,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAppPurchaseAmountSettingsByAppUserRequest) (*npool.GetAppPurchaseAmountSettingsByAppUserResponse, error) {
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
		AppPurchaseAmountSetting.
		Query().
		Where(
			apppurchaseamountsetting.And(
				apppurchaseamountsetting.AppID(appID),
				apppurchaseamountsetting.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app purchase amount setting: %v", err)
	}

	settings := []*npool.AppPurchaseAmountSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppPurchaseAmountSetting(info))
	}

	return &npool.GetAppPurchaseAmountSettingsByAppUserResponse{
		Infos: settings,
	}, nil
}
