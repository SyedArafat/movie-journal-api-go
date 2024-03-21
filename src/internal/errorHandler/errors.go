package errorhandler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"movie-journal-api/internal/utils"
)

func CustomFiberErrorHandler(ctx *fiber.Ctx, err error) error {
	//var errDetails string

	//er = err
	//errDetails = "Stack Trace:" + string(debug.Stack())
	//log.Println(err)
	//os.Exit(1)

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
		}
	}

	return utils.ErrorResponse(ctx, statusCode, errorMessage, errorDetails)
}
