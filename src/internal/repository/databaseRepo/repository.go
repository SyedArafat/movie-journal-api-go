package databaseRepo

import (
	"gorm.io/gorm"
	"movie-journal-api/bootstrap"
)

var db *gorm.DB

func (app *ApplicationRepo) getDB() *gorm.DB {
	if db == nil {
		db = app.InitializeDatabase().SQL
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
