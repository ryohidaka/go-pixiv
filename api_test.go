package pixiv_test

import (
	"go-pixiv"
	"go-pixiv/models"
	"go-pixiv/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestAuthenticate tests the Authenticate method in AuthSession.
func TestAuthenticate(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		err := testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth.json")
		if err != nil {
			t.Fatalf("Failed to mock response: %v", err)
		}

		// Prepare AuthSession and AuthParams
		authSession := testutil.CreateAuthSession("validRefreshToken")
		params := &models.AuthParams{
			ClientID:     "clientID",
			ClientSecret: "clientSecret",
			GrantType:    "refresh_token",
			RefreshToken: "validRefreshToken",
			GetSecureURL: 1,
		}

		// Call Authenticate
		authInfo, err := authSession.Authenticate(params)

		// Assert no error and correct values
		assert.NoError(t, err)
		assert.Equal(t, "validAccessToken", authSession.AccessToken)
		assert.Equal(t, "validRefreshToken", authSession.RefreshToken)
		assert.True(t, time.Now().Before(authSession.ExpiresAt))
		assert.Equal(t, "validAccessToken", authInfo.AccessToken)
		assert.Equal(t, "validRefreshToken", authInfo.RefreshToken)
	})
}

// TestRefreshAuth tests the RefreshAuth method in AuthSession.
func TestRefreshAuth(t *testing.T) {
	testutil.WithMockHTTP(t, func() {
		// Mock the Pixiv API response for refreshing the token.
		err := testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth.json")
		if err != nil {
			t.Fatalf("Failed to mock response: %v", err)
		}

		// Prepare AuthSession and set an initial refresh token
		authSession := testutil.CreateAuthSession("validRefreshToken")
		authSession.ExpiresAt = time.Now().Add(-time.Hour) // force refresh

		// Call RefreshAuth
		account, err := authSession.RefreshAuth(true)

		// Assert no error and correct values
		assert.NoError(t, err)
		assert.Equal(t, "validAccessToken", authSession.AccessToken)
		assert.Equal(t, "validRefreshToken", authSession.RefreshToken)
		assert.True(t, time.Now().Before(authSession.ExpiresAt))
		assert.Equal(t, "1", account.ID)
		assert.Equal(t, "username", account.Name)
	})
}
