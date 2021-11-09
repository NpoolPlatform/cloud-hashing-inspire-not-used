package defaultkpisetting

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/defaultkpisetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateDefaultKpiSetting(info *npool.DefaultKpiSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	return nil
}

func dbRowToDefaultKpiSetting(row *ent.DefaultKpiSetting) *npool.DefaultKpiSetting {
	return &npool.DefaultKpiSetting{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		GoodID:  row.GoodID.String(),
		Amount:  price.DBPriceToVisualPrice(row.Amount),
		Percent: row.Percent,
	}
}

func Create(ctx context.Context, in *npool.CreateDefaultKpiSettingRequest) (*npool.CreateDefaultKpiSettingResponse, error) {
	if err := validateDefaultKpiSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		DefaultKpiSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create default kpi setting: %v", err)
	}

	return &npool.CreateDefaultKpiSettingResponse{
		Info: dbRowToDefaultKpiSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetDefaultKpiSettingRequest) (*npool.GetDefaultKpiSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		DefaultKpiSetting.
		Query().
		Where(
			defaultkpisetting.And(
				defaultkpisetting.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query default kpi setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty default kpi setting")
	}

	return &npool.GetDefaultKpiSettingResponse{
		Info: dbRowToDefaultKpiSetting(infos[0]),
	}, nil
}

func GetByAppGood(ctx context.Context, in *npool.GetDefaultKpiSettingByAppGoodRequest) (*npool.GetDefaultKpiSettingByAppGoodResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	infos, err := db.Client().
		DefaultKpiSetting.
		Query().
		Where(
			defaultkpisetting.And(
				defaultkpisetting.AppID(appID),
				defaultkpisetting.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query default kpi setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty default kpi setting")
	}

	return &npool.GetDefaultKpiSettingByAppGoodResponse{
		Info: dbRowToDefaultKpiSetting(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateDefaultKpiSettingRequest) (*npool.UpdateDefaultKpiSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	if err := validateDefaultKpiSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		DefaultKpiSetting.
		UpdateOneID(id).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update default kpi setting: %v", err)
	}

	return &npool.UpdateDefaultKpiSettingResponse{
		Info: dbRowToDefaultKpiSetting(info),
	}, nil
}
