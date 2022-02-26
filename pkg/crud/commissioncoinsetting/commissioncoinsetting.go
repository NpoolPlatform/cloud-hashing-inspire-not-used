package commissioncoinsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateCommissionCoinSetting(info *npool.CommissionCoinSetting) error {
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	return nil
}

func dbRowToCommissionCoinSetting(row *ent.CommissionCoinSetting) *npool.CommissionCoinSetting {
	return &npool.CommissionCoinSetting{
		ID:         row.ID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Using:      row.Using,
	}
}

func Create(ctx context.Context, in *npool.CreateCommissionCoinSettingRequest) (*npool.CreateCommissionCoinSettingResponse, error) {
	if err := validateCommissionCoinSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		CommissionCoinSetting.
		Create().
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create commission coin setting: %v", err)
	}

	return &npool.CreateCommissionCoinSettingResponse{
		Info: dbRowToCommissionCoinSetting(info),
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetCommissionCoinSettingsRequest) (*npool.GetCommissionCoinSettingsResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		CommissionCoinSetting.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query commission coin setting: %v", err)
	}

	settings := []*npool.CommissionCoinSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToCommissionCoinSetting(info))
	}

	return &npool.GetCommissionCoinSettingsResponse{
		Infos: settings,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateCommissionCoinSettingRequest) (*npool.UpdateCommissionCoinSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	if err := validateCommissionCoinSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		CommissionCoinSetting.
		UpdateOneID(id).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetUsing(in.GetInfo().GetUsing()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update commission coin setting: %v", err)
	}

	return &npool.UpdateCommissionCoinSettingResponse{
		Info: dbRowToCommissionCoinSetting(info),
	}, nil
}
