package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/st0pcha/task-manager/backend/internal/config"
	"github.com/st0pcha/task-manager/backend/internal/routes"
)

func CreateAPI(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Task Manager",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Server.AllowOrigins,
		AllowHeaders:     []string{"Origin, Content-Type, Accept"},
		AllowMethods:     []string{"GET, POST, PUT, DELETE, OPTIONS"},
		AllowCredentials: false,
	}))

	routes.RegisterRoutes(app)

	host := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Fatal(app.Listen(host))
	return app
}
