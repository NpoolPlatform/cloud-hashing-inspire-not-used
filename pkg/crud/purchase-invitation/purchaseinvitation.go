package purchaseinvitation

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/purchaseinvitation"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validatePurchaseInvitation(info *npool.PurchaseInvitation) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	if _, err := uuid.Parse(info.GetInvitationCodeID()); err != nil {
		return xerrors.Errorf("invalid invitation code id: %v", err)
	}
	return nil
}

func dbRowToPurchaseInvitation(row *ent.PurchaseInvitation) *npool.PurchaseInvitation {
	return &npool.PurchaseInvitation{
		ID:               row.ID.String(),
		AppID:            row.AppID.String(),
		OrderID:          row.OrderID.String(),
		InvitationCodeID: row.InvitationCodeID.String(),
		Fulfilled:        row.Fulfilled,
	}
}

func Create(ctx context.Context, in *npool.CreatePurchaseInvitationRequest) (*npool.CreatePurchaseInvitationResponse, error) {
	if err := validatePurchaseInvitation(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		PurchaseInvitation.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetInvitationCodeID(uuid.MustParse(in.GetInfo().GetInvitationCodeID())).
		SetFulfilled(false).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create purchase invitation: %v", err)
	}

	return &npool.CreatePurchaseInvitationResponse{
		Info: dbRowToPurchaseInvitation(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdatePurchaseInvitationRequest) (*npool.UpdatePurchaseInvitationResponse, error) {
	if err := validatePurchaseInvitation(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	info, err := db.Client().
		PurchaseInvitation.
		UpdateOneID(id).
		SetFulfilled(in.GetInfo().GetFulfilled()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update purchase invitation: %v", err)
	}

	return &npool.UpdatePurchaseInvitationResponse{
		Info: dbRowToPurchaseInvitation(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetPurchaseInvitationRequest) (*npool.GetPurchaseInvitationResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	infos, err := db.Client().
		PurchaseInvitation.
		Query().
		Where(
			purchaseinvitation.And(
				purchaseinvitation.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query purchase invitation: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty purchase invitation")
	}

	return &npool.GetPurchaseInvitationResponse{
		Info: dbRowToPurchaseInvitation(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetPurchaseInvitationsByAppRequest) (*npool.GetPurchaseInvitationsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	infos, err := db.Client().
		PurchaseInvitation.
		Query().
		Where(
			purchaseinvitation.And(
				purchaseinvitation.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query purchase invitation: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty purchase invitation")
	}

	invitations := []*npool.PurchaseInvitation{}
	for _, info := range infos {
		invitations = append(invitations, dbRowToPurchaseInvitation(info))
	}

	return &npool.GetPurchaseInvitationsByAppResponse{
		Infos: invitations,
	}, nil
}

func GetByAppOrder(ctx context.Context, in *npool.GetPurchaseInvitationByAppOrderRequest) (*npool.GetPurchaseInvitationByAppOrderResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	orderID, err := uuid.Parse(in.GetOrderID())
	if err != nil {
		return nil, xerrors.Errorf("invalid order id: %v", err)
	}

	infos, err := db.Client().
		PurchaseInvitation.
		Query().
		Where(
			purchaseinvitation.And(
				purchaseinvitation.AppID(appID),
				purchaseinvitation.OrderID(orderID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query purchase invitation: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty purchase invitation")
	}

	return &npool.GetPurchaseInvitationByAppOrderResponse{
		Info: dbRowToPurchaseInvitation(infos[0]),
	}, nil
}
