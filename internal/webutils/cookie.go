package webutils

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// NewCookieJar returns a CookieJar with PHPSESSID set for the given base URL.
func NewCookieJar(phpsessid string, baseURL *url.URL) (http.CookieJar, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	jar.SetCookies(baseURL, []*http.Cookie{
		{
			Name:  "PHPSESSID",
			Value: phpsessid,
			Path:  "/",
		},
	})
	return jar, nil
}
