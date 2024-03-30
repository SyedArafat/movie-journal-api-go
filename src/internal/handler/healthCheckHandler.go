package handler

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/internal/services"
	"movie-journal-api/internal/utils"
)

func (ah *ApplicationHandler) HealthCheck(c *fiber.Ctx) error {
	as := services.NewApplicationService(ah.Application)
	err := as.CheckDBConnection()
	if err != nil {
		panic(err)
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Ok")
}
