// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/user-invitation-code" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserInvitationCode(ctx context.Context, in *npool.CreateUserInvitationCodeRequest) (*npool.CreateUserInvitationCodeResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %w", err)
		return &npool.CreateUserInvitationCodeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserInvitationCode(ctx context.Context, in *npool.GetUserInvitationCodeRequest) (*npool.GetUserInvitationCodeResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create user invitation code error: %w", err)
		return &npool.GetUserInvitationCodeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserInvitationCodeByAppUser(ctx context.Context, in *npool.GetUserInvitationCodeByAppUserRequest) (*npool.GetUserInvitationCodeByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get user invitation code by app user error: %w", err)
		return &npool.GetUserInvitationCodeByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserInvitationCodeByCode(ctx context.Context, in *npool.GetUserInvitationCodeByCodeRequest) (*npool.GetUserInvitationCodeByCodeResponse, error) {
	resp, err := crud.GetByCode(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get user invitation code by code error: %w", err)
		return &npool.GetUserInvitationCodeByCodeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
