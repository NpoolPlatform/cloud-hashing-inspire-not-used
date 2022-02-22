package api

import (
	"context"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"
	activitycrud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/activity"
	eventcouponcrud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/eventcoupon"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"golang.org/x/xerrors"
)

func validateEvent(event string) error {
	switch event {
	case constant.EventSharing:
	case constant.EventInvitationRegisteration:
	case constant.EventInvitationPurchase:
	case constant.EventRegister:
	case constant.EventKYCAuthenticate:
	case constant.EventTotalAmount:
	case constant.EventSingleAmount:
	default:
		return xerrors.Errorf("invalid event")
	}
	return nil
}

func Create(ctx context.Context, in *npool.CreateEventCouponRequest) (*npool.CreateEventCouponResponse, error) {
	if err := validateEvent(in.GetInfo().GetEvent()); err != nil {
		return nil, xerrors.Errorf("invalid event: %v", err)
	}

	resp, err := activitycrud.Get(ctx, &npool.GetActivityRequest{
		ID: in.GetInfo().GetActivityID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get activity: %v", err)
	}
	if resp.Info == nil {
		return nil, xerrors.Errorf("fail get activity")
	}

	return eventcouponcrud.Create(ctx, in)
}

func Update(ctx context.Context, in *npool.UpdateEventCouponRequest) (*npool.UpdateEventCouponResponse, error) {
	if err := validateEvent(in.GetInfo().GetEvent()); err != nil {
		return nil, xerrors.Errorf("invalid event: %v", err)
	}

	resp, err := activitycrud.Get(ctx, &npool.GetActivityRequest{
		ID: in.GetInfo().GetActivityID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get activity: %v", err)
	}
	if resp.Info == nil {
		return nil, xerrors.Errorf("fail get activity")
	}

	return eventcouponcrud.Update(ctx, in)
}
