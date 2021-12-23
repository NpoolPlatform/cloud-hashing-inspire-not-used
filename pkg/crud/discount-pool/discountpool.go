package discountpool

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/discountpool"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateDiscountPool(info *npool.DiscountPool) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetReleaseByUserID()); err != nil {
		return xerrors.Errorf("invalid release by user id: %v", err)
	}
	return nil
}

func dbRowToDiscountPool(row *ent.DiscountPool) *npool.DiscountPool {
	return &npool.DiscountPool{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		Discount:        row.Discount,
		ReleaseByUserID: row.ReleaseByUserID.String(),
		Start:           row.Start,
		DurationDays:    row.DurationDays,
		Message:         row.Message,
		Name:            row.Name,
	}
}

func Create(ctx context.Context, in *npool.CreateDiscountPoolRequest) (*npool.CreateDiscountPoolResponse, error) {
	if err := validateDiscountPool(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		DiscountPool.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetDiscount(in.GetInfo().GetDiscount()).
		SetReleaseByUserID(uuid.MustParse(in.GetInfo().GetReleaseByUserID())).
		SetStart(in.GetInfo().GetStart()).
		SetDurationDays(in.GetInfo().GetDurationDays()).
		SetMessage(in.GetInfo().GetMessage()).
		SetName(in.GetInfo().GetName()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create discount pool: %v", err)
	}

	return &npool.CreateDiscountPoolResponse{
		Info: dbRowToDiscountPool(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateDiscountPoolRequest) (*npool.UpdateDiscountPoolResponse, error) {
	if err := validateDiscountPool(in.GetInfo()); err != nil {
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
		DiscountPool.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		SetName(in.GetInfo().GetName()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update discount pool: %v", err)
	}

	return &npool.UpdateDiscountPoolResponse{
		Info: dbRowToDiscountPool(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetDiscountPoolRequest) (*npool.GetDiscountPoolResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		DiscountPool.
		Query().
		Where(
			discountpool.And(
				discountpool.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query discount pool: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty discount pool")
	}

	return &npool.GetDiscountPoolResponse{
		Info: dbRowToDiscountPool(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetDiscountPoolsByAppRequest) (*npool.GetDiscountPoolsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		DiscountPool.
		Query().
		Where(
			discountpool.And(
				discountpool.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query discount pool: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty discount pool")
	}

	coupons := []*npool.DiscountPool{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToDiscountPool(info))
	}

	return &npool.GetDiscountPoolsByAppResponse{
		Infos: coupons,
	}, nil
}

func GetByAppReleaser(ctx context.Context, in *npool.GetDiscountPoolsByAppReleaserRequest) (*npool.GetDiscountPoolsByAppReleaserResponse, error) {
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
		DiscountPool.
		Query().
		Where(
			discountpool.And(
				discountpool.AppID(appID),
				discountpool.ReleaseByUserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query discount pool: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty discount pool")
	}

	coupons := []*npool.DiscountPool{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToDiscountPool(info))
	}

	return &npool.GetDiscountPoolsByAppReleaserResponse{
		Infos: coupons,
	}, nil
}
