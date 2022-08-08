package commissioncoin

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/commissioncoinsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"
)

func CreateCommissionCoinSetting(ctx context.Context, in *npool.CreateCommissionCoinSettingRequest) (*npool.CreateCommissionCoinSettingResponse, error) {
	resp, err := crud.GetAll(ctx, &npool.GetCommissionCoinSettingsRequest{})
	if err != nil {
		return nil, fmt.Errorf("fail get settings: %v", err)
	}

	var myInfo *npool.CommissionCoinSetting

	for _, info := range resp.Infos {
		info.Using = false
		if info.CoinTypeID == in.GetInfo().GetCoinTypeID() {
			info.Using = true
			myInfo = info
		}
		_, err := crud.Update(ctx, &npool.UpdateCommissionCoinSettingRequest{
			Info: info,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update setting: %v", err)
		}
	}

	if myInfo == nil {
		return crud.Create(ctx, in)
	}

	return &npool.CreateCommissionCoinSettingResponse{
		Info: myInfo,
	}, nil
}
