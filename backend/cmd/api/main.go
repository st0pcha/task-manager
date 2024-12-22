package main

import (
	"github.com/st0pcha/task-manager/backend/internal/api"
	"github.com/st0pcha/task-manager/backend/internal/config"
)

func main() {
	cfg := config.Initialize()
	api.CreateAPI(cfg)
}
