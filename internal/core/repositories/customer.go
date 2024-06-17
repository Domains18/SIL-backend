package repositories

import (
	"database/sql"

	"github.com/Domains18/SIL-backend/internal/core/adapters"
	"github.com/Domains18/SIL-backend/internal/core/models"
)


type CustomerRepo struct {
	db *sql.DB
}


func (c CustomerRepo) CreateCustomer(customer models.Customer) error {
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

func (c CustomerRepo) GetCustomerByID(customerID int) (*models.Customer, error) {
	stmt, err := c.db.Prepare(getCustomerByIDQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var customer models.Customer
	err = stmt.QueryRow(customerID).Scan(&customer.ID, &customer.Name, &customer.Phone)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}


func (c CustomerRepo) UpdateCustomer(customer models.Customer) error {
	stmt, err := c.db.Prepare(updateCustomerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(customer.Name, customer.Code, customer.Phone, customer.ID)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepo(db *sql.DB) adapters.Customer{
	return &CustomerRepo{db: db}
}