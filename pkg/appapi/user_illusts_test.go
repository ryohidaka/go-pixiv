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

func ExampleAppPixivAPI_UserIllusts() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch user illusts for user ID 11 (Pixiv official account)
	illusts, _, _ := app.UserIllusts(11)

	for _, v := range illusts {
		// Print the illust title
		fmt.Println("Title:", v.Title)
	}
}

func TestUserIllust(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth_token")

		// Mock the user illusts response
		url := appapi.AppHosts + "v1/user/illusts?filter=for_ios&user_id=11"
		err := apptest.MockResponseFromFile("GET", url, "user_illusts")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Prepare options
		filter := "for_ios"
		opts := pixiv.UserIllustsOptions{
			Filter: &filter,
		}

		// Call the UserIllusts method
		illusts, next, err := api.UserIllusts(11, opts)
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
