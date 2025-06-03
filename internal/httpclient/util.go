package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SetHeaders sets custom HTTP headers to the provided request.
func SetHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

// reaReadResponsedResponse reads the body of an HTTP response and returns it as a byte slice.
func ReadResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return body, nil
}

// DecodeJSON unmarshals a JSON byte slice into the provided output structure.
func DecodeJSON(body []byte, out any) error {
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}
