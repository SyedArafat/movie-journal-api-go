package databaseRepo

import (
	"movie-journal-api/bootstrap"
	"movie-journal-api/internal/driver/database"
)

var db *database.DB

func (app *ApplicationRepo) getDB() *database.DB {
	if db == nil {
		db = app.InitializeDatabase()
	}

	return db
}

type ApplicationRepo struct {
	bootstrap.Application
}

func NewApplicationRepo(app bootstrap.Application) *ApplicationRepo {
	db = app.InitializeDatabase()
	return &ApplicationRepo{
		app,
	}
}
