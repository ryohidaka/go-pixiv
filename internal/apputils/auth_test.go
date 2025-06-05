package apputils_test

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"testing"
	"time"

	"github.com/ryohidaka/go-pixiv/internal/apputils"
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
)

// TestGenClientHash tests the genClientHash function.
func TestGenClientHash(t *testing.T) {
	clientTime := "2025-05-01T12:00:00Z"

	// Manually calculate the expected hash value using the known secret
	// ClientHashSecret is assumed to be "yourSecretValue"
	expectedSecret := appapi.ClientHashSecret // This will be imported from go-pixiv
	h := md5.New()
	io.WriteString(h, clientTime)
	io.WriteString(h, expectedSecret)
	expectedHash := hex.EncodeToString(h.Sum(nil)) // Expected MD5 hash

	// Call the function
	result := apputils.GenClientHash(clientTime, expectedSecret)

	// Check if the result matches the expected value
	if result != expectedHash {
		t.Errorf("expected %s, got %s", expectedHash, result)
	}
}

// TestGetExpiresAt tests the getExpiresAt function.
func TestGetExpiresAt(t *testing.T) {
	expiredIn := 3600 // 1 hour
	expirationTime := apputils.GetExpiresAt(expiredIn)

	// Check if the expiration time is approximately 1 hour from now
	if expirationTime.Before(time.Now().Add(time.Duration(expiredIn-5)*time.Second)) || expirationTime.After(time.Now().Add(time.Duration(expiredIn+5)*time.Second)) {
		t.Errorf("expected expiration time to be within 1 hour, got %v", expirationTime)
	}
}
