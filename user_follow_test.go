package pixiv_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleAppPixivAPI_UserFollowAdd() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Send a follow request to user ID 11 (Pixiv official account)
	app.UserFollowAdd(11)
}

func ExampleAppPixivAPI_UserFollowDelete() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Send a unfollow request to user ID 11 (Pixiv official account)
	app.UserFollowDelete(11)
}

// TestUserFollowAdd tests the UserFollowAdd method of AppPixivAPI
func TestUserFollowAdd(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock the user follow response
		url := pixiv.AppHosts + "v1/user/follow/add"
		err := testutil.MockResponseFromFile("POST", url, "empty")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Call UserFollowAdd with a specific userID and restrict mode
		userID := uint64(12345678)
		restrict := models.Private

		ok, err := api.UserFollowAdd(userID, restrict)
		assert.NoError(t, err)
		assert.True(t, ok, "UserFollowAdd should return true on success")

		// Verify request was made as expected
		info := httpmock.GetCallCountInfo()
		key := fmt.Sprintf("POST %s", url)
		if info[key] != 1 {
			t.Errorf("expected 1 POST request to %s, but got %d", url, info[key])
		}
	})
}

// TestUserFollowDelete tests the UserFollowDelete method of AppPixivAPI
func TestUserFollowDelete(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock the user unfollow response
		url := pixiv.AppHosts + "v1/user/follow/delete"
		err := testutil.MockResponseFromFile("POST", url, "empty")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Call UserFollowDelete with a specific userID
		userID := uint64(12345678)

		ok, err := api.UserFollowDelete(userID)
		assert.NoError(t, err)
		assert.True(t, ok, "UserFollowDelete should return true on success")

		// Verify request was made as expected
		info := httpmock.GetCallCountInfo()
		key := fmt.Sprintf("POST %s", url)
		if info[key] != 1 {
			t.Errorf("expected 1 POST request to %s, but got %d", url, info[key])
		}
	})
}
