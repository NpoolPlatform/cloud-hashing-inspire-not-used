package appinvitationsetting

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

func assertAppInvitationSetting(t *testing.T, actual, expected *npool.AppInvitationSetting) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.Count, expected.Count)
	assert.Equal(t, actual.Discount, expected.Discount)
	assert.Equal(t, actual.Title, expected.Title)
	assert.Equal(t, actual.BadgeLarge, expected.BadgeLarge)
	assert.Equal(t, actual.BadgeSmall, expected.BadgeSmall)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.AppInvitationSetting{
		AppID:      uuid.New().String(),
		Count:      10,
		Discount:   20,
		Title:      "test test",
		BadgeLarge: "NOT SET",
		BadgeSmall: "NOT SET",
	}

	resp, err := Create(context.Background(), &npool.CreateAppInvitationSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAppInvitationSetting(t, resp.Info, &setting)
	}

	resp1, err := Get(context.Background(), &npool.GetAppInvitationSettingRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppInvitationSetting(t, resp1.Info, &setting)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetAppInvitationSettingsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Infos[0].ID, resp.Info.ID)
		assertAppInvitationSetting(t, resp2.Infos[0], &setting)
	}

	setting.ID = resp.Info.ID

	resp3, err := Update(context.Background(), &npool.UpdateAppInvitationSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppInvitationSetting(t, resp3.Info, &setting)
	}
}
