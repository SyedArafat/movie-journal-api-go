package main

import (
	"log"
	"movie-journal-api/bootstrap"
	"movie-journal-api/config/fiberConfig"
	errorhandler "movie-journal-api/internal/errorHandler"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/routes"
)

func main() {
	app := bootstrap.Initialize()

	app.App.Use(errorhandler.GlobalErrorHandler)
	routes.InitRoutes(app)
	logger.Info("main", "AB Test MS successfully initiated")

	log.Println(app.App.Listen(":" + fiberConfig.Port))
}
