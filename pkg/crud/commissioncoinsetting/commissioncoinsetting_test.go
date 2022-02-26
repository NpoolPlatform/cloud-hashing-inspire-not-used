package commissioncoinsetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/test-init"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertCommissionCoinSetting(t *testing.T, actual, expected *npool.CommissionCoinSetting) {
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.CommissionCoinSetting{
		CoinTypeID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateCommissionCoinSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCommissionCoinSetting(t, resp.Info, &setting)
	}

	resp1, err := Get(context.Background(), &npool.GetCommissionCoinSettingRequest{})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCommissionCoinSetting(t, resp1.Info, &setting)
	}

	setting.ID = resp.Info.ID

	resp3, err := Update(context.Background(), &npool.UpdateCommissionCoinSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertCommissionCoinSetting(t, resp3.Info, &setting)
	}
}
