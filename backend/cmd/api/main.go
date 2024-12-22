package main

import (
	"log"

	"github.com/st0pcha/task-manager/backend/internal/api"
	"github.com/st0pcha/task-manager/backend/internal/config"
	"github.com/st0pcha/task-manager/backend/internal/dal"
	"github.com/st0pcha/task-manager/backend/internal/db"
)

func main() {
	log.Println("starting server...")
	cfg := config.Initialize()

	db.Connect(cfg.Postgres.DSN)
	if cfg.IsDev() {
		if err := db.AutoMigrate(&dal.User{}, &dal.Task{}); err != nil {
			log.Panic("failed to migrate database:", err)
		}
	}

	api.CreateAPI(cfg)
}
