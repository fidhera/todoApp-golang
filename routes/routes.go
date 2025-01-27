package routes

import (
	"github.com/fidhera/todo-app/controller"
	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app *fiber.App) {
	app.Get("/todos", controller.GetTodos)
	app.Get("/todos/:id", controller.GetTodoByID)
	app.Post("/todos", controller.CreateTodo)
	app.Put("/todos/:id", controller.UpdateTodo)
	app.Delete("/todos/:id", controller.DeleteTodo)
}
