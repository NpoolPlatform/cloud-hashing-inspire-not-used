package couponallocated

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	couponallocated "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-allocated" //nolint
	couponpool "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"           //nolint
	discountpool "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/discount-pool"       //nolint
)

func constructDetail(allocated *npool.CouponAllocated, pool *npool.CouponPool, discount *npool.DiscountPool) *npool.CouponAllocatedDetail {
	return &npool.CouponAllocatedDetail{
		Allocated: allocated,
		Coupon:    pool,
		Discount:  discount,
	}
}

func expandInfo(ctx context.Context, info *npool.CouponAllocated) (*npool.CouponAllocatedDetail, error) {
	var coupon *npool.CouponPool
	var discount *npool.DiscountPool

	switch info.Type {
	case constant.CouponTypeCoupon:
		couponPool, err := couponpool.Get(ctx, &npool.GetCouponPoolRequest{
			ID: info.CouponID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coupon allocated coupon pool: %v", err)
		}
		coupon = couponPool.Info
	case constant.CouponTypeDiscount:
		discountPool, err := discountpool.Get(ctx, &npool.GetDiscountPoolRequest{
			ID: info.CouponID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get discount allocated coupon pool: %v", err)
		}
		discount = discountPool.Info
	}

	if discount == nil && coupon == nil {
		return nil, fmt.Errorf("invalid coupon id")
	}

	return constructDetail(info, coupon, discount), nil
}

func Get(ctx context.Context, in *npool.GetCouponAllocatedDetailRequest) (*npool.GetCouponAllocatedDetailResponse, error) {
	info, err := couponallocated.Get(ctx, &npool.GetCouponAllocatedRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coupon allocated: %v", err)
	}

	detail, err := expandInfo(ctx, info.Info)
	if err != nil {
		return nil, fmt.Errorf("fail expand info: %v", err)
	}

	return &npool.GetCouponAllocatedDetailResponse{
		Info: detail,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetCouponsAllocatedDetailByAppRequest) (*npool.GetCouponsAllocatedDetailByAppResponse, error) {
	resp, err := couponallocated.GetByApp(ctx, &npool.GetCouponsAllocatedByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coupon allocated: %v", err)
	}

	details := []*npool.CouponAllocatedDetail{}
	for _, info := range resp.Infos {
		detail, err := expandInfo(ctx, info)
		if err != nil {
			return nil, fmt.Errorf("fail expand info: %v", err)
		}

		details = append(details, detail)
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
		return nil, fmt.Errorf("fail get coupon allocated: %v", err)
	}

	details := []*npool.CouponAllocatedDetail{}
	for _, info := range resp.Infos {
		detail, err := expandInfo(ctx, info)
		if err != nil {
			return nil, fmt.Errorf("fail expand info: %v", err)
		}

		details = append(details, detail)
	}

	return &npool.GetCouponsAllocatedDetailByAppUserResponse{
		Infos: details,
	}, nil
}
