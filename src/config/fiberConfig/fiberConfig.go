package fiberConfig

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	errorhandler "movie-journal-api/internal/errorHandler"
	"os"
	"strconv"
	"time"
)

var Port = os.Getenv("APP_SERVER_PORT")

func Get() fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	appName := os.Getenv("APP_NAME")
	return fiber.Config{
		AppName:        appName,
		CaseSensitive:  true,
		StrictRouting:  true,
		ReadTimeout:    time.Second * time.Duration(readTimeoutSecondsCount),
		ReadBufferSize: 5192,
		ErrorHandler:   errorhandler.GlobalFiberErrorHandler,
	}
}

func RecoveryConfig() recover.Config {
	return recover.Config{
		EnableStackTrace: false,
	}
}
