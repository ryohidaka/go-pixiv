package webapi

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/internal/webutils"
	"github.com/ryohidaka/go-pixiv/models/webmodel/bookmark"
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
)

type UserBookmarksIllustsOptions struct {
	Tag      *string       `url:"tag"` // Specify tags. Even if you do not specify a tag, you must include an empty tag in your query.
	Offset   uint16        `url:"offset"`
	Limit    uint8         `url:"limit"`
	Restrict core.Restrict `url:"rest"`
}

type userBookmarksIllustsParams struct {
	Tag      *string       `url:"tag"`
	Offset   uint16        `url:"offset"`
	Limit    uint8         `url:"limit"`
	Restrict core.Restrict `url:"rest"`
}

// UserLatestWorks returns bookmarks of users. ([Docs])
//
// [Docs]: https://github.com/daydreamer-json/pixiv-ajax-api-docs?tab=readme-ov-file#get-user-bookmarks
func (a *WebPixivAPI) UserBookmarksIllusts(uid uint64, opts ...UserBookmarksIllustsOptions) (*bookmark.BookmarkedIllusts, uint32, error) {
	// Construct request parameters
	params := &userBookmarksIllustsParams{
		Offset:   0,
		Limit:    48,
		Restrict: core.Show,
	}

	if opts != nil {
		opt := opts[0]
		params.Tag = opt.Tag
		params.Offset = opt.Offset
		params.Limit = opt.Limit
		params.Restrict = webutils.GetRestrict(&opt.Restrict)
	}

	path := fmt.Sprintf("user/%d/illusts/bookmarks", uid)

	res, err := Get[bookmark.BookmarkedIllustsRespponse](a, path, nil, params)
	if err != nil {
		return nil, 0, err
	}

	body := res.Body
	total := body.Total

	return &body, total, nil
}
