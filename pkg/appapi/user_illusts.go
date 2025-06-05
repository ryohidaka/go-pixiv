package appapi

import (
	"github.com/ryohidaka/go-pixiv/internal/urlutil"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/models/appmodel"
)

// UserIllustsOptions defines optional parameters for fetching a user's illustrations.
type UserIllustsOptions struct {
	Filter *string            // Optional filter string, e.g., "for_ios".
	Type   *models.IllustType // Optional type of illustration (e.g., "illust", "manga", "ugoira").
	Offset *int               // Optional offset for pagination.
}

// userIllustsParams is an internal struct for encoding URL parameters in API requests.
type userIllustsParams struct {
	UserID uint64             `url:"user_id,omitempty"` // ID of the target user.
	Filter *string            `url:"filter,omitempty"`  // Filter option.
	Type   *models.IllustType `url:"type,omitempty"`    // Type of illustration.
	Offset *int               `url:"offset,omitempty"`  // Pagination offset.
}

// UserIllusts retrieves a list of illustrations for a given user.
func (a *AppPixivAPI) UserIllusts(uid uint64, opts ...UserIllustsOptions) ([]models.Illust, int, error) {
	const path = "v1/user/illusts"

	// Prepare parameters with required UserID.
	params := &userIllustsParams{
		UserID: uid,
	}

	// Apply optional parameters if provided.
	if opts != nil {
		opt := opts[0]
		params.Filter = opt.Filter
		params.Type = opt.Type
		params.Offset = opt.Offset
	}

	// Initialize the response model
	data := &appmodel.IllustsResponse{}

	// Send the API request
	if err := a.Get(path, params, data); err != nil {
		return nil, 0, err
	}

	// Extract the offset for the next page from the NextURL field.
	next, err := urlutil.ParseNextPageOffset(data.NextURL, OffsetFieldOffset)

	return data.Illusts, next, err
}
