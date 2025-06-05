package pixiv

import (
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
