package pixiv_test

import (
	"go-pixiv"
	"go-pixiv/models"
	"go-pixiv/testutil"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// helper function to create AuthSession
func createAuthSession(refreshToken string) *pixiv.AuthSession {
	return &pixiv.AuthSession{
		BaseURL:      pixiv.AuthHosts,
		RefreshToken: refreshToken,
	}
}

// TestAuthenticate tests the Authenticate method in AuthSession.
func TestAuthenticate(t *testing.T) {
	// Initialize HTTP mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the Pixiv API response for authentication.
	err := testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth.json")
	if err != nil {
		t.Fatalf("Failed to mock response: %v", err)
	}

	// Prepare AuthSession and AuthParams
	authSession := createAuthSession("validRefreshToken")
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
}

// TestRefreshAuth tests the RefreshAuth method in AuthSession.
func TestRefreshAuth(t *testing.T) {
	// Initialize HTTP mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the Pixiv API response for refreshing the token.
	err := testutil.MockResponseFromFile("POST", pixiv.AuthHosts+"auth/token", "auth.json")
	if err != nil {
		t.Fatalf("Failed to mock response: %v", err)
	}

	// Prepare AuthSession and set an initial refresh token
	authSession := createAuthSession("validRefreshToken")
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
}
