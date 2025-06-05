package appapi_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/testutil/apptest"
	"github.com/stretchr/testify/assert"
)

func ExampleAppPixivAPI_IllustFollow() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch illust from user follows
	illusts, _, _ := app.IllustFollow()

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func TestIllustFollow(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth/token", "../../testutil")

		// Mock the illust follow response
		url := appapi.AppHosts + "v2/illust/follow?restrict=public"
		err := apptest.MockResponseFromFile("GET", url, "v2/illust/follow", "../../testutil")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		public := models.Public
		opts := pixiv.IllustFollowOptions{
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
