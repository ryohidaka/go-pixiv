package illust

import "github.com/ryohidaka/go-pixiv/models/webmodel/core"

type Illust struct {
	ID                      string                       `json:"id"`
	Title                   string                       `json:"title"`
	IllustType              int                          `json:"illustType"`
	XRestrict               int                          `json:"xRestrict"`
	Restrict                int                          `json:"restrict"`
	SL                      int                          `json:"sl"`
	URL                     string                       `json:"url"`
	Description             string                       `json:"description"`
	Tags                    []string                     `json:"tags"`
	UserID                  string                       `json:"userId"`
	UserName                string                       `json:"userName"`
	Width                   int                          `json:"width"`
	Height                  int                          `json:"height"`
	PageCount               int                          `json:"pageCount"`
	IsBookmarkable          bool                         `json:"isBookmarkable"`
	BookmarkData            interface{}                  `json:"bookmarkData"`
	Alt                     string                       `json:"alt"`
	TitleCaptionTranslation core.TitleCaptionTranslation `json:"titleCaptionTranslation"`
	CreateDate              string                       `json:"createDate"`
	UpdateDate              string                       `json:"updateDate"`
	IsUnlisted              bool                         `json:"isUnlisted"`
	IsMasked                bool                         `json:"isMasked"`
	AIType                  int                          `json:"aiType"`
	VisibilityScope         int                          `json:"visibilityScope"`
	ProfileImageUrl         string                       `json:"profileImageUrl"`
}
