package webapi_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/pkg/webapi"
	"github.com/ryohidaka/go-pixiv/testutil/webtest"
	"github.com/stretchr/testify/assert"
)

func ExampleWebPixivAPI_UserFollowers() {
	// Get the PHPSESSID used for authentication
	phpsessid := os.Getenv("PHPSESSID")

	// Create a new Pixiv Web API client
	app, _ := pixiv.NewWebApp(phpsessid)

	// Fetch user's followers.
	users, total, _ := app.UserFollowers(11)

	// Print the user
	fmt.Println("User: ", users[0].UserName)

	// Print the total
	fmt.Println("Total:", total)
}

func TestUserFollowers(t *testing.T) {
	webtest.WithMockHTTP(t, func() {
		// Mock the user following response
		url := webapi.ApiHosts + "user/11/followers"
		err := webtest.MockResponseFromFile("GET", url, "user_followers")
		assert.NoError(t, err)

		// Initialize the WebPixivAPI instance
		app, err := pixiv.NewWebApp("dummy-phpsessid")
		assert.NoError(t, err)

		// Call the UserFollowers method
		opts := pixiv.UserFollowersOptions{
			Offset:   0,
			Limit:    1,
			Restrict: core.Show,
		}

		users, total, err := app.UserFollowers(11, opts)
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.NotNil(t, total)

		// Check user
		assert.Equal(t, "11", users[0].UserID)
		assert.Equal(t, "pixiv事務局", users[0].UserName)

		// Check total
		assert.Equal(t, 1, int(total))
	})
}
