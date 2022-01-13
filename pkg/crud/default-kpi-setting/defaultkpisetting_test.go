package defaultkpisetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/test-init" //nolint
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

func assertDefaultKpiSetting(t *testing.T, actual, expected *npool.DefaultKpiSetting) {
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.Percent, expected.Percent)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.GoodID, expected.GoodID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.DefaultKpiSetting{
		AppID:   uuid.New().String(),
		GoodID:  uuid.New().String(),
		Amount:  10,
		Percent: 20,
	}

	resp, err := Create(context.Background(), &npool.CreateDefaultKpiSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertDefaultKpiSetting(t, resp.Info, &setting)
	}

	setting.ID = resp.Info.ID

	resp1, err := Get(context.Background(), &npool.GetDefaultKpiSettingRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertDefaultKpiSetting(t, resp1.Info, &setting)
	}

	resp2, err := GetByAppGood(context.Background(), &npool.GetDefaultKpiSettingByAppGoodRequest{
		AppID:  resp.Info.AppID,
		GoodID: resp.Info.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertDefaultKpiSetting(t, resp2.Info, &setting)
	}

	setting.Amount = 10.2

	resp3, err := Update(context.Background(), &npool.UpdateDefaultKpiSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertDefaultKpiSetting(t, resp3.Info, &setting)
	}
}
