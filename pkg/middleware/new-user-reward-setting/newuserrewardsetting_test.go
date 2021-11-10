package newuserrewardsetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/coupon-pool"             //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/crud/new-user-reward-setting" //nolint
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/test-init"                    //nolint

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

func assertNewUserRewardSettingDetail(t *testing.T, info *npool.NewUserRewardSettingDetail, setting *npool.NewUserRewardSetting, regCoupon, kycCoupon *npool.CouponPool) {
	assert.Equal(t, info.AppID, setting.AppID)
	assert.Equal(t, info.AutoGenerateInvitationCode, setting.AutoGenerateInvitationCode)

	assert.Equal(t, info.RegistrationCoupon.ID, regCoupon.ID)
	assert.Equal(t, info.RegistrationCoupon.Denomination, regCoupon.Denomination)
	assert.Equal(t, info.RegistrationCoupon.Circulation, regCoupon.Circulation)
	assert.Equal(t, info.RegistrationCoupon.ReleaseByUserID, regCoupon.ReleaseByUserID)
	assert.Equal(t, info.RegistrationCoupon.Start, regCoupon.Start)
	assert.Equal(t, info.RegistrationCoupon.DurationDays, regCoupon.DurationDays)
	assert.Equal(t, info.RegistrationCoupon.Message, regCoupon.Message)
	assert.Equal(t, info.RegistrationCoupon.Name, regCoupon.Name)

	assert.Equal(t, info.KycCoupon.ID, kycCoupon.ID)
	assert.Equal(t, info.KycCoupon.Denomination, kycCoupon.Denomination)
	assert.Equal(t, info.KycCoupon.Circulation, kycCoupon.Circulation)
	assert.Equal(t, info.KycCoupon.ReleaseByUserID, kycCoupon.ReleaseByUserID)
	assert.Equal(t, info.KycCoupon.Start, kycCoupon.Start)
	assert.Equal(t, info.KycCoupon.DurationDays, kycCoupon.DurationDays)
	assert.Equal(t, info.KycCoupon.Message, kycCoupon.Message)
	assert.Equal(t, info.KycCoupon.Name, kycCoupon.Name)
}

func TestGetDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appID := uuid.New().String()

	regCoupon := npool.CouponPool{
		AppID:           appID,
		Denomination:    10.9,
		Circulation:     10,
		ReleaseByUserID: uuid.New().String(),
		Start:           uint32(time.Now().Unix()),
		DurationDays:    10,
		Message:         "New User Registration coupon",
		Name:            fmt.Sprintf("Registration coupon %v", appID),
	}
	regCouponResp, err := couponpool.Create(context.Background(), &npool.CreateCouponPoolRequest{
		Info: &regCoupon,
	})
	assert.Nil(t, err)

	kycCoupon := npool.CouponPool{
		AppID:           appID,
		Denomination:    10.9,
		Circulation:     10,
		ReleaseByUserID: uuid.New().String(),
		Start:           uint32(time.Now().Unix()),
		DurationDays:    10,
		Message:         "New User Kyc coupon",
		Name:            fmt.Sprintf("Kyc coupon %v", appID),
	}
	kycCouponResp, err := couponpool.Create(context.Background(), &npool.CreateCouponPoolRequest{
		Info: &kycCoupon,
	})
	assert.Nil(t, err)

	setting := npool.NewUserRewardSetting{
		AppID:                      appID,
		RegistrationCouponID:       regCouponResp.Info.ID,
		KycCouponID:                kycCouponResp.Info.ID,
		AutoGenerateInvitationCode: false,
	}
	settingResp, err := newuserrewardsetting.Create(context.Background(), &npool.CreateNewUserRewardSettingRequest{
		Info: &setting,
	})
	assert.Nil(t, err)

	resp, err := Get(context.Background(), &npool.GetNewUserRewardSettingDetailRequest{
		ID: settingResp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info.ID, settingResp.Info.ID)
		assertNewUserRewardSettingDetail(t, resp.Info, settingResp.Info, regCouponResp.Info, kycCouponResp.Info)
	}
}
