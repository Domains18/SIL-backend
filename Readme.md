# Food Delivery Order Management

FoodDelivery  API is a REST API for managing orders and customers for your business. Built with Go and PostgreSQL, it features authentication, authorization, and SMS alerts.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Features

- Customer management
- Order management
- Authentication and authorization
- SMS alerts
- RESTful API

## Prerequisites

- Go (version 1.16 or higher)
- PostgreSQL (version 18 or higher)
- Docker (optional)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Domains18/SIL-backend.git
   cd SIL-backend
    ```

2. Install dependencies:
   ```bash
   go mod download
   go mod tidy
   ```

3. Setup the environment variables:
   ```bash
   cp .env.example .env
   ```

4. (optional) Build and run the Docker container:
   ```bash
   docker-compose up --build -t food-delivery
   docker run -p 8080:8080 food-delivery
   ```

5. Run the application locally:
   ```bash
    go run main.go
    ```

## Usage
- The application will be available at `http://localhost:8080`.

## Testing

1. Run the tests:
   ```bash
   go test ./...
   ```
