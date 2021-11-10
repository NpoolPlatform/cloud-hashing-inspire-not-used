package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/agency-setting"          //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/middleware/agency-setting" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAgencySetting(ctx context.Context, in *npool.CreateAgencySettingRequest) (*npool.CreateAgencySettingResponse, error) {
	resp, err := agencysetting.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create agency setting error: %w", err)
		return &npool.CreateAgencySettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAgencySetting(ctx context.Context, in *npool.GetAgencySettingRequest) (*npool.GetAgencySettingResponse, error) {
	resp, err := agencysetting.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get agency setting error: %w", err)
		return &npool.GetAgencySettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAgencySettingByApp(ctx context.Context, in *npool.GetAgencySettingByAppRequest) (*npool.GetAgencySettingByAppResponse, error) {
	resp, err := agencysetting.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get agency setting by app error: %w", err)
		return &npool.GetAgencySettingByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateAgencySetting(ctx context.Context, in *npool.UpdateAgencySettingRequest) (*npool.UpdateAgencySettingResponse, error) {
	resp, err := agencysetting.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get agency setting error: %w", err)
		return &npool.UpdateAgencySettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAgencySettingDetail(ctx context.Context, in *npool.GetAgencySettingDetailRequest) (*npool.GetAgencySettingDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get agency setting detail error: %w", err)
		return &npool.GetAgencySettingDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAgencySettingDetailByApp(ctx context.Context, in *npool.GetAgencySettingDetailByAppRequest) (*npool.GetAgencySettingDetailByAppResponse, error) {
	resp, err := mw.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get agency setting detail by app error: %w", err)
		return &npool.GetAgencySettingDetailByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
