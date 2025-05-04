package models

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
