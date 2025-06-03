package pixiv

import (
	"context"
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
