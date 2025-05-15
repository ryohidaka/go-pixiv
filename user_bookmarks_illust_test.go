package pixiv_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleAppPixivAPI_UserBookmarksIllust() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch user bookmarks illust for user ID 11 (Pixiv official account)
	illusts, _, _ := app.UserBookmarksIllust(11)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func ExampleAppPixivAPI_FetchAllBookmarkedIllusts() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch all user bookmarks illust for user ID 11 (Pixiv official account)
	illusts, _ := app.FetchAllBookmarkedIllusts(11, nil)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func TestUserBookmarksIllust(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock the user bookmarks illust response
		url := pixiv.AppHosts + "v1/user/bookmarks/illust?filter=for_ios&restrict=public&user_id=11"
		err := testutil.MockResponseFromFile("GET", url, "v1/user/bookmarks/illust")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		filter := "for_ios"
		public := models.Public
		opts := pixiv.UserBookmarksIllustOptions{
			Filter:   &filter,
			Restrict: &public,
		}

		// Call the UserBookmarksIllust method
		illusts, next, err := api.UserBookmarksIllust(11, opts)
		assert.NoError(t, err)
		assert.Len(t, illusts, 1)
		assert.Equal(t, 129899459, next)

		// Check contents of the first illustration
		illust := illusts[0]
		assert.Equal(t, uint64(129899459), illust.ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illust.Title)
		assert.Equal(t, models.IllustType("illust"), models.IllustType(illust.Type))
		assert.Equal(t, "https://i.pximg.net/c/600x1200_90/img-master/img/2025/05/01/11/19/11/129899459_p0_master1200.jpg", illust.ImageURLs.Large)
	})
}

func TestFetchAllBookmarkedIllusts(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock: authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock: page 1 of user bookmarks
		urlPage1 := pixiv.AppHosts + "v1/user/bookmarks/illust?filter=for_ios&restrict=public&user_id=11"
		err := testutil.MockResponseFromFile("GET", urlPage1, "v1/user/bookmarks/illust")
		assert.NoError(t, err)

		// Mock: page 2 of user bookmarks (with next_id)
		urlPage2 := pixiv.AppHosts + "v1/user/bookmarks/illust?filter=for_ios&max_bookmark_id=129899459&restrict=public&user_id=11"
		err = testutil.MockResponseFromFile("GET", urlPage2, "v1/user/bookmarks/illust_end")
		assert.NoError(t, err)

		// Initialize API instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Set request options
		filter := "for_ios"
		public := models.Public
		opts := &pixiv.UserBookmarksIllustOptions{
			Filter:   &filter,
			Restrict: &public,
		}

		// Call the main function (no sleep between requests)
		illusts, err := api.FetchAllBookmarkedIllusts(11, opts, 0)
		assert.NoError(t, err)
		assert.Len(t, illusts, 2) // One illustration per page

		// Verify contents of each illustration
		assert.Equal(t, uint64(129899459), illusts[0].ID)
		assert.Equal(t, uint64(129899459), illusts[1].ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[0].Title)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[1].Title)
	})
}
