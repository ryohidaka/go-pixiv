package user

import "github.com/ryohidaka/go-pixiv/models/webmodel/core"

type UserResponse struct {
	core.WebAPIResponse

	Body User `json:"body"`
}

type User struct {
	UserID         string       `json:"userId"`
	Name           string       `json:"name"`
	Image          string       `json:"image"`
	ImageBig       string       `json:"imageBig"`
	Premium        bool         `json:"premium"`
	IsFollowed     bool         `json:"isFollowed"`
	IsMypixiv      bool         `json:"isMypixiv"`
	IsBlocking     bool         `json:"isBlocking"`
	Background     *Background  `json:"background"`
	SketchLiveID   any          `json:"sketchLiveId"`
	Partial        int          `json:"partial"`
	SketchLives    []any        `json:"sketchLives"`
	Commission     any          `json:"commission"`
	Following      int          `json:"following"`
	MypixivCount   int          `json:"mypixivCount"`
	FollowedBack   bool         `json:"followedBack"`
	Comment        string       `json:"comment"`
	CommentHTML    string       `json:"commentHtml"`
	Webpage        string       `json:"webpage"`
	Social         *SocialLinks `json:"social"`
	CanSendMessage bool         `json:"canSendMessage"`
	Region         Region       `json:"region"`
	Age            Info         `json:"age"`
	BirthDay       Info         `json:"birthDay"`
	Gender         Info         `json:"gender"`
	Job            Info         `json:"job"`
	Workspace      any          `json:"workspace"`
	Official       bool         `json:"official"`
	Group          []UserGroup  `json:"group"`
}

type Background struct {
	Repeat    *string `json:"repeat"`
	Color     *string `json:"color"`
	URL       string  `json:"url"`
	IsPrivate bool    `json:"isPrivate"`
}

type SocialLinks struct {
	Twitter *SocialLink `json:"twitter"`
	Pawoo   *SocialLink `json:"pawoo"`
}

type SocialLink struct {
	URL string `json:"url"`
}

type Region struct {
	Name         string  `json:"name"`
	Region       string  `json:"region"`
	Prefecture   *string `json:"prefecture"`
	PrivacyLevel string  `json:"privacyLevel"`
}

type Info struct {
	Name         *string `json:"name"`
	PrivacyLevel *string `json:"privacyLevel"`
}

type UserGroup struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	IconURL string `json:"iconUrl"`
}
