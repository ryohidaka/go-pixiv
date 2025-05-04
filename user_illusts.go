package pixiv

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv/models"
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
//
// Parameters:
//   - uid: The ID of the user whose illustrations are being fetched.
//   - opts: Optional parameters for filtering and pagination.
//
// Returns:
//   - []models.Illust: A slice of illustrations returned by the API.
//   - int: The offset for the next page, if available (0 if not present).
//   - error: An error if the API request or offset parsing fails.
func (a *AppPixivAPI) UserIllusts(uid uint64, opts *UserIllustsOptions) ([]models.Illust, int, error) {
	const path = "v1/user/illusts"

	// Prepare parameters with required UserID.
	params := &userIllustsParams{
		UserID: uid,
	}

	// Apply optional parameters if provided.
	if opts != nil {
		params.Filter = opts.Filter
		params.Type = opts.Type
		params.Offset = opts.Offset
	}

	// Initialize the response model
	data := &models.IllustsResponse{}

	// Send the API request
	if err := a.Request(path, params, data); err != nil {
		return nil, 0, err
	}

	// Extract the offset for the next page from the NextURL field.
	next, err := parseNextPageOffset(data.NextURL, OffsetFieldOffset)

	return data.Illusts, next, err
}

// FetchAllUserIllusts retrieves all illustrations posted by a given user by paginating.
//
// Parameters:
//   - uid: Pixiv user ID whose illustrations should be fetched.
//   - opts: Optional filter parameters (Filter, Type). Offset will be overwritten for pagination.
//   - sleepMs: Optional sleep duration between requests in milliseconds (default: 1000ms).
//
// Returns:
//   - []models.Illust: A combined list of all retrieved illustrations.
//   - error: Any error encountered during the API request.
func (a *AppPixivAPI) FetchAllUserIllusts(uid uint64, opts *UserIllustsOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	// Use default sleep duration of 1000ms unless explicitly specified
	sleepDuration := 1000 * time.Millisecond
	if len(sleepMs) > 0 {
		sleepDuration = time.Duration(sleepMs[0]) * time.Millisecond
	}

	// Initialize the slog logger
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching user illustrations", "uid", uid)

	for {
		// Fetch a single page
		var illusts []models.Illust
		illusts, next, err = a.UserIllusts(uid, opts)

		logger.Info("Fetched illustrations", "count", len(illusts), "nextOffset", next)

		// Append successfully fetched data
		allIllusts = append(allIllusts, illusts...)

		if err != nil {
			logger.Error("Error fetching illustrations", "error", err)
			return allIllusts, err
		}

		// Exit if there is no next page
		if next == 0 {
			logger.Info("No more pages to fetch, exiting")
			break
		}

		// Update offset for pagination
		opts.Offset = &next

		logger.Info("Sleeping before next request", "sleepDuration", sleepDuration)
		time.Sleep(sleepDuration)
	}

	logger.Info("Total illustrations fetched", "total", len(allIllusts))

	return allIllusts, nil
}
