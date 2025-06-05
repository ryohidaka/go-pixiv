package coreutils

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// buildRequestURL constructs the complete URL with encoded query parameters from a query struct.
// If queryStruct is nil, no query parameters will be appended.
func BuildRequestURL(baseURL string, queryStruct any) (*url.URL, error) {
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if queryStruct != nil {
		values, err := query.Values(queryStruct)
		if err != nil {
			return nil, fmt.Errorf("failed to encode query parameters: %w", err)
		}
		reqURL.RawQuery = values.Encode()
	}

	return reqURL, nil
}
