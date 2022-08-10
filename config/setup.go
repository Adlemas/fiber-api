package config

import (
	"github.com/Adlemas/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", routes.WelcomeApi)

	// User endpoints
	app.Post("/api/users", routes.CreateUser)

	app.Get("/api/users", routes.GetUsers)

	app.Get("/api/users/:id", routes.GetUser)

	app.Put("/api/users/:id", routes.UpdateUser)

	app.Delete("/api/users/:id", routes.DeleteUser)

	// Products endpoints
	app.Post("/api/products", routes.CreateProduct)

	app.Get("/api/products", routes.GetProducts)

	app.Get("/api/products/:id", routes.GetProduct)

	app.Put("/api/products/:id", routes.UpdateProduct)

	app.Delete("/api/products/:id", routes.DeleteProduct)

	// Order endpoints
	app.Post("/api/orders", routes.CreateOrder)

	app.Get("/api/orders", routes.GetOrders)

	app.Get("/api/orders/:id", routes.GetOrder)
}
