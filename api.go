package pixiv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ryohidaka/go-pixiv/models"
)

// AuthSession holds the authentication state and provides methods to
// authenticate and refresh the access token.
type AuthSession struct {
	BaseURL      string                                // Base URL for the Pixiv API. If empty, defaults to the standard AuthHosts.
	AccessToken  string                                // The current access token.
	RefreshToken string                                // The current refresh token.
	ExpiresAt    time.Time                             // The time when the access token expires.
	AuthHook     func(string, string, time.Time) error // Optional hook called after successful authentication.
	HTTPClient   *http.Client                          // Optional custom HTTP client. Defaults to http.DefaultClient if nil.
}

// Authenticate performs authentication against the Pixiv API using the provided
// AuthParams. It updates the session with the new access and refresh tokens.
func (s *AuthSession) Authenticate(params *models.AuthParams) (*models.AuthInfo, error) {
	if s.BaseURL == "" {
		s.BaseURL = AuthHosts
	}

	clientTime := time.Now().Format(time.RFC3339)

	form := url.Values{
		"client_id":      {params.ClientID},
		"client_secret":  {params.ClientSecret},
		"grant_type":     {params.GrantType},
		"refresh_token":  {params.RefreshToken},
		"get_secure_url": {fmt.Sprintf("%d", params.GetSecureURL)},
	}

	req, err := http.NewRequest("POST", s.BaseURL+"auth/token", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	headers := map[string]string{
		"Content-Type":   "application/x-www-form-urlencoded",
		"User-Agent":     UserAgent,
		"X-Client-Time":  clientTime,
		"X-Client-Hash":  genClientHash(clientTime),
		"App-OS":         AppOS,
		"App-OS-Version": AppOSVersion,
	}
	setHeaders(req, headers)

	client := s.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("auth request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		var pixivErr models.PixivError
		if err := json.Unmarshal(body, &pixivErr); err == nil && pixivErr.HasError {
			for k, v := range pixivErr.Errors {
				return nil, fmt.Errorf("login %s error: %s", k, v.Message)
			}
		}

		return nil, fmt.Errorf("auth failed with status %d: %s", resp.StatusCode, string(body))
	}

	var res models.AuthResponse
	if err := decodeJSON(body, &res); err != nil {
		return nil, err
	}

	s.AccessToken = res.Response.AccessToken
	s.RefreshToken = res.Response.RefreshToken
	s.ExpiresAt = getExpiresAt(res.Response.ExpiresIn)

	if s.AuthHook != nil {
		if err := s.AuthHook(s.AccessToken, s.RefreshToken, s.ExpiresAt); err != nil {
			return nil, err
		}
	}

	return res.Response, nil
}

// RefreshAuth refreshes the access token if it has expired or if forced.
func (s *AuthSession) RefreshAuth(force bool) (*models.Account, error) {
	if s.RefreshToken == "" {
		return nil, fmt.Errorf("missing refresh token")
	}
	if !force && time.Now().Before(s.ExpiresAt) {
		return nil, nil
	}

	params := &models.AuthParams{
		GetSecureURL: 1,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		GrantType:    "refresh_token",
		RefreshToken: s.RefreshToken,
	}
	info, err := s.Authenticate(params)

	if err != nil {
		return nil, err
	}

	return &info.User, nil
}
