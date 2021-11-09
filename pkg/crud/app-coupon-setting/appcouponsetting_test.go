package appcouponsetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/test-init" //nolint

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

func assertAppCouponSetting(t *testing.T, actual, expected *npool.AppCouponSetting) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.DominationLimit, expected.DominationLimit)
	assert.Equal(t, actual.TotalLimit, expected.TotalLimit)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.AppCouponSetting{
		DominationLimit: 10,
		TotalLimit:      10,
		AppID:           uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAppCouponSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAppCouponSetting(t, resp.Info, &setting)
	}

	resp1, err := Get(context.Background(), &npool.GetAppCouponSettingRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppCouponSetting(t, resp1.Info, &setting)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetAppCouponSettingByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppCouponSetting(t, resp2.Info, &setting)
	}

	setting.ID = resp.Info.ID

	resp3, err := Update(context.Background(), &npool.UpdateAppCouponSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppCouponSetting(t, resp3.Info, &setting)
	}
}
