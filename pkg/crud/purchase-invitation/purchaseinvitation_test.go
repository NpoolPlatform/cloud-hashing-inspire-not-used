package purchaseinvitation

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

func assertPurchaseInvitation(t *testing.T, actual, expected *npool.PurchaseInvitation) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.OrderID, expected.OrderID)
	assert.Equal(t, actual.InvitationCodeID, expected.InvitationCodeID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	invitation := npool.PurchaseInvitation{
		AppID:            uuid.New().String(),
		OrderID:          uuid.New().String(),
		InvitationCodeID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreatePurchaseInvitationRequest{
		Info: &invitation,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertPurchaseInvitation(t, resp.Info, &invitation)
	}

	invitation.ID = resp.Info.ID

	resp1, err := Get(context.Background(), &npool.GetPurchaseInvitationRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertPurchaseInvitation(t, resp1.Info, &invitation)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetPurchaseInvitationsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetByAppOrder(context.Background(), &npool.GetPurchaseInvitationByAppOrderRequest{
		AppID:   resp.Info.AppID,
		OrderID: resp.Info.OrderID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertPurchaseInvitation(t, resp3.Info, &invitation)
	}

	resp4, err := Update(context.Background(), &npool.UpdatePurchaseInvitationRequest{
		Info: &invitation,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertPurchaseInvitation(t, resp4.Info, &invitation)
	}
}
