package model

type Vehicle struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}
