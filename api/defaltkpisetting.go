// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"
	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/default-kpi-setting" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDefaultKpiSetting(ctx context.Context, in *npool.CreateDefaultKpiSettingRequest) (*npool.CreateDefaultKpiSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create default kpi setting error: %w", err)
		return &npool.CreateDefaultKpiSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetDefaultKpiSetting(ctx context.Context, in *npool.GetDefaultKpiSettingRequest) (*npool.GetDefaultKpiSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create default kpi setting error: %w", err)
		return &npool.GetDefaultKpiSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetDefaultKpiSettingByAppGood(ctx context.Context, in *npool.GetDefaultKpiSettingByAppGoodRequest) (*npool.GetDefaultKpiSettingByAppGoodResponse, error) {
	resp, err := crud.GetByAppGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get default kpi setting by app good error: %w", err)
		return &npool.GetDefaultKpiSettingByAppGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateDefaultKpiSetting(ctx context.Context, in *npool.UpdateDefaultKpiSettingRequest) (*npool.UpdateDefaultKpiSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update default kpi setting error: %w", err)
		return &npool.UpdateDefaultKpiSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
