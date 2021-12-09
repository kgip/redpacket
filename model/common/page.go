package common

type Page struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
