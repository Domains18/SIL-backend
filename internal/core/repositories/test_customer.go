package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Domains18/SIL-backend/internal/core/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a new CustomerRepo instance with the mock database connection
	customerRepo := NewCustomerRepo(db)

	// Define the test customer data
	customer := models.Customer{
		Name:  "John Doe",
		Code:  "ABC123",
		Phone: "1234567890",
	}

	// Set up the expected SQL query and response
	mock.ExpectPrepare("INSERT INTO customers").ExpectExec().WithArgs(customer.Name, customer.Code, customer.Phone).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the CreateCustomer method
	err = customerRepo.CreateCustomer(customer)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCustomerByID(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a new CustomerRepo instance with the mock database connection
	customerRepo := NewCustomerRepo(db)

	// Define the test customer ID and expected customer data
	customerID := 1
	expectedCustomer := &models.Customer{
		ID:    1,
		Name:  "John Doe",
		Phone: "1234567890",
	}

	// Set up the expected SQL query and response
	rows := sqlmock.NewRows([]string{"id", "name", "phone"}).AddRow(expectedCustomer.ID, expectedCustomer.Name, expectedCustomer.Phone)
	mock.ExpectPrepare("SELECT id, name, phone FROM customers").ExpectQuery().WithArgs(customerID).WillReturnRows(rows)

	// Call the GetCustomerByID method
	customer, err := customerRepo.GetCustomerByID(customerID)

	// Assert that no error occurred and the returned customer matches the expected customer
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateCustomer(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a new CustomerRepo instance with the mock database connection
	customerRepo := NewCustomerRepo(db)

	// Define the test customer data
	customer := models.Customer{
		ID:    1,
		Name:  "John Doe",
		Code:  "ABC123",
		Phone: "1234567890",
	}

	// Set up the expected SQL query and response
	mock.ExpectPrepare("UPDATE customers").ExpectExec().WithArgs(customer.Name, customer.Code, customer.Phone, customer.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the UpdateCustomer method
	err = customerRepo.UpdateCustomer(customer)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
