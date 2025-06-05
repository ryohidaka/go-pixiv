package pixiv

import (
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
)

// AppPixivAPI handles Pixiv App API operations using OAuth authentication.
type AppPixivAPI = appapi.AppPixivAPI

// UserIllustsOptions defines optional parameters for fetching a user's illustrations.
type UserIllustsOptions = appapi.UserIllustsOptions

// UserFollowingOptions defines optional parameters for retrieving the list of followed users.
type UserFollowingOptions = appapi.UserFollowingOptions

// UserFollowerOptions defines optional parameters for retrieving the list of user followers.
type UserFollowerOptions = appapi.UserFollowerOptions

// UserDetailOptions defines optional parameters for retrieving user details.
type UserDetailOptions = appapi.UserDetailOptions

// UserBookmarksIllustOptions defines optional parameters for retrieving user bookmarks illust.
type UserBookmarksIllustOptions = appapi.UserBookmarksIllustOptions

// IllustFollowOptions defines optional parameters for the IllustFollow method.
type IllustFollowOptions = appapi.IllustFollowOptions

// NewApp initializes and returns a new AppPixivAPI instance using the provided refresh token.
func NewApp(refreshToken string) (*AppPixivAPI, error) {
	return appapi.NewApp(refreshToken)
}
