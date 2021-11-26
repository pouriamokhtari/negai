package handlers

import (
	"negai/database"
	"negai/models"

	"github.com/gofiber/fiber/v2"
)

// UserGet returns a user
func UserList(c *fiber.Ctx) error {
	var users []models.User
	database.Connection.Find(&users)
	return c.JSON(fiber.Map{
		"success": true,
		"user":    users,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	user := &models.User{
		Name: c.FormValue("user"),
	}
	database.Connection.Create(&user)
	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error": "object not found",
	})
}
