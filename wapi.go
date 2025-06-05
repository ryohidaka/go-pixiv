package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models/webmodel"
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

// UserShort returns a short user profile by user ID using the Pixiv Web API.
func (a *WebPixivAPI) UserShort(uid uint64) (*webmodel.UserShort, error) {
	return a.WebPixivAPI.UserShort(uid)
}
