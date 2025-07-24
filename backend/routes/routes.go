package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kngan6297/lab02-go-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	// API routes
	api := app.Group("/api")

	// Product routes
	products := api.Group("/products")
	products.Get("/", handlers.GetProducts)
	products.Get("/:id", handlers.GetProduct)
	products.Post("/", handlers.CreateProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)
}
