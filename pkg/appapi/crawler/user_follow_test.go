package crawler_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/pkg/appapi/crawler"
	"github.com/ryohidaka/go-pixiv/testutil/apptest"
	"github.com/stretchr/testify/assert"
)

func ExamplePixivCrawler_UserFollowAddMultiple() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv Crawler
	c, _ := crawler.NewCrawler(refreshToken)

	// Follow multiple users
	uids := []uint64{11}
	processed, _ := c.UserFollowAddMultiple(uids, nil)

	for _, v := range processed {
		// Print the user id
		fmt.Println("User ID:", v)
	}
}

func ExamplePixivCrawler_UserFollowDeleteMultiple() {
	// Get the refresh token used for authentication
	refreshToken := os.Getenv("PIXIV_REFRESH_TOKEN")

	// Create a new Pixiv Crawler
	c, _ := crawler.NewCrawler(refreshToken)

	// Unfollow multiple users
	uids := []uint64{11}
	processed, _ := c.UserFollowDeleteMultiple(uids)

	for _, v := range processed {
		// Print the user id
		fmt.Println("User ID:", v)
	}
}

func TestUserFollowAddMultiple(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth/token", "../../../testutil")

		// Mock the user follow response
		url := appapi.AppHosts + "v1/user/follow/add"
		err := apptest.MockResponseFromFile("POST", url, "empty", "../../../testutil")
		assert.NoError(t, err)

		// Initialize Crawler instance
		crawler, err := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)

		// Call UserFollowAddMultiple with a specific uids and restrict mode
		uids := []uint64{12345678}
		restrict := models.Private

		processed, err := crawler.UserFollowAddMultiple(uids, &restrict)
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

func TestUserFollowDeleteMultiple(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		// Mock the authentication response
		_ = apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth/token", "../../../testutil")

		// Mock the user unfollow response
		url := appapi.AppHosts + "v1/user/follow/delete"
		err := apptest.MockResponseFromFile("POST", url, "empty", "../../../testutil")
		assert.NoError(t, err)

		// Initialize Crawler instance
		crawler, err := crawler.NewCrawler("dummy-refresh-token")
		assert.NoError(t, err)

		// Call UserFollowDeleteMultiple with a specific uids and restrict mode
		uids := []uint64{12345678}

		processed, err := crawler.UserFollowDeleteMultiple(uids)
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
