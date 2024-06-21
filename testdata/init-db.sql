-- Create customers table
CREATE TABLE IF NOT EXISTS customers (
                                         id SERIAL PRIMARY KEY,
                                         name VARCHAR(50) NOT NULL,
    code VARCHAR(10) UNIQUE NOT NULL,
    phone VARCHAR(15) NOT NULL
    );

-- Insert sample data
INSERT INTO customers (name, code, phone) VALUES
                                              ('John Doe', 'JR001', '123456789'),
                                              ('Jane Smith', 'JP002', '987654321');

-- Create orders table
CREATE TABLE IF NOT EXISTS orders (
                                      id SERIAL PRIMARY KEY,
                                      customer_id INTEGER REFERENCES customers(id) NOT NULL,
    item VARCHAR(100) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );