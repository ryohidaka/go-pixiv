package crawler_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/pkg/appapi/crawler"
	"github.com/ryohidaka/go-pixiv/testutil/apptest"
	"github.com/stretchr/testify/assert"
)

// TestNewApp verifies that NewApp correctly initializes the API with mocked authentication.
func TestNewCrawler(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock the authentication response
		err := apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth_token")
		assert.NoError(t, err)

		// Create a new Pixiv Crawler
		c, _ := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
}
