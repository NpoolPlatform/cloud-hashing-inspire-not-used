package eventcoupon

import (
	"context"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/eventcoupon"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateEventCoupon(info *npool.EventCoupon) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invlaid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetActivityID()); err != nil {
		return xerrors.Errorf("invlaid activity by: %v", err)
	}
	if _, err := uuid.Parse(info.GetCouponID()); err != nil {
		return xerrors.Errorf("invlaid coupon by: %v", err)
	}
	if info.GetEvent() == "" {
		return xerrors.Errorf("invalid event")
	}
	if info.GetCount() == 0 {
		return xerrors.Errorf("invalid count")
	}
	if info.GetType() != constant.CouponTypeCoupon && info.GetType() != constant.CouponTypeDiscount {
		return xerrors.Errorf("invalid coupon type")
	}
	return nil
}

func dbRowToEventCoupon(row *ent.EventCoupon) *npool.EventCoupon {
	return &npool.EventCoupon{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		ActivityID: row.ActivityID.String(),
		CouponID:   row.CouponID.String(),
		Event:      row.Event,
		Type:       row.Type,
		Count:      row.Count,
	}
}

func Create(ctx context.Context, in *npool.CreateEventCouponRequest) (*npool.CreateEventCouponResponse, error) {
	if err := validateEventCoupon(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		EventCoupon.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetActivityID(uuid.MustParse(in.GetInfo().GetActivityID())).
		SetCouponID(uuid.MustParse(in.GetInfo().GetCouponID())).
		SetEvent(in.GetInfo().GetEvent()).
		SetType(in.GetInfo().GetType()).
		SetCount(in.GetInfo().GetCount()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create event coupon: %v", err)
	}

	return &npool.CreateEventCouponResponse{
		Info: dbRowToEventCoupon(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateEventCouponRequest) (*npool.UpdateEventCouponResponse, error) {
	if err := validateEventCoupon(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		EventCoupon.
		UpdateOneID(id).
		SetCount(in.GetInfo().GetCount()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update event coupon: %v", err)
	}

	return &npool.UpdateEventCouponResponse{
		Info: dbRowToEventCoupon(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetEventCouponRequest) (*npool.GetEventCouponResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		EventCoupon.
		Query().
		Where(
			eventcoupon.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event coupon: %v", err)
	}

	var coupon *npool.EventCoupon
	for _, info := range infos {
		coupon = dbRowToEventCoupon(info)
		break
	}

	return &npool.GetEventCouponResponse{
		Info: coupon,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetEventCouponsByAppRequest) (*npool.GetEventCouponsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		EventCoupon.
		Query().
		Where(
			eventcoupon.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event coupon: %v", err)
	}

	coupons := []*npool.EventCoupon{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToEventCoupon(info))
	}

	return &npool.GetEventCouponsByAppResponse{
		Infos: coupons,
	}, nil
}

func GetByAppActivityEvent(ctx context.Context, in *npool.GetEventCouponsByAppActivityEventRequest) (*npool.GetEventCouponsByAppActivityEventResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	actID, err := uuid.Parse(in.GetActivityID())
	if err != nil {
		return nil, xerrors.Errorf("invalid activity id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		EventCoupon.
		Query().
		Where(
			eventcoupon.And(
				eventcoupon.AppID(appID),
				eventcoupon.ActivityID(actID),
				eventcoupon.Event(in.GetEvent()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query event coupon: %v", err)
	}

	coupons := []*npool.EventCoupon{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToEventCoupon(info))
	}

	return &npool.GetEventCouponsByAppActivityEventResponse{
		Infos: coupons,
	}, nil
}
