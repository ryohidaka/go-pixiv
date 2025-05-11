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

func ExampleAppPixivAPI_IllustFollow() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch illust from user follows
	illusts, _, _ := app.IllustFollow(nil)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func ExampleAppPixivAPI_FetchAllIllustFollows() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch all illust from user follows
	illusts, _ := app.FetchAllIllustFollows(nil)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func TestIllustFollow(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock the illust follow response
		url := pixiv.AppHosts + "v2/illust/follow?restrict=public"
		err := testutil.MockResponseFromFile("GET", url, "v2/illust/follow")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		public := models.Public
		opts := &pixiv.IllustFollowOptions{
			Restrict: &public,
		}

		// Call the IllustFollow method
		illusts, next, err := api.IllustFollow(opts)
		assert.NoError(t, err)
		assert.Len(t, illusts, 1)
		assert.Equal(t, 30, next)

		// Check contents of the first illustration
		illust := illusts[0]
		assert.Equal(t, uint64(129899459), illust.ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illust.Title)
		assert.Equal(t, models.IllustType("illust"), models.IllustType(illust.Type))
		assert.Equal(t, "https://i.pximg.net/c/600x1200_90/img-master/img/2025/05/01/11/19/11/129899459_p0_master1200.jpg", illust.ImageURLs.Large)
	})
}

func TestFetchAllIllustFollows(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock: authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")

		// Mock: page 1 of illust follow
		urlPage1 := pixiv.AppHosts + "v2/illust/follow?restrict=public"
		err := testutil.MockResponseFromFile("GET", urlPage1, "v2/illust/follow")
		assert.NoError(t, err)

		// Mock: page 2 of illust follow (with offset)
		urlPage2 := pixiv.AppHosts + "v2/illust/follow?offset=30&restrict=public"
		err = testutil.MockResponseFromFile("GET", urlPage2, "v2/illust/follow_end")
		assert.NoError(t, err)

		// Initialize API instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		public := models.Public
		opts := &pixiv.IllustFollowOptions{
			Restrict: &public,
		}

		// Call the main function (no sleep between requests)
		illusts, err := api.FetchAllIllustFollows(opts)
		assert.NoError(t, err)
		assert.Len(t, illusts, 2) // One illustration per page

		// Verify contents of each illustration
		assert.Equal(t, uint64(129899459), illusts[0].ID)
		assert.Equal(t, uint64(129899459), illusts[1].ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[0].Title)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[1].Title)
	})
}
