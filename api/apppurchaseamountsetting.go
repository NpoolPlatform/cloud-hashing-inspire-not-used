// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/apppurchaseamountsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppPurchaseAmountSetting(ctx context.Context, in *npool.CreateAppPurchaseAmountSettingRequest) (*npool.CreateAppPurchaseAmountSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create app purchase amount setting error: %v", err)
		return &npool.CreateAppPurchaseAmountSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppPurchaseAmountSettingForOtherApp(ctx context.Context, in *npool.CreateAppPurchaseAmountSettingForOtherAppRequest) (*npool.CreateAppPurchaseAmountSettingForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateAppPurchaseAmountSettingRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create app purchase amount setting error: %v", err)
		return &npool.CreateAppPurchaseAmountSettingForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppPurchaseAmountSettingForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateAppPurchaseAmountSetting(ctx context.Context, in *npool.UpdateAppPurchaseAmountSettingRequest) (*npool.UpdateAppPurchaseAmountSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update app purchase amount setting error: %v", err)
		return &npool.UpdateAppPurchaseAmountSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppPurchaseAmountSetting(ctx context.Context, in *npool.GetAppPurchaseAmountSettingRequest) (*npool.GetAppPurchaseAmountSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app purchase amount setting error: %v", err)
		return &npool.GetAppPurchaseAmountSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppPurchaseAmountSettingsByApp(ctx context.Context, in *npool.GetAppPurchaseAmountSettingsByAppRequest) (*npool.GetAppPurchaseAmountSettingsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app purchase amount settings by app error: %v", err)
		return &npool.GetAppPurchaseAmountSettingsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppPurchaseAmountSettingsByOtherApp(ctx context.Context, in *npool.GetAppPurchaseAmountSettingsByOtherAppRequest) (*npool.GetAppPurchaseAmountSettingsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAppPurchaseAmountSettingsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app purchase amount settings by app error: %v", err)
		return &npool.GetAppPurchaseAmountSettingsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppPurchaseAmountSettingsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
