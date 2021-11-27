package handlers

import (
	"negai/database"
	"negai/helpers"
	"negai/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type RegisterParams struct {
	Email    string
	Password string
}

type LoginParams struct {
	Email    string
	Password string
}

func Register(c *fiber.Ctx) error {
	params := &RegisterParams{}

	if err := c.BodyParser(params); err != nil {
		return c.JSON(fiber.Map{
			"Error": "validation error",
		})
	}

	passwordDigest, err := helpers.HashPassword(params.Password)
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Email:          params.Email,
		PasswordDigest: passwordDigest,
	}
	database.Connection.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	params := &LoginParams{}

	if err := c.BodyParser(params); err != nil {
		return c.JSON(fiber.Map{
			"Error": "validation error",
		})
	}

	user := &models.User{}

	result := database.Connection.Where("email = ?", params.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Error": "Unauthorized",
		})
	}

	if !helpers.CheckPasssword(params.Password, user.PasswordDigest) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Error": "Unauthorized",
		})
	}

	claims := jwt.MapClaims{
		"Email":   user.Email,
		"Expires": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": encodedToken,
	})
}
