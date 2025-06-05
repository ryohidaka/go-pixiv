package apputils

import (
	"fmt"
	"net/url"
	"strconv"
)

// ParseNextPageOffset extracts the offset value from the query parameters of a given URL string.
func ParseNextPageOffset(s, field string) (int, error) {
	// If the input string is empty, return 0 as a default offset.
	if s == "" {
		return 0, nil
	}

	// Parse the input string into a URL structure.
	u, err := url.Parse(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse URL: %s {%s}", s, err)
	}

	// Parse the raw query string from the URL.
	queryParams, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to parse query parameters: %s {%s}", s, err)
	}

	// Retrieve the offset value from the query parameters using the specified field name.
	offsetParam := queryParams.Get(field)
	if offsetParam == "" {
		return 0, fmt.Errorf("missing query parameter: %s", field)
	}

	// Convert the offset parameter from string to integer.
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return 0, fmt.Errorf("invalid offset value: %s {%s}", offsetParam, err)
	}

	return offset, nil
}
