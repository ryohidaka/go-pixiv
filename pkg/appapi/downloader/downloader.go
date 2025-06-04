package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ryohidaka/go-pixiv/pkg/appapi"
)

// Downloader handles downloading files and bytes with context and timeout.
type Downloader struct {
	client  *http.Client
	timeout time.Duration
	ctx     context.Context
	cancel  context.CancelFunc
	referer string
}

// NewDownloader creates a new Downloader instance.
// If no context is provided, context.Background() is used.
func NewDownloader(ctx context.Context) *Downloader {
	ctx, cancel := context.WithCancel(ctx)

	return &Downloader{
		client:  http.DefaultClient,
		timeout: 10 * time.Second,
		ctx:     ctx,
		cancel:  cancel,
		referer: appapi.AppHosts,
	}
}

// Close cancels the internal context to abort operations.
func (d *Downloader) Close() {
	d.cancel()
}

// DownloadBytes downloads the content from the specified URL and returns the data bytes.
func (d *Downloader) DownloadBytes(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(d.ctx, d.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Referer", d.referer)

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download failed with status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// DownloadFileOptions holds optional parameters for DownloadFile.
type DownloadFileOptions struct {
	Dir     string // Directory to save the file (optional)
	Name    string // Filename to save as (optional, default: base name from URL)
	Replace bool   // Whether to overwrite existing file (default false)
}

// DownloadFile downloads the file from the given URL and saves it according to opts.
// It returns true if the file was successfully downloaded and saved, or false if skipped.
// If opts is nil, default options are applied (no directory, default filename, no overwrite).
func (d *Downloader) DownloadFile(url string, opts *DownloadFileOptions) (bool, error) {
	dir := ""
	name := ""
	replace := false

	if opts != nil {
		dir = opts.Dir
		name = opts.Name
		replace = opts.Replace
	}

	if name == "" {
		name = filepath.Base(url)
	}
	fullPath := filepath.Join(dir, name)

	if !replace {
		if _, err := os.Stat(fullPath); err == nil {
			// File exists and replace is false; skip download
			return false, nil
		}
	}

	data, err := d.DownloadBytes(url)
	if err != nil {
		return false, err
	}

	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return false, fmt.Errorf("failed to create directory: %w", err)
		}
	}

	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return false, fmt.Errorf("failed to save file: %w", err)
	}

	return true, nil
}
