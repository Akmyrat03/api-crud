package main

import (
	"log"

	"github.com/Akmyrat03/api/crud/internal/api"
	"github.com/Akmyrat03/api/crud/internal/config"
	"github.com/gofiber/fiber/v2"
)

// crud = create read update delete
func main() {
	app := fiber.New()

	api.Routes(app)

	cfg := config.LoadConfig()

	port := cfg.App.Port

	log.Printf("port = %v", port)

	app.Listen(port)
}
