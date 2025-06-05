package apptest

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

// MockResponseFromFile mocks an API response from an external JSON file
// for the specified URL path and HTTP method (GET, POST, etc.).
// It allows optional customization of the base path used to locate the JSON fixture.
func MockResponseFromFile(method, url, path string, basePath ...string) error {
	root := "testutil"
	if len(basePath) > 0 && basePath[0] != "" {
		root = basePath[0]
	}

	// Build full file path
	filePath := fmt.Sprintf("%s/fixtures/json/%s.json", root, path)

	// Open the JSON file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open mock response file: %v", err)
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read mock response file: %v", err)
	}

	// Register the mocked response for the specified URL path and HTTP method
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
