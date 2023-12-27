package main

import (
	"github.com/aunjaffery/teamsprint/config"
	"github.com/aunjaffery/teamsprint/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadConfig("dev.env")
	config.ConnectDB(config.Envs.DNS)
	app := fiber.New()
	app.Use(cors.New())
	routes.SetupRoutes(app)
	app.Listen(":8000")
}

