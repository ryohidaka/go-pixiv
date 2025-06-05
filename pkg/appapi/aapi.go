package appapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ryohidaka/go-pixiv/internal/apputils"
	"github.com/ryohidaka/go-pixiv/internal/coreutils"

	"github.com/ryohidaka/go-pixiv/models/appmodel"
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

	params := &appmodel.AuthParams{
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
	auth.ExpiresAt = apputils.GetExpiresAt(authInfo.ExpiresIn)

	return &AppPixivAPI{
		httpClient: http.DefaultClient,
		auth:       auth,
	}, nil
}

// Get sends a GET request to the specified path with query parameters,
// and decodes the JSON response into the provided output structure.
func Get(a *AppPixivAPI, path string, queryStruct any, out any) error {
	return request(a, "GET", path, queryStruct, nil, out)
}

// Post sends a POST request to the specified path with query parameters and body,
// and decodes the JSON response into the provided output structure.
func Post(a *AppPixivAPI, path string, queryStruct any, body io.Reader, out any) error {
	return request(a, "POST", path, queryStruct, body, out)
}

// Request sends an HTTP request (GET, POST, etc.) to the specified Pixiv API endpoint,
// optionally including OAuth authorization.
func request(a *AppPixivAPI, method, path string, queryStruct any, body io.Reader, out any) error {
	if err := refreshTokenIfNeeded(a); err != nil {
		return err
	}
	baseUrl := AppHosts + path
	reqURL, err := coreutils.BuildRequestURL(baseUrl, queryStruct)
	if err != nil {
		return err
	}

	req, err := createRequest(a, method, reqURL, body)
	if err != nil {
		return err
	}

	return handleResponse(a, req, out)
}

// refreshTokenIfNeeded refreshes the access token if it is expired or about to expire.
func refreshTokenIfNeeded(a *AppPixivAPI) error {
	if _, err := a.auth.RefreshAuth(false); err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	return nil
}

// createRequest creates a new HTTP request with required headers.
func createRequest(a *AppPixivAPI, method string, reqURL *url.URL, body io.Reader) (*http.Request, error) {
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
	if body != nil && (method == "POST" || method == "PUT" || method == "PATCH") {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	apputils.SetHeaders(req, headers)
	return req, nil
}

// handleResponse sends the HTTP request and decodes the response body into `out`.
func handleResponse(a *AppPixivAPI, req *http.Request, out any) error {
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
