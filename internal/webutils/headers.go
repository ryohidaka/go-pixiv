package webutils

import (
	"net/http"
)

// RoundTripperFunc is a helper type that allows defining RoundTripper with a function.
type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// WithDefaultHeaders wraps a RoundTripper and adds default headers.
func WithDefaultHeaders(rt http.RoundTripper, userAgent string, referer string) http.RoundTripper {
	return RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("Referer", referer)
		return rt.RoundTrip(req)
	})
}
