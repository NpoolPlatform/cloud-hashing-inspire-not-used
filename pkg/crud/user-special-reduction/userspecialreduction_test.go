// +build !codeanalysis

package userspecialreduction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertUserSpecialReduction(t *testing.T, actual, expected *npool.UserSpecialReduction) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.ReleaseByUserID, expected.ReleaseByUserID)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.DurationDays, expected.DurationDays)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coupon := npool.UserSpecialReduction{
		AppID:           uuid.New().String(),
		UserID:          uuid.New().String(),
		Amount:          5,
		ReleaseByUserID: uuid.New().String(),
		Start:           uint32(time.Now().Unix()),
		DurationDays:    10,
		Message:         "For test coupon",
	}

	resp, err := Create(context.Background(), &npool.CreateUserSpecialReductionRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUserSpecialReduction(t, resp.Info, &coupon)
	}

	coupon.ID = resp.Info.ID

	resp1, err := Get(context.Background(), &npool.GetUserSpecialReductionRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertUserSpecialReduction(t, resp1.Info, &coupon)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetUserSpecialReductionsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetByAppReleaser(context.Background(), &npool.GetUserSpecialReductionsByAppReleaserRequest{
		AppID:  resp.Info.AppID,
		UserID: resp.Info.ReleaseByUserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByAppUser(context.Background(), &npool.GetUserSpecialReductionsByAppUserRequest{
		AppID:  resp.Info.AppID,
		UserID: resp.Info.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp4.Infos), 1)
	}

	resp5, err := Update(context.Background(), &npool.UpdateUserSpecialReductionRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.Info.ID, resp.Info.ID)
		assertUserSpecialReduction(t, resp5.Info, &coupon)
	}
}
