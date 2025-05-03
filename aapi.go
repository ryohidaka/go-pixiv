package pixiv

import (
	"encoding/json"
	"fmt"
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
//
// It performs the initial token refresh and sets up authentication state.
//
// Parameters:
//   - refreshToken: The OAuth refresh token to authenticate the user.
//
// Returns:
//   - *AppPixivAPI: A pointer to the initialized AppPixivAPI instance.
//   - error: An error if authentication fails.
func NewApp(refreshToken string) (*AppPixivAPI, error) {
	slog.Debug("Initializing AppPixivAPI", slog.String("refresh_token", refreshToken))

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
		slog.Error("Authentication failed", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	auth.AccessToken = authInfo.AccessToken
	auth.RefreshToken = authInfo.RefreshToken
	auth.ExpiresAt = getExpiresAt(authInfo.ExpiresIn)

	slog.Debug("Authentication successful",
		slog.String("access_token", auth.AccessToken),
		slog.Time("expires_at", auth.ExpiresAt),
	)

	return &AppPixivAPI{
		httpClient: http.DefaultClient,
		auth:       auth,
	}, nil
}

// request sends a GET request to the specified Pixiv API endpoint,
// optionally including OAuth authorization.
//
// Parameters:
//   - path: API path (e.g., "/v1/user/detail").
//   - queryStruct: Struct containing query parameters, encoded via `query.Values`.
//   - out: A pointer to a variable where the response will be decoded.
//
// Returns:
//   - error: An error if request or decoding fails.
func (a *AppPixivAPI) Request(path string, queryStruct any, out any) error {
	slog.Debug("Request start", slog.String("path", path))
	if err := a.refreshTokenIfNeeded(); err != nil {
		return err
	}

	reqURL, err := a.buildRequestURL(path, queryStruct)
	if err != nil {
		slog.Error("Failed to build request URL", slog.String("error", err.Error()))
		return err
	}

	req, err := a.createRequest(reqURL)
	if err != nil {
		slog.Error("Failed to create request", slog.String("error", err.Error()))
		return err
	}

	return a.handleResponse(req, out)
}

// refreshTokenIfNeeded refreshes the access token if it is expired or about to expire.
//
// Returns:
//   - error: An error if the token refresh fails.
func (a *AppPixivAPI) refreshTokenIfNeeded() error {
	slog.Debug("Checking if token refresh is needed")
	if _, err := a.auth.RefreshAuth(false); err != nil {
		slog.Error("Token refresh failed", slog.String("error", err.Error()))
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	slog.Debug("Token refresh succeeded")
	return nil
}

// buildRequestURL constructs the complete URL with encoded query parameters.
//
// Parameters:
//   - path: The API endpoint path (e.g., "/v1/user/detail").
//   - queryStruct: A struct to be encoded as URL query parameters.
//
// Returns:
//   - *url.URL: The complete request URL.
//   - error: An error if URL parsing or query encoding fails.
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
	slog.Debug("Built request URL", slog.String("url", reqURL.String()))
	return reqURL, nil
}

// createRequest creates a new HTTP GET request with required headers.
//
// Parameters:
//   - reqURL: The full request URL.
//
// Returns:
//   - *http.Request: The constructed HTTP request.
//   - error: An error if request creation fails.
func (a *AppPixivAPI) createRequest(reqURL *url.URL) (*http.Request, error) {
	req, err := http.NewRequest("GET", reqURL.String(), nil)
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

	setHeaders(req, headers)
	slog.Debug("Created request with headers", slog.Any("headers", headers))
	return req, nil
}

// handleResponse sends the HTTP request and decodes the response body into `out`.
//
// Parameters:
//   - req: The HTTP request to execute.
//   - out: A pointer to the value to decode the JSON response into.
//
// Returns:
//   - error: An error if the request fails or the response is invalid.
func (a *AppPixivAPI) handleResponse(req *http.Request, out any) error {
	resp, err := a.httpClient.Do(req)
	if err != nil {
		slog.Error("API request failed", slog.String("error", err.Error()))
		return fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	slog.Debug("Received response", slog.Int("status", resp.StatusCode))

	if resp.StatusCode >= 400 {
		slog.Error("API error", slog.Int("status", resp.StatusCode), slog.String("status_text", http.StatusText(resp.StatusCode)))
		return fmt.Errorf("API error: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
