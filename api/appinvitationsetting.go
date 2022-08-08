//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/appinvitationsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppInvitationSetting(ctx context.Context, in *npool.CreateAppInvitationSettingRequest) (*npool.CreateAppInvitationSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create app invitation setting error: %v", err)
		return &npool.CreateAppInvitationSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppInvitationSettingForOtherApp(ctx context.Context, in *npool.CreateAppInvitationSettingForOtherAppRequest) (*npool.CreateAppInvitationSettingForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateAppInvitationSettingRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create app invitation setting error: %v", err)
		return &npool.CreateAppInvitationSettingForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppInvitationSettingForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetAppInvitationSetting(ctx context.Context, in *npool.GetAppInvitationSettingRequest) (*npool.GetAppInvitationSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app invitation setting error: %v", err)
		return &npool.GetAppInvitationSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppInvitationSettingsByApp(ctx context.Context, in *npool.GetAppInvitationSettingsByAppRequest) (*npool.GetAppInvitationSettingsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app invitation setting by app error: %v", err)
		return &npool.GetAppInvitationSettingsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppInvitationSettingsByOtherApp(ctx context.Context, in *npool.GetAppInvitationSettingsByOtherAppRequest) (*npool.GetAppInvitationSettingsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAppInvitationSettingsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app invitation setting by app error: %v", err)
		return &npool.GetAppInvitationSettingsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppInvitationSettingsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) UpdateAppInvitationSetting(ctx context.Context, in *npool.UpdateAppInvitationSettingRequest) (*npool.UpdateAppInvitationSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update app invitation setting error: %v", err)
		return &npool.UpdateAppInvitationSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
