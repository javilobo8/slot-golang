package app

import (
	"slot-golang/pkg/handlers"
	"slot-golang/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CreateApp() {
	app := fiber.New(fiber.Config{})
	app.Use(middlewares.PreMiddleware)
	app.Get("/bet", handlers.BetHandler)
	app.Listen(":8080")
}
