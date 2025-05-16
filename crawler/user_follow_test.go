package crawler_test

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/crawler"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserFollowAddMultiple(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth/token", true)

		// Mock the user follower response
		url := pixiv.AppHosts + "v1/user/follow/add"
		err := testutil.MockResponseFromFile("POST", url, "empty", true)
		assert.NoError(t, err)

		// Initialize Crawler instance
		crawler, err := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)

		// Call UserFollowAddMultiple with a specific uids and restrict mode
		uids := []uint64{12345678}
		restrict := models.Private

		processed, err := crawler.UserFollowAddMultiple(uids, restrict)
		assert.NoError(t, err)
		assert.Equal(t, []uint64{12345678}, processed)

		// Verify request was made as expected
		info := httpmock.GetCallCountInfo()
		key := fmt.Sprintf("POST %s", url)
		if info[key] != 1 {
			t.Errorf("expected 1 POST request to %s, but got %d", url, info[key])
		}
	})
}
