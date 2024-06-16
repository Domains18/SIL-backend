package pkg

import (
    "database/sql"
    "fmt"
    "os"
)

var DB *sql.DB

func InitializeDatabase() {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = sql.Open("postgres", dbConnectionString)
    if err != nil {
        panic(fmt.Errorf("failed to connect to database: %v", err))
    }
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    err = DB.Ping()
    if err != nil {
        panic(fmt.Errorf("failed to ping database: %v", err))
    }

    err = CreateTables()
    if err != nil {
        panic(fmt.Errorf("failed to create tables: %v", err))
    }
}

func CreateTables() error {
    createCustomersTable := `
        CREATE TABLE IF NOT EXISTS customers (
            id SERIAL PRIMARY KEY,
            name VARCHAR(10) NOT NULL,
            code VARCHAR(10) UNIQUE NOT NULL,
            phone VARCHAR(10) NOT NULL
        )`
    _, err := DB.Exec(createCustomersTable)
    if err != nil {
        return fmt.Errorf("failed to create customers table: %v", err)
    }

    createOrdersTable := `
        CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            customer_id INTEGER REFERENCES customers(id) ON DELETE CASCADE,
            item VARCHAR(10) NOT NULL,
            amount DECIMAL(10, 2) NOT NULL,
            time TIMESTAMP NOT NULL
        )`
    _, err = DB.Exec(createOrdersTable)
    if err != nil {
        return fmt.Errorf("failed to create orders table: %v", err)
    }

    return nil
}