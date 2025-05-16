package pixiv_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/testutil"

	"github.com/stretchr/testify/assert"
)

// TestNewApp verifies that NewApp correctly initializes the API with mocked authentication.
func TestNewApp(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		err := testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")
		assert.NoError(t, err)

		// Create a new AppPixivAPI instance
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)
		assert.NotNil(t, api)
	})
}

// TestAppPixivAPIRequest verifies that API requests return the expected mocked data.
func TestAppPixivAPIRequest(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		err := testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token")
		assert.NoError(t, err)

		// Mock API endpoint response
		apiURL := pixiv.AppHosts + "v1/user/detail?user_id=123"
		err = testutil.MockResponseFromFile("GET", apiURL, "v1/user/detail")
		assert.NoError(t, err)

		// Initialize AppPixivAPI
		api, err := pixiv.NewApp("dummy-refresh-token")
		assert.NoError(t, err)

		// Call the GET method wrapper
		type response struct {
			User struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		}
		var out response
		err = api.Get("v1/user/detail", struct {
			UserID int `url:"user_id"`
		}{UserID: 123}, &out)
		assert.NoError(t, err)
		assert.Equal(t, 11, out.User.ID)
		assert.Equal(t, "pixiv事務局", out.User.Name)
	})
}
