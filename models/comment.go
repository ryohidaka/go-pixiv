package models

type IllustComments struct {
	TotalComments uint64    `json:"total_comments"`
	Comments      []Comment `json:"comments"`
	NextURL       string    `json:"next_url"`
}

type Comment struct {
	ID             uint64   `json:"id"`
	CommentComment string   `json:"comment"`
	Date           string   `json:"date"`
	User           *User    `json:"user"`
	HasReplies     bool     `json:"has_replies"`
	ParentComment  *Comment `json:"parent_comment"`
}
