package databaseRepo

import (
	"movie-journal-api/bootstrap"
	"movie-journal-api/internal/driver/database"
	"movie-journal-api/internal/logger"
)

var db *database.DB

func (app *ApplicationRepo) getDB() *database.DB {
	if db == nil {
		logger.Info("db is nil", "")
		db = app.InitializeDatabase()
	}

	return db
}

type ApplicationRepo struct {
	bootstrap.Application
}

func NewApplicationRepo(app bootstrap.Application) *ApplicationRepo {
	app.InitializeDatabase()
	return &ApplicationRepo{
		app,
	}
}
