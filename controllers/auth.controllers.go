package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Welcome, you have been authenticated successfully",
	})
}