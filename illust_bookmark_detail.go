package pixiv

import "github.com/ryohidaka/go-pixiv/models"

// illustBookmarkDetailParams represents the request parameters
// for retrieving illustration bookmark detail from the Pixiv API.
type illustBookmarkDetailParams struct {
	IllustID uint64 `url:"illust_id,omitempty"` // The ID of the illustration
}

// IllustBookmarkDetail retrieves the bookmark detail information
// for a specific illustration by its ID.
//
// Parameters:
//   - id: The illustration ID to fetch the bookmark details for.
//
// Returns:
//   - A pointer to IllustBookmarkDetail containing the bookmark info.
//   - An error if the API request fails.
func (a *AppPixivAPI) IllustBookmarkDetail(id uint64) (*models.IllustBookmarkDetail, error) {
	const path = "v2/illust/bookmark/detail"

	// Construct request parameters
	params := &illustBookmarkDetailParams{
		IllustID: id, // Set the illustration ID parameter
	}

	// Initialize the response model
	data := &models.IllustBookmarkDetail{}

	// Perform the API request with the specified path and parameters
	if err := a.Request(path, params, data); err != nil {
		return nil, err
	}

	// Return the illustration bookmark detail data from the response
	return data, nil
}
