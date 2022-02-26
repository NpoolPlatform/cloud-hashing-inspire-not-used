package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/appcommissionsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppCommissionSetting(ctx context.Context, in *npool.CreateAppCommissionSettingRequest) (*npool.CreateAppCommissionSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create app commission setting error: %v", err)
		return &npool.CreateAppCommissionSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppCommissionSettingForOtherApp(ctx context.Context, in *npool.CreateAppCommissionSettingForOtherAppRequest) (*npool.CreateAppCommissionSettingForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateAppCommissionSettingRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create app commission setting error: %v", err)
		return &npool.CreateAppCommissionSettingForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppCommissionSettingForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetAppCommissionSetting(ctx context.Context, in *npool.GetAppCommissionSettingRequest) (*npool.GetAppCommissionSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app commission setting error: %v", err)
		return &npool.GetAppCommissionSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppCommissionSettingByApp(ctx context.Context, in *npool.GetAppCommissionSettingByAppRequest) (*npool.GetAppCommissionSettingByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app commission setting by app error: %v", err)
		return &npool.GetAppCommissionSettingByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppCommissionSettingByOtherApp(ctx context.Context, in *npool.GetAppCommissionSettingByOtherAppRequest) (*npool.GetAppCommissionSettingByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAppCommissionSettingByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app commission setting by app error: %v", err)
		return &npool.GetAppCommissionSettingByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppCommissionSettingByOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateAppCommissionSetting(ctx context.Context, in *npool.UpdateAppCommissionSettingRequest) (*npool.UpdateAppCommissionSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app commission setting error: %v", err)
		return &npool.UpdateAppCommissionSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
