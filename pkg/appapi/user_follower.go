package appapi

import (
	"github.com/ryohidaka/go-pixiv/internal/modelutil"
	"github.com/ryohidaka/go-pixiv/internal/urlutil"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/models/appmodel"
)

// UserFollowerOptions defines optional parameters for retrieving the list of user followers.
type UserFollowerOptions struct {
	Restrict *models.Restrict // Restrict visibility: public or private
	Offset   *int             // Offset for pagination
}

type userFollowerParams struct {
	UserID   uint64          `url:"user_id,omitempty"`
	Restrict models.Restrict `url:"restrict,omitempty"`
	Offset   *int            `url:"offset,omitempty"`
}

// UserFollower fetches the list of users who follow the specified user.
func (a *AppPixivAPI) UserFollower(uid uint64, opts ...UserFollowerOptions) ([]models.UserPreview, int, error) {
	const path = "v1/user/follower"

	// Construct request parameters
	params := &userFollowerParams{
		UserID:   uid,
		Restrict: models.Public,
	}

	// Populate optional parameters if opts is provided
	if opts != nil {
		opt := opts[0]
		params.Restrict = modelutil.GetRestrict(opt.Restrict)
		params.Offset = opt.Offset
	}

	// Initialize the response model
	data := &appmodel.UserFollowList{}

	// Send the API request
	if err := a.Get(path, params, data); err != nil {
		return nil, 0, err
	}

	// Parse the next page offset from the response's NextURL field
	next, err := urlutil.ParseNextPageOffset(data.NextURL, OffsetFieldOffset)
	return data.UserPreviews, next, err
}
