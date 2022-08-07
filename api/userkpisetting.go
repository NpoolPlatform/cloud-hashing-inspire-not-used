//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/user-kpi-setting" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserKpiSetting(ctx context.Context, in *npool.CreateUserKpiSettingRequest) (*npool.CreateUserKpiSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user kpi setting error: %w", err)
		return &npool.CreateUserKpiSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserKpiSetting(ctx context.Context, in *npool.GetUserKpiSettingRequest) (*npool.GetUserKpiSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user kpi setting error: %w", err)
		return &npool.GetUserKpiSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserKpiSettingByAppGood(ctx context.Context, in *npool.GetUserKpiSettingByAppGoodRequest) (*npool.GetUserKpiSettingByAppGoodResponse, error) {
	resp, err := crud.GetByAppGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get user kpi setting by app good error: %w", err)
		return &npool.GetUserKpiSettingByAppGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateUserKpiSetting(ctx context.Context, in *npool.UpdateUserKpiSettingRequest) (*npool.UpdateUserKpiSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update user kpi setting error: %w", err)
		return &npool.UpdateUserKpiSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
