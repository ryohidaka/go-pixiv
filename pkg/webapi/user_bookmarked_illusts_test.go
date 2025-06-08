package webapi_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/pkg/webapi"
	"github.com/ryohidaka/go-pixiv/testutil/webtest"
	"github.com/stretchr/testify/assert"
)

func ExampleWebPixivAPI_UserBookmarksIllusts() {
	// Get the PHPSESSID used for authentication
	phpsessid := os.Getenv("PHPSESSID")

	// Create a new Pixiv Web API client
	app, _ := pixiv.NewWebApp(phpsessid)

	opts := pixiv.UserBookmarksIllustsOptions{
		Offset: 0,
		Limit:  10,
	}

	// Fetch bookmarks of users.
	data, total, _ := app.UserBookmarksIllusts(11, opts)

	// Print the bookmark illusts title
	for _, illust := range data.Works {
		fmt.Println("Title: ", illust.Title)
	}

	// Print the total
	fmt.Println("Total:", total)
}

func TestUserBookmarksIllusts(t *testing.T) {
	webtest.WithMockHTTP(t, func() {
		// Mock the user following response
		url := webapi.ApiHosts + "user/11/illusts/bookmarks?limit=3&offset=0&rest=show&tag="
		err := webtest.MockResponseFromFile("GET", url, "user_bookmarked_illusts")
		assert.NoError(t, err)

		// Initialize the WebPixivAPI instance
		app, err := pixiv.NewWebApp("dummy-phpsessid")
		assert.NoError(t, err)

		// Call the UserBookmarksIllusts method
		opts := pixiv.UserBookmarksIllustsOptions{
			Offset:   0,
			Limit:    3,
			Restrict: core.Show,
		}

		data, total, err := app.UserBookmarksIllusts(11, opts)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		// Check total
		assert.Equal(t, 3, int(total))

		expected := []struct {
			id    string
			title string
		}{
			{"130967689", "comic POOL5月コミックス発売情報"},
			{"131093702", "マンガ投稿企画『ツキコミ〜裏を見ちゃった！〜』開催"},
			{"131202453", "pixivが大阪・関西万博で期間限定展示を開催"},
		}

		for i, exp := range expected {
			assert.Equal(t, exp.id, data.Works[i].ID)
			assert.Equal(t, exp.title, data.Works[i].Title)
		}
	})
}
