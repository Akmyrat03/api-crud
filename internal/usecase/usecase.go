package usecase

import (
	"context"

	"github.com/Akmyrat03/api/crud/internal/domain"
	"github.com/Akmyrat03/api/crud/internal/repository"
)

type BookUC interface {
	Create(ctx context.Context, book domain.Book) error
}

type bookUC struct {
	bookRepo repository.BookRepository
}

func NewBookUC(bookRepo repository.BookRepository) *bookUC {
	return &bookUC{bookRepo: bookRepo}
}

func (b bookUC) Create(ctx context.Context, book domain.Book) error {
	if err := b.bookRepo.Create(ctx, book); err != nil {
		return err
	}

	return nil
}
