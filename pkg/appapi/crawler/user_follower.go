package crawler

import (
	"time"

	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
)

// FetchAllUserFollowers retrieves all followers of the specified user by paginating.
func (c *PixivCrawler) FetchAllUserFollowers(uid uint64, opts *appapi.UserFollowerOptions, sleepMs ...int) ([]models.UserPreview, error) {
	var allFollowers []models.UserPreview
	var next int
	var err error

	for {
		var followers []models.UserPreview
		followers, next, err = c.app.UserFollower(uid, []appapi.UserFollowerOptions{*opts}...)

		allFollowers = append(allFollowers, followers...)
		if err != nil {
			return allFollowers, err
		}

		if next == 0 {
			break
		}

		opts.Offset = &next

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		time.Sleep(sleepDuration)
	}

	return allFollowers, nil
}
