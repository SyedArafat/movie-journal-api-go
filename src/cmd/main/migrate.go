package main

import (
	"flag"
	"movie-journal-api/bootstrap"
	"movie-journal-api/internal/driver/database"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/model"
)

func main() {
	migrate := flag.Bool("migrate", false, "Run migrations")
	flag.Parse()

	if *migrate {
		db := bootstrap.Initialize().InitializeDatabase()
		logger.Info("test", "Migration started")
		runMigrations(db)
		logger.Info("migration_done", "Migration completed successfully")

	}
}

func runMigrations(db *database.DB) {
	err := db.SQL.AutoMigrate(&model.User{})
	if err != nil {
		logger.Error("migration_error", "Migration failed", "", err)
	}
}
