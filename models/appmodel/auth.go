package appmodel

type AuthParams struct {
	GetSecureURL int    `url:"get_secure_url,omitempty"`
	ClientID     string `url:"client_id,omitempty"`
	ClientSecret string `url:"client_secret,omitempty"`
	GrantType    string `url:"grant_type,omitempty"`
	RefreshToken string `url:"refresh_token,omitempty"`
}

type AuthResponse struct {
	Response *AuthInfo `json:"response"`
}

type AuthInfo struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    int     `json:"expires_in"`
	TokenType    string  `json:"token_type"`
	Scope        string  `json:"scope"`
	RefreshToken string  `json:"refresh_token"`
	User         Account `json:"user"`
	DeviceToken  string  `json:"device_token"`
}

type Account struct {
	ProfileImage     AccountProfileImages `json:"profile_image_urls"`
	ID               string               `json:"id"`
	Name             string               `json:"name"`
	Account          string               `json:"account"`
	MailAddress      string               `json:"mail_address"`
	IsPremium        bool                 `json:"is_premium"`
	XRestrict        int                  `json:"x_restrict"`
	IsMailAuthorized bool                 `json:"is_mail_authorized"`
}

type AccountProfileImages struct {
	Px16  string `json:"px_16x16"`
	Px50  string `json:"px_50x50"`
	Px170 string `json:"px_170x170"`
}
