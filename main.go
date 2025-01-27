package main

import (
	"github.com/fidhera/todo-app/database"
	"github.com/fidhera/todo-app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.TodoRoutes(app)
	app.Listen(":2597")
}
