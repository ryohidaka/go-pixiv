package pixiv

import "github.com/ryohidaka/go-pixiv/models"

// illustDetailParams represents the query parameters for fetching illustration details.
// It includes the ID of the illustration to be retrieved.
type illustDetailParams struct {
	IllustID uint64 `url:"illust_id,omitempty"` // ID of the illustration
}

// IllustDetail retrieves detailed information about a specific illustration by its ID.
//
// Parameters:
//   - id: The unique ID of the illustration to fetch.
//
// Returns:
//   - A pointer to an Illust object containing detailed information about the illustration.
//   - An error if the request fails or the response is invalid.
func (a *AppPixivAPI) IllustDetail(id uint64) (*models.Illust, error) {
	path := "v1/illust/detail"

	// Construct request parameters
	params := &illustDetailParams{
		IllustID: id, // Set the illustration ID parameter
	}

	// Initialize the response model
	data := &models.IllustResponse{}

	// Perform the API request with the specified path and parameters
	if err := a.Request(path, params, data); err != nil {
		return nil, err
	}

	// Return the illustration data from the response
	return &data.Illust, nil
}
