package errorhandler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/utils"
	"runtime/debug"
)

func GlobalFiberErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	var statusCode int
	var errorMessage string
	var errorDetails interface{}

	// Check the type of error and set status code and message accordingly
	var e *fiber.Error
	switch {
	case errors.As(err, &e):
		statusCode = e.Code
		errorMessage = e.Message
		errorDetails = e // Fiber error details
	default:
		statusCode = fiber.StatusInternalServerError
		errorMessage = "Internal Server Errors"
		errorDetails = map[string]interface{}{
			"reason": err.Error(),
			"trace":  string(debug.Stack()),
		}
	}
	logger.PrintError("global_error", err.Error(), err)

	return utils.ErrorResponse(ctx, statusCode, errorMessage, errorDetails)
}
