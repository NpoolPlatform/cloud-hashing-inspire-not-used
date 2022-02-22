// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/activity" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateActivity(ctx context.Context, in *npool.CreateActivityRequest) (*npool.CreateActivityResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create activity error: %v", err)
		return &npool.CreateActivityResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateActivityForOtherApp(ctx context.Context, in *npool.CreateActivityForOtherAppRequest) (*npool.CreateActivityForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreateActivityRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("create activity error: %v", err)
		return &npool.CreateActivityForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateActivityForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateActivity(ctx context.Context, in *npool.UpdateActivityRequest) (*npool.UpdateActivityResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update activity error: %v", err)
		return &npool.UpdateActivityResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetActivity(ctx context.Context, in *npool.GetActivityRequest) (*npool.GetActivityResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get activity error: %v", err)
		return &npool.GetActivityResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetActivitiesByApp(ctx context.Context, in *npool.GetActivitiesByAppRequest) (*npool.GetActivitiesByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get activity by app error: %v", err)
		return &npool.GetActivitiesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetActivitiesByOtherApp(ctx context.Context, in *npool.GetActivitiesByOtherAppRequest) (*npool.GetActivitiesByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetActivitiesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get activity by app error: %v", err)
		return &npool.GetActivitiesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetActivitiesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetActivityByAppName(ctx context.Context, in *npool.GetActivityByAppNameRequest) (*npool.GetActivityByAppNameResponse, error) {
	resp, err := crud.GetByAppName(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get activity by app activity name error: %v", err)
		return &npool.GetActivityByAppNameResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
