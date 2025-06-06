package crawler_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"

	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/pkg/appapi/crawler"
	"github.com/ryohidaka/go-pixiv/testutil/apptest"
	"github.com/stretchr/testify/assert"
)

func ExamplePixivCrawler_FetchAllUserFollowing() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv Crawler
	c, _ := crawler.NewCrawler(refreshToken)

	// Fetch all user following for user ID 11 (Pixiv official account)
	users, _ := c.FetchAllUserFollowing(11, nil)

	for _, v := range users {
		// Print the user name
		fmt.Println("Name:", v.User.Name)
	}
}

func TestFetchAllUserFollowing(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock: authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth_token")

		// Mock: page 1 of user following
		urlPage1 := appapi.AppHosts + "v1/user/following?restrict=public&user_id=11"
		err := apptest.MockResponseFromFile("GET", urlPage1, "user_following")
		assert.NoError(t, err)

		// Mock: page 2 of user following (with offset)
		urlPage2 := appapi.AppHosts + "v1/user/following?offset=30&restrict=public&user_id=11"
		err = apptest.MockResponseFromFile("GET", urlPage2, "user_following_end")
		assert.NoError(t, err)

		// Initialize Crawler instance
		crawler, err := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		public := models.Public
		opts := &pixiv.UserFollowingOptions{
			Restrict: &public,
		}

		// Call the main function (no sleep between requests)
		users, err := crawler.FetchAllUserFollowing(11, opts)
		assert.NoError(t, err)
		assert.Len(t, users, 2) // One user per page

		// Verify contents of each user
		assert.Equal(t, uint64(11), users[0].User.ID)
		assert.Equal(t, uint64(11), users[1].User.ID)
		assert.Equal(t, "pixiv事務局", users[0].User.Name)
		assert.Equal(t, "pixiv事務局", users[1].User.Name)
	})
}
