package adapters

import "github.com/Domains18/SIL-backend/internal/core/models"


type Order interface {
	CreateOrder(order models.Order) (int, error)
	GetCustomerByID(customerID int) (models.Customer, error)
}

type Customer interface {
	GetCustomerByID(customerID int) (models.Customer, error)
	AddCustomer(customer models.Customer) error
}

type SMS  interface {
	SendSMS(message string, phone string) error
}