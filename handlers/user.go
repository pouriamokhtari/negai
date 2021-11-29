package handlers

import (
	"negai/database"
	"negai/helpers"
	"negai/models"

	"github.com/gofiber/fiber/v2"
)

type NewUserParams struct {
	Email    string `validate:"required,email"`
	FullName string `validate:"required"`
	Role     string `validate:"oneof=admin member,omitempty"`
	Password string `validate:"required,min=8"`
}

type UpdateUserParams struct {
	Email    string `validate:"email,omitempty"`
	FullName string `validate:"omitempty"`
	Role     string `validate:"oneof=admin member,omitempty"`
	Password string `validate:"min=8,omitempty"`
}

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
	params := &NewUserParams{}

	if err := c.BodyParser(params); err != nil {
		return BadRequest(c)
	}

	if err := helpers.ValidateStruct(params); err != nil {
		return ValidationError(c, err)
	}

	passwordDigest, err := helpers.HashPassword(params.Password)
	if err != nil {
		return InternalServerError(c)
	}

	user := &models.User{
		Email:          params.Email,
		FullName:       params.FullName,
		Role:           models.NewRoleFromString(params.Role),
		PasswordDigest: passwordDigest,
	}

	database.Connection.Create(user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	params := &UpdateUserParams{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return NotFound(c)
	}

	if err := c.BodyParser(params); err != nil {
		return BadRequest(c)
	}

	if err := helpers.ValidateStruct(params); err != nil {
		return ValidationError(c, err)
	}

	var user models.User
	database.Connection.First(&user, id)

	user = models.User{
		Email:    params.Email,
		FullName: params.FullName,
		Role:     models.NewRoleFromString(params.Role),
	}

	user.PasswordDigest, err = helpers.HashPassword(params.Password)
	if err != nil {
		return InternalServerError(c)
	}

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
