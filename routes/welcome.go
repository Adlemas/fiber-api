package routes

import "github.com/gofiber/fiber/v2"

func WelcomeApi(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}
