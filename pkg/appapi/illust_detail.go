package appapi

import (
	"github.com/ryohidaka/go-pixiv/models"
	"github.com/ryohidaka/go-pixiv/models/appmodel"
)

// illustDetailParams represents the query parameters for fetching illustration details.
// It includes the ID of the illustration to be retrieved.
type illustDetailParams struct {
	IllustID uint64 `url:"illust_id,omitempty"` // ID of the illustration
}

// IllustDetail retrieves detailed information about a specific illustration by its ID.
func (a *AppPixivAPI) IllustDetail(id uint64) (*models.Illust, error) {
	const path = "v1/illust/detail"

	// Construct request parameters
	params := &illustDetailParams{
		IllustID: id, // Set the illustration ID parameter
	}

	// Initialize the response model
	data := &appmodel.IllustResponse{}

	// Send the API request
	if err := a.Get(path, params, data); err != nil {
		return nil, err
	}

	// Return the illustration data from the response
	return &data.Illust, nil
}
