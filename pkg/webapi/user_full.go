package webapi

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/models/webmodel/user"
)

type userFullParams struct {
	Full int `url:"full"`
}

// UserFull returns a full user information. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-information-full
func (a *WebPixivAPI) UserFull(uid uint64) (*user.User, error) {
	path := fmt.Sprintf("user/%d", uid)
	referer := fmt.Sprintf(AppHosts+"member.php?id=%d", uid)

	params := userFullParams{
		Full: 1,
	}

	res, err := Get[user.UserResponse](a, path, &referer, params)
	if err != nil {
		return nil, err
	}

	return &res.Body, nil
}
