package defaultkpisetting

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/defaultkpisetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"
)

func validateDefaultKpiSetting(info *npool.DefaultKpiSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return fmt.Errorf("invalid good id: %v", err)
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
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		DefaultKpiSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create default kpi setting: %v", err)
	}

	return &npool.CreateDefaultKpiSettingResponse{
		Info: dbRowToDefaultKpiSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetDefaultKpiSettingRequest) (*npool.GetDefaultKpiSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		DefaultKpiSetting.
		Query().
		Where(
			defaultkpisetting.And(
				defaultkpisetting.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query default kpi setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("empty default kpi setting")
	}

	return &npool.GetDefaultKpiSettingResponse{
		Info: dbRowToDefaultKpiSetting(infos[0]),
	}, nil
}

func GetByAppGood(ctx context.Context, in *npool.GetDefaultKpiSettingByAppGoodRequest) (*npool.GetDefaultKpiSettingByAppGoodResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, fmt.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
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
		return nil, fmt.Errorf("fail query default kpi setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("empty default kpi setting")
	}

	return &npool.GetDefaultKpiSettingByAppGoodResponse{
		Info: dbRowToDefaultKpiSetting(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateDefaultKpiSettingRequest) (*npool.UpdateDefaultKpiSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	if err := validateDefaultKpiSetting(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		DefaultKpiSetting.
		UpdateOneID(id).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPercent(in.GetInfo().GetPercent()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update default kpi setting: %v", err)
	}

	return &npool.UpdateDefaultKpiSettingResponse{
		Info: dbRowToDefaultKpiSetting(info),
	}, nil
}
