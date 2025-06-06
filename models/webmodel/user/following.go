package user

import (
	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
	"github.com/ryohidaka/go-pixiv/models/webmodel/illust"
)

type UserFolowingResponse struct {
	core.WebAPIResponse

	Body UserFollowing `json:"body"`
}

type UserFollowing struct {
	Users          []FollowingUser `json:"users"`
	Total          uint32          `json:"total"`
	FollowUserTags []any           `json:"followUserTags"`
}

type FollowingUser struct {
	UserID               string          `json:"userId"`
	UserName             string          `json:"userName"`
	ProfileImageURL      string          `json:"profileImageUrl"`
	ProfileImageSmallURL string          `json:"profileImageSmallUrl"`
	UserComment          *string         `json:"userComment"`
	Premium              bool            `json:"premium"`
	Following            bool            `json:"following"`
	Followed             bool            `json:"followed"`
	IsBlocking           bool            `json:"isBlocking"`
	IsMypixiv            bool            `json:"isMypixiv"`
	Illusts              []illust.Illust `json:"illusts"`
	Commission           any             `json:"commission"`
}
