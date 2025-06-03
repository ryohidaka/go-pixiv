package crawler

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

// FetchAllUserIllusts retrieves all illustrations posted by a given user by paginating.
func (c *PixivCrawler) FetchAllUserIllusts(uid uint64, opts *pixiv.UserIllustsOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	// Initialize the slog logger
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching user illustrations", "uid", uid)

	for {
		// Fetch a single page
		var illusts []models.Illust
		illusts, next, err = c.app.UserIllusts(uid, []pixiv.UserIllustsOptions{*opts}...)

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

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		logger.Info("Sleeping before next request", "sleepDuration", sleepDuration)
		time.Sleep(sleepDuration)
	}

	logger.Info("Total illustrations fetched", "total", len(allIllusts))

	return allIllusts, nil
}
