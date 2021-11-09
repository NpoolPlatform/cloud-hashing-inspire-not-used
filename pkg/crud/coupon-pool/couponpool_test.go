package couponpool

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertCouponPool(t *testing.T, actual, expected *npool.CouponPool) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.Denomination, expected.Denomination)
	assert.Equal(t, actual.Circulation, expected.Circulation)
	assert.Equal(t, actual.ReleaseByUserID, expected.ReleaseByUserID)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.DurationDays, expected.DurationDays)
	assert.Equal(t, actual.Message, expected.Message)
	assert.Equal(t, actual.Name, expected.Name)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coupon := npool.CouponPool{
		AppID:           uuid.New().String(),
		Denomination:    5,
		Circulation:     5,
		ReleaseByUserID: uuid.New().String(),
		Start:           uint32(time.Now().Unix()),
		DurationDays:    10,
		Message:         "For test coupon",
		Name:            fmt.Sprintf("Test Coupon-%v", uuid.New().String()),
	}

	resp, err := Create(context.Background(), &npool.CreateCouponPoolRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCouponPool(t, resp.Info, &coupon)
	}

	coupon.ID = resp.Info.ID

	resp1, err := Get(context.Background(), &npool.GetCouponPoolRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCouponPool(t, resp1.Info, &coupon)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetCouponPoolsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetByAppReleaser(context.Background(), &npool.GetCouponPoolsByAppReleaserRequest{
		AppID:  resp.Info.AppID,
		UserID: resp.Info.ReleaseByUserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := Update(context.Background(), &npool.UpdateCouponPoolRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertCouponPool(t, resp4.Info, &coupon)
	}
}
