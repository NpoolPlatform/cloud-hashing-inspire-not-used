// +build !codeanalysis

package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"
	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/purchase-invitation" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePurchaseInvitation(ctx context.Context, in *npool.CreatePurchaseInvitationRequest) (*npool.CreatePurchaseInvitationResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create purchase invitation error: %w", err)
		return &npool.CreatePurchaseInvitationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdatePurchaseInvitation(ctx context.Context, in *npool.UpdatePurchaseInvitationRequest) (*npool.UpdatePurchaseInvitationResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update purchase invitation error: %w", err)
		return &npool.UpdatePurchaseInvitationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPurchaseInvitation(ctx context.Context, in *npool.GetPurchaseInvitationRequest) (*npool.GetPurchaseInvitationResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get purchase invitation error: %w", err)
		return &npool.GetPurchaseInvitationResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPurchaseInvitationsByApp(ctx context.Context, in *npool.GetPurchaseInvitationsByAppRequest) (*npool.GetPurchaseInvitationsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get purchase invitations by app error: %w", err)
		return &npool.GetPurchaseInvitationsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPurchaseInvitationByAppOrder(ctx context.Context, in *npool.GetPurchaseInvitationByAppOrderRequest) (*npool.GetPurchaseInvitationByAppOrderResponse, error) {
	resp, err := crud.GetByAppOrder(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get purchase invitation by app order error: %w", err)
		return &npool.GetPurchaseInvitationByAppOrderResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
