package discountpool

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/discountpool"

	"github.com/google/uuid"
)

func validateDiscountPool(info *npool.DiscountPool) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetReleaseByUserID()); err != nil {
		return fmt.Errorf("invalid release by user id: %v", err)
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
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
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
		return nil, fmt.Errorf("fail create discount pool: %v", err)
	}

	return &npool.CreateDiscountPoolResponse{
		Info: dbRowToDiscountPool(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateDiscountPoolRequest) (*npool.UpdateDiscountPoolResponse, error) {
	if err := validateDiscountPool(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		DiscountPool.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		SetName(in.GetInfo().GetName()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update discount pool: %v", err)
	}

	return &npool.UpdateDiscountPoolResponse{
		Info: dbRowToDiscountPool(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetDiscountPoolRequest) (*npool.GetDiscountPoolResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		DiscountPool.
		Query().
		Where(
			discountpool.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query discount pool: %v", err)
	}

	var discount *npool.DiscountPool
	for _, info := range infos {
		discount = dbRowToDiscountPool(info)
		break
	}

	return &npool.GetDiscountPoolResponse{
		Info: discount,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetDiscountPoolsByAppRequest) (*npool.GetDiscountPoolsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		DiscountPool.
		Query().
		Where(
			discountpool.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query discount pool: %v", err)
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
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, fmt.Errorf("invlaid releaser id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
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
		return nil, fmt.Errorf("fail query discount pool: %v", err)
	}

	coupons := []*npool.DiscountPool{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToDiscountPool(info))
	}

	return &npool.GetDiscountPoolsByAppReleaserResponse{
		Infos: coupons,
	}, nil
}
