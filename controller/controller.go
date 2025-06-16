package controller

import (
	"intern_template_v1/middleware"
	"intern_template_v1/model"

	"github.com/gofiber/fiber/v2"
)

// SampleController is an example endpoint which returns a
// simple string message.
func SampleController(c *fiber.Ctx) error {
	return c.SendString("Hello, Golang World!")
}

func GetUsers(c *fiber.Ctx) error {
	var users []model.User
	if err := middleware.DBConn.Table("users").Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := middleware.DBConn.Table("users").Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	if err := middleware.DBConn.Table("users").First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := middleware.DBConn.Table("users").Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := middleware.DBConn.Table("users").Delete(&model.User{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}
	return c.SendStatus(204)
}
