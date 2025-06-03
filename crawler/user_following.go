package crawler

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

// FetchAllUserFollowing retrieves all users followed by the specified user by paginating.
func (c *PixivCrawler) FetchAllUserFollowing(uid uint64, opts *pixiv.UserFollowingOptions, sleepMs ...int) ([]models.UserPreview, error) {
	var allUsers []models.UserPreview
	var next int
	var err error

	// Logger setup
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching all followed users", "userID", uid)

	for {
		var users []models.UserPreview
		users, next, err = c.app.UserFollowing(uid, []pixiv.UserFollowingOptions{*opts}...)

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

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		logger.Info("Sleeping before next request", "sleepDuration", sleepDuration)
		time.Sleep(sleepDuration)
	}

	logger.Info("Total followed users fetched", "total", len(allUsers))

	return allUsers, nil
}
