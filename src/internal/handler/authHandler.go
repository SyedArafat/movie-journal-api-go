package handler

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/services"
	"movie-journal-api/internal/utils"
	"movie-journal-api/internal/validator"
)

func (ah *ApplicationHandler) Register(ctx *fiber.Ctx) error {
	as := services.NewApplicationService(ah.Application)

	var user validator.User
	err := ctx.BodyParser(&user)
	if err != nil {
		logger.Error("parse_error", "user body parse error", "", err)
	}
	//// Validate request body
	if err := ah.Validator.Struct(user); err != nil {
		logger.Error("validation_error", "Signup failed due to invalid input", "", err)
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error(), err)
	}

	err = as.AddUser(user)

	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error(), err)
	}

	return ctx.Format(user)
}
