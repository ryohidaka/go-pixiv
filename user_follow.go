package pixiv

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ryohidaka/go-pixiv/models"
)

// UserFollowAdd sends a follow request to a user on Pixiv.
func (a *AppPixivAPI) UserFollowAdd(uid uint64, restrict ...models.Restrict) (bool, error) {
	const path = "v1/user/follow/add"

	// Set default restriction to Public unless a custom one is provided
	r := models.Public
	if len(restrict) > 0 {
		r = restrict[0]
	}

	// Construct form-encoded data
	data := url.Values{}
	data.Set("user_id", fmt.Sprintf("%d", uid))
	data.Set("restrict", string(r)) // Ensure models.Restrict is castable to string

	// Convert form data to io.Reader
	body := strings.NewReader(data.Encode())

	// Send POST request with form-encoded body
	if err := a.Post(path, nil, body, nil); err != nil {
		return false, err
	}

	return true, nil
}

// UserFollowAdd sends a unfollow request to a user on Pixiv.
func (a *AppPixivAPI) UserFollowDelete(uid uint64) (bool, error) {
	const path = "v1/user/follow/delete"

	// Construct form-encoded data
	data := url.Values{}
	data.Set("user_id", fmt.Sprintf("%d", uid))

	// Convert form data to io.Reader
	body := strings.NewReader(data.Encode())

	// Send POST request with form-encoded body
	if err := a.Post(path, nil, body, nil); err != nil {
		return false, err
	}

	return true, nil
}
