package appapi_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleAppPixivAPI_IllustBookmarkDetail() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch illust bookmark details for illust ID 129899459
	bookmark, _ := app.IllustBookmarkDetail(129899459)

	// Print the illust title and type
	fmt.Println("IsBookmarked:", bookmark.BookmarkDetail.IsBookmarked)
}

func TestIllustBookmarkDetail(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth/token", "../../testutil")

		// Mock the user illusts response
		url := appapi.AppHosts + "v2/illust/bookmark/detail?illust_id=129899459"
		err := testutil.MockResponseFromFile("GET", url, "v2/illust/bookmark/detail", "../../testutil")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Call the IllustBookmarkDetail method
		bookmark, err := api.IllustBookmarkDetail(129899459)
		assert.NoError(t, err)

		// Check contents of the bookmark
		assert.Equal(t, true, bookmark.BookmarkDetail.IsBookmarked)

		tag := bookmark.BookmarkDetail.Tags[0]
		assert.Equal(t, "漫画", tag.Name)
		assert.Equal(t, false, tag.IsRegistered)
	})
}
