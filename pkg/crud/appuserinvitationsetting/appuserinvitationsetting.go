package appuserinvitationsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appuserinvitationsetting"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppUserInvitationSetting(info *npool.AppUserInvitationSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	return nil
}

func dbRowToAppUserInvitationSetting(row *ent.AppUserInvitationSetting) *npool.AppUserInvitationSetting {
	return &npool.AppUserInvitationSetting{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		Count:      row.Count,
		Discount:   row.Discount,
		Title:      row.Title,
		BadgeLarge: row.BadgeLarge,
		BadgeSmall: row.BadgeSmall,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserInvitationSettingRequest) (*npool.CreateAppUserInvitationSettingResponse, error) {
	if err := validateAppUserInvitationSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserInvitationSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetCount(in.GetInfo().GetCount()).
		SetDiscount(in.GetInfo().GetDiscount()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app user invitation setting: %v", err)
	}

	return &npool.CreateAppUserInvitationSettingResponse{
		Info: dbRowToAppUserInvitationSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserInvitationSettingRequest) (*npool.GetAppUserInvitationSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserInvitationSetting.
		Query().
		Where(
			appuserinvitationsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user invitation setting: %v", err)
	}

	var setting *npool.AppUserInvitationSetting
	for _, info := range infos {
		setting = dbRowToAppUserInvitationSetting(info)
		break
	}

	return &npool.GetAppUserInvitationSettingResponse{
		Info: setting,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppUserInvitationSettingsByAppRequest) (*npool.GetAppUserInvitationSettingsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUserInvitationSetting.
		Query().
		Where(
			appuserinvitationsetting.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user invitation setting: %v", err)
	}

	settings := []*npool.AppUserInvitationSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppUserInvitationSetting(info))
	}

	return &npool.GetAppUserInvitationSettingsByAppResponse{
		Infos: settings,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAppUserInvitationSettingsByAppUserRequest) (*npool.GetAppUserInvitationSettingsByAppUserResponse, error) {
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
		AppUserInvitationSetting.
		Query().
		Where(
			appuserinvitationsetting.And(
				appuserinvitationsetting.AppID(appID),
				appuserinvitationsetting.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user invitation setting: %v", err)
	}

	settings := []*npool.AppUserInvitationSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppUserInvitationSetting(info))
	}

	return &npool.GetAppUserInvitationSettingsByAppUserResponse{
		Infos: settings,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserInvitationSettingRequest) (*npool.UpdateAppUserInvitationSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	if err := validateAppUserInvitationSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUserInvitationSetting.
		UpdateOneID(id).
		SetCount(in.GetInfo().GetCount()).
		SetDiscount(in.GetInfo().GetDiscount()).
		SetTitle(in.GetInfo().GetTitle()).
		SetBadgeLarge(in.GetInfo().GetBadgeLarge()).
		SetBadgeSmall(in.GetInfo().GetBadgeSmall()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app user invitation setting: %v", err)
	}

	return &npool.UpdateAppUserInvitationSettingResponse{
		Info: dbRowToAppUserInvitationSetting(info),
	}, nil
}
