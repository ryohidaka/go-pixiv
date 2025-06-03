package pixiv

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
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
func NewDownloader(ctxs ...context.Context) *Downloader {
	var ctx context.Context
	if len(ctxs) > 0 && ctxs[0] != nil {
		ctx = ctxs[0]
	} else {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)

	return &Downloader{
		client:  http.DefaultClient,
		timeout: 10 * time.Second,
		ctx:     ctx,
		cancel:  cancel,
		referer: AppHosts,
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
