package crawler

import (
	"fmt"

	"github.com/ryohidaka/go-pixiv/models"
)

// UserFollowAddMultiple sends follow requests to multiple users on Pixiv.
//
// If the optional restrict parameter is provided, it will be used for all users;
// otherwise, the default is models.Public.
//
// Parameters:
//   - uids: A slice of user IDs to follow.
//   - restrict (optional): Restriction level of the follow (e.g., Public or Private).
//
// Returns:
//   - []uint64: A list of user IDs that have been processed (success or failure).
//   - error: An error object if any request fails; otherwise, nil.
func (c *PixivCrawler) UserFollowAddMultiple(uids []uint64, restrict ...models.Restrict) ([]uint64, error) {
	var processed []uint64

	for _, uid := range uids {
		success, err := c.app.UserFollowAdd(uid, restrict...)
		processed = append(processed, uid)

		if err != nil {
			return processed, err // Exit loop and return error with processed IDs
		}
		if !success {
			return processed, fmt.Errorf("failed to follow user %d", uid)
		}
	}

	return processed, nil
}

// UserFollowDeleteMultiple sends unfollow requests to multiple users on Pixiv.
//
// Parameters:
//   - uids: A slice of user IDs to follow.
//
// Returns:
//   - []uint64: A list of user IDs that have been processed (success or failure).
//   - error: An error object if any request fails; otherwise, nil.
func (c *PixivCrawler) UserFollowDeleteMultiple(uids []uint64) ([]uint64, error) {
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
	}

	return processed, nil
}
