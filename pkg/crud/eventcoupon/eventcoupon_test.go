package eventcoupon

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

func assertEventCoupon(t *testing.T, actual, expected *npool.EventCoupon) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.ActivityID, expected.ActivityID)
	assert.Equal(t, actual.Event, expected.Event)
	assert.Equal(t, actual.CouponID, expected.CouponID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coupon := npool.EventCoupon{
		AppID:      uuid.New().String(),
		ActivityID: uuid.New().String(),
		Event:      "sharing",
		CouponID:   uuid.New().String(),
		Count:      1,
		Type:       "discount",
	}

	resp, err := Create(context.Background(), &npool.CreateEventCouponRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertEventCoupon(t, resp.Info, &coupon)
	}

	coupon.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateEventCouponRequest{
		Info: &coupon,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertEventCoupon(t, resp1.Info, &coupon)
	}

	resp2, err := Get(context.Background(), &npool.GetEventCouponRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertEventCoupon(t, resp2.Info, &coupon)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetEventCouponsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByAppActivityEvent(context.Background(), &npool.GetEventCouponsByAppActivityEventRequest{
		AppID:      resp.Info.AppID,
		ActivityID: resp.Info.ActivityID,
		Event:      resp.Info.Event,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp4.Infos), 1)
	}
}
