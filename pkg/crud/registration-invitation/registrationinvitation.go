package registrationinvitation

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateRegistrationInvitation(info *npool.RegistrationInvitation) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetInviterID()); err != nil {
		return xerrors.Errorf("invalid inviter id: %v", err)
	}
	if _, err := uuid.Parse(info.GetInviteeID()); err != nil {
		return xerrors.Errorf("invalid invitee id: %v", err)
	}
	return nil
}

func dbRowToRegistrationInvitation(row *ent.RegistrationInvitation) *npool.RegistrationInvitation {
	return &npool.RegistrationInvitation{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		InviterID: row.InviterID.String(),
		InviteeID: row.InviteeID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateRegistrationInvitationRequest) (*npool.CreateRegistrationInvitationResponse, error) {
	if err := validateRegistrationInvitation(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		RegistrationInvitation.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetInviterID(uuid.MustParse(in.GetInfo().GetInviterID())).
		SetInviteeID(uuid.MustParse(in.GetInfo().GetInviteeID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create registration invitation: %v", err)
	}

	return &npool.CreateRegistrationInvitationResponse{
		Info: dbRowToRegistrationInvitation(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateRegistrationInvitationRequest) (*npool.UpdateRegistrationInvitationResponse, error) {
	if err := validateRegistrationInvitation(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	info, err := db.Client().
		RegistrationInvitation.
		UpdateOneID(id).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update registration invitation: %v", err)
	}

	return &npool.UpdateRegistrationInvitationResponse{
		Info: dbRowToRegistrationInvitation(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetRegistrationInvitationRequest) (*npool.GetRegistrationInvitationResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		RegistrationInvitation.
		Query().
		Where(
			registrationinvitation.And(
				registrationinvitation.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query registration invitation: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty registration invitation")
	}

	return &npool.GetRegistrationInvitationResponse{
		Info: dbRowToRegistrationInvitation(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetRegistrationInvitationsByAppRequest) (*npool.GetRegistrationInvitationsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	infos, err := db.Client().
		RegistrationInvitation.
		Query().
		Where(
			registrationinvitation.And(
				registrationinvitation.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query registration invitation: %v", err)
	}

	invitations := []*npool.RegistrationInvitation{}
	for _, info := range infos {
		invitations = append(invitations, dbRowToRegistrationInvitation(info))
	}

	return &npool.GetRegistrationInvitationsByAppResponse{
		Infos: invitations,
	}, nil
}

func GetByAppInviter(ctx context.Context, in *npool.GetRegistrationInvitationsByAppInviterRequest) (*npool.GetRegistrationInvitationsByAppInviterResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	inviterID, err := uuid.Parse(in.GetInviterID())
	if err != nil {
		return nil, xerrors.Errorf("invalid inviter id: %v", err)
	}

	infos, err := db.Client().
		RegistrationInvitation.
		Query().
		Where(
			registrationinvitation.And(
				registrationinvitation.AppID(appID),
				registrationinvitation.InviterID(inviterID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query registration invitation: %v", err)
	}

	invitations := []*npool.RegistrationInvitation{}
	for _, info := range infos {
		invitations = append(invitations, dbRowToRegistrationInvitation(info))
	}

	return &npool.GetRegistrationInvitationsByAppInviterResponse{
		Infos: invitations,
	}, nil
}

func GetByAppInvitee(ctx context.Context, in *npool.GetRegistrationInvitationByAppInviteeRequest) (*npool.GetRegistrationInvitationByAppInviteeResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	inviteeID, err := uuid.Parse(in.GetInviteeID())
	if err != nil {
		return nil, xerrors.Errorf("invalid invitee id: %v", err)
	}

	infos, err := db.Client().
		RegistrationInvitation.
		Query().
		Where(
			registrationinvitation.And(
				registrationinvitation.AppID(appID),
				registrationinvitation.InviteeID(inviteeID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query registration invitation: %v", err)
	}

	return &npool.GetRegistrationInvitationByAppInviteeResponse{
		Info: dbRowToRegistrationInvitation(infos[0]),
	}, nil
}
