package pixiv

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv/models"
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

	// Initialize the response model
	data := &models.IllustsResponse{}

	// Send the API request
	if err := a.Request(path, params, data); err != nil {
		return nil, 0, err
	}

	// Parse the next page offset from the response's NextURL field
	next, err := parseNextPageOffset(data.NextURL, OffsetFieldOffset)
	return data.Illusts, next, err
}

// FetchAllIllustFollows retrieves all illustrations from followed users by paginating.
//
// Parameters:
//   - opts: Optional parameters for the follow illust request (e.g. Restrict).
//   - sleepMs: Optional sleep duration between requests in milliseconds (default: 1000ms).
//
// Returns:
//   - []models.Illust: A combined list of all retrieved follow illustrations.
//   - error: Any error encountered during the API request.
func (a *AppPixivAPI) FetchAllIllustFollows(opts *IllustFollowOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	// Set up the logger
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching followed users' illustrations")

	for {
		var illusts []models.Illust
		illusts, next, err = a.IllustFollow(opts)

		logger.Info("Fetched illustrations", "count", len(illusts), "nextOffset", next)

		allIllusts = append(allIllusts, illusts...)
		if err != nil {
			logger.Error("Error fetching follow illustrations", "error", err)
			return allIllusts, err
		}

		if next == 0 {
			logger.Info("No more pages to fetch, exiting")
			break
		}

		opts.Offset = &next

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		logger.Info("Sleeping before next request", "sleepDuration", sleepDuration)
		time.Sleep(sleepDuration)
	}

	logger.Info("Total illustrations fetched", "total", len(allIllusts))

	return allIllusts, nil
}
