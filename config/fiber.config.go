package config

import (
	"github.com/daviesoyasor/fiber-tester/middlewares"
	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configurating Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Return Fiber configuration.
	return fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	}
}