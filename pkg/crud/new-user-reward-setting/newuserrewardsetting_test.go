package newuserrewardsetting

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

func assertNewUserRewardSetting(t *testing.T, actual, expected *npool.NewUserRewardSetting) {
	assert.Equal(t, actual.RegistrationCouponID, expected.RegistrationCouponID)
	assert.Equal(t, actual.KycCouponID, expected.KycCouponID)
	assert.Equal(t, actual.AppID, expected.AppID)
}

func TestCRUD(t *testing.T) {
	setting := npool.NewUserRewardSetting{
		RegistrationCouponID: uuid.New().String(),
		KycCouponID:          uuid.New().String(),
		AppID:                uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateNewUserRewardSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertNewUserRewardSetting(t, resp.Info, &setting)
	}
}
