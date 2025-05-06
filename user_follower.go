package pixiv

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv/models"
)

// UserFollowerOptions defines optional parameters for retrieving the list of user followers.
type UserFollowerOptions struct {
	Restrict *models.Restrict // Restrict visibility: public or private
	Offset   *int             // Offset for pagination
}

type userFollowerParams struct {
	UserID   uint64          `url:"user_id,omitempty"`
	Restrict models.Restrict `url:"restrict,omitempty"`
	Offset   *int            `url:"offset,omitempty"`
}

// UserFollower fetches the list of users who follow the specified user.
//
// Parameters:
//   - uid: Pixiv user ID of the target user.
//   - opts: Optional parameters for restricting visibility or paginating results.
//
// Returns:
//   - A pointer to models.UserFollowList containing the list of followers.
//   - An integer indicating the offset for the next page, or 0 if there is no next page.
//   - An error if the request fails.
func (a *AppPixivAPI) UserFollower(uid uint64, opts *UserFollowerOptions) ([]models.UserPreview, int, error) {
	const path = "v1/user/follower"

	// Construct request parameters
	params := &userFollowerParams{
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

// FetchAllUserFollowers retrieves all followers of the specified user by paginating.
//
// Parameters:
//   - uid: Pixiv user ID of the target user.
//   - opts: Optional parameters such as Restrict. Offset will be managed internally.
//   - sleepMs: Optional sleep duration between requests in milliseconds (default: 1000ms).
//
// Returns:
//   - []models.UserPreview: A complete list of followers.
//   - error: Any error encountered during the request.
func (a *AppPixivAPI) FetchAllUserFollowers(uid uint64, opts *UserFollowerOptions, sleepMs ...int) ([]models.UserPreview, error) {
	var allFollowers []models.UserPreview
	var next int
	var err error

	// Logger setup
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching all followers", "userID", uid)

	for {
		var followers []models.UserPreview
		followers, next, err = a.UserFollower(uid, opts)

		logger.Info("Fetched followers", "count", len(followers), "nextOffset", next)

		allFollowers = append(allFollowers, followers...)
		if err != nil {
			logger.Error("Error fetching followers", "error", err)
			return allFollowers, err
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

	logger.Info("Total followers fetched", "total", len(allFollowers))

	return allFollowers, nil
}
