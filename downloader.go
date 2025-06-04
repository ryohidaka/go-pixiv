package pixiv

import (
	"context"

	"github.com/ryohidaka/go-pixiv/pkg/appapi/downloader"
)

// Downloader handles downloading files and bytes with context and timeout.
type Downloader = downloader.Downloader

// NewDownloader creates a new Downloader instance.
// If no context is provided, context.Background() is used.
func NewDownloader(ctxs ...context.Context) *downloader.Downloader {
	var ctx context.Context
	if len(ctxs) > 0 && ctxs[0] != nil {
		ctx = ctxs[0]
	} else {
		ctx = context.Background()
	}

	return downloader.NewDownloader(ctx)
}

// DownloadFileOptions holds optional parameters for DownloadFile.
type DownloadFileOptions = downloader.DownloadFileOptions
