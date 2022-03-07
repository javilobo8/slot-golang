package app

import (
	"slot-golang/pkg/handlers"
	"slot-golang/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func initializeMiddlewares(app *fiber.App) {
	app.Use(middlewares.PreMiddleware)
}

func initializeHandlers(app *fiber.App) {
	app.Get("/bet", handlers.BetHandler)
}

func CreateApp() {
	app := fiber.New(fiber.Config{})

	initializeMiddlewares(app)
	initializeHandlers(app)

	app.Listen(":8080")
}
