package adapters

import "github.com/Domains18/SIL-backend/internal/core/models"


type Order interface {
	CreateOrder(order *models.Order) error
	GetCustomerOrders(customerID int) ([]models.Order, error)
}

type Customer interface {
	GetCustomerByID(customerID int) (*models.Customer, error)
	CreateCustomer(customer *models.Customer) error
}

type SMS interface {
	SendSMS(phoneNumber, message string) error
}