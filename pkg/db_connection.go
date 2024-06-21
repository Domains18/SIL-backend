package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(4)

	CreateTables()
}
func CreateTables() error {
	createCustomersTable := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		code VARCHAR(10) UNIQUE NOT NULL,
		phone VARCHAR(15) NOT NULL
	)
	`

	_, err := DB.Exec(createCustomersTable)
	if err != nil {
		return fmt.Errorf("could not create customers table: %v", err)
	}

	createOrdersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_id INTEGER REFERENCES customers(id) ON DELETE CASCADE,
		item VARCHAR(50) NOT NULL,
		amount DECIMAL(10, 2) NOT NULL,
		time TIMESTAMP NOT NULL
	)
	`
	_, err = DB.Exec(createOrdersTable)
	if err != nil {
		return fmt.Errorf("could not create orders table: %v", err)
	}

	return nil
}
