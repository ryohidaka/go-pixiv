package crawler

import (
	"github.com/ryohidaka/go-pixiv/pkg/appapi"
)

// PixivCrawler is a wrapper around the AppPixivAPI client.
// It provides methods to interact with Pixiv using the authenticated app client.
type PixivCrawler struct {
	app *appapi.AppPixivAPI
}

// NewCrawler initializes and returns a new PixivCrawler instance using the provided refresh token.
func NewCrawler(refreshToken string) (*PixivCrawler, error) {
	app, err := appapi.NewApp(refreshToken)
	if err != nil {
		return nil, err
	}

	return &PixivCrawler{app: app}, nil
}
