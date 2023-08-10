package v1

import (
	"github.com/daviesoyasor/fiber-tester/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	route := router.Group("/auth") // Create a sub-router for /api/v1/auth
	route.Get("/login", controllers.Login)
}
