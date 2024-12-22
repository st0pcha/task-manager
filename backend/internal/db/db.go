package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panic("failed to connect database:", err)
	}

	DB = db
	log.Println("connected to database")
}

func AutoMigrate(tables ...interface{}) error {
	for _, table := range tables {
		if err := DB.AutoMigrate(table); err != nil {
			log.Printf("failed to migrate table: %v", table)
			return err
		}
	}

	log.Printf("migrated tables: %v", tables)
	return nil
}
