package bookmark

import (
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/models/webmodel/illust"
)

type BookmarkedIllustsRespponse struct {
	core.WebAPIResponse

	Body BookmarkedIllusts `json:"Body"`
}

type BookmarkedIllusts struct {
	Works        []illust.Illust     `json:"works"`
	Total        uint32              `json:"total"`
	BookmarkTags map[string][]string `json:"bookmarkTags"`
}
