package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Domains18/SIL-backend/internal/core/adapters"
	"github.com/Domains18/SIL-backend/internal/core/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) CreateOrder(order models.Order) (int, error) {
	args := m.Called(order)
	return args.Int(0), args.Error(1)
}

func (m *MockOrderRepo) GetCustomerByID(customerID int) (models.Customer, error) {
	args := m.Called(customerID)
	return args.Get(0).(models.Customer), args.Error(1)
}

type MockCustomerRepo struct {
	mock.Mock
}

func (m *MockCustomerRepo) CreateCustomer(customer models.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepo) GetCustomerByID(id int) (models.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Customer), args.Error(1)
}

// TestHandler_AddCustomerHandler tests the AddCustomerHandler function
func TestHandler_AddCustomerHandler(t *testing.T) {
	orderRepo := new(MockOrderRepo)
	customerRepo := new(MockCustomerRepo)
	handler := Handler{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}

	//  sample customer
	customer := models.Customer{
		ID:    1,
		Name:  "John",
		Phone: "+254712345678",
	}

	body, err := json.Marshal(customer)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/customers", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	customerRepo.On("AddCustomer", customer).Return(nil)

	handler.AddCustomerHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "Customer added successfully", rr.Body.String())
	customerRepo.AssertCalled(t, "AddCustomer", customer)
}

// TestHandler_AddCustomerHandler_BadMethod tests the AddCustomerHandler function with a bad method
func TestHandler_AddCustomerHandler_BadMethod(t *testing.T) {
	orderRepo := new(MockOrderRepo)
	customerRepo := new(MockCustomerRepo)
	handler := Handler{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}

	// Create a test request with a GET method
	req, err := http.NewRequest(http.MethodGet, "/customers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.AddCustomerHandler(rr, req)

	// Assert that the status code is 405 Method Not Allowed
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)

	assert.Contains(t, rr.Body.String(), "method_not_allowed")
	assert.Contains(t, rr.Body.String(), "Method not allowed")
}

// TestHandler_AddCustomerHandler_BadRequest tests the AddCustomerHandler function with a bad request
func TestHandler_AddCustomerHandler_BadRequest(t *testing.T) {
	orderRepo := new(MockOrderRepo)
	customerRepo := new(MockCustomerRepo)
	handler := Handler{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}

	req, err := http.NewRequest(http.MethodPost, "/customers", bytes.NewReader([]byte("invalid")))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.AddCustomerHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "bad_request")
	assert.Contains(t, rr.Body.String(), "Failed to parse request body")
}

func TestNewHandler(t *testing.T) {
	type args struct {
		orderRepo    adapters.Order
		customerRepo adapters.Customer
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.orderRepo, tt.args.customerRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeErrorResponse(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		status  int
		code    string
		message string
		details string
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeErrorResponse(tt.args.w, tt.args.status, tt.args.code, tt.args.message, tt.args.details)
		})
	}
}