package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/models/appmodel"
)

// UserDetailOptions defines optional parameters for retrieving user details.
type UserDetailOptions struct {
	Filter *string // Optional filter to narrow down the response
}

type userDetailParams struct {
	UserID uint64  `url:"user_id,omitempty"`
	Filter *string `url:"filter,omitempty"`
}

// UserDetail retrieves detailed information about a Pixiv user by their user ID.
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
		User: &appmodel.User{},
	}

	// Send the API request
	if err := a.Get(path, params, detail); err != nil {
		return nil, err
	}
	return detail, nil
}
