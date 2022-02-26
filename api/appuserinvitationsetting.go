// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/appuserinvitationsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppUserInvitationSetting(ctx context.Context, in *npool.CreateAppUserInvitationSettingRequest) (*npool.CreateAppUserInvitationSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create app user invitation setting error: %v", err)
		return &npool.CreateAppUserInvitationSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserInvitationSettingForOtherAppUser(ctx context.Context, in *npool.CreateAppUserInvitationSettingForOtherAppUserRequest) (*npool.CreateAppUserInvitationSettingForOtherAppUserResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()
	info.UserID = in.GetTargetUserID()

	resp, err := crud.Create(ctx, &npool.CreateAppUserInvitationSettingRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create app user invitation setting error: %v", err)
		return &npool.CreateAppUserInvitationSettingForOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppUserInvitationSettingForOtherAppUserResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetAppUserInvitationSetting(ctx context.Context, in *npool.GetAppUserInvitationSettingRequest) (*npool.GetAppUserInvitationSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app user invitation setting error: %v", err)
		return &npool.GetAppUserInvitationSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserInvitationSettingsByApp(ctx context.Context, in *npool.GetAppUserInvitationSettingsByAppRequest) (*npool.GetAppUserInvitationSettingsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app user invitation setting by app error: %v", err)
		return &npool.GetAppUserInvitationSettingsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserInvitationSettingsByOtherApp(ctx context.Context, in *npool.GetAppUserInvitationSettingsByOtherAppRequest) (*npool.GetAppUserInvitationSettingsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAppUserInvitationSettingsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app user invitation setting by app error: %v", err)
		return &npool.GetAppUserInvitationSettingsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppUserInvitationSettingsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetAppUserInvitationSettingsByAppUser(ctx context.Context, in *npool.GetAppUserInvitationSettingsByAppUserRequest) (*npool.GetAppUserInvitationSettingsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app user invitation setting by app error: %v", err)
		return &npool.GetAppUserInvitationSettingsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppUserInvitationSetting(ctx context.Context, in *npool.UpdateAppUserInvitationSettingRequest) (*npool.UpdateAppUserInvitationSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update app user invitation setting error: %v", err)
		return &npool.UpdateAppUserInvitationSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
