package webapi

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/models/webmodel/user"
)

// UserShort returns a user information in a simplified format. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-information-short
func (a *WebPixivAPI) UserShort(uid uint64) (*user.UserShort, error) {
	path := fmt.Sprintf("user/%d", uid)
	referer := fmt.Sprintf(AppHosts+"member.php?id=%d", uid)

	res, err := Get[user.UserShortResponse](a, path, &referer, nil)
	if err != nil {
		return nil, err
	}

	return &res.Body, nil
}
