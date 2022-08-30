//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/google/uuid"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/user-invitation-code" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserInvitationCode(ctx context.Context, in *npool.CreateUserInvitationCodeRequest) (*npool.CreateUserInvitationCodeResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.CreateUserInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateUserInvitationCode(ctx context.Context, in *npool.UpdateUserInvitationCodeRequest) (*npool.UpdateUserInvitationCodeResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update user invitation code error: %v", err)
		return &npool.UpdateUserInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateUserInvitationCodeForOtherAppUser(ctx context.Context, in *npool.CreateUserInvitationCodeForOtherAppUserRequest) (*npool.CreateUserInvitationCodeForOtherAppUserResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()
	info.UserID = in.GetTargetUserID()

	resp, err := crud.Create(ctx, &npool.CreateUserInvitationCodeRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.CreateUserInvitationCodeForOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateUserInvitationCodeForOtherAppUserResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) CreateUserInvitationCodeForAppOtherUser(ctx context.Context, in *npool.CreateUserInvitationCodeForAppOtherUserRequest) (*npool.CreateUserInvitationCodeForAppOtherUserResponse, error) {
	info := in.GetInfo()
	info.UserID = in.GetTargetUserID()

	resp, err := crud.Create(ctx, &npool.CreateUserInvitationCodeRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.CreateUserInvitationCodeForAppOtherUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateUserInvitationCodeForAppOtherUserResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetUserInvitationCode(ctx context.Context, in *npool.GetUserInvitationCodeRequest) (*npool.GetUserInvitationCodeResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.GetUserInvitationCodeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetManyUserInvitationCodes(ctx context.Context, in *npool.GetManyUserInvitationCodesRequest) (*npool.GetManyUserInvitationCodesResponse, error) {
	if len(in.GetUserIDs()) == 0 {
		logger.Sugar().Errorw("UserIDs is empty")
		return &npool.GetManyUserInvitationCodesResponse{}, status.Error(codes.InvalidArgument, "UserIDs is empty")
	}

	userIDs := []uuid.UUID{}
	for _, val := range in.GetUserIDs() {
		userID, err := uuid.Parse(val)
		if err != nil {
			logger.Sugar().Errorw("get many user invitation code error: %v", err)
			return &npool.GetManyUserInvitationCodesResponse{}, status.Error(codes.InvalidArgument, "UserIDs is invalid")
		}
		userIDs = append(userIDs, userID)
	}

	infos, err := crud.GetByManyUser(ctx, userIDs)
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.GetManyUserInvitationCodesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetManyUserInvitationCodesResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetUserInvitationCodes(ctx context.Context, in *npool.GetUserInvitationCodesRequest) (*npool.GetUserInvitationCodesResponse, error) {
	resp, err := crud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.GetUserInvitationCodesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserInvitationCodesByApp(ctx context.Context, in *npool.GetUserInvitationCodesByAppRequest) (*npool.GetUserInvitationCodesByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.GetUserInvitationCodesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserInvitationCodesByOtherApp(ctx context.Context, in *npool.GetUserInvitationCodesByOtherAppRequest) (*npool.GetUserInvitationCodesByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetUserInvitationCodesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %v", err)
		return &npool.GetUserInvitationCodesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetUserInvitationCodesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetUserInvitationCodeByAppUser(ctx context.Context, in *npool.GetUserInvitationCodeByAppUserRequest) (*npool.GetUserInvitationCodeByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get user invitation code by app user error: %v", err)
		return &npool.GetUserInvitationCodeByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserInvitationCodeByCode(ctx context.Context, in *npool.GetUserInvitationCodeByCodeRequest) (*npool.GetUserInvitationCodeByCodeResponse, error) {
	resp, err := crud.GetByCode(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get user invitation code by code error: %v", err)
		return &npool.GetUserInvitationCodeByCodeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
