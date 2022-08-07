package activity

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/activity"

	"github.com/google/uuid"
)

func validateActivity(info *npool.Activity) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invlaid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCreatedBy()); err != nil {
		return fmt.Errorf("invlaid created by: %v", err)
	}
	if info.GetName() == "" {
		return fmt.Errorf("invalid name")
	}
	if info.GetEnd() <= info.GetStart() {
		return fmt.Errorf("invalid expiration")
	}
	if info.GetStart() <= uint32(time.Now().Unix()) {
		return fmt.Errorf("invalid start %v < %v", info.GetStart(), time.Now().Unix())
	}
	return nil
}

func dbRowToActivity(row *ent.Activity) *npool.Activity {
	return &npool.Activity{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		CreatedBy:      row.CreatedBy.String(),
		Name:           row.Name,
		Start:          row.Start,
		End:            row.End,
		SystemActivity: row.SystemActivity,
	}
}

func Create(ctx context.Context, in *npool.CreateActivityRequest) (*npool.CreateActivityResponse, error) {
	if err := validateActivity(in.GetInfo()); err != nil {
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		Activity.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetCreatedBy(uuid.MustParse(in.GetInfo().GetCreatedBy())).
		SetName(in.GetInfo().GetName()).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		SetSystemActivity(in.GetInfo().GetSystemActivity()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create activity: %v", err)
	}

	return &npool.CreateActivityResponse{
		Info: dbRowToActivity(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateActivityRequest) (*npool.UpdateActivityResponse, error) {
	if err := validateActivity(in.GetInfo()); err != nil {
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
		Activity.
		UpdateOneID(id).
		SetName(in.GetInfo().GetName()).
		SetStart(in.GetInfo().GetStart()).
		SetEnd(in.GetInfo().GetEnd()).
		SetSystemActivity(in.GetInfo().GetSystemActivity()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update activity: %v", err)
	}

	return &npool.UpdateActivityResponse{
		Info: dbRowToActivity(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetActivityRequest) (*npool.GetActivityResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Activity.
		Query().
		Where(
			activity.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query activity: %v", err)
	}

	var act *npool.Activity
	for _, info := range infos {
		act = dbRowToActivity(info)
		break
	}

	return &npool.GetActivityResponse{
		Info: act,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetActivitiesByAppRequest) (*npool.GetActivitiesByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Activity.
		Query().
		Where(
			activity.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query activity: %v", err)
	}

	acts := []*npool.Activity{}
	for _, info := range infos {
		acts = append(acts, dbRowToActivity(info))
	}

	return &npool.GetActivitiesByAppResponse{
		Infos: acts,
	}, nil
}

func GetByAppName(ctx context.Context, in *npool.GetActivityByAppNameRequest) (*npool.GetActivityByAppNameResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		Activity.
		Query().
		Where(
			activity.And(
				activity.AppID(appID),
				activity.Name(in.GetName()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query activity: %v", err)
	}

	var act *npool.Activity
	for _, info := range infos {
		act = dbRowToActivity(info)
		break
	}

	return &npool.GetActivityByAppNameResponse{
		Info: act,
	}, nil
}
