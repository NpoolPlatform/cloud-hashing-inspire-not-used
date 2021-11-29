package couponallocated

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

func assertCouponAllocated(t *testing.T, actual, expected *npool.CouponAllocated) {
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.CouponID, expected.CouponID)
	assert.Equal(t, actual.Used, expected.Used)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coupon := npool.CouponAllocated{
		AppID:    uuid.New().String(),
		UserID:   uuid.New().String(),
		Type:     "discount",
		CouponID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateCouponAllocatedRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCouponAllocated(t, resp.Info, &coupon)
	}

	coupon.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateCouponAllocatedRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCouponAllocated(t, resp1.Info, &coupon)
	}

	resp2, err := Get(context.Background(), &npool.GetCouponAllocatedRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertCouponAllocated(t, resp2.Info, &coupon)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetCouponsAllocatedByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByAppUser(context.Background(), &npool.GetCouponsAllocatedByAppUserRequest{
		AppID:  resp.Info.AppID,
		UserID: resp.Info.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp4.Infos), 1)
	}
}
