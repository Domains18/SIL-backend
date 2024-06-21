package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Domains18/food-delivery/internal/core/adapters"
	"github.com/Domains18/food-delivery/internal/core/models"
	Usecase "github.com/Domains18/food-delivery/internal/core/usecase"
	"go.opentelemetry.io/otel"
)

type Handler struct {
	orderRepo    adapters.OrderRepository
	customerRepo adapters.CustomerRepository
}

func (h *Handler) AddCustomerHandler(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer("customer-service").Start(r.Context(), "AddCustomerHandler")
	defer span.End()
	if r.Method != http.MethodPost {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "method_not_allowed", "Method not allowed", "")
		return
	}

	// Parse the request body
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "bad_request", "Failed to parse request body", err.Error())
		span.RecordError(err)
		return
	}

	// Add the customer to the database
	err = h.customerRepo.AddCustomer(customer)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "internal_error", "Failed to add customer to the database", err.Error())
		span.RecordError(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Customer added successfully"))
}

// AddOrderHandler handles adds an order to the database

func (h *Handler) AddOrderHandler(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer("order-service").Start(r.Context(), "AddOrderHandler")
	defer span.End()
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		span.RecordError(errors.New("Method not allowed"))

		return
	}

	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		span.RecordError(err)

		return
	}

	if order.Time == "" {
		order.Time = time.Now().UTC().Format(time.RFC3339)
	}

	_, err = h.orderRepo.CreateOrder(order)
	if err != nil {
		http.Error(w, "Failed to add order to the database", http.StatusInternalServerError)
		return
	}

	// Send SMS alert to the customer
	err = Usecase.SendSmsAlert(order)
	if err != nil {
		log.Println("Failed to send SMS alert:", err)
		http.Error(w, "Failed to send SMS alert", http.StatusInternalServerError)
		span.RecordError(err)
		return
	}

	response := struct {
		Message string `json:"message"`
		ID      int    `json:"order_id"`
	}{
		Message: "Order added successfully",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		span.RecordError(err)

		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}

// a custom error response structure
type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Write the error response to the response
func writeErrorResponse(w http.ResponseWriter, status int, code string, message string, details string) {
	errorResponse := ErrorResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Details: details,
	}

	// Set the content type and status code of the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Encode and write the error response to the response
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		// If encoding fails, use the http.Error function as a fallback
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func NewHandler(orderRepo adapters.OrderRepository, customerRepo adapters.CustomerRepository) *Handler {
	return &Handler{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}
}
