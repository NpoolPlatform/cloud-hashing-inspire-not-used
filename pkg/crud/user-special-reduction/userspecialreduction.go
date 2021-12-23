package userspecialreduction

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userspecialreduction"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateUserSpecialReduction(info *npool.UserSpecialReduction) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetReleaseByUserID()); err != nil {
		return xerrors.Errorf("invalid release by user id: %v", err)
	}
	return nil
}

func dbRowToUserSpecialReduction(row *ent.UserSpecialReduction) *npool.UserSpecialReduction {
	return &npool.UserSpecialReduction{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		Amount:          price.DBPriceToVisualPrice(row.Amount),
		ReleaseByUserID: row.ReleaseByUserID.String(),
		Start:           row.Start,
		DurationDays:    row.DurationDays,
		Message:         row.Message,
	}
}

func Create(ctx context.Context, in *npool.CreateUserSpecialReductionRequest) (*npool.CreateUserSpecialReductionResponse, error) {
	if err := validateUserSpecialReduction(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserSpecialReduction.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetReleaseByUserID(uuid.MustParse(in.GetInfo().GetReleaseByUserID())).
		SetStart(in.GetInfo().GetStart()).
		SetDurationDays(in.GetInfo().GetDurationDays()).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user special reduction: %v", err)
	}

	return &npool.CreateUserSpecialReductionResponse{
		Info: dbRowToUserSpecialReduction(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserSpecialReductionRequest) (*npool.UpdateUserSpecialReductionResponse, error) {
	if err := validateUserSpecialReduction(in.GetInfo()); err != nil {
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
		UserSpecialReduction.
		UpdateOneID(id).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update user special reduction: %v", err)
	}

	return &npool.UpdateUserSpecialReductionResponse{
		Info: dbRowToUserSpecialReduction(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserSpecialReductionRequest) (*npool.GetUserSpecialReductionResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserSpecialReduction.
		Query().
		Where(
			userspecialreduction.And(
				userspecialreduction.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user special reduction: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty user special reduction")
	}

	return &npool.GetUserSpecialReductionResponse{
		Info: dbRowToUserSpecialReduction(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetUserSpecialReductionsByAppRequest) (*npool.GetUserSpecialReductionsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserSpecialReduction.
		Query().
		Where(
			userspecialreduction.And(
				userspecialreduction.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user special reduction: %v", err)
	}

	coupons := []*npool.UserSpecialReduction{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToUserSpecialReduction(info))
	}

	return &npool.GetUserSpecialReductionsByAppResponse{
		Infos: coupons,
	}, nil
}

func GetByAppReleaser(ctx context.Context, in *npool.GetUserSpecialReductionsByAppReleaserRequest) (*npool.GetUserSpecialReductionsByAppReleaserResponse, error) {
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
		UserSpecialReduction.
		Query().
		Where(
			userspecialreduction.And(
				userspecialreduction.AppID(appID),
				userspecialreduction.ReleaseByUserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user special reduction: %v", err)
	}

	coupons := []*npool.UserSpecialReduction{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToUserSpecialReduction(info))
	}

	return &npool.GetUserSpecialReductionsByAppReleaserResponse{
		Infos: coupons,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserSpecialReductionsByAppUserRequest) (*npool.GetUserSpecialReductionsByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid user id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserSpecialReduction.
		Query().
		Where(
			userspecialreduction.And(
				userspecialreduction.AppID(appID),
				userspecialreduction.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user special reduction: %v", err)
	}

	coupons := []*npool.UserSpecialReduction{}
	for _, info := range infos {
		coupons = append(coupons, dbRowToUserSpecialReduction(info))
	}

	return &npool.GetUserSpecialReductionsByAppUserResponse{
		Infos: coupons,
	}, nil
}
