package crawler_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/crawler"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

// TestNewApp verifies that NewApp correctly initializes the API with mocked authentication.
func TestNewCrawler(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		err := testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token", "../testutil")
		assert.NoError(t, err)

		// Create a new Pixiv Crawler
		c, _ := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
}
