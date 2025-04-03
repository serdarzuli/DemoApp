# Library Management System

A RESTful API for managing books in a library, built with Go and Gin framework.

## Features

- CRUD operations for books
- RESTful API endpoints
- Swagger documentation
- Debug mode support
- Clean architecture following SOLID principles

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | /api/books | Create a new book |
| GET    | /api/books | Get all books |
| GET    | /api/books/{id} | Get a specific book |
| PUT    | /api/books/{id} | Update a book |
| DELETE | /api/books/{id} | Delete a book |

## Prerequisites

- Go 1.16 or higher
- Swag CLI

## Installation

1. Clone the repository:
```bash
git clone https://github.com/serdarzuli/golang-library-management
cd golang-library-management
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate Swagger documentation:
```bash
swag init
```

## Running the Application

1. Start the server:
```bash
go run main.go
```

2. Access the API:
- API Base URL: `http://localhost:8080/api`
- Swagger UI: `http://localhost:8080/swagger/index.html`

## Debug Mode

To run in debug mode, uncomment the following line in `main.go`:
```go
gin.SetMode(gin.DebugMode)
```

## Project Structure

```
.
├── handlers/     # HTTP request handlers
├── models/       # Data models
├── repository/   # Data access layer
├── service/      # Business logic
├── interfaces/   # Interface definitions
└── main.go       # Application entry point
```

## Testing the API

### Using Postman

1. Create a book (POST):
```
URL: http://localhost:8080/api/books
Headers: Content-Type: application/json
Body:
{
    "id": 1,
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "isbn": "978-0134190440",
    "quantity": 5
}
```

2. Get all books (GET):
```
URL: http://localhost:8080/api/books
```

3. Get a specific book (GET):
```
URL: http://localhost:8080/api/books/1
```

4. Update a book (PUT):
```
URL: http://localhost:8080/api/books/1
Headers: Content-Type: application/json
Body:
{
    "title": "Updated Title",
    "author": "Updated Author",
    "isbn": "978-0134190440",
    "quantity": 10
}
```

5. Delete a book (DELETE):
```
URL: http://localhost:8080/api/books/1
```

## License

[MIT License](LICENSE) 