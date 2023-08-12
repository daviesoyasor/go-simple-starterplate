package main

import (
	"time"

	"github.com/daviesoyasor/fiber-tester/config"
	"github.com/daviesoyasor/fiber-tester/middlewares"
	"github.com/daviesoyasor/fiber-tester/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	// Define Fiber config.
	config := config.FiberConfig()

	// Start a new fiber app (Initialize a fiber instance)
	app := fiber.New(config)

	// Set up routes
	prefixUrl := app.Group("/api/v1")
	routes.SetUpRoutes(prefixUrl)

	// === Middlewares ====
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))
	
	app.Use(middlewares.NotFoundRoute)
	// === End Middlewares ====

	// Listen on PORT 3000
	app.Listen(":3000")
}
