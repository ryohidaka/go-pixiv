package pixiv_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryohidaka/go-pixiv"
)

func TestDownloadBytes(t *testing.T) {
	// Set up a test HTTP server that returns fixed data
	const testData = "hello pixiv"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check Referer header present
		if r.Header.Get("Referer") == "" {
			t.Errorf("missing Referer header")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testData))
	}))
	defer ts.Close()

	d := pixiv.NewDownloader(context.Background())
	defer d.Close()

	data, err := d.DownloadBytes(ts.URL)
	if err != nil {
		t.Fatalf("DownloadBytes failed: %v", err)
	}
	if string(data) != testData {
		t.Errorf("DownloadBytes returned %q, want %q", string(data), testData)
	}
}
