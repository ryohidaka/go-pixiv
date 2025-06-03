package pixiv_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryohidaka/go-pixiv"
)

func ExampleDownloader_DownloadBytes() {
	d := pixiv.NewDownloader()
	defer d.Close()

	url := "https://i.pximg.net/c/600x1200_90/img-master/img/2025/05/01/11/19/11/129899459_p0_master1200.jpg"

	data, err := d.DownloadBytes(url)
	if err != nil {
		log.Printf("Download failed: %v", err)
		return
	}

	// Output only length for demo
	fmt.Printf("Downloaded %d bytes\n", len(data))
}

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
