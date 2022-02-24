package couponallocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-allocated" //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"      //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/test-init"             //nolint
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

func assertCouponAllocatedDetail(t *testing.T, info *npool.CouponAllocatedDetail, coupon *npool.CouponAllocated, couponPool *npool.CouponPool) {
	assert.Equal(t, info.Allocated.AppID, coupon.AppID)
	assert.Equal(t, info.Allocated.UserID, coupon.UserID)
	assert.Equal(t, info.Allocated.Type, coupon.Type)

	assert.Equal(t, info.Coupon.ID, couponPool.ID)
	assert.Equal(t, info.Coupon.Denomination, couponPool.Denomination)
	assert.Equal(t, info.Coupon.Circulation, couponPool.Circulation)
	assert.Equal(t, info.Coupon.ReleaseByUserID, couponPool.ReleaseByUserID)
	assert.Equal(t, info.Coupon.Start, couponPool.Start)
	assert.Equal(t, info.Coupon.DurationDays, couponPool.DurationDays)
	assert.Equal(t, info.Coupon.Message, couponPool.Message)
	assert.Equal(t, info.Coupon.Name, couponPool.Name)
}

func TestGetDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appID := uuid.New().String()

	coupon := npool.CouponPool{
		AppID:           appID,
		Denomination:    10.9,
		Circulation:     10,
		ReleaseByUserID: uuid.New().String(),
		Start:           uint32(time.Now().Unix()),
		DurationDays:    10,
		Message:         "5-1 Promotion",
		Name:            fmt.Sprintf("5-1 Promotion-%v", appID),
	}
	couponResp, err := couponpool.Create(context.Background(), &npool.CreateCouponPoolRequest{
		Info: &coupon,
	})
	assert.Nil(t, err)

	couponAllocated := npool.CouponAllocated{
		AppID:    coupon.AppID,
		UserID:   uuid.New().String(),
		Type:     "coupon",
		CouponID: couponResp.Info.ID,
	}
	allocatedResp, err := couponallocated.Create(context.Background(), &npool.CreateCouponAllocatedRequest{
		Info: &couponAllocated,
	})
	assert.Nil(t, err)

	resp, err := Get(context.Background(), &npool.GetCouponAllocatedDetailRequest{
		ID: allocatedResp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.Allocated.ID, uuid.UUID{}.String())
		assertCouponAllocatedDetail(t, resp.Info, allocatedResp.Info, couponResp.Info)
	}

	resp1, err := GetByApp(context.Background(), &npool.GetCouponsAllocatedDetailByAppRequest{
		AppID: allocatedResp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp1.Infos), 1)
	}

	resp2, err := GetByAppUser(context.Background(), &npool.GetCouponsAllocatedDetailByAppUserRequest{
		AppID:  allocatedResp.Info.AppID,
		UserID: allocatedResp.Info.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}
}
