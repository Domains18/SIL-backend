package routes

import (
	"fmt"
	"net/http"

	"github.com/Domains18/food-delivery/internal/core/adapters"
	"github.com/Domains18/food-delivery/internal/routes/handlers"
	"github.com/Domains18/food-delivery/pkg/authenticator"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// RegisterRoutes --->registers all routes for the application
func RegisterRoutes(mux *http.ServeMux, orderRepo adapters.OrderRepository, customerRepo adapters.CustomerRepository, auth *authenticator.Authenticator) {
	handler := handlers.NewHandler(orderRepo, customerRepo)

	store := sessions.NewCookieStore([]byte("John"))
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the FoodDelivery API!")
	}), "Index"))

	mux.Handle("/login", otelhttp.NewHandler(handlers.LoginHandler(auth, store), "Login"))
	mux.Handle("/logout", otelhttp.NewHandler(http.HandlerFunc(handlers.LogoutHandler), "Logout"))
	mux.Handle("/callback", otelhttp.NewHandler(http.HandlerFunc(handlers.CallBackHandler(auth, store)), "Callback"))
	mux.Handle("/user", otelhttp.NewHandler(http.HandlerFunc(handlers.UserHandler), "User"))

	//mux.HandleFunc("/user", middleware.IsAuthenticated(user.Handler))
	mux.Handle("/signup", otelhttp.NewHandler(http.HandlerFunc(handler.AddCustomerHandler), "AddCustomer"))
	mux.Handle("/order", otelhttp.NewHandler(http.HandlerFunc(handler.AddOrderHandler), "AddOrder"))

}
