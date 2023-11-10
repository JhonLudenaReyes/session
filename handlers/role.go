package handlers

import (
	"github/JhonLudenaReyes/session/config"
	"github/JhonLudenaReyes/session/entities"

	"github.com/gofiber/fiber/v2"
)

// ...

func AddRole(c *fiber.Ctx) error {
	role := new(entities.Role)

	if err := c.BodyParser(role); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&role)
	return c.Status(201).JSON(role)
}
