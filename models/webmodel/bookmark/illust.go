package bookmark

import (
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/models/webmodel/illust"
)

type BookmarkedIllustsResponse struct {
	core.WebAPIResponse

	Body BookmarkedIllusts `json:"body"`
}

type BookmarkedIllusts struct {
	Works []illust.Illust `json:"works"`
	Total uint32          `json:"total"`
	// NOTE: No response due to indefinite structure
	// BookmarkTags []interface{}  `json:"bookmarkTags"`
}
