package crawler

import (
	"log/slog"

	"github.com/ryohidaka/go-pixiv"
)

// PixivCrawler is a wrapper around the AppPixivAPI client.
// It provides methods to interact with Pixiv using the authenticated app client.
type PixivCrawler struct {
	app *pixiv.AppPixivAPI
}

// NewCrawler initializes and returns a new PixivCrawler instance using the provided refresh token.
//
// Parameters:
//   - refreshToken: A string containing the Pixiv OAuth2 refresh token.
//
// Returns:
//   - *PixivCrawler: A pointer to the newly created PixivCrawler instance.
//   - error: An error object if initialization fails; otherwise nil.
func NewCrawler(refreshToken string) (*PixivCrawler, error) {
	slog.Debug("Initializing PixivCrawler", slog.String("refresh_token", refreshToken))

	app, err := pixiv.NewApp(refreshToken)
	if err != nil {
		return nil, err
	}

	return &PixivCrawler{app: app}, nil
}
