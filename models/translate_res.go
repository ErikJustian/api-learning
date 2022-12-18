package models

type TranslateRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Val     string `json:"val"`
}
