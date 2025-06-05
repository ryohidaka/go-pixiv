package apptest

import (
	"embed"
	"fmt"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

//go:embed fixtures/json/*.json
var fixtures embed.FS

// MockResponseFromFile mocks an API response from an embedded JSON file
// for the specified URL path and HTTP method (GET, POST, etc.).
func MockResponseFromFile(method, url, path string) error {
	// Ensure `.json` extension
	if !strings.HasSuffix(path, ".json") {
		path += ".json"
	}

	filePath := fmt.Sprintf("fixtures/json/%s", path)

	// Read embedded file
	data, err := fixtures.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read embedded mock response file: %v", err)
	}

	// Register the mocked response
	httpmock.RegisterResponder(method, url,
		httpmock.NewStringResponder(200, string(data)))

	return nil
}

// WithMockHTTP activates httpmock for the duration of the test function and then deactivates it.
// It ensures that HTTP requests are intercepted and handled using registered mock responders.
func WithMockHTTP(t *testing.T, testFunc func()) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testFunc()
}
