package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"Proj-GO/database"
	"Proj-GO/models"
)

func Register(c *fiber.Ctx) error {
    var data map[string]string

    if err := c.BodyParser(&data); err != nil {
      return err
    }

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.Users{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

    return c.JSON(user)
}
