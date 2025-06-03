package crawler

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

// FetchAllIllustFollows retrieves all illustrations from followed users by paginating.
func (c *PixivCrawler) FetchAllIllustFollows(opts *pixiv.IllustFollowOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	// Set up the logger
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(handler)

	logger.Info("Fetching followed users' illustrations")

	for {
		var illusts []models.Illust
		illusts, next, err = c.app.IllustFollow([]pixiv.IllustFollowOptions{*opts}...)

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
