package appinvitationsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appinvitationsetting"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppInvitationSetting(info *npool.AppInvitationSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToAppInvitationSetting(row *ent.AppInvitationSetting) *npool.AppInvitationSetting {
	return &npool.AppInvitationSetting{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		Count:      row.Count,
		Discount:   row.Discount,
		Title:      row.Title,
		BadgeLarge: row.BadgeLarge,
		BadgeSmall: row.BadgeSmall,
	}
}

func Create(ctx context.Context, in *npool.CreateAppInvitationSettingRequest) (*npool.CreateAppInvitationSettingResponse, error) {
	if err := validateAppInvitationSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppInvitationSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetCount(in.GetInfo().GetCount()).
		SetDiscount(in.GetInfo().GetDiscount()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app invitation setting: %v", err)
	}

	return &npool.CreateAppInvitationSettingResponse{
		Info: dbRowToAppInvitationSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppInvitationSettingRequest) (*npool.GetAppInvitationSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppInvitationSetting.
		Query().
		Where(
			appinvitationsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app invitation setting: %v", err)
	}

	var setting *npool.AppInvitationSetting
	for _, info := range infos {
		setting = dbRowToAppInvitationSetting(info)
		break
	}

	return &npool.GetAppInvitationSettingResponse{
		Info: setting,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppInvitationSettingsByAppRequest) (*npool.GetAppInvitationSettingsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppInvitationSetting.
		Query().
		Where(
			appinvitationsetting.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app invitation setting: %v", err)
	}

	settings := []*npool.AppInvitationSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppInvitationSetting(info))
	}

	return &npool.GetAppInvitationSettingsByAppResponse{
		Infos: settings,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppInvitationSettingRequest) (*npool.UpdateAppInvitationSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	if err := validateAppInvitationSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppInvitationSetting.
		UpdateOneID(id).
		SetCount(in.GetInfo().GetCount()).
		SetDiscount(in.GetInfo().GetDiscount()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app invitation setting: %v", err)
	}

	return &npool.UpdateAppInvitationSettingResponse{
		Info: dbRowToAppInvitationSetting(info),
	}, nil
}
