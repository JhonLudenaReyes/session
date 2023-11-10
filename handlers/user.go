package handlers

import (
	"fmt"
	"github/JhonLudenaReyes/session/config"
	"github/JhonLudenaReyes/session/entities"

	"github.com/gofiber/fiber/v2"
)

// ...

func GetLoginUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	result := config.Database.Joins("Person").Joins("Role").Find(&user, "user= ? and password= ?", user.User, user.Password)

	if result.RowsAffected == 0 {
		return c.SendString("¡Su usuario y/o contraseña son incorrectos!")
	}

	return c.Status(200).JSON(user)
}

// ...

func GetUsers(c *fiber.Ctx) error {
	var users []entities.User

	config.Database.Find(&users)

	fmt.Println(users)
	return c.Status(200).JSON(users)
}

// ...

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.User

	result := config.Database.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

// ...

func AddUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&user)
	return c.Status(201).JSON(user)
}

// ...

func UpdateUser(c *fiber.Ctx) error {
	user := new(entities.User)
	userId := c.Params("userId")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("user_id = ?", userId).Updates(&user)
	return c.Status(200).JSON(user)
}

// ...

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.User

	result := config.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// ...
