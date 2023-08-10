package routes

import (
	"github.com/daviesoyasor/fiber-tester/routes/v1"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(router fiber.Router) {
	v1.AuthRoutes(router)
	// Add more route groups as needed
}