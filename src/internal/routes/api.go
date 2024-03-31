package routes

import (
	"movie-journal-api/bootstrap"
	"movie-journal-api/internal/handler"
)

func InitRoutes(app bootstrap.Application) {
	applicationHandler := handler.NewApplicationHandler(app)

	app.App.Get("/", applicationHandler.Base)
	app.App.Get("/health-check", applicationHandler.HealthCheck)

	api := app.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", applicationHandler.Register)
}
