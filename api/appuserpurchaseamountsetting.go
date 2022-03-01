// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/appuserpurchaseamountsetting"
	mw "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/middleware/purchaseamount"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppUserPurchaseAmountSetting(ctx context.Context, in *npool.CreateAppUserPurchaseAmountSettingRequest) (*npool.CreateAppUserPurchaseAmountSettingResponse, error) {
	resp, err := mw.CreateAppUserPurchaseAmountSetting(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create app user purchase amount setting error: %v", err)
		return &npool.CreateAppUserPurchaseAmountSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppUserPurchaseAmountSettingForOtherAppUser(ctx context.Context, in *npool.CreateAppUserPurchaseAmountSettingForOtherAppUserRequest) (*npool.CreateAppUserPurchaseAmountSettingForOtherAppUserResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()
	info.UserID = in.GetTargetUserID()

	resp, err := crud.Create(ctx, &npool.CreateAppUserPurchaseAmountSettingRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create app user purchase amount setting error: %v", err)
		return &npool.CreateAppUserPurchaseAmountSettingForOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateAppUserPurchaseAmountSettingForOtherAppUserResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateAppUserPurchaseAmountSetting(ctx context.Context, in *npool.UpdateAppUserPurchaseAmountSettingRequest) (*npool.UpdateAppUserPurchaseAmountSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update app user purchase amount setting error: %v", err)
		return &npool.UpdateAppUserPurchaseAmountSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserPurchaseAmountSetting(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingRequest) (*npool.GetAppUserPurchaseAmountSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app user purchase amount setting error: %v", err)
		return &npool.GetAppUserPurchaseAmountSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserPurchaseAmountSettingsByApp(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingsByAppRequest) (*npool.GetAppUserPurchaseAmountSettingsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app user purchase amount settings by app error: %v", err)
		return &npool.GetAppUserPurchaseAmountSettingsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserPurchaseAmountSettingsByOtherApp(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingsByOtherAppRequest) (*npool.GetAppUserPurchaseAmountSettingsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAppUserPurchaseAmountSettingsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app user purchase amount settings by app error: %v", err)
		return &npool.GetAppUserPurchaseAmountSettingsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppUserPurchaseAmountSettingsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetAppUserPurchaseAmountSettingsByAppUser(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingsByAppUserRequest) (*npool.GetAppUserPurchaseAmountSettingsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app user purchase amount settings by app error: %v", err)
		return &npool.GetAppUserPurchaseAmountSettingsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppUserPurchaseAmountSettingsByOtherAppUser(ctx context.Context, in *npool.GetAppUserPurchaseAmountSettingsByOtherAppUserRequest) (*npool.GetAppUserPurchaseAmountSettingsByOtherAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, &npool.GetAppUserPurchaseAmountSettingsByAppUserRequest{
		AppID:  in.GetTargetAppID(),
		UserID: in.GetTargetUserID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app user purchase amount settings by app error: %v", err)
		return &npool.GetAppUserPurchaseAmountSettingsByOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppUserPurchaseAmountSettingsByOtherAppUserResponse{
		Infos: resp.Infos,
	}, nil
}
