package pixiv_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserFollower(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock the user follower response
		url := pixiv.AppHosts + "v1/user/follower?restrict=public&user_id=11"
		err := testutil.MockResponseFromFile("GET", url, "v1/user/follower")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		public := models.Public
		opts := &pixiv.UserFollowerOptions{
			Restrict: &public,
		}

		// Call the UserFollower method
		users, next, err := api.UserFollower(11, opts)
		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, 30, next)

		// Check contents of the first user
		user := users[0]
		assert.Equal(t, uint64(11), user.User.ID)
		assert.Equal(t, "pixiv事務局", user.User.Name)
	})
}
