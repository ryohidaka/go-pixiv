package pixiv

import "github.com/ryohidaka/go-pixiv/pkg/appapi/crawler"

// NewCrawler initializes and returns a new PixivCrawler instance using the provided refresh token.
func NewCrawler(refreshToken string) (*crawler.PixivCrawler, error) {
	return crawler.NewCrawler(refreshToken)
}
