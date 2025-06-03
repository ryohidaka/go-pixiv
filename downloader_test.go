package pixiv_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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

func ExampleDownloader_DownloadFile() {
	d := pixiv.NewDownloader()
	defer d.Close()

	url := "https://i.pximg.net/c/600x1200_90/img-master/img/2025/05/01/11/19/11/129899459_p0_master1200.jpg"

	opts := &pixiv.DownloadFileOptions{
		Dir:     "test_downloads",
		Name:    "sample.jpg",
		Replace: true,
	}

	ok, err := d.DownloadFile(url, opts)
	if err != nil {
		fmt.Println("Download failed:", err)
		return
	}
	if ok {
		fmt.Println("Downloaded successfully")
	} else {
		fmt.Println("Download skipped (already exists)")
	}
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

func TestDownloadFile(t *testing.T) {
	const testData = "pixiv image data"

	// Set up a test HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testData))
	}))
	defer ts.Close()

	d := pixiv.NewDownloader(context.Background())
	defer d.Close()

	tmpDir := t.TempDir()
	fileName := "testfile.txt"
	filePath := filepath.Join(tmpDir, fileName)

	// First download (should succeed)
	ok, err := d.DownloadFile(ts.URL, &pixiv.DownloadFileOptions{
		Dir:  tmpDir,
		Name: fileName,
	})
	if err != nil {
		t.Fatalf("DownloadFile failed: %v", err)
	}
	if !ok {
		t.Errorf("expected download to occur, got ok=false")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read downloaded file: %v", err)
	}
	if string(content) != testData {
		t.Errorf("downloaded content = %q, want %q", string(content), testData)
	}

	// Second download (should be skipped due to existing file and Replace=false)
	ok, err = d.DownloadFile(ts.URL, &pixiv.DownloadFileOptions{
		Dir:  tmpDir,
		Name: fileName,
	})
	if err != nil {
		t.Fatalf("DownloadFile second call failed: %v", err)
	}
	if ok {
		t.Errorf("expected second call to skip download, got ok=true")
	}

	// Third download (overwrite enabled)
	ok, err = d.DownloadFile(ts.URL, &pixiv.DownloadFileOptions{
		Dir:     tmpDir,
		Name:    fileName,
		Replace: true,
	})
	if err != nil {
		t.Fatalf("DownloadFile overwrite failed: %v", err)
	}
	if !ok {
		t.Errorf("expected overwrite download to occur, got ok=false")
	}
}
