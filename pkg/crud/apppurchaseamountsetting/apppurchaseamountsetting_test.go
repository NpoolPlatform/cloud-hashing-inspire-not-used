package apppurchaseamountsetting

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

func assertAppPurchaseAmountSetting(t *testing.T, actual, expected *npool.AppPurchaseAmountSetting) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.Percent, expected.Percent)
	assert.Equal(t, actual.Title, expected.Title)
	assert.Equal(t, actual.BadgeLarge, expected.BadgeLarge)
	assert.Equal(t, actual.BadgeSmall, expected.BadgeSmall)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.AppPurchaseAmountSetting{
		AppID:      uuid.New().String(),
		Amount:     1000,
		Percent:    20,
		Title:      "test test",
		BadgeLarge: "NOT SET",
		BadgeSmall: "NOT SET",
	}

	resp, err := Create(context.Background(), &npool.CreateAppPurchaseAmountSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAppPurchaseAmountSetting(t, resp.Info, &setting)
	}

	setting.ID = resp.Info.ID

	resp1, err := Get(context.Background(), &npool.GetAppPurchaseAmountSettingRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppPurchaseAmountSetting(t, resp1.Info, &setting)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetAppPurchaseAmountSettingsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := Update(context.Background(), &npool.UpdateAppPurchaseAmountSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppPurchaseAmountSetting(t, resp3.Info, &setting)
	}
}
