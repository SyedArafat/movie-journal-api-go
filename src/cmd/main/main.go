package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"movie-journal-api/bootstrap"
	"movie-journal-api/config/fiberConfig"
	"movie-journal-api/internal/routes"
)

func main() {
	app := bootstrap.Initialize()
	app.App.Use(recover.New(fiberConfig.RecoveryConfig()))
	app.App.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	routes.InitRoutes(app)

	fmt.Println("AB Test MS successfully initiated")

	log.Println(app.App.Listen(":" + fiberConfig.Port))
}
