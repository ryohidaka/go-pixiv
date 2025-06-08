package webapi

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/internal/webutils"
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/models/webmodel/user"
)

type UserFollowersOptions struct {
	Offset   uint16        `url:"offset"`
	Limit    uint8         `url:"limit"`
	Restrict core.Restrict `url:"rest"`
}

type userFollowersParams struct {
	Offset   uint16        `url:"offset,omitempty"`
	Limit    uint8         `url:"limit,omitempty"`
	Restrict core.Restrict `url:"rest,omitempty"`
}

// UserFollowers returns user's followers. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-following-users
func (a *WebPixivAPI) UserFollowers(uid uint64, opts ...UserFollowersOptions) ([]user.FollowerUser, uint32, error) {
	// Construct request parameters
	params := &userFollowersParams{
		Offset:   0,
		Limit:    48,
		Restrict: core.Show,
	}

	if opts != nil {
		opt := opts[0]
		params.Offset = opt.Offset
		params.Limit = opt.Limit
		params.Restrict = webutils.GetRestrict(&opt.Restrict)
	}

	path := fmt.Sprintf("user/%d/followers", uid)

	res, err := Get[user.UserFolowersResponse](a, path, nil, params)
	if err != nil {
		return nil, 0, err
	}

	body := res.Body
	users := body.Users
	total := body.Total

	return users, total, nil
}
