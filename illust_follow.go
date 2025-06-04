package pixiv

import (
	"github.com/ryohidaka/go-pixiv/internal/modelutil"
	"github.com/ryohidaka/go-pixiv/internal/urlutil"
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/models/appmodel"
)

// IllustFollowOptions defines optional parameters for the IllustFollow method.
type IllustFollowOptions struct {
	Restrict *models.Restrict // Optional restriction setting (e.g., public or private)
	Offset   *int             // Optional pagination offset
}

type illustFollowParams struct {
	Restrict models.Restrict `url:"restrict,omitempty"` // Visibility restriction
	Offset   *int            `url:"offset,omitempty"`   // Pagination offset
}

// IllustFollow retrieves a list of illustrations from users that the authenticated user follows.
// It supports optional parameters such as restriction (public/private) and pagination offset.
func (a *AppPixivAPI) IllustFollow(opts ...IllustFollowOptions) ([]models.Illust, int, error) {
	const path = "v2/illust/follow"

	// Construct request parameters
	params := &illustFollowParams{
		Restrict: models.Private, // Default to private if not specified
	}

	// Override default parameters if options are provided
	if opts != nil {
		opt := opts[0]
		params.Restrict = modelutil.GetRestrict(opt.Restrict)
		params.Offset = opt.Offset
	}

	// Initialize the response model
	data := &appmodel.IllustsResponse{}

	// Send the API request
	if err := a.Get(path, params, data); err != nil {
		return nil, 0, err
	}

	// Parse the next page offset from the response's NextURL field
	next, err := urlutil.ParseNextPageOffset(data.NextURL, OffsetFieldOffset)
	return data.Illusts, next, err
}
