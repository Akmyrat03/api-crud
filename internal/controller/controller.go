package controller

import (
	"github.com/Akmyrat03/api/crud/internal/controller/request"
	"github.com/Akmyrat03/api/crud/internal/domain"
	"github.com/Akmyrat03/api/crud/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

// crud-> create, read, update, delete

type BookController interface {
	CreateBook(c *fiber.Ctx) error
}

type bookController struct {
	bookUC usecase.BookUC
}

func NewBookController(bookUC usecase.BookUC) *bookController {
	return &bookController{bookUC: bookUC}
}

func (b bookController) CreateBook(c *fiber.Ctx) error {
	var req request.BookRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	book := domain.Book{
		Author:      req.Author,
		Title:       req.Title,
		Description: req.Description,
	}

	if err := b.bookUC.Create(c.Context(), book); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
