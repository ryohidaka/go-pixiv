package pixiv

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv/models"
)

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
//   - uid: Pixiv user ID of the target user.
//   - opts: Optional parameters for restricting visibility or paginating results.
//
// Returns:
//   - A pointer to models.UserFollowList containing the list of followed users.
//   - An integer indicating the offset for the next page, or 0 if there is no next page.
//   - An error if the request fails.
func (a *AppPixivAPI) UserFollowing(uid uint64, opts *UserFollowingOptions) ([]models.UserPreview, int, error) {
	const path = "v1/user/following"

	// Construct request parameters
	params := &userFollowingParams{
		UserID:   uid,
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

// FetchAllUserFollowing retrieves all users followed by the specified user by paginating.
//
// Parameters:
//   - uid: Pixiv user ID of the target user.
//   - opts: Optional parameters such as Restrict. Offset will be managed internally.
//   - sleepMs: Optional sleep duration between requests in milliseconds (default: 1000ms).
//
// Returns:
//   - []models.UserPreview: A complete list of followed users.
//   - error: Any error encountered during the request.
func (a *AppPixivAPI) FetchAllUserFollowing(uid uint64, opts *UserFollowingOptions, sleepMs ...int) ([]models.UserPreview, error) {
	var allUsers []models.UserPreview
	var next int
	var err error

	// Use default sleep duration of 1000ms unless specified
	sleepDuration := 1000 * time.Millisecond
	if len(sleepMs) > 0 {
		sleepDuration = time.Duration(sleepMs[0]) * time.Millisecond
	}

	// Logger setup
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching all followed users", "userID", uid)

	for {
		var users []models.UserPreview
		users, next, err = a.UserFollowing(uid, opts)

		logger.Info("Fetched users", "count", len(users), "nextOffset", next)

		allUsers = append(allUsers, users...)
		if err != nil {
			logger.Error("Error fetching followed users", "error", err)
			return allUsers, err
		}

		if next == 0 {
			logger.Info("No more pages to fetch, exiting")
			break
		}

		opts.Offset = &next

		logger.Info("Sleeping before next request", "sleepDuration", sleepDuration)
		time.Sleep(sleepDuration)
	}

	logger.Info("Total followed users fetched", "total", len(allUsers))

	return allUsers, nil
}
