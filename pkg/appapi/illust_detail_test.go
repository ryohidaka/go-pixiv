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

func ExampleAppPixivAPI_IllustDetail() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch illust details for illust ID 129899459
	illust, _ := app.IllustDetail(129899459)

	// Print the illust title and type
	fmt.Println("Title:", illust.Title)
	fmt.Println("Type:", illust.Type)

	// Output:
	// Title: 「出張版！アクション月例漫画賞」いよいよ開催！
	// Type: illust
}

func TestIllustDetail(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth/token", "../../testutil")

		// Mock the user illusts response
		url := appapi.AppHosts + "v1/illust/detail?illust_id=129899459"
		err := apptest.MockResponseFromFile("GET", url, "v1/illust/detail", "../../testutil")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Call the IllustDetail method
		illust, err := api.IllustDetail(129899459)
		assert.NoError(t, err)

		// Check contents of the illustration
		assert.Equal(t, uint64(129899459), illust.ID)
		assert.Equal(t, "「出張版！アクション月例漫画賞」いよいよ開催！", illust.Title)
		assert.Equal(t, models.IllustType("illust"), models.IllustType(illust.Type))
		assert.Equal(t, "https://i.pximg.net/c/600x1200_90/img-master/img/2025/05/01/11/19/11/129899459_p0_master1200.jpg", illust.ImageURLs.Large)
	})
}
