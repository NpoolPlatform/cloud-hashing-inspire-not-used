package userkpisetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userkpisetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateUserKpiSetting(info *npool.UserKpiSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	return nil
}

func dbRowToUserKpiSetting(row *ent.UserKpiSetting) *npool.UserKpiSetting {
	return &npool.UserKpiSetting{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		UserID:  row.UserID.String(),
		GoodID:  row.GoodID.String(),
		Amount:  price.DBPriceToVisualPrice(row.Amount),
		Percent: row.Percent,
	}
}

func Create(ctx context.Context, in *npool.CreateUserKpiSettingRequest) (*npool.CreateUserKpiSettingResponse, error) {
	if err := validateUserKpiSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserKpiSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create default kpi setting: %v", err)
	}

	return &npool.CreateUserKpiSettingResponse{
		Info: dbRowToUserKpiSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserKpiSettingRequest) (*npool.GetUserKpiSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserKpiSetting.
		Query().
		Where(
			userkpisetting.And(
				userkpisetting.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query default kpi setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty default kpi setting")
	}

	return &npool.GetUserKpiSettingResponse{
		Info: dbRowToUserKpiSetting(infos[0]),
	}, nil
}

func GetByAppGood(ctx context.Context, in *npool.GetUserKpiSettingByAppGoodRequest) (*npool.GetUserKpiSettingByAppGoodResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserKpiSetting.
		Query().
		Where(
			userkpisetting.And(
				userkpisetting.AppID(appID),
				userkpisetting.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query default kpi setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty default kpi setting")
	}

	return &npool.GetUserKpiSettingByAppGoodResponse{
		Info: dbRowToUserKpiSetting(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserKpiSettingRequest) (*npool.UpdateUserKpiSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	if err := validateUserKpiSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserKpiSetting.
		UpdateOneID(id).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update default kpi setting: %v", err)
	}

	return &npool.UpdateUserKpiSettingResponse{
		Info: dbRowToUserKpiSetting(info),
	}, nil
}
