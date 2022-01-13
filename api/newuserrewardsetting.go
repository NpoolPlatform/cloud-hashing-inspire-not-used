// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/new-user-reward-setting"     //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/middleware/new-user-reward-setting" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateNewUserRewardSetting(ctx context.Context, in *npool.CreateNewUserRewardSettingRequest) (*npool.CreateNewUserRewardSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create new user reward setting error: %w", err)
		return &npool.CreateNewUserRewardSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetNewUserRewardSetting(ctx context.Context, in *npool.GetNewUserRewardSettingRequest) (*npool.GetNewUserRewardSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get new user reward setting error: %w", err)
		return &npool.GetNewUserRewardSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetNewUserRewardSettingDetail(ctx context.Context, in *npool.GetNewUserRewardSettingDetailRequest) (*npool.GetNewUserRewardSettingDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get new user reward setting detail error: %w", err)
		return &npool.GetNewUserRewardSettingDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetNewUserRewardSettingByApp(ctx context.Context, in *npool.GetNewUserRewardSettingByAppRequest) (*npool.GetNewUserRewardSettingByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get new user reward setting by app error: %w", err)
		return &npool.GetNewUserRewardSettingByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateNewUserRewardSetting(ctx context.Context, in *npool.UpdateNewUserRewardSettingRequest) (*npool.UpdateNewUserRewardSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update new user reward setting error: %w", err)
		return &npool.UpdateNewUserRewardSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
