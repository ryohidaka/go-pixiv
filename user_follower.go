package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models"
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
//
// Parameters:
//   - uid: Pixiv user ID of the target user.
//   - opts: Optional parameters for restricting visibility or paginating results.
//
// Returns:
//   - A pointer to models.UserFollowList containing the list of followers.
//   - An integer indicating the offset for the next page, or 0 if there is no next page.
//   - An error if the request fails.
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
