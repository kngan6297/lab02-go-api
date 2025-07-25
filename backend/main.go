package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kngan6297/lab02-go-api/database"
	"github.com/kngan6297/lab02-go-api/routes"
)

func main() {
	// Load environment variables
	env := os.Getenv("ENV")
	if env == "production" {
		if err := godotenv.Load("config.production.env"); err != nil {
			log.Println("No production config file found, using system environment variables")
		}
	} else {
		if err := godotenv.Load("config.env"); err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	}

	// Connect to database
	database.ConnectDB()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Configure CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:3000, https://pawfectfriends.xyz, https://www.pawfectfriends.xyz",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Health check endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Lab02 Go API with Fiber and PostgreSQL",
			"status":  "running",
		})
	})

	// Get port from environment or use default
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
