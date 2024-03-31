package services

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/bootstrap"
	"movie-journal-api/internal/utils"
)

type ApplicationService struct {
	bootstrap.Application
}

func NewApplicationService(app bootstrap.Application) *ApplicationService {
	return &ApplicationService{
		app,
	}
}

func (as *ApplicationService) BaseEndPoint(c *fiber.Ctx) error {
	return utils.SuccessResponse(
		c,
		fiber.StatusOK,
		"Movie Journal API",
	)
}
