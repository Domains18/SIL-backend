package models

type Order struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Item       string  `json:"item"`
	Amount     float64 `json:"amount"`
	Time       string  `json:"time"`
}
