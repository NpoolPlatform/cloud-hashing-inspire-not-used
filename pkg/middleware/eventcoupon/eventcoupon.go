package api

import (
	"context"

	activitycrud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/activity"
	eventcouponcrud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/eventcoupon"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"golang.org/x/xerrors"
)

func Create(ctx context.Context, in *npool.CreateEventCouponRequest) (*npool.CreateEventCouponResponse, error) {
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
