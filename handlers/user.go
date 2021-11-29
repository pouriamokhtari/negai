package handlers

import (
	"negai/helpers"
	"negai/models"

	"github.com/gofiber/fiber/v2"
)

type NewUserParams struct {
	Email    string `validate:"required,email"`
	FullName string `validate:"required"`
	Role     string `validate:"oneof=admin member|eq="`
	Password string `validate:"required,min=8"`
}

type UpdateUserParams struct {
	Email    string `validate:"email|eq="`
	FullName string
	Role     string `validate:"oneof=admin member|eq="`
	Password string `validate:"min=8|eq="`
}

func GetUser(c *fiber.Ctx) error {
	user := &models.User{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return NotFound(c)
	}
	user.Find(uint(id))
	return c.JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := models.GetAllUsers()

	if err != nil {
		InternalServerError(c, err)
	}
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

	user := &models.User{
		Email:    params.Email,
		FullName: params.FullName,
		Role:     models.RoleFromString(params.Role),
		Password: params.Password,
	}

	if err := user.Create(); err != nil {
		return InternalServerError(c, err)
	}
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

	user := &models.User{}
	user.Find(uint(id))

	err = user.Update(models.User{
		Email:    params.Email,
		FullName: params.FullName,
		Role:     models.RoleFromString(params.Role),
		Password: params.Password,
	})
	if err != nil {
		return InternalServerError(c, err)
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return NotFound(c)
	}

	user := &models.User{}
	user.Find(uint(id))
	if err := user.Delete(); err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
