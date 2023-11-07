// @/main.go
package main

import (
	"log"

	"github/JhonLudenaReyes/session/config"
	"github/JhonLudenaReyes/session/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	// Initialize default config
	app.Use(cors.New())

	config.Connect()

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:userId", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)

	app.Post("/people", handlers.AddPerson)

	log.Fatal(app.Listen(":4449"))
}
