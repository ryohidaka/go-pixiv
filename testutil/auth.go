package testutil

import (
	"github.com/ryohidaka/go-pixiv"
)

// CreateAuthSession returns a new instance of AuthSession,
// initialized with the given refresh token and the default Pixiv base URL.
// This helper is intended for use in unit tests.
func CreateAuthSession(refreshToken string) *pixiv.AuthSession {
	return &pixiv.AuthSession{
		BaseURL:      pixiv.AuthHosts,
		RefreshToken: refreshToken,
	}
}
