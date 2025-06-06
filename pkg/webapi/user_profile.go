package webapi

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/models/webmodel/user"
)

// UserProfile returns a user information along with information about artwork posted by the user. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-information-focus-on-artwork
func (a *WebPixivAPI) UserProfile(uid uint64) (*user.UserProfile, error) {
	path := fmt.Sprintf("user/%d/profile/all", uid)
	referer := fmt.Sprintf(AppHosts+"member.php?id=%d", uid)

	res, err := Get[user.UserProfileResponse](a, path, &referer, nil)
	if err != nil {
		return nil, err
	}

	return &res.Body, nil
}
