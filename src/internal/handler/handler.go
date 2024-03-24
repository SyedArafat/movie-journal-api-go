package handler

import (
	"errors"
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
	//return fmt.Errorf("What an idea")
	return errors.New("asasdasdasdasd")
	return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
	logger.PrintInfo("movie-journal-api", "Movie journal API")
	return as.BaseEndPoint(c)
}
