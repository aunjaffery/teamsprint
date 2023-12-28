package main

import (
	"log"

	"github.com/aunjaffery/teamsprint/config"
	"github.com/aunjaffery/teamsprint/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadConfig("dev.env")
	config.ConnectDB(config.Envs.DNS)
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8088"))
}
