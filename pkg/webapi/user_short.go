package webapi

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/models/webmodel"
	"github.com/ryohidaka/go-pixiv/models/webmodel/user"
)

// UserShort returns a short user profile by user ID using the Pixiv Web API.
func (a *WebPixivAPI) UserShort(uid uint64) (*webmodel.UserShort, error) {
	path := fmt.Sprintf("user/%d", uid)
	referer := fmt.Sprintf(AppHosts+"member.php?id=%d", uid)

	res, err := Get[user.UserShortResponse](a, path, &referer, nil)
	if err != nil {
		return nil, err
	}

	return &res.Body, nil
}
