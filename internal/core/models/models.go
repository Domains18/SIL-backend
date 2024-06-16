package models

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}



type Order struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Item       string  `json:"item"`
	Amount     float64 `json:"amount"`
	Time       string  `json:"time"`
}