package user

import "github.com/ryohidaka/go-pixiv/models/webmodel/core"

type UserProfileResponse struct {
	core.WebAPIResponse

	Body UserProfile `json:"body"`
}

type UserProfile struct {
	Illusts                   map[string]interface{}  `json:"illusts"` // map of illustration IDs to null (interface{}でnull対応)
	Manga                     map[string]interface{}  `json:"manga"`   // map of manga IDs to null
	Novels                    map[string]interface{}  `json:"novels"`  // map of novel IDs to null
	MangaSeries               []interface{}           `json:"mangaSeries"`
	NovelSeries               []NovelSeriesItem       `json:"novelSeries"`
	Pickup                    []PickupItem            `json:"pickup"`
	BookmarkCount             BookmarkCount           `json:"bookmarkCount"`
	ExternalSiteWorksStatus   ExternalSiteWorksStatus `json:"externalSiteWorksStatus"`
	Request                   Request                 `json:"request"`
	ShouldShowSensitiveNotice bool                    `json:"shouldShowSensitiveNotice"`
}

type NovelSeriesItem struct {
	ID                            string      `json:"id"`
	UserID                        string      `json:"userId"`
	UserName                      string      `json:"userName"`
	ProfileImageURL               string      `json:"profileImageUrl"`
	XRestrict                     int         `json:"xRestrict"`
	IsOriginal                    bool        `json:"isOriginal"`
	IsConcluded                   bool        `json:"isConcluded"`
	GenreID                       string      `json:"genreId"`
	Title                         string      `json:"title"`
	Caption                       string      `json:"caption"`
	Language                      string      `json:"language"`
	Tags                          []string    `json:"tags"`
	PublishedContentCount         int         `json:"publishedContentCount"`
	PublishedTotalCharacterCount  int         `json:"publishedTotalCharacterCount"`
	PublishedTotalWordCount       int         `json:"publishedTotalWordCount"`
	PublishedReadingTime          int         `json:"publishedReadingTime"`
	UseWordCount                  bool        `json:"useWordCount"`
	LastPublishedContentTimestamp int64       `json:"lastPublishedContentTimestamp"`
	CreatedTimestamp              int64       `json:"createdTimestamp"`
	UpdatedTimestamp              int64       `json:"updatedTimestamp"`
	CreateDate                    string      `json:"createDate"`
	UpdateDate                    string      `json:"updateDate"`
	FirstNovelID                  string      `json:"firstNovelId"`
	LatestNovelID                 string      `json:"latestNovelId"`
	DisplaySeriesContentCount     int         `json:"displaySeriesContentCount"`
	ShareText                     string      `json:"shareText"`
	Total                         int         `json:"total"`
	FirstEpisode                  Episode     `json:"firstEpisode"`
	WatchCount                    *int        `json:"watchCount"`
	MaxXRestrict                  *int        `json:"maxXRestrict"`
	Cover                         Cover       `json:"cover"`
	CoverSettingData              interface{} `json:"coverSettingData"`
	IsWatched                     bool        `json:"isWatched"`
	IsNotifying                   bool        `json:"isNotifying"`
	AiType                        int         `json:"aiType"`
}

type Episode struct {
	URL string `json:"url"`
}

type Cover struct {
	URLs CoverURLs `json:"urls"`
}

type CoverURLs struct {
	Mw240       string `json:"240mw"`
	Mw480       string `json:"480mw"`
	Px1200x1200 string `json:"1200x1200"`
	Px128x128   string `json:"128x128"`
	Original    string `json:"original"`
}

type PickupItem struct {
	ID                      string                  `json:"id"`
	Title                   string                  `json:"title"`
	IllustType              int                     `json:"illustType"`
	XRestrict               int                     `json:"xRestrict"`
	Restrict                int                     `json:"restrict"`
	Sl                      int                     `json:"sl"`
	URL                     string                  `json:"url"`
	Description             string                  `json:"description"`
	Tags                    []string                `json:"tags"`
	UserID                  string                  `json:"userId"`
	UserName                string                  `json:"userName"`
	Width                   int                     `json:"width"`
	Height                  int                     `json:"height"`
	PageCount               int                     `json:"pageCount"`
	IsBookmarkable          bool                    `json:"isBookmarkable"`
	BookmarkData            interface{}             `json:"bookmarkData"`
	Alt                     string                  `json:"alt"`
	TitleCaptionTranslation TitleCaptionTranslation `json:"titleCaptionTranslation"`
	CreateDate              string                  `json:"createDate"`
	UpdateDate              string                  `json:"updateDate"`
	IsUnlisted              bool                    `json:"isUnlisted"`
	IsMasked                bool                    `json:"isMasked"`
	AiType                  int                     `json:"aiType"`
	VisibilityScope         int                     `json:"visibilityScope"`
	Type                    string                  `json:"type"`
	Deletable               bool                    `json:"deletable"`
	Draggable               bool                    `json:"draggable"`
	ContentURL              string                  `json:"contentUrl"`
}

type TitleCaptionTranslation struct {
	WorkTitle   *string `json:"workTitle"`
	WorkCaption *string `json:"workCaption"`
}

type BookmarkCount struct {
	Public  BookmarkCountType `json:"public"`
	Private BookmarkCountType `json:"private"`
}

type BookmarkCountType struct {
	Illust int `json:"illust"`
	Novel  int `json:"novel"`
}

type ExternalSiteWorksStatus struct {
	Booth    bool `json:"booth"`
	Sketch   bool `json:"sketch"`
	VroidHub bool `json:"vroidHub"`
}

type Request struct {
	ShowRequestTab     bool      `json:"showRequestTab"`
	ShowRequestSentTab bool      `json:"showRequestSentTab"`
	PostWorks          PostWorks `json:"postWorks"`
}

type PostWorks struct {
	Artworks []interface{} `json:"artworks"`
	Novels   []interface{} `json:"novels"`
}
