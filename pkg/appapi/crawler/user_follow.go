package crawler

import (
	"fmt"
	"time"

	"github.com/ryohidaka/go-pixiv/models"
)

// UserFollowAddMultiple sends follow requests to multiple users on Pixiv.
func (c *PixivCrawler) UserFollowAddMultiple(uids []uint64, restrict *models.Restrict, sleepMs ...int) ([]uint64, error) {
	var processed []uint64

	for _, uid := range uids {
		success, err := c.app.UserFollowAdd(uid, []models.Restrict{*restrict}...)
		processed = append(processed, uid)

		if err != nil {
			return processed, err // Exit loop and return error with processed IDs
		}
		if !success {
			return processed, fmt.Errorf("failed to follow user %d", uid)
		}

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		time.Sleep(sleepDuration)
	}

	return processed, nil
}

// UserFollowDeleteMultiple sends unfollow requests to multiple users on Pixiv.
func (c *PixivCrawler) UserFollowDeleteMultiple(uids []uint64, sleepMs ...int) ([]uint64, error) {
	var processed []uint64

	for _, uid := range uids {
		success, err := c.app.UserFollowDelete(uid)
		processed = append(processed, uid)

		if err != nil {
			return processed, err // Exit loop and return error with processed IDs
		}
		if !success {
			return processed, fmt.Errorf("failed to follow user %d", uid)
		}

		// Sleep between requests to avoid rate limits
		sleepDuration := getSleepDuration(sleepMs...)
		time.Sleep(sleepDuration)
	}

	return processed, nil
}
