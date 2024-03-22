package utils

import (
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/internal/dto"
)

func SuccessResponse(ctx *fiber.Ctx, code int, data any) error {
	return ctx.Status(code).JSON(dto.SuccessResponseDto{
		Status:     "Success",
		StatusCode: code,
		Data:       data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, message any, details any) error {
	return ctx.Status(code).JSON(dto.ErrorResponseDto{
		Status:       "Error",
		StatusCode:   code,
		ErrorMessage: message,
		ErrorDetails: details,
	})
}
