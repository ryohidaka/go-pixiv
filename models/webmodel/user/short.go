package user

import "github.com/ryohidaka/go-pixiv/models/webmodel/core"

type UserShortResponse struct {
	core.WebAPIResponse

	Body UserShort `json:"body"`
}

type UserShort struct {
	UserID       string         `json:"userId"`
	Name         string         `json:"name"`
	Image        string         `json:"image"`
	ImageBig     string         `json:"imageBig"`
	Premium      bool           `json:"premium"`
	IsFollowed   bool           `json:"isFollowed"`
	IsMypixiv    bool           `json:"isMypixiv"`
	IsBlocking   bool           `json:"isBlocking"`
	Background   UserBackground `json:"background"`
	SketchLiveID interface{}    `json:"sketchLiveId"`
	Partial      int            `json:"partial"`
	SketchLives  []interface{}  `json:"sketchLives"`
	Commission   interface{}    `json:"commission"`
}

type UserBackground struct {
	Repeat    *string `json:"repeat"`
	Color     *string `json:"color"`
	URL       string  `json:"url"`
	IsPrivate bool    `json:"isPrivate"`
}
