package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models/webmodel/bookmark"
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/models/webmodel/user"
	"github.com/ryohidaka/go-pixiv/pkg/webapi"
)

// WebPixivAPI wraps webapi.WebPixivAPI to expose its methods in this package.
type WebPixivAPI struct {
	*webapi.WebPixivAPI
}

// NewWebApp initializes and returns a new instance of WebPixivAPI.
func NewWebApp(phpsessid string) (*WebPixivAPI, error) {
	api, err := webapi.NewWebApp(phpsessid)
	if err != nil {
		return nil, err
	}
	return &WebPixivAPI{WebPixivAPI: api}, nil
}

// UserShort returns a user information in a simplified format. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-information-short
func (a *WebPixivAPI) UserShort(uid uint64) (*user.UserShort, error) {
	return a.WebPixivAPI.UserShort(uid)
}

// UserFull returns a full user information. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-information-full
func (a *WebPixivAPI) UserFull(uid uint64) (*user.User, error) {
	return a.WebPixivAPI.UserFull(uid)
}

// UserProfile returns a user information along with information about artwork posted by the user. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-information-focus-on-artwork
func (a *WebPixivAPI) UserProfile(uid uint64) (*user.UserProfile, error) {
	return a.WebPixivAPI.UserProfile(uid)
}

type WebRestrict = core.Restrict

const (
	Show = core.Show
	Hide = core.Hide
)

type WebUserFollowingOptions = webapi.UserFollowingOptions

// UserFull returns following users. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-following-users
func (a *WebPixivAPI) UserFollowing(uid uint64, opts ...WebUserFollowingOptions) ([]user.FollowingUser, uint32, error) {
	return a.WebPixivAPI.UserFollowing(uid, opts...)
}

type UserFollowersOptions = webapi.UserFollowersOptions

// UserFollowers returns user's followers. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-users-followers
func (a *WebPixivAPI) UserFollowers(uid uint64, opts ...UserFollowersOptions) ([]user.FollowerUser, uint32, error) {
	return a.WebPixivAPI.UserFollowers(uid, opts...)
}

// UserLatestWorks returns the latest artworks of users. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-latest-artworks
func (a *WebPixivAPI) UserLatestWorks(uid uint64) (*user.UserWorks, error) {
	return a.WebPixivAPI.UserLatestWorks(uid)
}

type UserBookmarksIllustsOptions = webapi.UserBookmarksIllustsOptions

// UserLatestWorks returns bookmarks of users. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-bookmarks
func (a *WebPixivAPI) UserBookmarksIllusts(uid uint64, opts ...UserBookmarksIllustsOptions) (*bookmark.BookmarkedIllusts, uint32, error) {
	return a.WebPixivAPI.UserBookmarksIllusts(uid, opts...)
}
