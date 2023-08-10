package main

import (
	"github.com/daviesoyasor/fiber-tester/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Set up routes
	prefixUrl := app.Group("/api/v1")
	routes.SetUpRoutes(prefixUrl)

	// Listen on PORT 3000
	app.Listen(":3000")
}
