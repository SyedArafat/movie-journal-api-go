package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"movie-journal-api/bootstrap"
	"movie-journal-api/config/fiberConfig"
	errorhandler "movie-journal-api/internal/errorHandler"
	"movie-journal-api/internal/routes"
)

func main() {
	app := bootstrap.Initialize()

	app.App.Use(errorhandler.GlobalErrorHandler)

	//app.App.Use(func(c *fiber.Ctx) error {
	//	defer func() {
	//		if r := recover(); r != nil {
	//			panicMsg := fmt.Sprintf("Panic: %v\n\nStack Trace:\n%s", r, debug.Stack())
	//			c.Status(fiber.StatusInternalServerError).SendString(panicMsg)
	//			//
	//			//// Check if the panic message is of type *net.OpError
	//			//if opErr, ok := r.(*net.OpError); ok {
	//			//	panicMsg := fmt.Sprintf("Net Error: %v", opErr.Error())
	//			//	c.Status(fiber.StatusInternalServerError).SendString(panicMsg)
	//			//} else {
	//			//	// If not, handle it accordingly (possibly convert it to string)
	//			//	panicMsg := fmt.Sprintf("Panic: %v", r)
	//			//	c.Status(fiber.StatusInternalServerError).SendString(panicMsg)
	//			//}
	//		}
	//	}()
	//	return c.Next()
	//})
	app.App.Use(recover.New(fiberConfig.RecoveryConfig()))
	//app.App.Use(logger.New(logger.Config{
	//	Format: "${status} - ${method} ${path}\n",
	//}))

	routes.InitRoutes(app)

	fmt.Println("AB Test MS successfully initiated")

	log.Println(app.App.Listen(":" + fiberConfig.Port))
}
