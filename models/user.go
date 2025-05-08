package models

type User struct {
	ID            uint64      `json:"id"`
	Name          string      `json:"name"`
	Account       string      `json:"account"`
	ProfileImages *UserImages `json:"profile_image_urls"`
	Comment       *string     `json:"comment,omitempty"` // `Comment` is used only in UserDetail
	IsFollowed    bool        `json:"is_followed"`
}

type UserImages struct {
	Medium string `json:"medium"`
}

type UserFollowList struct {
	UserPreviews []UserPreview `json:"user_previews"`
	NextURL      string        `json:"next_url"`
}

type UserPreview struct {
	User    User          `json:"user"`
	Illusts []Illust      `json:"illusts"`
	Novels  []interface{} `json:"novels"`
	IsMuted bool          `json:"is_muted"`
}
