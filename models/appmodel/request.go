package appmodel

type Request struct {
	RequestInfo  RequestInfo `json:"request_info"`
	RequestUsers []User      `json:"request_users"`
}

type RequestInfo struct {
	FanUserID         uint64            `json:"fan_user_id"`
	CollaborateStatus CollaborateStatus `json:"collaborate_status"`
	Role              string            `json:"role"`
}

type CollaborateStatus struct {
	Collaborating            bool     `json:"collaborating"`
	CollaborateAnonymousFlag bool     `json:"collaborate_anonymous_flag"`
	CollaborateUserSamples   []string `json:"collaborate_user_samples"`
}
