package models

type User struct {
	ID            uint64      `json:"id"`
	Name          string      `json:"name"`
	Account       string      `json:"account"`
	Comment       string      `json:"comment"`
	IsFollowed    bool        `json:"is_followed"`
	ProfileImages *UserImages `json:"profile_image_urls"`
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
