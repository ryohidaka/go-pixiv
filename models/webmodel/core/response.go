package core

type WebAPIResponse struct {
	Error   bool    `json:"error"`
	Message *string `json:"message"`
	Body    interface{}
}
