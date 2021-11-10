package couponallocated

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-allocated" //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"      //nolint

	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetCouponAllocatedDetailRequest) (*npool.GetCouponAllocatedDetailResponse, error) {
	info, err := couponallocated.Get(ctx, &npool.GetCouponAllocatedRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get coupon allocated: %v", err)
	}

	couponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
		ID: info.Info.CouponID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get coupon allocated coupon pool: %v", err)
	}

	return &npool.GetCouponAllocatedDetailResponse{
		Info: &npool.CouponAllocatedDetail{
			ID:     info.Info.ID,
			AppID:  info.Info.AppID,
			UserID: info.Info.UserID,
			Used:   info.Info.Used,
			Coupon: couponPool.Info,
		},
	}, nil
}
