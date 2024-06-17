package routes

import (
	"fmt"
	"net/http"

	"github.com/Domains18/SIL-backend/internal/core/adapters"
	"github.com/Domains18/SIL-backend/internal/routes/handlers"
	"github.com/Domains18/SIL-backend/pkg/auth"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// RegisterRoutes --->registers all routes for the application
// RegisterRoutes --->registers all routes for the application
func RegisterRoutes(mux *http.ServeMux, orderRepo adapters.Order, customerRepo adapters.Customer, auth *authenticator.Authenticator) {
	handler := handlers.NewHandler(orderRepo, customerRepo)

	store := sessions.NewCookieStore([]byte("mike"))
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Savannah OrderManagement API!")
	}), "Index"))

	mux.Handle("/login", otelhttp.NewHandler(handlers.LoginHandler(*auth, store), "Login"))
	mux.Handle("/logout", otelhttp.NewHandler(http.HandlerFunc(handlers.LogoutHandler), "Logout"))
	mux.Handle("/callback", otelhttp.NewHandler(http.HandlerFunc(handlers.CallBackHandler(*auth, store)), "Callback"))
	mux.Handle("/user", otelhttp.NewHandler(http.HandlerFunc(handlers.UserHandler), "User"))

	//mux.HandleFunc("/user", middleware.IsAuthenticated(user.Handler))
	mux.Handle("/signup", otelhttp.NewHandler(http.HandlerFunc(handler.AddCustomerHandler), "AddCustomer"))
	mux.Handle("/order", otelhttp.NewHandler(http.HandlerFunc(handler.AddOrderHandler), "AddOrder"))

}