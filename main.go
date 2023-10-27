package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	UserId   string
	User     string
	LastName string
	Address  string
}

func loadding(c *fiber.Ctx) error {

	return c.SendString("Bienvenido Jhon Ludeña!")
}

func userById(c *fiber.Ctx) error {

	users := User{
		User:     "Jhon",
		LastName: "Ludeña",
		Address:  "Salinas barrio 6 de Junio",
	}

	fmt.Println(users)

	return c.Status(fiber.StatusOK).JSON(users)
}

func UserCreate(c *fiber.Ctx) error {
	user := User{}

	c.BodyParser(user)

	user.UserId = uuid.NewString()

	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Get("/", loadding)

	app.Post("/users/user-by-id", userById)

	log.Fatal(app.Listen(":3000"))
}
