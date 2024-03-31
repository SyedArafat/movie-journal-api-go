package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/config/fiberConfig"
	"movie-journal-api/internal/driver/database"
	"movie-journal-api/internal/logger"
)

var App Application

type Application struct {
	App       *fiber.App
	db        *database.DB
	Validator *validator.Validate
}

func Initialize() Application {
	logger.Initialize()
	application := fiber.New(fiberConfig.Get())
	val := validator.New()
	return initializeApplication(application, val)
}

func initializeApplication(a *fiber.App, v *validator.Validate) Application {
	App = Application{
		App:       a,
		Validator: v,
	}
	return App
}

func (App Application) InitializeDatabase() *database.DB {
	if App.db == nil {
		App.db, _ = database.ConnectSQL()
	}
	return App.db
}
