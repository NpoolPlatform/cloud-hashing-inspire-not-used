package appcommissionsetting

import (
	"context"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcommissionsetting"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppCommissionSetting(info *npool.AppCommissionSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if info.GetType() != constant.CommissionByAmount {
		return xerrors.Errorf("invalid commission type")
	}
	return nil
}

func dbRowToAppCommissionSetting(row *ent.AppCommissionSetting) *npool.AppCommissionSetting {
	return &npool.AppCommissionSetting{
		ID:                 row.ID.String(),
		AppID:              row.AppID.String(),
		Type:               row.Type,
		Level:              row.Level,
		UniqueSetting:      row.UniqueSetting,
		InvitationDiscount: row.InvitationDiscount,
	}
}

func Create(ctx context.Context, in *npool.CreateAppCommissionSettingRequest) (*npool.CreateAppCommissionSettingResponse, error) {
	if err := validateAppCommissionSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppCommissionSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetType(in.GetInfo().GetType()).
		SetLevel(in.GetInfo().GetLevel()).
		SetUniqueSetting(in.GetInfo().GetUniqueSetting()).
		SetInvitationDiscount(in.GetInfo().GetInvitationDiscount()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app commission setting: %v", err)
	}

	return &npool.CreateAppCommissionSettingResponse{
		Info: dbRowToAppCommissionSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppCommissionSettingRequest) (*npool.GetAppCommissionSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppCommissionSetting.
		Query().
		Where(
			appcommissionsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app commission setting: %v", err)
	}

	var setting *npool.AppCommissionSetting
	for _, info := range infos {
		setting = dbRowToAppCommissionSetting(info)
		break
	}

	return &npool.GetAppCommissionSettingResponse{
		Info: setting,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppCommissionSettingByAppRequest) (*npool.GetAppCommissionSettingByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppCommissionSetting.
		Query().
		Where(
			appcommissionsetting.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app commission setting: %v", err)
	}

	var setting *npool.AppCommissionSetting
	for _, info := range infos {
		setting = dbRowToAppCommissionSetting(info)
		break
	}

	return &npool.GetAppCommissionSettingByAppResponse{
		Info: setting,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppCommissionSettingRequest) (*npool.UpdateAppCommissionSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	if err := validateAppCommissionSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppCommissionSetting.
		UpdateOneID(id).
		SetType(in.GetInfo().GetType()).
		SetLevel(in.GetInfo().GetLevel()).
		SetUniqueSetting(in.GetInfo().GetUniqueSetting()).
		SetInvitationDiscount(in.GetInfo().GetInvitationDiscount()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app commission setting: %v", err)
	}

	return &npool.UpdateAppCommissionSettingResponse{
		Info: dbRowToAppCommissionSetting(info),
	}, nil
}
