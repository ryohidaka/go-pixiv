package crawler

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

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
func (c *PixivCrawler) FetchAllUserFollowers(uid uint64, opts *pixiv.UserFollowerOptions, sleepMs ...int) ([]models.UserPreview, error) {
	var allFollowers []models.UserPreview
	var next int
	var err error

	// Logger setup
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching all followers", "userID", uid)

	for {
		var followers []models.UserPreview
		followers, next, err = c.app.UserFollower(uid, []pixiv.UserFollowerOptions{*opts}...)

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
