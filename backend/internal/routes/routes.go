package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/controllers"
)

func RegisterRoutes(app *fiber.App) {
	mainController(app)
	authController(app)
}

func mainController(app *fiber.App) {
	app.Get("/", controllers.GetHelloWorld)
}

func authController(app *fiber.App) {
	g := app.Group("/auth")
	g.Post("/register", controllers.Register)
	g.Post("/login", controllers.Register)
}
