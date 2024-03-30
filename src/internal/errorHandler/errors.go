package errorhandler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"movie-journal-api/config"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/utils"
	"runtime/debug"
)

func FiberErrorHandler(ctx *fiber.Ctx, err error) error {
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
		errorMessage = config.ServerErrorMessage
		errorDetails = map[string]interface{}{
			"reason": err.Error(),
			//"trace":  string(debug.Stack()),
		}
	}
	log.Println("----------- Error log ------------")
	logger.Error("global_error", err.Error(), "", err)

	return utils.ErrorResponse(ctx, statusCode, errorMessage, errorDetails)
}

func GlobalErrorHandler(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			reason := fmt.Sprintf("%v\n", r)
			statusCode := fiber.StatusInternalServerError
			errorMessage := config.ServerErrorMessage
			errorDetails := map[string]interface{}{
				"reason": reason,
				"trace":  string(debug.Stack()),
			}
			logger.Error("panic", reason, string(debug.Stack()), nil)
			err := utils.ErrorResponse(c, statusCode, errorMessage, errorDetails)
			if err != nil {
				return
			}

		}
	}()
	return c.Next()
}
