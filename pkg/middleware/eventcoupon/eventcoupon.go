package api

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"
	activitycrud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/activity"
	eventcouponcrud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/eventcoupon"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"
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
		return fmt.Errorf("invalid event")
	}
	return nil
}

func Create(ctx context.Context, in *npool.CreateEventCouponRequest) (*npool.CreateEventCouponResponse, error) {
	if err := validateEvent(in.GetInfo().GetEvent()); err != nil {
		return nil, fmt.Errorf("invalid event: %v", err)
	}

	resp, err := activitycrud.Get(ctx, &npool.GetActivityRequest{
		ID: in.GetInfo().GetActivityID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get activity: %v", err)
	}
	if resp.Info == nil {
		return nil, fmt.Errorf("fail get activity")
	}

	return eventcouponcrud.Create(ctx, in)
}

func Update(ctx context.Context, in *npool.UpdateEventCouponRequest) (*npool.UpdateEventCouponResponse, error) {
	if err := validateEvent(in.GetInfo().GetEvent()); err != nil {
		return nil, fmt.Errorf("invalid event: %v", err)
	}

	resp, err := activitycrud.Get(ctx, &npool.GetActivityRequest{
		ID: in.GetInfo().GetActivityID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get activity: %v", err)
	}
	if resp.Info == nil {
		return nil, fmt.Errorf("fail get activity")
	}

	return eventcouponcrud.Update(ctx, in)
}
