package pixiv_test

import (
	"fmt"
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleAppPixivAPI_UserFollowing() {
	// Get the refresh token used for authentication
	refreshToken := testutil.GetRefreshToken()

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch user following for user ID 11 (Pixiv official account)
	users, _, _ := app.UserFollowing(11, nil)

	for _, v := range users {
		// Print the user name
		fmt.Println("Name:", v.User.Name)
	}
}

func TestUserFollowing(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock the user illusts response
		url := pixiv.AppHosts + "v1/user/following?restrict=public&user_id=11"
		err := testutil.MockResponseFromFile("GET", url, "v1/user/following")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		public := models.Public
		opts := &pixiv.UserFollowingOptions{
			Restrict: &public,
		}

		// Call the UserFollowing method
		users, next, err := api.UserFollowing(11, opts)
		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, 30, next)

		// Check contents of the first user
		user := users[0]
		assert.Equal(t, uint64(11), user.User.ID)
		assert.Equal(t, "pixiv事務局", user.User.Name)
	})
}
