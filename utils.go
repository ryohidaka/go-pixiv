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
)

// genClientHash returns an MD5 hash generated from the given clientTime and a predefined secret.
// This is typically used for authentication or request validation.
//
// Parameters:
//   - clientTime: A timestamp string used as part of the hash input.
//
// Returns:
//   - A hexadecimal string representing the MD5 hash.
func genClientHash(clientTime string) string {
	h := md5.New()
	io.WriteString(h, clientTime)
	io.WriteString(h, ClientHashSecret)
	return hex.EncodeToString(h.Sum(nil))
}

// setHeaders sets custom HTTP headers to the provided request.
//
// Parameters:
//   - req: The HTTP request to which the headers should be applied.
//   - headers: A map of header keys and values to set.
func setHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

// readResponse reads the body of an HTTP response and returns it as a byte slice.
//
// Parameters:
//   - resp: The HTTP response to read from.
//
// Returns:
//   - A byte slice containing the response body.
//   - An error if reading the body fails.
func readResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	return body, nil
}

// decodeJSON unmarshals a JSON byte slice into the provided output structure.
//
// Parameters:
//   - body: The raw JSON response body.
//   - out: A pointer to the structure where the data should be decoded.
//
// Returns:
//   - An error if JSON unmarshalling fails.
func decodeJSON(body []byte, out any) error {
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return nil
}

// getExpiresAt calculates and returns the time when the access token will expire.
//
// Parameters:
//   - expiredIn: The number of seconds from the current time until expiration.
//
// Returns:
//   - time.Time: The time when the access token will expire.
func getExpiresAt(expiredIn int) time.Time {
	return time.Now().Add(time.Duration(expiredIn) * time.Second)
}

// parseNextPageOffset extracts the offset value from the query parameters of a given URL string.
//
// Parameters:
//   - s: A string representing the full URL (e.g., "https://example.com?page=2&offset=100").
//   - field: The name of the query parameter to extract the offset from (e.g., "offset").
//
// Returns:
//   - int: The extracted offset value. Returns 0 if the input URL is empty.
//   - error: An error if the URL is invalid, the parameter is missing, or the value cannot be converted to an integer.
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
