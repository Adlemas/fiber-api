package main

import (
	"log"

	"github.com/Adlemas/fiber-api/config"
	"github.com/Adlemas/fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	config.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
