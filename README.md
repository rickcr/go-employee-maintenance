# Employee Maintenance API

A REST API for managing employees and departments, built with Go.

## Project Structure

```
├── cmd/            # Application entry point
├── client/         # API client (not used but could be for tests, etc)
├── models/         # Data models (Employee, Department)
├── server/         # HTTP handlers and routing
└── services/       # Business logic
```

## Running the Server

```bash
go run cmd/main.go
```

The server starts on `http://localhost:8080`.

## API Documentation

Swagger UI is available at `http://localhost:8080/swagger` when the server is running.

## API Endpoints

### Departments

| Method | Endpoint           | Description            |
|--------|-------------------|------------------------|
| GET    | /departments      | Get all departments    |
| POST   | /departments      | Create a department    |
| GET    | /departments/{id} | Get department by ID   |
| PUT    | /departments/{id} | Update a department    |
| DELETE | /departments/{id} | Delete a department    |

### Employees

| Method | Endpoint         | Description          |
|--------|-----------------|----------------------|
| GET    | /employees      | Get all employees    |
| POST   | /employees      | Create an employee   |
| GET    | /employees/{id} | Get employee by ID   |
| PUT    | /employees/{id} | Update an employee   |
| DELETE | /employees/{id} | Delete an employee   |

## Running Tests

Tests are located in the `services/` directory alongside the service implementations.

```bash
go test ./services/...
```
