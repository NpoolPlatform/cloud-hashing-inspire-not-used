package appcommissionsetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"

	testinit "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/test-init"

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

func assertAppCommissionSetting(t *testing.T, actual, expected *npool.AppCommissionSetting) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.Type, expected.Type)
	assert.Equal(t, actual.Level, expected.Level)
	assert.Equal(t, actual.InvitationDiscount, expected.InvitationDiscount)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.AppCommissionSetting{
		AppID:              uuid.New().String(),
		Type:               constant.CommissionByAmount,
		Level:              5,
		InvitationDiscount: true,
	}

	resp, err := Create(context.Background(), &npool.CreateAppCommissionSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAppCommissionSetting(t, resp.Info, &setting)
	}

	resp1, err := Get(context.Background(), &npool.GetAppCommissionSettingRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppCommissionSetting(t, resp1.Info, &setting)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetAppCommissionSettingByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppCommissionSetting(t, resp2.Info, &setting)
	}

	setting.ID = resp.Info.ID

	resp3, err := Update(context.Background(), &npool.UpdateAppCommissionSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppCommissionSetting(t, resp3.Info, &setting)
	}
}
