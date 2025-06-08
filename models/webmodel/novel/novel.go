package novel

import "github.com/ryohidaka/go-pixiv/models/webmodel/core"

type Novel struct {
	ID                      string                       `json:"id"`
	Title                   string                       `json:"title"`
	Genre                   string                       `json:"genre"`
	XRestrict               int                          `json:"xRestrict"`
	Restrict                int                          `json:"restrict"`
	URL                     string                       `json:"url"`
	Tags                    []string                     `json:"tags"`
	UserID                  string                       `json:"userId"`
	UserName                string                       `json:"userName"`
	ProfileImageURL         string                       `json:"profileImageUrl"`
	TextCount               int                          `json:"textCount"`
	WordCount               int                          `json:"wordCount"`
	ReadingTime             int                          `json:"readingTime"`
	UseWordCount            bool                         `json:"useWordCount"`
	Description             string                       `json:"description"`
	IsBookmarkable          bool                         `json:"isBookmarkable"`
	BookmarkData            interface{}                  `json:"bookmarkData"`
	BookmarkCount           int                          `json:"bookmarkCount"`
	IsOriginal              bool                         `json:"isOriginal"`
	Marker                  interface{}                  `json:"marker"`
	TitleCaptionTranslation core.TitleCaptionTranslation `json:"titleCaptionTranslation"`
	CreateDate              string                       `json:"createDate"`
	UpdateDate              string                       `json:"updateDate"`
	IsMasked                bool                         `json:"isMasked"`
	AIType                  int                          `json:"aiType"`
	IsUnlisted              bool                         `json:"isUnlisted"`
	VisibilityScope         int                          `json:"visibilityScope"`
}
