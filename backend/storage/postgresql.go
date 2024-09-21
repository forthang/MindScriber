package storage

import (
	"fmt"
	"github.com/forthang/esketit/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Username,
		cfg.Storage.Password, cfg.Storage.Database, cfg.Storage.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable connect to DB: %v", err)
	}

	return db
}
