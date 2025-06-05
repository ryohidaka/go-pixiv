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

func ExamplePixivCrawler_FetchAllBookmarkedIllusts() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv Crawler
	c, _ := crawler.NewCrawler(refreshToken)

	// Fetch all user bookmarks illust for user ID 11 (Pixiv official account)
	illusts, _ := c.FetchAllBookmarkedIllusts(11, nil)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func TestFetchAllBookmarkedIllusts(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock: authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth_token")

		// Mock: page 1 of user bookmarks
		urlPage1 := appapi.AppHosts + "v1/user/bookmarks/illust?filter=for_ios&restrict=public&user_id=11"
		err := apptest.MockResponseFromFile("GET", urlPage1, "bookmarks_illust")
		assert.NoError(t, err)

		// Mock: page 2 of user bookmarks (with next_id)
		urlPage2 := appapi.AppHosts + "v1/user/bookmarks/illust?filter=for_ios&max_bookmark_id=129899459&restrict=public&user_id=11"
		err = apptest.MockResponseFromFile("GET", urlPage2, "bookmarks_illust_end")
		assert.NoError(t, err)

		// Initialize Crawler instance
		crawler, err := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)

		// Set request options
		filter := "for_ios"
		public := models.Public
		opts := &pixiv.UserBookmarksIllustOptions{
			Filter:   &filter,
			Restrict: &public,
		}

		// Call the main function (no sleep between requests)
		illusts, err := crawler.FetchAllBookmarkedIllusts(11, opts, 0)
		assert.NoError(t, err)
		assert.Len(t, illusts, 2) // One illustration per page

		// Verify contents of each illustration
		assert.Equal(t, uint64(129899459), illusts[0].ID)
		assert.Equal(t, uint64(129899459), illusts[1].ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[0].Title)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[1].Title)
	})
}
