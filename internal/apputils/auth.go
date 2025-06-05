package apputils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
)

// GenClientHash returns an MD5 hash generated from the given clientTime and a predefined secret.
// This is typically used for authentication or request validation.
func GenClientHash(clientTime string, clientHashSecret string) string {
	h := md5.New()
	io.WriteString(h, clientTime)
	io.WriteString(h, clientHashSecret)
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

// GetExpiresAt calculates and returns the time when the access token will expire.
func GetExpiresAt(expiredIn int) time.Time {
	expiresAt := time.Now().Add(time.Duration(expiredIn) * time.Second)

	return expiresAt
}
