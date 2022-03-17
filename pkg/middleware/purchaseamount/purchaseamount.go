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

	for _, info := range resp.Infos {
		if info.Amount != in.GetInfo().GetAmount() || info.End != 0 {
			continue
		}

		if info.Percent == in.GetInfo().GetPercent() {
			return &npool.CreateAppPurchaseAmountSettingResponse{
				Info: info,
			}, nil
		}

		info.End = start
		_, err := appsetting.Update(ctx, &npool.UpdateAppPurchaseAmountSettingRequest{
			Info: info,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail update setting: %v", err)
		}
		break
	}

	info := in.GetInfo()
	info.Start = start
	info.End = 0

	return appsetting.Create(ctx, &npool.CreateAppPurchaseAmountSettingRequest{
		Info: info,
	})
}
