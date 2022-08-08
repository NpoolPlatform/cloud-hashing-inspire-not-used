//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/discount-pool" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDiscountPool(ctx context.Context, in *npool.CreateDiscountPoolRequest) (*npool.CreateDiscountPoolResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create discount pool error: %v", err)
		return &npool.CreateDiscountPoolResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateDiscountPoolForOtherApp(ctx context.Context, in *npool.CreateDiscountPoolForOtherAppRequest) (*npool.CreateDiscountPoolForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateDiscountPoolRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("create discount pool error: %v", err)
		return &npool.CreateDiscountPoolForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateDiscountPoolForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateDiscountPool(ctx context.Context, in *npool.UpdateDiscountPoolRequest) (*npool.UpdateDiscountPoolResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update discount pool error: %v", err)
		return &npool.UpdateDiscountPoolResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetDiscountPool(ctx context.Context, in *npool.GetDiscountPoolRequest) (*npool.GetDiscountPoolResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get discount pool error: %v", err)
		return &npool.GetDiscountPoolResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetDiscountPoolsByApp(ctx context.Context, in *npool.GetDiscountPoolsByAppRequest) (*npool.GetDiscountPoolsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get discount pool by app error: %v", err)
		return &npool.GetDiscountPoolsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetDiscountPoolsByOtherApp(ctx context.Context, in *npool.GetDiscountPoolsByOtherAppRequest) (*npool.GetDiscountPoolsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetDiscountPoolsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get discount pool by app error: %v", err)
		return &npool.GetDiscountPoolsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetDiscountPoolsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetDiscountPoolsByAppReleaser(ctx context.Context, in *npool.GetDiscountPoolsByAppReleaserRequest) (*npool.GetDiscountPoolsByAppReleaserResponse, error) {
	resp, err := crud.GetByAppReleaser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get discount pool by app releaser error: %v", err)
		return &npool.GetDiscountPoolsByAppReleaserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetDiscountPoolsByOtherAppReleaser(ctx context.Context, in *npool.GetDiscountPoolsByOtherAppReleaserRequest) (*npool.GetDiscountPoolsByOtherAppReleaserResponse, error) {
	resp, err := crud.GetByAppReleaser(ctx, &npool.GetDiscountPoolsByAppReleaserRequest{
		AppID:  in.GetTargetAppID(),
		UserID: in.GetTargetUserID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get discount pool by app releaser error: %v", err)
		return &npool.GetDiscountPoolsByOtherAppReleaserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetDiscountPoolsByOtherAppReleaserResponse{
		Infos: resp.Infos,
	}, nil
}
