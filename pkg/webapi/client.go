package webapi

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ryohidaka/go-pixiv/internal/coreutils"
	"github.com/ryohidaka/go-pixiv/internal/webutils"
)

type WebPixivAPI struct {
	httpClient *http.Client
}

// NewWebApp initializes and returns a new instance of WebPixivAPI.
func NewWebApp(phpsessid string) (*WebPixivAPI, error) {
	baseURL, err := url.Parse(ApiHosts)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	jar, err := webutils.NewCookieJar(phpsessid, baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	client := &http.Client{
		Jar: jar,
		Transport: webutils.WithDefaultHeaders(
			http.DefaultTransport,
			UserAgent,
			baseURL.String(),
		),
	}

	return &WebPixivAPI{
		httpClient: client,
	}, nil
}

// Get sends a GET request to the given API path and decodes the response into the specified type T.
func Get[T any](a *WebPixivAPI, path string, referer *string, queryParams any) (*T, error) {
	baseURL := ApiHosts + path

	// Construct full URL with query parameters
	fullURL, err := coreutils.BuildRequestURL(baseURL, queryParams)
	if err != nil {
		return nil, err
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	// Set referer header if provided
	if referer != nil {
		req.Header.Set("Referer", *referer)
	}

	// Execute request using the shared HTTP client
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("GET failed with status %d: %s", resp.StatusCode, resp.Status)
	}

	// Decode response body into type T
	return webutils.DecodeJSON[T](resp.Body)
}
