package activity

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertActivity(t *testing.T, actual, expected *npool.Activity) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.CreatedBy, expected.CreatedBy)
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.End, expected.End)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	act := npool.Activity{
		AppID:     uuid.New().String(),
		CreatedBy: uuid.New().String(),
		Name:      uuid.New().String(),
		Start:     uint32(time.Now().Unix() + 1000),
		End:       uint32(time.Now().Unix() + 2000),
	}

	resp, err := Create(context.Background(), &npool.CreateActivityRequest{
		Info: &act,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertActivity(t, resp.Info, &act)
	}

	act.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateActivityRequest{
		Info: &act,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertActivity(t, resp1.Info, &act)
	}

	resp2, err := Get(context.Background(), &npool.GetActivityRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertActivity(t, resp2.Info, &act)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetActivitiesByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByAppName(context.Background(), &npool.GetActivityByAppNameRequest{
		AppID: resp.Info.AppID,
		Name:  resp.Info.Name,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertActivity(t, resp4.Info, &act)
	}
}
