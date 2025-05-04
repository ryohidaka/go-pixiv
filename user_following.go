package pixiv

import "github.com/ryohidaka/go-pixiv/models"

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
//
// Parameters:
//   - userID: Pixiv user ID of the target user.
//   - opts: Optional parameters for restricting visibility or paginating results.
//
// Returns:
//   - A pointer to models.UserFollowList containing the list of followed users.
//   - An integer indicating the offset for the next page, or 0 if there is no next page.
//   - An error if the request fails.
func (a *AppPixivAPI) UserFollowing(userID uint64, opts *UserFollowingOptions) ([]models.UserPreview, int, error) {
	const path = "v1/user/following"

	// Construct request parameters
	params := &userFollowingParams{
		UserID:   userID,
		Restrict: models.Public,
	}

	// Populate optional parameters if opts is provided
	if opts != nil {
		params.Restrict = getRestrict(opts.Restrict)
		params.Offset = opts.Offset
	}

	// Initialize the response model
	data := &models.UserFollowList{}

	// Send the API request
	if err := a.Request(path, params, data); err != nil {
		return nil, 0, err
	}

	// Parse the next page offset from the response's NextURL field
	next, err := parseNextPageOffset(data.NextURL, OffsetFieldOffset)
	return data.UserPreviews, next, err
}
