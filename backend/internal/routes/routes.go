package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/st0pcha/task-manager/backend/internal/controllers"
)

func RegisterRoutes(app *fiber.App) {
	mainController(app)
}

func mainController(app *fiber.App) {
	app.Get("/", controllers.GetHelloWorld)
}
