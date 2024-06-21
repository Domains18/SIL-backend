package repository

import (
	"database/sql"

	"github.com/Domains18/food-delivery/internal/core/adapters"
	"github.com/Domains18/food-delivery/internal/core/models"
)

type OrderRepo struct {
	db *sql.DB
}

// CreateOrder inserts a new order into the database and returns its ID
func (o OrderRepo) CreateOrder(order models.Order) (int, error) {
	stmt, err := o.db.Prepare(createOrderQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var orderID int
	err = stmt.QueryRow(order.CustomerID, order.Item, order.Amount, order.Time).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

// GetCustomerByID queries the database for a customer by its ID and returns a customer object
func (o OrderRepo) GetCustomerByID(customerID int) (models.Customer, error) {
	stmt, err := o.db.Prepare(getCustomerByIDQuery)
	if err != nil {
		return models.Customer{}, err
	}
	defer stmt.Close()

	var customer models.Customer
	err = stmt.QueryRow(customerID).Scan(&customer.ID, &customer.Name, &customer.Phone)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func NewOrderRepo(db *sql.DB) adapters.OrderRepository {
	return &OrderRepo{db: db}
}
