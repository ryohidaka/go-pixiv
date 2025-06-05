package webutils_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ryohidaka/go-pixiv/internal/webutils"
	"github.com/stretchr/testify/assert"
)

func TestWithDefaultHeaders(t *testing.T) {
	userAgent := "TestUserAgent/1.0"
	referer := "https://example.com"

	// Capture the request passed to the RoundTripper
	var capturedReq *http.Request
	baseRT := webutils.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		capturedReq = req
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("OK")),
			Header:     make(http.Header),
		}, nil
	})

	rt := webutils.WithDefaultHeaders(baseRT, userAgent, referer)

	req, err := http.NewRequest("GET", "http://dummy", nil)
	assert.NoError(t, err)

	resp, err := rt.RoundTrip(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	assert.NotNil(t, capturedReq)
	assert.Equal(t, userAgent, capturedReq.Header.Get("User-Agent"))
	assert.Equal(t, referer, capturedReq.Header.Get("Referer"))
}
