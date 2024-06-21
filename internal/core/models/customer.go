package models

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
