package user

import (
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/models/webmodel/illust"
	"github.com/ryohidaka/go-pixiv/models/webmodel/novel"
)

type UserWorksResponse struct {
	core.WebAPIResponse

	Body UserWorks `json:"body"`
}

type UserWorks struct {
	Illusts map[string]*illust.Illust `json:"illusts"`
	Novels  map[string]*novel.Novel   `json:"novels"`
}
