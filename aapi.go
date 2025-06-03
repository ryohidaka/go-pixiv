package pixiv

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ryohidaka/go-pixiv/models"
)

// AppPixivAPI handles Pixiv App API operations using OAuth authentication.
type AppPixivAPI struct {
	httpClient *http.Client
	auth       *AuthSession
}

// NewApp initializes and returns a new AppPixivAPI instance using the provided refresh token.
func NewApp(refreshToken string) (*AppPixivAPI, error) {
	auth := &AuthSession{
		RefreshToken: refreshToken,
		HTTPClient:   http.DefaultClient,
	}

	params := &models.AuthParams{
		GetSecureURL: 1,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}

	authInfo, err := auth.Authenticate(params)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	auth.AccessToken = authInfo.AccessToken
	auth.RefreshToken = authInfo.RefreshToken
	auth.ExpiresAt = getExpiresAt(authInfo.ExpiresIn)

	return &AppPixivAPI{
		httpClient: http.DefaultClient,
		auth:       auth,
	}, nil
}

// Get sends a GET request to the specified path with query parameters,
// and decodes the JSON response into the provided output structure.
func (a *AppPixivAPI) Get(path string, queryStruct any, out any) error {
	return a.request("GET", path, queryStruct, nil, out)
}

// Post sends a POST request to the specified path with query parameters and body,
// and decodes the JSON response into the provided output structure.
func (a *AppPixivAPI) Post(path string, queryStruct any, body io.Reader, out any) error {
	return a.request("POST", path, queryStruct, body, out)
}

// Request sends an HTTP request (GET, POST, etc.) to the specified Pixiv API endpoint,
// optionally including OAuth authorization.
func (a *AppPixivAPI) request(method, path string, queryStruct any, body io.Reader, out any) error {
	if err := a.refreshTokenIfNeeded(); err != nil {
		return err
	}

	reqURL, err := a.buildRequestURL(path, queryStruct)
	if err != nil {
		slog.Error("Failed to build request URL", slog.String("error", err.Error()))
		return err
	}

	req, err := a.createRequest(method, reqURL, body)
	if err != nil {
		slog.Error("Failed to create request", slog.String("error", err.Error()))
		return err
	}

	return a.handleResponse(req, out)
}

// refreshTokenIfNeeded refreshes the access token if it is expired or about to expire.
func (a *AppPixivAPI) refreshTokenIfNeeded() error {
	if _, err := a.auth.RefreshAuth(false); err != nil {
		slog.Error("Token refresh failed", slog.String("error", err.Error()))
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	return nil
}

// buildRequestURL constructs the complete URL with encoded query parameters.
func (a *AppPixivAPI) buildRequestURL(path string, queryStruct any) (*url.URL, error) {
	reqURL, err := url.Parse(AppHosts + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	if queryStruct != nil {
		values, err := query.Values(queryStruct)
		if err != nil {
			return nil, fmt.Errorf("failed to encode query parameters: %w", err)
		}
		reqURL.RawQuery = values.Encode()
	}

	return reqURL, nil
}

// createRequest creates a new HTTP request with required headers.
func (a *AppPixivAPI) createRequest(method string, reqURL *url.URL, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, reqURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	headers := map[string]string{
		"User-Agent":     UserAgent,
		"App-Version":    AppVersion,
		"App-OS-VERSION": AppOSVersion,
		"App-OS":         AppOS,
		"Authorization":  "Bearer " + a.auth.AccessToken,
	}

	// Content-Type is set for methods with body
	if body != nil && (method == "POST" || method == "PUT" || method == "PATCH") {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	setHeaders(req, headers)

	return req, nil
}

// handleResponse sends the HTTP request and decodes the response body into `out`.
func (a *AppPixivAPI) handleResponse(req *http.Request, out any) error {
	resp, err := a.httpClient.Do(req)
	if err != nil {
		slog.Error("API request failed", slog.String("error", err.Error()))
		return fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		slog.Error("API error", slog.Int("status", resp.StatusCode), slog.String("status_text", http.StatusText(resp.StatusCode)))
		return fmt.Errorf("API error: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Skip decoding if output destination is nil
	if out == nil {
		return nil
	}

	// Decode JSON response into the provided output variable
	return json.NewDecoder(resp.Body).Decode(out)
}
