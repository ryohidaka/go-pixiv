package pixiv

import (
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
