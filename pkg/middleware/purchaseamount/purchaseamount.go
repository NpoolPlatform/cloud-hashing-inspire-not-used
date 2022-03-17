package purchaseamount

import (
	"context"
	"time"

	appsetting "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/apppurchaseamountsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"golang.org/x/xerrors"
)

func CreateAppPurchaseAmountSetting(ctx context.Context, in *npool.CreateAppPurchaseAmountSettingRequest) (*npool.CreateAppPurchaseAmountSettingResponse, error) {
	resp, err := appsetting.GetByApp(ctx, &npool.GetAppPurchaseAmountSettingsByAppRequest{
		AppID: in.GetInfo().GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get settings: %v", err)
	}

	start := uint32(time.Now().Unix())
	var setting *npool.AppPurchaseAmountSetting

	for _, info := range resp.Infos {
		if info.Start == 0 {
			info.Start = start
		}
		if info.Amount == in.GetInfo().GetAmount() {
			if info.End == 0 {
				info.End = start
			}
		} else {
			info.End = 0
			info.Percent = in.GetInfo().GetPercent()
			info.BadgeLarge = in.GetInfo().GetBadgeLarge()
			info.BadgeSmall = in.GetInfo().GetBadgeSmall()
			setting = info
		}
		_, err := appsetting.Update(ctx, &npool.UpdateAppPurchaseAmountSettingRequest{
			Info: info,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail update setting: %v", err)
		}
	}

	if setting != nil {
		return &npool.CreateAppPurchaseAmountSettingResponse{
			Info: setting,
		}, nil
	}

	info := in.GetInfo()
	info.Start = start
	info.End = 0

	return appsetting.Create(ctx, &npool.CreateAppPurchaseAmountSettingRequest{
		Info: info,
	})
}
