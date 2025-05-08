package models

type UserDetail struct {
	User             *User             `json:"user"`
	Profile          *Profile          `json:"profile"`
	ProfilePublicity *ProfilePublicity `json:"profile_publicity"`
	Workspace        *Workspace        `json:"workspace"`
}

type Profile struct {
	Webpage                    interface{} `json:"webpage"`
	Gender                     string      `json:"gender"`
	Birth                      string      `json:"birth"`
	BirthDay                   string      `json:"birth_day"`
	BirthYear                  uint64      `json:"birth_year"`
	Region                     string      `json:"region"`
	AddressID                  uint64      `json:"address_id"`
	CountryCode                string      `json:"country_code"`
	Job                        string      `json:"job"`
	JobID                      uint64      `json:"job_id"`
	TotalFollowUsers           uint64      `json:"total_follow_users"`
	TotalMypixivUsers          uint64      `json:"total_mypixiv_users"`
	TotalIllusts               uint64      `json:"total_illusts"`
	TotalManga                 uint64      `json:"total_manga"`
	TotalNovels                uint64      `json:"total_novels"`
	TotalIllustBookmarksPublic uint64      `json:"total_illust_bookmarks_public"`
	TotalIllustSeries          uint64      `json:"total_illust_series"`
	TotalNovelSeries           uint64      `json:"total_novel_series"`
	BackgroundImageURL         string      `json:"background_image_url"`
	TwitterAccount             string      `json:"twitter_account"`
	TwitterURL                 string      `json:"twitter_url"`
	PawooURL                   string      `json:"pawoo_url"`
	IsPremium                  bool        `json:"is_premium"`
	IsUsingCustomProfileImage  bool        `json:"is_using_custom_profile_image"`
}

type ProfilePublicity struct {
	Gender    string `json:"gender"`
	Region    string `json:"region"`
	BirthDay  string `json:"birth_day"`
	BirthYear string `json:"birth_year"`
	Job       string `json:"job"`
	Pawoo     bool   `json:"pawoo"`
}

type Workspace struct {
	Pc                string `json:"pc"`
	Monitor           string `json:"monitor"`
	Tool              string `json:"tool"`
	Scanner           string `json:"scanner"`
	Tablet            string `json:"tablet"`
	Mouse             string `json:"mouse"`
	Printer           string `json:"printer"`
	Desktop           string `json:"desktop"`
	Music             string `json:"music"`
	Desk              string `json:"desk"`
	Chair             string `json:"chair"`
	Comment           string `json:"comment"`
	WorkspaceImageURL string `json:"workspace_image_url"`
}
