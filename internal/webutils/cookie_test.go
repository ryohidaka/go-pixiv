package webutils_test

import (
	"net/url"
	"testing"

	"github.com/ryohidaka/go-pixiv/internal/webutils"
)

func TestNewCookieJar(t *testing.T) {
	baseURL, _ := url.Parse("https://example.com")
	jar, err := webutils.NewCookieJar("session123", baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cookies := jar.Cookies(baseURL)
	if len(cookies) != 1 || cookies[0].Name != "PHPSESSID" || cookies[0].Value != "session123" {
		t.Errorf("unexpected cookies: %+v", cookies)
	}
}
