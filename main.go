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

	/*app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))*/

	config.Connect()

	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:userId", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)

	// API SESSION LOGIN
	app.Post("/users/session/login", handlers.GetLoginUser)

	app.Post("/people", handlers.AddPerson)

	log.Fatal(app.Listen(":4449"))
}
