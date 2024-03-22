package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/config/fiberConfig"
	"movie-journal-api/internal/driver/database"
	"movie-journal-api/internal/logger"
)

var App Application

type Application struct {
	App *fiber.App
	db  *database.DB
}

func Initialize() Application {
	logger.Initialize()
	application := fiber.New(fiberConfig.Get())
	return initializeApplication(application)
}

func initializeApplication(a *fiber.App) Application {
	App = Application{
		App: a,
	}
	return App
}

func (App Application) InitializeDatabase() *database.DB {
	if App.db == nil {
		App.db, _ = database.ConnectSQL()
	}
	return App.db
}
