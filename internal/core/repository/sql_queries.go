package repository

const (
	createCustomerQuery  = `INSERT INTO customers (name, code, phone) VALUES ($1, $2, $3)`
	getCustomerByIDQuery = `SELECT id, name, phone FROM customers WHERE id = $1`
	updateCustomerQuery  = `UPDATE customers SET name = $1, code = $2, phone = $3 WHERE id = $4`
	deleteCustomerQuery  = `DELETE FROM customers WHERE id = $1`
	createOrderQuery     = `INSERT INTO orders (customer_id, item, amount, time) VALUES ($1, $2, $3, $4) RETURNING id`
)
