package pixiv

import (
	"github.com/ryohidaka/go-pixiv/models"
)

// illustBookmarkDetailParams represents the request parameters
// for retrieving illustration bookmark detail from the Pixiv API.
type illustBookmarkDetailParams struct {
	IllustID uint64 `url:"illust_id,omitempty"` // The ID of the illustration
}

// IllustBookmarkDetail retrieves the bookmark detail information
// for a specific illustration by its ID.
func (a *AppPixivAPI) IllustBookmarkDetail(id uint64) (*models.IllustBookmarkDetail, error) {
	const path = "v2/illust/bookmark/detail"

	// Construct request parameters
	params := &illustBookmarkDetailParams{
		IllustID: id, // Set the illustration ID parameter
	}

	// Initialize the response model
	data := &models.IllustBookmarkDetail{}

	// Send the API request
	if err := a.Get(path, params, data); err != nil {
		return nil, err
	}

	// Return the illustration bookmark detail data from the response
	return data, nil
}
