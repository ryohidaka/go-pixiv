package appmodel

type IllustBookmarkDetail struct {
	BookmarkDetail BookmarkDetail `json:"bookmark_detail"`
}

type BookmarkDetail struct {
	IsBookmarked bool                `json:"is_bookmarked"`
	Tags         []BookmarkDetailTag `json:"tags"`
	Restrict     string              `json:"restrict"`
}

type BookmarkDetailTag struct {
	Name         string `json:"name"`
	IsRegistered bool   `json:"is_registered"`
}
