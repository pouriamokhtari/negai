package handlers

import (
	"negai/database"
	"negai/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return NotFound(c)
	}
	database.Connection.First(&user, id)
	return c.JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Connection.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return BadRequest(c)
	}

	database.Connection.Create(user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return NotFound(c)
	}

	if err := c.BodyParser(user); err != nil {
		return BadRequest(c)
	}

	database.Connection.First(&user, id)
	database.Connection.Updates(user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return NotFound(c)
	}

	database.Connection.First(&user, id)
	database.Connection.Delete(&user)
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
