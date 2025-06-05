package appapi_test

import (
	"testing"
	"time"

	"github.com/ryohidaka/go-pixiv/models/appmodel"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
	"github.com/ryohidaka/go-pixiv/testutil/apptest"

	"github.com/stretchr/testify/assert"
)

// TestAuthenticate tests the Authenticate method in AuthSession.
func TestAuthenticate(t *testing.T) {
	apptest.WithMockHTTP(t, func() {
		err := apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth_token")
		if err != nil {
			t.Fatalf("Failed to mock response: %v", err)
		}

		// Prepare AuthSession and AuthParams
		authSession := apptest.CreateAuthSession("validRefreshToken")
		params := &appmodel.AuthParams{
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
	apptest.WithMockHTTP(t, func() {
		// Mock the Pixiv API response for refreshing the token.
		err := apptest.MockResponseFromFile("POST", appapi.AuthHosts+"auth/token", "auth_token")
		if err != nil {
			t.Fatalf("Failed to mock response: %v", err)
		}

		// Prepare AuthSession and set an initial refresh token
		authSession := apptest.CreateAuthSession("validRefreshToken")
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
