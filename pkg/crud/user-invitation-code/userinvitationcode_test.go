package userinvitationcode

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

func assertUserInvitationCode(t *testing.T, actual, expected *npool.UserInvitationCode) {
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.NotEqual(t, actual.InvitationCode, "")
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userInvitationCode := npool.UserInvitationCode{
		AppID:  uuid.New().String(),
		UserID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateUserInvitationCodeRequest{
		Info: &userInvitationCode,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUserInvitationCode(t, resp.Info, &userInvitationCode)
	}
}
