package pixiv

import "github.com/ryohidaka/go-pixiv/models"

// UserDetailOptions defines optional parameters for retrieving user details.
type UserDetailOptions struct {
	Filter *string // Optional filter to narrow down the response
}

type userDetailParams struct {
	UserID uint64  `url:"user_id,omitempty"`
	Filter *string `url:"filter,omitempty"`
}

// UserDetail retrieves detailed information about a Pixiv user by their user ID.
//
// Parameters:
//   - uid: The user ID of the Pixiv user to retrieve details for.
//   - opts: Optional parameters for the request (e.g., filter conditions).
//
// Returns:
//   - A pointer to a models.UserDetail struct containing the user's information.
//   - An error if the request fails or the response cannot be parsed.
func (a *AppPixivAPI) UserDetail(uid uint64, opts ...UserDetailOptions) (*models.UserDetail, error) {
	const path = "v1/user/detail"

	// Construct request parameters
	params := userDetailParams{
		UserID: uid,
	}
	if opts != nil {
		opt := opts[0]
		params.Filter = opt.Filter
	}

	// Initialize the response model
	detail := &models.UserDetail{
		User: &models.User{},
	}

	// Send the API request
	if err := a.Request(path, params, detail); err != nil {
		return nil, err
	}
	return detail, nil
}
