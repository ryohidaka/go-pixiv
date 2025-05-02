package testutil

import (
	"fmt"
	"io"
	"os"

	"github.com/jarcoal/httpmock"
)

// MockResponseFromFile mocks an API response from an external JSON file
// for the specified URL path and HTTP method (GET, POST, etc.)
// It reads the JSON file and registers the response using the httpmock package.
//
// Parameters:
// - method: The HTTP method (e.g., "GET", "POST")
// - url: The URL path for which the mock response will be registered
// - fileName: The name of the JSON file that contains the mock response data
//
// Returns:
// - An error if there is an issue reading the file or registering the mock response
func MockResponseFromFile(method, url, fileName string) error {
	// Open the JSON file for reading
	file, err := os.Open("testutil/fixtures/json/" + fileName)
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
