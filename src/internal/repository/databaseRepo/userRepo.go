package databaseRepo

import (
	"movie-journal-api/internal/model"
)

func (app *ApplicationRepo) AddUser(user model.User) error {
	db := app.getDB().SQL
	if err := db.Create(&user).Error; err != nil {
		// Handle GORM errors
		return err
	}
	db.Create(&user)
}
