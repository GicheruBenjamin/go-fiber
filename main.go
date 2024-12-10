package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// Parameterized route
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("User ID: " + id)
	})

	// JSON response
	app.Get("/json", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, Fiber!",
		})
	})

	// Start server
	app.Listen(":3000")
}
