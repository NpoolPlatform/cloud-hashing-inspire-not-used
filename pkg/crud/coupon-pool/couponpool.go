package couponpool

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateCouponPool(info *npool.CouponPool) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetReleaseByUserID()); err != nil {
		return xerrors.Errorf("invalid release by user id: %v", err)
	}
	return nil
}

func dbRowToCouponPool(row *ent.CouponPool) *npool.CouponPool {
	return &npool.CouponPool{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		Denomination:    price.DBPriceToVisualPrice(row.Denomination),
		Circulation:     row.Circulation,
		ReleaseByUserID: row.ReleaseByUserID.String(),
		Start:           row.Start,
		DurationDays:    row.DurationDays,
		Message:         row.Message,
		Name:            row.Name,
	}
}

func Create(ctx context.Context, in *npool.CreateCouponPoolRequest) (*npool.CreateCouponPoolResponse, error) {
	if err := validateCouponPool(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		CouponPool.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetDenomination(price.VisualPriceToDBPrice(in.GetInfo().GetDenomination())).
		SetCirculation(in.GetInfo().GetCirculation()).
		SetReleaseByUserID(uuid.MustParse(in.GetInfo().GetReleaseByUserID())).
		SetStart(in.GetInfo().GetStart()).
		SetDurationDays(in.GetInfo().GetDurationDays()).
		SetMessage(in.GetInfo().GetMessage()).
		SetName(in.GetInfo().GetName()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create coupon pool: %v", err)
	}

	return &npool.CreateCouponPoolResponse{
		Info: dbRowToCouponPool(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateCouponPoolRequest) (*npool.UpdateCouponPoolResponse, error) {
	if err := validateCouponPool(in.GetInfo()); err != nil {
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
		CouponPool.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		SetName(in.GetInfo().GetName()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update coupon pool: %v", err)
	}

	return &npool.UpdateCouponPoolResponse{
		Info: dbRowToCouponPool(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCouponPoolRequest) (*npool.GetCouponPoolResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		CouponPool.
		Query().
		Where(
			couponpool.And(
				couponpool.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coupon pool: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty coupon pool")
	}

	return &npool.GetCouponPoolResponse{
		Info: dbRowToCouponPool(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetCouponPoolsByAppRequest) (*npool.GetCouponPoolsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		CouponPool.
		Query().
		Where(
			couponpool.And(
				couponpool.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coupon pool: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty coupon pool")
	}

	coupons := []*npool.CouponPool{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToCouponPool(info))
	}

	return &npool.GetCouponPoolsByAppResponse{
		Infos: coupons,
	}, nil
}

func GetByAppReleaser(ctx context.Context, in *npool.GetCouponPoolsByAppReleaserRequest) (*npool.GetCouponPoolsByAppReleaserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid releaser id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		CouponPool.
		Query().
		Where(
			couponpool.And(
				couponpool.AppID(appID),
				couponpool.ReleaseByUserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coupon pool: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty coupon pool")
	}

	coupons := []*npool.CouponPool{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToCouponPool(info))
	}

	return &npool.GetCouponPoolsByAppReleaserResponse{
		Infos: coupons,
	}, nil
}
