package webutils

import (
	"encoding/json"
	"fmt"
	"io"
)

// DecodeJSON decodes a JSON body into the given type T.
func DecodeJSON[T any](body io.Reader) (*T, error) {
	raw, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var result T
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}
