// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/eventcoupon"     //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/middleware/eventcoupon" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateEventCoupon(ctx context.Context, in *npool.CreateEventCouponRequest) (*npool.CreateEventCouponResponse, error) {
	resp, err := mw.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create event coupon error: %v", err)
		return &npool.CreateEventCouponResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateEventCouponForOtherApp(ctx context.Context, in *npool.CreateEventCouponForOtherAppRequest) (*npool.CreateEventCouponForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := mw.Create(ctx, &npool.CreateEventCouponRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorw("create event coupon error: %v", err)
		return &npool.CreateEventCouponForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateEventCouponForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateEventCoupon(ctx context.Context, in *npool.UpdateEventCouponRequest) (*npool.UpdateEventCouponResponse, error) {
	resp, err := mw.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update event coupon error: %v", err)
		return &npool.UpdateEventCouponResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetEventCoupon(ctx context.Context, in *npool.GetEventCouponRequest) (*npool.GetEventCouponResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get event coupon error: %v", err)
		return &npool.GetEventCouponResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetEventCouponsByApp(ctx context.Context, in *npool.GetEventCouponsByAppRequest) (*npool.GetEventCouponsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get event coupon by app error: %v", err)
		return &npool.GetEventCouponsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetEventCouponsByOtherApp(ctx context.Context, in *npool.GetEventCouponsByOtherAppRequest) (*npool.GetEventCouponsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetEventCouponsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorw("get event coupon by app error: %v", err)
		return &npool.GetEventCouponsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetEventCouponsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetEventCouponsByAppActivityEvent(ctx context.Context, in *npool.GetEventCouponsByAppActivityEventRequest) (*npool.GetEventCouponsByAppActivityEventResponse, error) {
	resp, err := crud.GetByAppActivityEvent(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get event coupon by app event coupon name error: %v", err)
		return &npool.GetEventCouponsByAppActivityEventResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
