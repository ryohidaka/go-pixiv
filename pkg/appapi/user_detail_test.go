package appapi_test

import (
	"fmt"
	"os"

	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/testutil/apptest"

	"github.com/stretchr/testify/assert"
)

func ExampleAppPixivAPI_UserDetail() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv App API client
	app, _ := pixiv.NewApp(refreshToken)

	// Fetch user details for user ID 11 (Pixiv official account)
	user, _ := app.UserDetail(11)

	// Print the user's name and account
	fmt.Println("Name:", user.User.Name)
	fmt.Println("Account:", user.User.Account)

	// Output:
	// Name: pixiv事務局
	// Account: pixiv
}

func TestUserDetail(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth_token")

		// Mock the user detail response
		url := appapi.AppHosts + "v1/user/detail?filter=for_ios&user_id=11"
		err := apptest.MockResponseFromFile("GET", url, "user_detail")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, _ := pixiv.NewApp("dummy-refresh-token")

		filter := "for_ios"

		// Call the UserDetail method
		detail, err := api.UserDetail(11, pixiv.UserDetailOptions{
			Filter: &filter,
		})
		assert.NoError(t, err)
		assert.NotNil(t, detail)

		// Check user details
		assert.NotNil(t, detail.User)
		assert.Equal(t, uint64(11), detail.User.ID)
		assert.Equal(t, "pixiv事務局", detail.User.Name)
		assert.Equal(t, "pixiv", detail.User.Account)

		// Check profile details
		assert.NotNil(t, detail.Profile)
		assert.Equal(t, "", detail.Profile.Gender)
		assert.Equal(t, "日本 東京都", detail.Profile.Region)

		// Check profile publicity details
		assert.NotNil(t, detail.ProfilePublicity)
		assert.Equal(t, "public", detail.ProfilePublicity.Gender)

		// Check workspace details
		assert.NotNil(t, detail.Workspace)
		assert.Equal(t, "", detail.Workspace.Tablet)
	})
}
