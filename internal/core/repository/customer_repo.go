package repository

import (
	"database/sql"
	"github.com/Domains18/food-delivery/internal/core/adapters"
	"github.com/Domains18/food-delivery/internal/core/models"
)

type CustomerRepo struct {
	db *sql.DB
}

func (c CustomerRepo) AddCustomer(customer models.Customer) error {
	stmt, err := c.db.Prepare(createCustomerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, customer.Code, customer.Phone)
	if err != nil {
		return err
	}

	return nil
}

// GetCustomerByID queries the database for a customer by its ID and returns a customer object

func (c CustomerRepo) GetCustomerByID(customerID int) (models.Customer, error) {
	stmt, err := c.db.Prepare(getCustomerByIDQuery)
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

func NewCustomerRepo(db *sql.DB) adapters.CustomerRepository {
	return &CustomerRepo{db: db}
}
