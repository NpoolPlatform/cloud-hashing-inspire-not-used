package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-allocated"     //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/middleware/coupon-allocated" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCouponAllocated(ctx context.Context, in *npool.CreateCouponAllocatedRequest) (*npool.CreateCouponAllocatedResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create coupon allocated error: %w", err)
		return &npool.CreateCouponAllocatedResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponAllocated(ctx context.Context, in *npool.GetCouponAllocatedRequest) (*npool.GetCouponAllocatedResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon allocated error: %w", err)
		return &npool.GetCouponAllocatedResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponsAllocatedByApp(ctx context.Context, in *npool.GetCouponsAllocatedByAppRequest) (*npool.GetCouponsAllocatedByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon allocated by app error: %w", err)
		return &npool.GetCouponsAllocatedByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponsAllocatedByAppUser(ctx context.Context, in *npool.GetCouponsAllocatedByAppUserRequest) (*npool.GetCouponsAllocatedByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon allocated by app user error: %w", err)
		return &npool.GetCouponsAllocatedByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateCouponAllocated(ctx context.Context, in *npool.UpdateCouponAllocatedRequest) (*npool.UpdateCouponAllocatedResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update coupon allocated error: %w", err)
		return &npool.UpdateCouponAllocatedResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponAllocatedDetail(ctx context.Context, in *npool.GetCouponAllocatedDetailRequest) (*npool.GetCouponAllocatedDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon allocated detail error: %w", err)
		return &npool.GetCouponAllocatedDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponsAllocatedDetailByApp(ctx context.Context, in *npool.GetCouponsAllocatedDetailByAppRequest) (*npool.GetCouponsAllocatedDetailByAppResponse, error) {
	resp, err := mw.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("app coupon allocated detail by app error: %w", err)
		return &npool.GetCouponsAllocatedDetailByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponsAllocatedDetailByAppUser(ctx context.Context, in *npool.GetCouponsAllocatedDetailByAppUserRequest) (*npool.GetCouponsAllocatedDetailByAppUserResponse, error) {
	resp, err := mw.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon allocated detail by app user error: %w", err)
		return &npool.GetCouponsAllocatedDetailByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
