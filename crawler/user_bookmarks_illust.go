package crawler

import (
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

// FetchAllBookmarkedIllusts retrieves all bookmarked illustrations for a given user by paginating
func (c *PixivCrawler) FetchAllBookmarkedIllusts(uid uint64, opts *pixiv.UserBookmarksIllustOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	// Start the fetch process
	for {
		// Retrieve a page of bookmarked illustrations
		var illusts []models.Illust
		illusts, next, err = c.app.UserBookmarksIllust(uid, []pixiv.UserBookmarksIllustOptions{*opts}...)

		// Always append the successfully fetched results before returning on error
		allIllusts = append(allIllusts, illusts...)
		if err != nil {
			return allIllusts, err
		}

		// Exit loop if there are no more pages
		if next == 0 {
			break
		}

		// Set the next max_bookmark_id for pagination
		opts.MaxBookmarkID = &next

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		time.Sleep(sleepDuration)
	}

	return allIllusts, nil
}
