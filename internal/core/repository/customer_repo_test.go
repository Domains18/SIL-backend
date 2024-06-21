package repository_test

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/Domains18/food-delivery/internal/core/models"
	"github.com/Domains18/food-delivery/internal/core/repository"
)

func TestCustomerRepository_AddCustomerAndGetCustomerByID(t *testing.T) {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("../../../testdata", "init-db.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	}()

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	// Create a new database connection
	db, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)

	customerRepo := repository.NewCustomerRepo(db)

	// Test AddCustomer method
	testCustomer := models.Customer{
		Name:  "John Doe",
		Code:  "JD001",
		Phone: "123456789",
	}
	err = customerRepo.AddCustomer(testCustomer)
	assert.NoError(t, err)

	// Test GetCustomerByID method
	customerID := 1
	customer, err := customerRepo.GetCustomerByID(customerID)
	assert.NoError(t, err)

	expectedCustomer := models.Customer{
		ID:    customerID,
		Name:  "John Doe",
		Phone: "123456789",
	}

	assert.Equal(t, expectedCustomer, customer)
}
