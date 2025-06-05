package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/models/appmodel"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
)

// AppPixivAPI wraps appapi.AppPixivAPI to expose its methods in this package.
type AppPixivAPI struct {
	*appapi.AppPixivAPI
}

// NewApp initializes and returns a new AppPixivAPI instance using the provided refresh token.
func NewApp(refreshToken string) (*AppPixivAPI, error) {
	api, err := appapi.NewApp(refreshToken)
	if err != nil {
		return nil, err
	}
	return &AppPixivAPI{AppPixivAPI: api}, nil
}

// UserDetailOptions defines optional parameters for retrieving user details.
type UserDetailOptions = appapi.UserDetailOptions

// UserDetail retrieves detailed information about a Pixiv user by their user ID.
func (a *AppPixivAPI) UserDetail(uid uint64, opts ...UserDetailOptions) (*models.UserDetail, error) {
	return a.AppPixivAPI.UserDetail(uid, opts...)
}

// UserIllustsOptions defines optional parameters for fetching a user's illustrations.
type UserIllustsOptions = appapi.UserIllustsOptions

// UserIllusts retrieves a list of illustrations for a given user.
func (a *AppPixivAPI) UserIllusts(uid uint64, opts ...UserIllustsOptions) ([]models.Illust, int, error) {
	return a.AppPixivAPI.UserIllusts(uid, opts...)
}

// UserBookmarksIllustOptions defines optional parameters for retrieving user bookmarks illust.
type UserBookmarksIllustOptions = appapi.UserBookmarksIllustOptions

// UserBookmarksIllust retrieves a list of bookmarked illustrations for a given user.
// It allows optional parameters such as restrict level, filter, max bookmark ID, and tag.
func (a *AppPixivAPI) UserBookmarksIllust(uid uint64, opts ...UserBookmarksIllustOptions) ([]models.Illust, int, error) {
	return a.AppPixivAPI.UserBookmarksIllust(uid, opts...)
}

// IllustFollowOptions defines optional parameters for the IllustFollow method.
type IllustFollowOptions = appapi.IllustFollowOptions

// IllustFollow retrieves a list of illustrations from users that the authenticated user follows.
// It supports optional parameters such as restriction (public/private) and pagination offset.
func (a *AppPixivAPI) IllustFollow(opts ...IllustFollowOptions) ([]models.Illust, int, error) {
	return a.AppPixivAPI.IllustFollow(opts...)
}

// IllustDetail retrieves detailed information about a specific illustration by its ID.
func (a *AppPixivAPI) IllustDetail(id uint64) (*appmodel.Illust, error) {
	return a.AppPixivAPI.IllustDetail(id)
}

// IllustBookmarkDetail retrieves the bookmark detail information
// for a specific illustration by its ID.
func (a *AppPixivAPI) IllustBookmarkDetail(id uint64) (*models.IllustBookmarkDetail, error) {
	return a.AppPixivAPI.IllustBookmarkDetail(id)
}

// UserFollowingOptions defines optional parameters for retrieving the list of followed users.
type UserFollowingOptions = appapi.UserFollowingOptions

// UserFollowing fetches the list of users followed by the specified user.
func (a *AppPixivAPI) UserFollowing(uid uint64, opts ...UserFollowingOptions) ([]models.UserPreview, int, error) {
	return a.AppPixivAPI.UserFollowing(uid, opts...)
}

// UserFollowerOptions defines optional parameters for retrieving the list of user followers.
type UserFollowerOptions = appapi.UserFollowerOptions

// UserFollower fetches the list of users who follow the specified user.
func (a *AppPixivAPI) UserFollower(uid uint64, opts ...UserFollowerOptions) ([]models.UserPreview, int, error) {
	return a.AppPixivAPI.UserFollower(uid, opts...)
}

type Restrict = models.Restrict

const (
	Public  = models.Public
	Private = models.Private
)

// UserFollowAdd sends a follow request to a user on Pixiv.
func (a *AppPixivAPI) UserFollowAdd(uid uint64, restrict ...Restrict) (bool, error) {
	return a.AppPixivAPI.UserFollowAdd(uid, restrict...)
}

// UserFollowAdd sends a unfollow request to a user on Pixiv.
func (a *AppPixivAPI) UserFollowDelete(uid uint64) (bool, error) {
	return a.AppPixivAPI.UserFollowDelete(uid)
}
