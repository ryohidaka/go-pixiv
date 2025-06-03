package crawler

import (
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

// FetchAllUserIllusts retrieves all illustrations posted by a given user by paginating.
func (c *PixivCrawler) FetchAllUserIllusts(uid uint64, opts *pixiv.UserIllustsOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	for {
		// Fetch a single page
		var illusts []models.Illust
		illusts, next, err = c.app.UserIllusts(uid, []pixiv.UserIllustsOptions{*opts}...)

		// Append successfully fetched data
		allIllusts = append(allIllusts, illusts...)

		if err != nil {
			return allIllusts, err
		}

		// Exit if there is no next page
		if next == 0 {
			break
		}

		// Update offset for pagination
		opts.Offset = &next

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		time.Sleep(sleepDuration)
	}

	return allIllusts, nil
}
