package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).
		JSON(fiber.Map{
			"Error": "object not found",
		})
}

func BadRequest(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"Error": "malformed request",
		})
}
