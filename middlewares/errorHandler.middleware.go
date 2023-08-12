package middlewares

import "github.com/gofiber/fiber/v2"

// [c *fiber.Ctx] represents the context of an HTTP request and response object.
// Just like in Node Js where you have (req, res)
func ErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}