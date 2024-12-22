package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st0pcha/task-manager/backend/internal/controllers"
	"github.com/st0pcha/task-manager/backend/pkg/middleware"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.RequestLogger)

	mainController(api)
	authController(api)
}

func mainController(r fiber.Router) {
	r.Get("/", controllers.GetHelloWorld)
}

func authController(r fiber.Router) {
	g := r.Group("/auth")
	g.Post("/register", controllers.Register)
	g.Post("/login", controllers.Login)
}
