package pixiv_test

import (
	"go-pixiv"
	"go-pixiv/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDetail(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth.json")

		// Mock the user detail response
		url := pixiv.AppHosts + "v1/user/detail?filter=for_ios&user_id=11"
		err := testutil.MockResponseFromFile("GET", url, "user-detail.json")
		assert.NoError(t, err)

		// Initialize the AppPixivAPI instance
		api, _ := pixiv.NewApp("dummy-refresh-token")

		filter := "for_ios"

		// Call the UserDetail method
		detail, err := api.UserDetail(11, &pixiv.UserDetailOptions{
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
