package repositories

import (
	"database/sql"

	"github.com/Domains18/SIL-backend/internal/core/adapters"
	"github.com/Domains18/SIL-backend/internal/core/models"
)

type OrderRepo struct {
	db *sql.DB
}


func (o *OrderRepo) CreateOrder(order models.Order) (int, error){
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

// func (o *OrderRepo) GetOrderByID(orderID int) (*models.Order, error){
// 	stmt, err := o.db.Prepare(getOrderByIDQuery)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	order := &models.Order{}
// 	err = stmt.QueryRow(orderID).Scan(&order.ID, &order.CustomerID, &order.Item, &order.Amount, &order.Time)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return order, nil
// }

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

func NewOrderRepo(db *sql.DB) adapters.Order {
	return &OrderRepo{db: db}
}