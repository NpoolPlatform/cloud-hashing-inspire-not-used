package discountpool

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

func assertDiscountPool(t *testing.T, actual, expected *npool.DiscountPool) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.Discount, expected.Discount)
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

	discount := npool.DiscountPool{
		AppID:           uuid.New().String(),
		Discount:        5,
		ReleaseByUserID: uuid.New().String(),
		Start:           uint32(time.Now().Unix()),
		DurationDays:    10,
		Message:         "For test discount",
		Name:            fmt.Sprintf("Test Coupon-%v", uuid.New().String()),
	}

	resp, err := Create(context.Background(), &npool.CreateDiscountPoolRequest{
		Info: &discount,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertDiscountPool(t, resp.Info, &discount)
	}

	discount.ID = resp.Info.ID

	resp1, err := Get(context.Background(), &npool.GetDiscountPoolRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertDiscountPool(t, resp1.Info, &discount)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetDiscountPoolsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetByAppReleaser(context.Background(), &npool.GetDiscountPoolsByAppReleaserRequest{
		AppID:  resp.Info.AppID,
		UserID: resp.Info.ReleaseByUserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := Update(context.Background(), &npool.UpdateDiscountPoolRequest{
		Info: &discount,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertDiscountPool(t, resp4.Info, &discount)
	}
}
