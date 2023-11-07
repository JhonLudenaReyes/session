package handlers

import (
	"github/JhonLudenaReyes/session/config"
	"github/JhonLudenaReyes/session/entities"

	"github.com/gofiber/fiber/v2"
)

// ...

func AddPerson(c *fiber.Ctx) error {
	person := new(entities.Person)

	if err := c.BodyParser(person); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&person)
	return c.Status(201).JSON(person)
}
