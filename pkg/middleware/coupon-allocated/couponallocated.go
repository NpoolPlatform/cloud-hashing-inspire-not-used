package couponallocated

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-allocated" //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"      //nolint

	"golang.org/x/xerrors"
)

func constructDetail(allocated *npool.CouponAllocated, pool *npool.CouponPool) *npool.CouponAllocatedDetail {
	return &npool.CouponAllocatedDetail{
		ID:     allocated.ID,
		AppID:  allocated.AppID,
		UserID: allocated.UserID,
		Coupon: pool,
	}
}

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
		Info: constructDetail(info.Info, couponPool.Info),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetCouponsAllocatedDetailByAppRequest) (*npool.GetCouponsAllocatedDetailByAppResponse, error) {
	resp, err := couponallocated.GetByApp(ctx, &npool.GetCouponsAllocatedByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get coupon allocated: %v", err)
	}

	details := []*npool.CouponAllocatedDetail{}
	for _, info := range resp.Infos {
		couponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
			ID: info.CouponID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get coupon allocated coupon pool: %v", err)
		}

		details = append(details, constructDetail(info, couponPool.Info))
	}

	return &npool.GetCouponsAllocatedDetailByAppResponse{
		Infos: details,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetCouponsAllocatedDetailByAppUserRequest) (*npool.GetCouponsAllocatedDetailByAppUserResponse, error) {
	resp, err := couponallocated.GetByAppUser(ctx, &npool.GetCouponsAllocatedByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get coupon allocated: %v", err)
	}

	details := []*npool.CouponAllocatedDetail{}
	for _, info := range resp.Infos {
		couponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
			ID: info.CouponID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get coupon allocated coupon pool: %v", err)
		}

		details = append(details, constructDetail(info, couponPool.Info))
	}

	return &npool.GetCouponsAllocatedDetailByAppUserResponse{
		Infos: details,
	}, nil
}
