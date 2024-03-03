package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/users", GetUsers)
	app.Get("/user/:id", GetUser)
	app.Post("/user", SaveUser)
	app.Delete("/user/:id", DeleteUser)
	app.Put("/user/:id", UpdateUser)

}

func main() {
	fmt.Println("main <------>")
	InitialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":3000")
}
