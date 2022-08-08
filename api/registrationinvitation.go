//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/dtm-labs/dtmcli"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/registration-invitation" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRegistrationInvitation(ctx context.Context, in *npool.CreateRegistrationInvitationRequest) (*npool.CreateRegistrationInvitationResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create registration invitation error: %w", err)
		return &npool.CreateRegistrationInvitationResponse{}, status.New(codes.Aborted, dtmcli.ResultFailure).Err()
	}
	return resp, nil
}

func (s *Server) CreateRegistrationInvitationRevert(ctx context.Context, in *npool.CreateRegistrationInvitationRequest) (*npool.CreateRegistrationInvitationResponse, error) {
	resp, err := crud.CreateRevert(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create registration invitation error: %w", err)
		return &npool.CreateRegistrationInvitationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateRegistrationInvitation(ctx context.Context, in *npool.UpdateRegistrationInvitationRequest) (*npool.UpdateRegistrationInvitationResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update registration invitation error: %w", err)
		return &npool.UpdateRegistrationInvitationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRegistrationInvitation(ctx context.Context, in *npool.GetRegistrationInvitationRequest) (*npool.GetRegistrationInvitationResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get registration invitation error: %w", err)
		return &npool.GetRegistrationInvitationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRegistrationInvitationsByApp(ctx context.Context, in *npool.GetRegistrationInvitationsByAppRequest) (*npool.GetRegistrationInvitationsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get registration invitations by app error: %w", err)
		return &npool.GetRegistrationInvitationsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRegistrationInvitationsByOtherApp(ctx context.Context, in *npool.GetRegistrationInvitationsByOtherAppRequest) (*npool.GetRegistrationInvitationsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetRegistrationInvitationsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get registration invitations by app error: %w", err)
		return &npool.GetRegistrationInvitationsByOtherAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &npool.GetRegistrationInvitationsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetRegistrationInvitationsByAppInviter(ctx context.Context, in *npool.GetRegistrationInvitationsByAppInviterRequest) (*npool.GetRegistrationInvitationsByAppInviterResponse, error) {
	resp, err := crud.GetByAppInviter(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get registration invitations by app inviter error: %w", err)
		return &npool.GetRegistrationInvitationsByAppInviterResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetRegistrationInvitationByAppInvitee(ctx context.Context, in *npool.GetRegistrationInvitationByAppInviteeRequest) (*npool.GetRegistrationInvitationByAppInviteeResponse, error) {
	resp, err := crud.GetByAppInvitee(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get registration invitation by app invitee error: %w", err)
		return &npool.GetRegistrationInvitationByAppInviteeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
