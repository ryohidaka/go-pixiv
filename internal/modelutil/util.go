package modelutil

import (
	"github.com/ryohidaka/go-pixiv/models"
)

// GetRestrict safely dereferences a *Restrict pointer and returns its value.
// If the pointer is nil or the value is an empty string, it returns models.Public.
func GetRestrict(r *models.Restrict) models.Restrict {
	// Check if the pointer is non-nil and the value is not an empty string
	if r != nil && *r != "" {
		return *r
	}
	// Return the default value when the pointer is nil or empty
	return models.Public
}
