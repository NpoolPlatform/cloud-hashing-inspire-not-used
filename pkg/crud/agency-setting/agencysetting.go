package agencysetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/agencysetting"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAgencySetting(info *npool.AgencySetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if info.GetTotalPurchaseRewardPercent() >= 100 {
		return xerrors.Errorf("too much purchase reward")
	}
	return nil
}

func dbRowToAgencySetting(row *ent.AgencySetting) *npool.AgencySetting {
	return &npool.AgencySetting{
		ID:                          row.ID.String(),
		AppID:                       row.AppID.String(),
		RegistrationRewardThreshold: row.RegistrationRewardThreshold,
		RegistrationCouponID:        row.RegistrationCouponID.String(),
		KycRewardThreshold:          row.KycRewardThreshold,
		KycCouponID:                 row.KycCouponID.String(),
		TotalPurchaseRewardPercent:  row.TotalPurchaseRewardPercent,
		PurchaseRewardChainLevels:   row.PurchaseRewardChainLevels,
		LevelPurchaseRewardPercent:  row.LevelPurchaseRewardPercent,
	}
}

func Create(ctx context.Context, in *npool.CreateAgencySettingRequest) (*npool.CreateAgencySettingResponse, error) {
	if err := validateAgencySetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AgencySetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetRegistrationRewardThreshold(in.GetInfo().GetRegistrationRewardThreshold()).
		SetRegistrationCouponID(uuid.MustParse(in.GetInfo().GetRegistrationCouponID())).
		SetKycRewardThreshold(in.GetInfo().GetKycRewardThreshold()).
		SetKycCouponID(uuid.MustParse(in.GetInfo().GetKycCouponID())).
		SetTotalPurchaseRewardPercent(in.GetInfo().GetTotalPurchaseRewardPercent()).
		SetPurchaseRewardChainLevels(in.GetInfo().GetPurchaseRewardChainLevels()).
		SetLevelPurchaseRewardPercent(in.GetInfo().GetLevelPurchaseRewardPercent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create agency setting: %v", err)
	}

	return &npool.CreateAgencySettingResponse{
		Info: dbRowToAgencySetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAgencySettingRequest) (*npool.GetAgencySettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AgencySetting.
		Query().
		Where(
			agencysetting.And(
				agencysetting.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query agency setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty agency setting")
	}

	return &npool.GetAgencySettingResponse{
		Info: dbRowToAgencySetting(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAgencySettingByAppRequest) (*npool.GetAgencySettingByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AgencySetting.
		Query().
		Where(
			agencysetting.And(
				agencysetting.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query agency setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty agency setting")
	}

	return &npool.GetAgencySettingByAppResponse{
		Info: dbRowToAgencySetting(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAgencySettingRequest) (*npool.UpdateAgencySettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	if err := validateAgencySetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AgencySetting.
		UpdateOneID(id).
		SetRegistrationRewardThreshold(in.GetInfo().GetRegistrationRewardThreshold()).
		SetRegistrationCouponID(uuid.MustParse(in.GetInfo().GetRegistrationCouponID())).
		SetKycRewardThreshold(in.GetInfo().GetKycRewardThreshold()).
		SetKycCouponID(uuid.MustParse(in.GetInfo().GetKycCouponID())).
		SetTotalPurchaseRewardPercent(in.GetInfo().GetTotalPurchaseRewardPercent()).
		SetPurchaseRewardChainLevels(in.GetInfo().GetPurchaseRewardChainLevels()).
		SetLevelPurchaseRewardPercent(in.GetInfo().GetLevelPurchaseRewardPercent()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update agency setting: %v", err)
	}

	return &npool.UpdateAgencySettingResponse{
		Info: dbRowToAgencySetting(info),
	}, nil
}
