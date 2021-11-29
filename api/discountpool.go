// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"
	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/discount-pool" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDiscountPool(ctx context.Context, in *npool.CreateDiscountPoolRequest) (*npool.CreateDiscountPoolResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create discount pool error: %w", err)
		return &npool.CreateDiscountPoolResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateDiscountPool(ctx context.Context, in *npool.UpdateDiscountPoolRequest) (*npool.UpdateDiscountPoolResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update discount pool error: %w", err)
		return &npool.UpdateDiscountPoolResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetDiscountPool(ctx context.Context, in *npool.GetDiscountPoolRequest) (*npool.GetDiscountPoolResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get discount pool error: %w", err)
		return &npool.GetDiscountPoolResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetDiscountPoolsByApp(ctx context.Context, in *npool.GetDiscountPoolsByAppRequest) (*npool.GetDiscountPoolsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get discount pool by app error: %w", err)
		return &npool.GetDiscountPoolsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetDiscountPoolsByAppReleaser(ctx context.Context, in *npool.GetDiscountPoolsByAppReleaserRequest) (*npool.GetDiscountPoolsByAppReleaserResponse, error) {
	resp, err := crud.GetByAppReleaser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get discount pool by app releaser error: %w", err)
		return &npool.GetDiscountPoolsByAppReleaserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
