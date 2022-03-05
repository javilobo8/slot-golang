package main

import (
	"flag"
	"math/rand"
	"slot-golang/handlers"
	"slot-golang/middlewares"
	"slot-golang/rtp"
	"time"

	"github.com/gofiber/fiber/v2"
)

func app() {
	app := fiber.New(fiber.Config{})
	app.Use(middlewares.PreMiddleware)
	app.Get("/bet", handlers.BetHandler)
	app.Listen(":8080")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mode := flag.String("mode", "server", "run mode, server or rtp")
	flag.Parse()

	switch *mode {
	case "rtp":
		rtp.CalculateRTP()
	case "server":
		app()
	}
}
