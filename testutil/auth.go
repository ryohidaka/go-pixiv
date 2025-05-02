package testutil

import (
	"go-pixiv"
)

// CreateAuthSession returns a new instance of AuthSession,
// initialized with the given refresh token and the default Pixiv base URL.
// This helper is intended for use in unit tests.
//
// Parameters:
//   - refreshToken: The refresh token string used to initialize the AuthSession.
//
// Returns:
//   - A pointer to a pixiv.AuthSession initialized with the provided refresh token.
func CreateAuthSession(refreshToken string) *pixiv.AuthSession {
	return &pixiv.AuthSession{
		BaseURL:      pixiv.AuthHosts,
		RefreshToken: refreshToken,
	}
}
