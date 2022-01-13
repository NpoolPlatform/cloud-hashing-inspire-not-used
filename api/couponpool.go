// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCouponPool(ctx context.Context, in *npool.CreateCouponPoolRequest) (*npool.CreateCouponPoolResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create coupon pool error: %w", err)
		return &npool.CreateCouponPoolResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateCouponPool(ctx context.Context, in *npool.UpdateCouponPoolRequest) (*npool.UpdateCouponPoolResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update coupon pool error: %w", err)
		return &npool.UpdateCouponPoolResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponPool(ctx context.Context, in *npool.GetCouponPoolRequest) (*npool.GetCouponPoolResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon pool error: %w", err)
		return &npool.GetCouponPoolResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponPoolsByApp(ctx context.Context, in *npool.GetCouponPoolsByAppRequest) (*npool.GetCouponPoolsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon pool by app error: %w", err)
		return &npool.GetCouponPoolsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCouponPoolsByAppReleaser(ctx context.Context, in *npool.GetCouponPoolsByAppReleaserRequest) (*npool.GetCouponPoolsByAppReleaserResponse, error) {
	resp, err := crud.GetByAppReleaser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get coupon pool by app releaser error: %w", err)
		return &npool.GetCouponPoolsByAppReleaserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
