package crawler

import (
	"time"

	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
)

// FetchAllIllustFollows retrieves all illustrations from followed users by paginating.
func (c *PixivCrawler) FetchAllIllustFollows(opts *appapi.IllustFollowOptions, sleepMs ...int) ([]models.Illust, error) {
	var allIllusts []models.Illust
	var next int
	var err error

	for {
		var illusts []models.Illust
		illusts, next, err = c.app.IllustFollow([]appapi.IllustFollowOptions{*opts}...)

		allIllusts = append(allIllusts, illusts...)
		if err != nil {
			return allIllusts, err
		}

		if next == 0 {
			break
		}

		opts.Offset = &next

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		time.Sleep(sleepDuration)
	}

	return allIllusts, nil
}
