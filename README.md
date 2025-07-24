# Lab02 Go API

A simple REST API built with Go and the Gorilla Mux router for learning purposes.

## Features

- RESTful API endpoints
- JSON request/response handling
- Request logging middleware
- In-memory data storage (for demo purposes)
- Health check endpoint

## Prerequisites

- Go 1.24.5 or later
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/kngan6297/lab02-go-api.git
cd lab02-go-api
```

2. Install dependencies:

```bash
go mod tidy
```

## Running the Application

Start the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### 1. Home

- **GET** `/`
- Returns a welcome message with timestamp

**Response:**

```json
{
  "message": "Welcome to Lab02 Go API",
  "timestamp": "2024-01-01T12:00:00Z",
  "status": "success"
}
```

### 2. Health Check

- **GET** `/api/health`
- Returns the health status of the API

**Response:**

```json
{
  "message": "API is healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "status": "healthy"
}
```

### 3. Get All Users

- **GET** `/api/users`
- Returns all users in the system

**Response:**

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "username": "johndoe"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com",
    "username": "janesmith"
  }
]
```

### 4. Get User by ID

- **GET** `/api/users/{id}`
- Returns a specific user by ID (currently returns first user for demo)

**Response:**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "username": "johndoe"
}
```

### 5. Create User

- **POST** `/api/users`
- Creates a new user

**Request Body:**

```json
{
  "name": "New User",
  "email": "newuser@example.com",
  "username": "newuser"
}
```

**Response:**

```json
{
  "id": 4,
  "name": "New User",
  "email": "newuser@example.com",
  "username": "newuser"
}
```

## Testing the API

You can test the API using curl or any HTTP client:

```bash
# Test home endpoint
curl http://localhost:8080/

# Test health endpoint
curl http://localhost:8080/api/health

# Get all users
curl http://localhost:8080/api/users

# Get user by ID
curl http://localhost:8080/api/users/1

# Create a new user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","username":"testuser"}'
```

## Project Structure

```
lab02-go-api/
├── main.go          # Main application file
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
└── README.md        # This file
```

## Dependencies

- `github.com/gorilla/mux` - HTTP router and URL matcher

## Development

To add new endpoints or modify existing ones, edit the `main.go` file. The application uses in-memory storage for simplicity, but in a production environment, you would want to integrate with a database.

## License

This project is for educational purposes.
