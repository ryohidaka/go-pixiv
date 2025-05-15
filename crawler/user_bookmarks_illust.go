package crawler

import (
	"log/slog"
	"os"
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

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
func (c *PixivCrawler) FetchAllBookmarkedIllusts(uid uint64, opts *pixiv.UserBookmarksIllustOptions, sleepMs ...int) ([]models.Illust, error) {
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
		illusts, next, err = c.app.UserBookmarksIllust(uid, []pixiv.UserBookmarksIllustOptions{*opts}...)

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
