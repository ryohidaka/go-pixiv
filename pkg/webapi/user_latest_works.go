package webapi

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/models/webmodel/user"
)

// UserLatestWorks returns the latest artworks of users. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-latest-artworks
func (a *WebPixivAPI) UserLatestWorks(uid uint64) (*user.UserWorks, error) {
	path := fmt.Sprintf("user/%d/works/latest", uid)

	res, err := Get[user.UserWorksResponse](a, path, nil, nil)
	if err != nil {
		return nil, err
	}

	return &res.Body, nil
}
