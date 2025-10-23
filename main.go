package main

import (
	"github.com/Akmyrat03/api/crud/internal/api"
	"github.com/gofiber/fiber/v2"
)

// crud = create read update delete
func main() {
	app := fiber.New()

	api.Routes(app)

	app.Listen(":3000")
}
