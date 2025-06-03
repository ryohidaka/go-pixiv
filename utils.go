package pixiv

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ryohidaka/go-pixiv/models"
)

// genClientHash returns an MD5 hash generated from the given clientTime and a predefined secret.
// This is typically used for authentication or request validation.
func genClientHash(clientTime string) string {
	h := md5.New()
	io.WriteString(h, clientTime)
	io.WriteString(h, ClientHashSecret)
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

// setHeaders sets custom HTTP headers to the provided request.
func setHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

// readResponse reads the body of an HTTP response and returns it as a byte slice.
func readResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return body, nil
}

// decodeJSON unmarshals a JSON byte slice into the provided output structure.
func decodeJSON(body []byte, out any) error {
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

// getExpiresAt calculates and returns the time when the access token will expire.
func getExpiresAt(expiredIn int) time.Time {
	expiresAt := time.Now().Add(time.Duration(expiredIn) * time.Second)

	return expiresAt
}

// parseNextPageOffset extracts the offset value from the query parameters of a given URL string.
func parseNextPageOffset(s, field string) (int, error) {
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

// getRestrict safely dereferences a *Restrict pointer and returns its value.
// If the pointer is nil or the value is an empty string, it returns models.Public.
func getRestrict(r *models.Restrict) models.Restrict {
	// Check if the pointer is non-nil and the value is not an empty string
	if r != nil && *r != "" {
		return *r
	}
	// Return the default value when the pointer is nil or empty
	return models.Public
}
