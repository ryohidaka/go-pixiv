package illust

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ryohidaka/go-pixiv/models/webmodel/core"
)

type Illust struct {
	ID                      IllustID                     `json:"id"`
	Title                   string                       `json:"title"`
	IllustType              int                          `json:"illustType"`
	XRestrict               int                          `json:"xRestrict"`
	Restrict                int                          `json:"restrict"`
	SL                      int                          `json:"sl"`
	URL                     string                       `json:"url"`
	Description             string                       `json:"description"`
	Tags                    []string                     `json:"tags"`
	UserID                  UserID                       `json:"userId"`
	UserName                string                       `json:"userName"`
	Width                   int                          `json:"width"`
	Height                  int                          `json:"height"`
	PageCount               int                          `json:"pageCount"`
	IsBookmarkable          bool                         `json:"isBookmarkable"`
	BookmarkData            *core.BookmarkData           `json:"bookmarkData"`
	Alt                     string                       `json:"alt"`
	TitleCaptionTranslation core.TitleCaptionTranslation `json:"titleCaptionTranslation"`
	CreateDate              string                       `json:"createDate"`
	UpdateDate              string                       `json:"updateDate"`
	IsUnlisted              bool                         `json:"isUnlisted"`
	IsMasked                bool                         `json:"isMasked"`
	AIType                  int                          `json:"aiType"`
	VisibilityScope         int                          `json:"visibilityScope"`
	ProfileImageUrl         string                       `json:"profileImageUrl"`
}

// IllustID allows unmarshaling JSON numbers or strings into a Go string.
// NOTE: Deleted illustration IDs were found to have a pattern of numeric type, so convert them to strings.
type IllustID string

func (i *IllustID) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch val := v.(type) {
	case float64:
		*i = IllustID(strconv.FormatInt(int64(val), 10))
	case string:
		*i = IllustID(val)
	default:
		return fmt.Errorf("unexpected type for StringID: %T", val)
	}
	return nil
}

// UserID allows unmarshaling JSON numbers or strings into a Go string.
// NOTE: Deleted user IDs were found to have a pattern of numeric type, so convert them to strings.
type UserID string

func (i *UserID) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch val := v.(type) {
	case float64:
		*i = UserID(strconv.FormatInt(int64(val), 10))
	case string:
		*i = UserID(val)
	default:
		return fmt.Errorf("unexpected type for StringID: %T", val)
	}
	return nil
}
