package pixiv

import (
	"fmt"
	"go-pixiv/models"
	"net/http"
)

// AppPixivAPI handles Pixiv app API operations.
type AppPixivAPI struct {
	httpClient *http.Client
	auth       *AuthSession
}

// NewApp creates a new instance of AppPixivAPI and initializes authentication.
//
// Parameters:
//   - refreshToken: the refresh token
//
// Returns:
//   - *AppPixivAPI: initialized API client
//   - error: if token is invalid
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
