package errorhandler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
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
	log.Println("----------- logg ------------")
	stack := debug.Stack()
	panicMsg := string(stack)
	//fmt.Println("Panic occurred:", panicMsg)
	log.Println(panicMsg)
	logger.PrintError("global_error", err.Error(), err)

	return utils.ErrorResponse(ctx, statusCode, errorMessage, errorDetails)
}

func GlobalErrorHandler(c *fiber.Ctx) error {
	var statusCode int
	var errorMessage string
	var errorDetails interface{}
	defer func() {
		if r := recover(); r != nil {
			statusCode = fiber.StatusInternalServerError
			errorMessage = "Internal Server Errors"
			//errorDetails := debug.Stack()
			errorDetails = map[string]interface{}{
				"reason": fmt.Sprintf("Panic: %v\n", r),
				"trace":  string(debug.Stack()),
			}

			err := utils.ErrorResponse(c, statusCode, errorMessage, errorDetails)
			if err != nil {
				return
			}

		}
	}()
	return c.Next()

}
