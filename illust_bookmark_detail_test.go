package pixiv_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func TestIllustBookmarkDetail(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth.json")

		// Mock the user illusts response
		url := pixiv.AppHosts + "v2/illust/bookmark/detail?illust_id=129899459"
		err := testutil.MockResponseFromFile("GET", url, "illust-bookmark-detail.json")
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
