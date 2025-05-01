package pixiv

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
