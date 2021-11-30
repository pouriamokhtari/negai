package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func InternalServerError(c *fiber.Ctx, err error) error {
	log.Println(err)
	return c.Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{
			"Error": "internal server error",
		})
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

func InvalidJWT(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{
			"Error": "missing or malformed jwt",
		})
}

func ValidationError(c *fiber.Ctx, err []string) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"Error":      "validation error",
			"Validation": err,
		})
}
