package crawler

import (
	"time"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/models"
)

// FetchAllUserFollowing retrieves all users followed by the specified user by paginating.
func (c *PixivCrawler) FetchAllUserFollowing(uid uint64, opts *pixiv.UserFollowingOptions, sleepMs ...int) ([]models.UserPreview, error) {
	var allUsers []models.UserPreview
	var next int
	var err error

	for {
		var users []models.UserPreview
		users, next, err = c.app.UserFollowing(uid, []pixiv.UserFollowingOptions{*opts}...)

		allUsers = append(allUsers, users...)
		if err != nil {
			return allUsers, err
		}

		if next == 0 {
			break
		}

		opts.Offset = &next

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		time.Sleep(sleepDuration)
	}

	return allUsers, nil
}
