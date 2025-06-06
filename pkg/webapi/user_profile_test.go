package webapi_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/pkg/webapi"
	"github.com/ryohidaka/go-pixiv/testutil/webtest"
	"github.com/stretchr/testify/assert"
)

func TestUserProfile(t *testing.T) {
	webtest.WithMockHTTP(t, func() {

		// Mock the user detail response
		url := webapi.ApiHosts + "user/11/profile/all"
		err := webtest.MockResponseFromFile("GET", url, "user_profile")
		assert.NoError(t, err)

		// Initialize the WebPixivAPI instance
		app, err := pixiv.NewWebApp("dummy-phpsessid")
		assert.NoError(t, err)

		// Call the UserProfile method
		user, err := app.UserProfile(11)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		// Check illusts
		assert.NotNil(t, user.Illusts)

		// Check manga
		assert.NotNil(t, user.Manga)

		// Check novels
		assert.NotNil(t, user.Novels)
		assert.Equal(t, "1462193", user.NovelSeries[0].ID)

		// Check pickup
		assert.NotNil(t, user.Pickup)
		assert.Equal(t, "129155745", user.Pickup[0].ID)
	})
}
