package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/commissioncoinsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCommissionCoinSetting(ctx context.Context, in *npool.CreateCommissionCoinSettingRequest) (*npool.CreateCommissionCoinSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create commission coin setting error: %v", err)
		return &npool.CreateCommissionCoinSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCommissionCoinSetting(ctx context.Context, in *npool.GetCommissionCoinSettingRequest) (*npool.GetCommissionCoinSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get commission coin setting error: %v", err)
		return &npool.GetCommissionCoinSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateCommissionCoinSetting(ctx context.Context, in *npool.UpdateCommissionCoinSettingRequest) (*npool.UpdateCommissionCoinSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get commission coin setting error: %v", err)
		return &npool.UpdateCommissionCoinSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
