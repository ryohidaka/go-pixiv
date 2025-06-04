package crawler_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/crawler"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func ExamplePixivCrawler_FetchAllUserIllusts() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv Crawler
	c, _ := crawler.NewCrawler(refreshToken)

	// Fetch all user illusts for user ID 11 (Pixiv official account)
	illusts, _ := c.FetchAllUserIllusts(11, nil)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func TestFetchAllUserIllusts(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock: authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token", "../testutil")

		// Mock: page 1 of user illustrations
		urlPage1 := pixiv.AppHosts + "v1/user/illusts?filter=for_ios&user_id=11"
		err := testutil.MockResponseFromFile("GET", urlPage1, "v1/user/illusts", "../testutil")
		assert.NoError(t, err)

		// Mock: page 2 of user illustrations (with offset)
		urlPage2 := pixiv.AppHosts + "v1/user/illusts?filter=for_ios&offset=30&user_id=11"
		err = testutil.MockResponseFromFile("GET", urlPage2, "v1/user/illusts_end", "../testutil")
		assert.NoError(t, err)

		// Initialize Crawler instance
		crawler, err := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)

		// Set request options
		filter := "for_ios"
		opts := &pixiv.UserIllustsOptions{
			Filter: &filter,
		}

		// Call the main function (no sleep between requests)
		illusts, err := crawler.FetchAllUserIllusts(11, opts, 0)
		assert.NoError(t, err)
		assert.Len(t, illusts, 2) // One illustration per page

		// Verify contents of each illustration
		assert.Equal(t, uint64(129899459), illusts[0].ID)
		assert.Equal(t, uint64(129899459), illusts[1].ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[0].Title)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illusts[1].Title)
	})
}
