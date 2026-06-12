package database

import (
	"fmt"
	"log/slog"
	"time"

	"arthemis-watcher/internal/env"
	"arthemis-watcher/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(logger *slog.Logger) (*gorm.DB, error) {
	host := env.GetEnv("DB_HOST", "arthemis-watcher-postgres")
	port := env.GetEnv("DB_PORT", "5432")
	user := env.GetEnv("DB_USER", "watcher_user")
	password := env.GetEnv("DB_PASSWORD", "watcher_password")
	databaseName := env.GetEnv("DB_NAME", "watcher_db")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		databaseName,
	)

	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	if err := db.AutoMigrate(&models.AuditLog{}); err != nil {
		logger.Error("Error on AutoMigrate: AuditLog")
		return nil, err
	}

	return db, nil
}
