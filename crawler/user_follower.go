package crawler

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

// FetchAllUserFollowers retrieves all followers of the specified user by paginating.
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
