package client

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get inspire connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewCloudHashingInspireClient(conn)

	return fn(_ctx, cli)
}

func CreateInvitationCode(ctx context.Context, in *npool.UserInvitationCode) (*npool.UserInvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.CreateUserInvitationCode(ctx, &npool.CreateUserInvitationCodeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create user invitation code: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create user invitation code: %v", err)
	}
	return info.(*npool.UserInvitationCode), nil
}

func CreateAmountSetting(ctx context.Context, in *npool.AppPurchaseAmountSetting) (*npool.AppPurchaseAmountSetting, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.CreateAppPurchaseAmountSetting(ctx, &npool.CreateAppPurchaseAmountSettingRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create amount setting: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create amount setting: %v", err)
	}
	return info.(*npool.AppPurchaseAmountSetting), nil
}

func GetAmountSettings(ctx context.Context, appID, userID string) ([]*npool.AppPurchaseAmountSetting, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.GetAppPurchaseAmountSettingsByAppUser(ctx, &npool.GetAppPurchaseAmountSettingsByAppUserRequest{
			AppID:  appID,
			UserID: userID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get amount settings: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get amount settings: %v", err)
	}
	return infos.([]*npool.AppPurchaseAmountSetting), nil
}

func GetInvitation(ctx context.Context, appID, inviteeID string) (*npool.RegistrationInvitation, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.GetRegistrationInvitationByAppInvitee(ctx, &npool.GetRegistrationInvitationByAppInviteeRequest{
			AppID:     appID,
			InviteeID: inviteeID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registration invitation: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get registration invitation: %v", err)
	}
	return info.(*npool.RegistrationInvitation), nil
}

func CreateInvitation(ctx context.Context, appID, inviterID, inviteeID string) (*npool.RegistrationInvitation, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.CreateRegistrationInvitation(ctx, &npool.CreateRegistrationInvitationRequest{
			Info: &npool.RegistrationInvitation{
				AppID:     appID,
				InviterID: inviterID,
				InviteeID: inviteeID,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registration invitation: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get registration invitation: %v", err)
	}
	return info.(*npool.RegistrationInvitation), nil
}

func CreateInvitationRevert(ctx context.Context, appID, inviterID, inviteeID string) (*npool.RegistrationInvitation, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.CreateRegistrationInvitationRevert(ctx, &npool.CreateRegistrationInvitationRequest{
			Info: &npool.RegistrationInvitation{
				AppID:     appID,
				InviterID: inviterID,
				InviteeID: inviteeID,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("fail get registration invitation: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get registration invitation: %v", err)
	}
	return info.(*npool.RegistrationInvitation), nil
}

func GetUserInvitationCodeOnly(ctx context.Context, conds *npool.Conds) (*npool.UserInvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.GetUserInvitationCodeByCode(ctx, &npool.GetUserInvitationCodeByCodeRequest{
			Code: conds.GetInvitationCode().GetValue(),
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.UserInvitationCode), nil
}

func GetUserInvitationCodeByAppUser(ctx context.Context, appID, userID string) (*npool.UserInvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.GetUserInvitationCodeByAppUser(ctx, &npool.GetUserInvitationCodeByAppUserRequest{
			AppID:  appID,
			UserID: userID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.UserInvitationCode), nil
}

func UpdateUserInvitationCode(ctx context.Context, id string, confirmed bool) (*npool.UserInvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.UpdateUserInvitationCode(ctx, &npool.UpdateUserInvitationCodeRequest{
			Info: &npool.UserInvitationCode{
				ID:        id,
				Confirmed: confirmed,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.UserInvitationCode), nil
}

func GetManyUserInvitationCodes(ctx context.Context, userIDs []string) ([]*npool.UserInvitationCode, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.GetManyUserInvitationCodes(ctx, &npool.GetManyUserInvitationCodesRequest{
			UserIDs: userIDs,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.UserInvitationCode), nil
}
