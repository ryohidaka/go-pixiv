package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models"
)

// UserFollowingOptions defines optional parameters for retrieving the list of followed users.
type UserFollowingOptions struct {
	Restrict *models.Restrict // Restrict visibility: public or private
	Offset   *int             // Offset for pagination
}

type userFollowingParams struct {
	UserID   uint64          `url:"user_id,omitempty"`
	Restrict models.Restrict `url:"restrict,omitempty"`
	Offset   *int            `url:"offset,omitempty"`
}

// UserFollowing fetches the list of users followed by the specified user.
func (a *AppPixivAPI) UserFollowing(uid uint64, opts ...UserFollowingOptions) ([]models.UserPreview, int, error) {
	const path = "v1/user/following"

	// Construct request parameters
	params := &userFollowingParams{
		UserID:   uid,
		Restrict: models.Public,
	}

	// Populate optional parameters if opts is provided
	if opts != nil {
		opt := opts[0]
		params.Restrict = getRestrict(opt.Restrict)
		params.Offset = opt.Offset
	}

	// Initialize the response model
	data := &models.UserFollowList{}

	// Send the API request
	if err := a.Get(path, params, data); err != nil {
		return nil, 0, err
	}

	// Parse the next page offset from the response's NextURL field
	next, err := parseNextPageOffset(data.NextURL, OffsetFieldOffset)
	return data.UserPreviews, next, err
}
