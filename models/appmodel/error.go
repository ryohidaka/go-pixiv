package appmodel

type PixivError struct {
	HasError bool              `json:"has_error"`
	Errors   map[string]Perror `json:"errors"`
}

type Perror struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
