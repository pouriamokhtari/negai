package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	return InternalServerError(c)
}

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

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{
			"Error": "unauthorized",
		})
}

func InternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"Error": "internal server error",
		})
}
