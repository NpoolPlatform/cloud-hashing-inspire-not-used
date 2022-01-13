// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/app-coupon-setting" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppCouponSetting(ctx context.Context, in *npool.CreateAppCouponSettingRequest) (*npool.CreateAppCouponSettingResponse, error) {
	resp, err := appcouponsetting.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create app coupon setting error: %w", err)
		return &npool.CreateAppCouponSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAppCouponSetting(ctx context.Context, in *npool.GetAppCouponSettingRequest) (*npool.GetAppCouponSettingResponse, error) {
	resp, err := appcouponsetting.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get app coupon setting error: %w", err)
		return &npool.GetAppCouponSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetAppCouponSettingByApp(ctx context.Context, in *npool.GetAppCouponSettingByAppRequest) (*npool.GetAppCouponSettingByAppResponse, error) {
	resp, err := appcouponsetting.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get app coupon setting by app error: %w", err)
		return &npool.GetAppCouponSettingByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateAppCouponSetting(ctx context.Context, in *npool.UpdateAppCouponSettingRequest) (*npool.UpdateAppCouponSettingResponse, error) {
	resp, err := appcouponsetting.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update app coupon setting error: %w", err)
		return &npool.UpdateAppCouponSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
