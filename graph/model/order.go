package model

type Order struct {
	ID          string `json:"id"`
	RequestDate string `json:"requestDate"`
	Status      string `json:"status"`
}
