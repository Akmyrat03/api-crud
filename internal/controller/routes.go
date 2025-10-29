package controller

import "github.com/gofiber/fiber/v2"

func MapBookRoutes(group fiber.Router, c bookController) {
	group.Post("/books", c.CreateBook)
}

// group = ""
// 192.168.1.1:3000/books

// group = "/api/v1"
// 192.168.1.1:3000/api/v1/books
