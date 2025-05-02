package pixiv

import (
	"go-pixiv/testutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestNewApp verifies that NewApp correctly initializes the API with mocked authentication.
func TestNewApp(t *testing.T) {
	// Activate HTTP mocking
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the authentication response
	err := testutil.MockResponseFromFile("POST", AuthHosts+"auth/token", "auth.json")
	assert.NoError(t, err)

	// Create a new AppPixivAPI instance
	api, err := NewApp("dummy-refresh-token")
	assert.NoError(t, err)
	assert.NotNil(t, api)
	assert.NotEmpty(t, api.auth.AccessToken)
}

// TestAppPixivAPIRequest verifies that API requests return the expected mocked data.
func TestAppPixivAPIRequest(t *testing.T) {
	// Activate HTTP mocking
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the authentication response
	err := testutil.MockResponseFromFile("POST", AuthHosts+"auth/token", "auth.json")
	assert.NoError(t, err)

	// Mock API endpoint response
	apiURL := AppHosts + "v1/user/detail?user_id=123"
	err = testutil.MockResponseFromFile("GET", apiURL, "user-detail.json")
	assert.NoError(t, err)

	// Initialize AppPixivAPI
	api, err := NewApp("dummy-refresh-token")
	assert.NoError(t, err)

	// Call the request method
	type response struct {
		User struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"user"`
	}
	var out response
	err = api.request("v1/user/detail", struct {
		UserID int `url:"user_id"`
	}{UserID: 123}, &out)
	assert.NoError(t, err)
	assert.Equal(t, 11, out.User.ID)
	assert.Equal(t, "pixiv事務局", out.User.Name)
}
