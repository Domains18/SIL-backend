package repository_test

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"
	"time"

	"github.com/Domains18/food-delivery/internal/core/models"
	"github.com/Domains18/food-delivery/internal/core/repository"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestOrderRepository_CreateOrderAndGetCustomerByID(t *testing.T) {
	ctx := context.Background()

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

	db, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)

	orderRepo := repository.NewOrderRepo(db)

	// Test CreateOrder method
	testOrder := models.Order{
		CustomerID: 1,
		Item:       "Soap",
		Amount:     100,
		Time:       "2021-12-01 12:00:00",
	}
	orderID, err := orderRepo.CreateOrder(testOrder)
	assert.NoError(t, err)
	assert.NotZero(t, orderID)

	// Test GetCustomerByID method
	customerID := 1
	customer, err := orderRepo.GetCustomerByID(customerID)
	assert.NoError(t, err)

	expectedCustomer := models.Customer{
		ID:    customerID,
		Name:  "John Doe",
		Phone: "123456789",
	}

	assert.Equal(t, expectedCustomer, customer)
}
