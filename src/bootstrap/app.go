package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/config"
	"movie-journal-api/internal/driver/database"
)

var App Application

type Application struct {
	App *fiber.App
	db  *database.DB
}

func Initialize() Application {
	application := fiber.New(config.FiberConfig())
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
