package pixiv

import (
	"log/slog"
	"os"
	"time"

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
func (a *AppPixivAPI) UserBookmarksIllust(uid uint64, opts *UserBookmarksIllustOptions) ([]models.Illust, int, error) {
	const path = "v1/user/bookmarks/illust"

	// Construct request parameters
	params := userBookmarksIllustParams{
		UserID:   uid,
		Restrict: models.Public,
	}

	// Populate optional parameters if opts is provided
	if opts != nil {
		params.Restrict = getRestrict(opts.Restrict)
		params.Filter = opts.Filter
		params.MaxBookmarkID = opts.MaxBookmarkID
		params.Tag = opts.Tag
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

// FetchAllBookmarkedIllusts retrieves all bookmarked illustrations for a given user by paginating
//
// Parameters:
//   - uid: Pixiv user ID to retrieve bookmarks for.
//   - baseOpts: Optional filter parameters (Restrict, Filter, Tag). MaxBookmarkID will be overwritten for pagination.
//   - sleepMs: Optional sleep duration between requests in milliseconds (default: 1000ms).
//
// Returns:
//   - []models.Illust: A combined list of all retrieved bookmarked illustrations (even if an error occurs mid-fetch).
//   - error: Any error encountered during the API request.
func (a *AppPixivAPI) FetchAllBookmarkedIllusts(uid uint64, opts *UserBookmarksIllustOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	// Initialize the slog logger with options
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	// Start the fetch process
	logger.Info("Fetching bookmarked illustrations", "uid", uid)

	for {
		// Retrieve a page of bookmarked illustrations
		var illusts []models.Illust
		illusts, next, err = a.UserBookmarksIllust(uid, opts)

		// Log the number of illustrations fetched in this request
		logger.Info("Fetched illustrations", "count", len(illusts), "nextBookmarkID", next)

		// Always append the successfully fetched results before returning on error
		allIllusts = append(allIllusts, illusts...)
		if err != nil {
			// Log the error
			logger.Error("Error fetching illustrations", "error", err)
			return allIllusts, err
		}

		// Exit loop if there are no more pages
		if next == 0 {
			logger.Info("No more pages to fetch, exiting")
			break
		}

		// Set the next max_bookmark_id for pagination
		opts.MaxBookmarkID = &next

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		logger.Info("Sleeping before next request", "sleepDuration", sleepDuration)
		time.Sleep(sleepDuration)
	}

	// Log the total number of illustrations fetched
	logger.Info("Total illustrations fetched", "total", len(allIllusts))

	return allIllusts, nil
}
