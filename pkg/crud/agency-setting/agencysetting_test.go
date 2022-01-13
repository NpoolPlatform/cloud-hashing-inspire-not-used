package agencysetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"
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

func assertAgencySetting(t *testing.T, actual, expected *npool.AgencySetting) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.RegistrationRewardThreshold, expected.RegistrationRewardThreshold)
	assert.Equal(t, actual.RegistrationCouponID, expected.RegistrationCouponID)
	assert.Equal(t, actual.KycRewardThreshold, expected.KycRewardThreshold)
	assert.Equal(t, actual.KycCouponID, expected.KycCouponID)
	assert.Equal(t, actual.TotalPurchaseRewardPercent, expected.TotalPurchaseRewardPercent)
	assert.Equal(t, actual.PurchaseRewardChainLevels, expected.PurchaseRewardChainLevels)
	assert.Equal(t, actual.LevelPurchaseRewardPercent, expected.LevelPurchaseRewardPercent)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.AgencySetting{
		RegistrationRewardThreshold: 30,
		RegistrationCouponID:        uuid.New().String(),
		KycRewardThreshold:          5,
		KycCouponID:                 uuid.New().String(),
		TotalPurchaseRewardPercent:  19,
		PurchaseRewardChainLevels:   2,
		LevelPurchaseRewardPercent:  5,
		AppID:                       uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateAgencySettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAgencySetting(t, resp.Info, &setting)
	}

	resp1, err := Get(context.Background(), &npool.GetAgencySettingRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAgencySetting(t, resp1.Info, &setting)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetAgencySettingByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAgencySetting(t, resp2.Info, &setting)
	}

	setting.ID = resp.Info.ID

	resp3, err := Update(context.Background(), &npool.UpdateAgencySettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAgencySetting(t, resp3.Info, &setting)
	}
}
