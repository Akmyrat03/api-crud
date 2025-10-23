package api

import (
	"strconv"

	"github.com/Akmyrat03/api/crud/internal/models"
	"github.com/gofiber/fiber/v2"
)

func BookList(c *fiber.Ctx) error {
	bookStore := models.BookStore

	return c.Status(fiber.StatusOK).JSON(bookStore)
}

func BookDetail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "id should be valid string",
		})
	}

	books := models.BookStore

	response := models.Book{}

	for _, book := range books {
		if int(book.ID) == id {
			response = book
		}
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// func BookCreate(c *fiber.Ctx) error {
// 	var req models.BookRequest
// 	if err:= c.BodyParser(&req); err!=nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			""
// 		})
// 	}
// }
