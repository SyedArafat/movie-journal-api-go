package handler

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/bootstrap"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/services"
)

type ApplicationHandler struct {
	bootstrap.Application
}

func NewApplicationHandler(app bootstrap.Application) *ApplicationHandler {
	return &ApplicationHandler{
		Application: app,
	}
}

func (ah *ApplicationHandler) Base(c *fiber.Ctx) error {
	as := services.NewApplicationService(ah.Application)
	logger.PrintInfo("test-database", "Testing database. Inside repo")
	return as.BaseEndPoint(c)
}
