package appmodel

import "time"

type IllustsResponse struct {
	Illusts []Illust `json:"illusts"`
	NextURL string   `json:"next_url"`
}

type IllustResponse struct {
	Illust Illust `json:"illust"`
}

type Illust struct {
	ID                   uint64          `json:"id"`
	Title                string          `json:"title"`
	Type                 IllustType      `json:"type"`
	ImageURLs            *Images         `json:"image_urls"`
	Caption              *string         `json:"caption"`
	Restrict             int             `json:"restrict"`
	User                 *User           `json:"user"`
	Tags                 []Tag           `json:"tags"`
	Tools                []string        `json:"tools"`
	CreateDate           time.Time       `json:"create_date"`
	PageCount            int             `json:"page_count"`
	Width                int             `json:"width"`
	Height               int             `json:"height"`
	SanityLevel          int             `json:"sanity_level"`
	XRestrict            int             `json:"x_restrict"`
	Series               *Series         `json:"series"`
	MetaSinglePage       *MetaSinglePage `json:"meta_single_page"`
	MetaPages            []MetaPage      `json:"meta_pages"`
	TotalView            int             `json:"total_view"`
	TotalBookmarks       int             `json:"total_bookmarks"`
	IsBookmarked         bool            `json:"is_bookmarked"`
	Visible              bool            `json:"visible"`
	IsMuted              bool            `json:"is_muted"`
	TotalComments        int             `json:"total_comments"`
	IllustAIType         IllustAIType    `json:"illust_ai_type"`
	IllustBookStyle      int             `json:"illust_book_style"`
	Request              *Request        `json:"request"`
	CommentAccessControl *int            `json:"comment_access_control"` // `CommentAccessControl` is used only in IllustDetail
}

type Images struct {
	SquareMedium string `json:"square_medium"`
	Medium       string `json:"medium"`
	Large        string `json:"large"`
	Original     string `json:"original"`
}

type Tag struct {
	Name           string  `json:"name"`
	TranslatedName *string `json:"translated_name"`
}

type Series struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type MetaSinglePage struct {
	OriginalImageURL string `json:"original_image_url"`
}

type MetaPage struct {
	Images Images `json:"image_urls"`
}

type IllustAIType int

const (
	IllustAITypeNone IllustAIType = iota
	IllustAITypeOriginal
	IllustAITypeAIGenerated
)

type IllustType string

const (
	IllustTypeIllust IllustType = "illust"
	IllustTypeManga  IllustType = "manga"
	IllustTypeUgoira IllustType = "ugoira"
)
