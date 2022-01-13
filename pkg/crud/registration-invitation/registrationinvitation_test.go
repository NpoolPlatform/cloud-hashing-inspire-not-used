package registrationinvitation

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

func assertRegistrationInvitation(t *testing.T, actual, expected *npool.RegistrationInvitation) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.InviterID, expected.InviterID)
	assert.Equal(t, actual.InviteeID, expected.InviteeID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	invitation := npool.RegistrationInvitation{
		AppID:     uuid.New().String(),
		InviterID: uuid.New().String(),
		InviteeID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateRegistrationInvitationRequest{
		Info: &invitation,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertRegistrationInvitation(t, resp.Info, &invitation)
	}

	invitation.ID = resp.Info.ID

	resp1, err := Get(context.Background(), &npool.GetRegistrationInvitationRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertRegistrationInvitation(t, resp1.Info, &invitation)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetRegistrationInvitationsByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetByAppInviter(context.Background(), &npool.GetRegistrationInvitationsByAppInviterRequest{
		AppID:     resp.Info.AppID,
		InviterID: resp.Info.InviterID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByAppInvitee(context.Background(), &npool.GetRegistrationInvitationByAppInviteeRequest{
		AppID:     resp.Info.AppID,
		InviteeID: resp.Info.InviteeID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertRegistrationInvitation(t, resp4.Info, &invitation)
	}

	resp5, err := Update(context.Background(), &npool.UpdateRegistrationInvitationRequest{
		Info: &invitation,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.Info.ID, resp.Info.ID)
		assertRegistrationInvitation(t, resp5.Info, &invitation)
	}
}
