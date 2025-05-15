package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models"
)

// UserBookmarksIllustOptions defines optional parameters for retrieving user bookmarks illust.
type UserBookmarksIllustOptions struct {
	Restrict      *models.Restrict
	Filter        *string
	MaxBookmarkID *int
	Tag           *string
}

type userBookmarksIllustParams struct {
	UserID        uint64          `url:"user_id,omitempty"`
	Restrict      models.Restrict `url:"restrict,omitempty"`
	Filter        *string         `url:"filter,omitempty"`
	MaxBookmarkID *int            `url:"max_bookmark_id,omitempty"`
	Tag           *string         `url:"tag,omitempty"`
}

// UserBookmarksIllust retrieves a list of bookmarked illustrations for a given user.
// It allows optional parameters such as restrict level, filter, max bookmark ID, and tag.
//
// Parameters:
//   - uid: The Pixiv user ID whose bookmarks should be fetched.
//   - opts: Optional parameters for filtering the results. Can be nil.
//
// Returns:
//   - []models.Illust: A list of illustrations.
//   - int: The offset for the next page (based on max_bookmark_id), or 0 if there is no next page.
//   - error: Any error encountered during the API request or pagination parsing.
func (a *AppPixivAPI) UserBookmarksIllust(uid uint64, opts ...UserBookmarksIllustOptions) ([]models.Illust, int, error) {
	const path = "v1/user/bookmarks/illust"

	// Construct request parameters
	params := userBookmarksIllustParams{
		UserID:   uid,
		Restrict: models.Public,
	}

	// Populate optional parameters if opts is provided
	if opts != nil {
		opt := opts[0]
		params.Restrict = getRestrict(opt.Restrict)
		params.Filter = opt.Filter
		params.MaxBookmarkID = opt.MaxBookmarkID
		params.Tag = opt.Tag
	}

	// Initialize the response model
	data := &models.IllustsResponse{}

	// Send the API request
	if err := a.Request(path, params, data); err != nil {
		return nil, 0, err
	}

	// Parse the next page offset from the NextURL (if any)
	next, err := parseNextPageOffset(data.NextURL, "max_bookmark_id")
	return data.Illusts, next, err
}
