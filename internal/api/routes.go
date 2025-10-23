package api

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/books", BookList)
	v1.Get("/books/:id", BookDetail)
}
