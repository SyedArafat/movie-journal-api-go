package handler

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/services"
)

func (ah *ApplicationHandler) HealthCheck(c *fiber.Ctx) error {
	logger.PrintInfo("test", "coming")
	as := services.NewApplicationService(ah.Application)
	err := as.CheckDBConnection()
	if err != nil {
		panic(err)
	}
	return c.SendString("Ok!")
}
