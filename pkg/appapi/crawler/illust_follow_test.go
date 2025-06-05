package crawler_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"

	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/pkg/appapi/crawler"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func ExamplePixivCrawler_FetchAllIllustFollows() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv Crawler
	c, _ := crawler.NewCrawler(refreshToken)

	// Fetch all illust from user follows
	illusts, _ := c.FetchAllIllustFollows(nil)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func TestFetchAllIllustFollows(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock: authentication response
		_ = testutil.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth/token", "../../../testutil")

		// Mock: page 1 of illust follow
		urlPage1 := appapi.AppHosts + "v2/illust/follow?restrict=public"
		err := testutil.MockResponseFromFile("GET", urlPage1, "v2/illust/follow", "../../../testutil")
		assert.NoError(t, err)

		// Mock: page 2 of illust follow (with offset)
		urlPage2 := appapi.AppHosts + "v2/illust/follow?offset=30&restrict=public"
		err = testutil.MockResponseFromFile("GET", urlPage2, "v2/illust/follow_end", "../../../testutil")
		assert.NoError(t, err)

		// Initialize Crawler instance
		crawler, err := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		public := models.Public
		opts := &pixiv.IllustFollowOptions{
			Restrict: &public,
		}

		// Call the main function (no sleep between requests)
		illusts, err := crawler.FetchAllIllustFollows(opts)
		assert.NoError(t, err)
		assert.Len(t, illusts, 2) // One illustration per page

		// Verify contents of each illustration
		assert.Equal(t, uint64(129899459), illusts[0].ID)
		assert.Equal(t, uint64(129899459), illusts[1].ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[0].Title)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[1].Title)
	})
}
