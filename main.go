package main

import (
	"slot-golang/handlers"
	"slot-golang/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(middlewares.PreMiddleware)

	app.Get("/test", handlers.BetHandler)

	app.Listen(":8080")
}
