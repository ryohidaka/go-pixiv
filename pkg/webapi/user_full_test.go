package webapi_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/pkg/webapi"
	"github.com/ryohidaka/go-pixiv/testutil/webtest"
	"github.com/stretchr/testify/assert"
)

func TestUserFull(t *testing.T) {
	webtest.WithMockHTTP(t, func() {

		// Mock the user detail response
		url := webapi.ApiHosts + "user/11?full=1"
		err := webtest.MockResponseFromFile("GET", url, "user_full")
		assert.NoError(t, err)

		// Initialize the WebPixivAPI instance
		app, err := pixiv.NewWebApp("dummy-phpsessid")
		assert.NoError(t, err)

		// Call the UserFull method
		user, err := app.UserFull(11)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		// Check user
		assert.Equal(t, "11", user.UserID)
		assert.Equal(t, "pixiv事務局", user.Name)
		assert.Equal(t, "https://www.pixiv.net/", user.Webpage)
	})
}
