package pixiv

import "github.com/ryohidaka/go-pixiv/models"

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
//
// Parameters:
//   - opts: Pointer to IllustFollowOptions containing optional request parameters.
//
// Returns:
//   - A slice of Illust representing the retrieved illustrations.
//   - An integer indicating the offset for the next page, or 0 if there is no next page.
//   - An error if the request fails or the response cannot be parsed.
func (a *AppPixivAPI) IllustFollow(opts *IllustFollowOptions) ([]models.Illust, int, error) {
	const path = "v2/illust/follow"

	// Construct request parameters
	params := &illustFollowParams{
		Restrict: models.Private, // Default to private if not specified
	}

	// Override default parameters if options are provided
	if opts != nil {
		params.Restrict = getRestrict(opts.Restrict)
		params.Offset = opts.Offset
	}

	data := &models.IllustsResponse{}

	// Send the API request with the constructed parameters
	if err := a.Request(path, params, data); err != nil {
		return nil, 0, err
	}

	// Parse the next page offset from the response's NextURL field
	next, err := parseNextPageOffset(data.NextURL, OffsetFieldOffset)
	return data.Illusts, next, err
}
