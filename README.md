# Lab02 - Full Stack Product Management System

A complete full-stack web application built with Go (Fiber) backend and React (Vite) frontend for managing products.

## 🏗️ Tech Stack

### Backend

- **Go** with Fiber framework
- **PostgreSQL** database
- **GORM** for database operations
- **Environment variables** for configuration
- **CORS** enabled for frontend communication

### Frontend

- **React 18** with Vite
- **Tailwind CSS** for styling
- **Axios** for API communication
- **Modern UI** with loading and error states

## 📁 Project Structure

```
lab02-go-api/
├── backend/
│   ├── config.env          # Environment variables
│   ├── main.go             # Main application entry point
│   ├── go.mod              # Go dependencies
│   ├── go.sum              # Go dependencies checksum
│   ├── models/
│   │   └── product.go      # Product model
│   ├── database/
│   │   └── database.go     # Database connection
│   ├── handlers/
│   │   └── product_handler.go  # API handlers
│   └── routes/
│       └── routes.go       # Route definitions
├── frontend/
│   ├── package.json        # Node.js dependencies
│   ├── vite.config.js      # Vite configuration
│   ├── tailwind.config.js  # Tailwind CSS configuration
│   ├── postcss.config.js   # PostCSS configuration
│   ├── index.html          # Main HTML file
│   └── src/
│       ├── main.jsx        # React entry point
│       ├── App.jsx         # Main React component
│       ├── App.css         # Component styles
│       └── index.css       # Global styles
└── README.md               # This file
```

## 🚀 Quick Start

### Prerequisites

- Go 1.24.5+
- Node.js 16+
- PostgreSQL 17+
- Git

### 1. Database Setup

#### Option A: Using pgAdmin (Recommended)

1. Open pgAdmin from your Start menu
2. Connect to PostgreSQL server using your password
3. Create a new database named `lab02shop`
4. The table will be created automatically when you run the backend

#### Option B: Using psql command line

```bash
psql -U postgres
CREATE DATABASE lab02shop;
\q
```

### 2. Backend Setup

1. **Navigate to backend directory:**

   ```bash
   cd backend
   ```

2. **Update database configuration:**
   Edit `config.env` and update the password:

   ```env
   DB_PASSWORD=your_actual_password
   ```

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

4. **Run the backend:**

   ```bash
   go run main.go
   ```

   The backend will start on `http://localhost:8080`

### 3. Frontend Setup

1. **Navigate to frontend directory:**

   ```bash
   cd frontend
   ```

2. **Install dependencies:**

   ```bash
   npm install
   ```

3. **Run the frontend:**

   ```bash
   npm run dev
   ```

   The frontend will start on `http://localhost:5173`

## 📡 API Endpoints

| Method | Endpoint            | Description        |
| ------ | ------------------- | ------------------ |
| GET    | `/api/products`     | Get all products   |
| GET    | `/api/products/:id` | Get product by ID  |
| POST   | `/api/products`     | Create new product |
| PUT    | `/api/products/:id` | Update product     |
| DELETE | `/api/products/:id` | Delete product     |

### Example API Usage

```bash
# Get all products
curl http://localhost:8080/api/products

# Create a product
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","price":999}'

# Update a product
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Laptop","price":1099}'

# Delete a product
curl -X DELETE http://localhost:8080/api/products/1
```

## 🎨 Frontend Features

- **Product List**: Display all products in a clean table
- **Add Product**: Form to create new products
- **Edit Product**: Inline editing with form
- **Delete Product**: Remove products with confirmation
- **Loading States**: Show loading indicators
- **Error Handling**: Display error messages
- **Responsive Design**: Works on desktop and mobile

## 🔧 Configuration

### Backend Environment Variables

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=lab02shop
DB_SSLMODE=disable
SERVER_PORT=8080
```

### Frontend Configuration

The frontend is configured to connect to `http://localhost:8080/api` by default. You can modify this in `src/App.jsx` if needed.

## 🛠️ Development

### Backend Development

- The backend uses Fiber framework for fast HTTP routing
- GORM handles database operations with automatic migrations
- CORS is configured to allow frontend communication
- Environment variables are loaded from `config.env`

### Frontend Development

- Vite provides fast development server with hot reload
- Tailwind CSS for utility-first styling
- Axios for HTTP requests with error handling
- React hooks for state management

## 🚀 Production Build

### Backend

```bash
cd backend
go build -o main
./main
```

### Frontend

```bash
cd frontend
npm run build
```

The built files will be in `frontend/dist/`

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Error**

   - Ensure PostgreSQL is running
   - Check password in `config.env`
   - Verify database `lab02shop` exists

2. **CORS Errors**

   - Backend CORS is configured for `localhost:5173` and `localhost:3000`
   - Ensure frontend is running on the correct port

3. **Port Already in Use**

   - Backend: Change `SERVER_PORT` in `config.env`
   - Frontend: Change port in `vite.config.js`

4. **Module Not Found Errors**
   - Run `go mod tidy` in backend directory
   - Run `npm install` in frontend directory

## 📝 License

This project is for educational purposes.

## 🤝 Contributing

This is a lab assignment. Feel free to modify and extend the functionality as needed.
