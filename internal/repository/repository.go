package repository

import (
	"context"

	"github.com/Akmyrat03/api/crud/internal/domain"
	"github.com/Akmyrat03/api/crud/pkg/connection"
)

// struct and interface strongly connected
var _ BookRepository = (*bookRepository)(nil)

type BookRepository interface {
	Create(ctx context.Context, book domain.Book) error
}

type bookRepository struct {
	psqlDB connection.DB
}

func NewBookRepository(psqlDB connection.DB) *bookRepository {
	return &bookRepository{psqlDB: psqlDB}
}

// CRUD = CREATE, READ, UPDATE, DELETE
func (b bookRepository) Create(ctx context.Context, book domain.Book) error {
	query := `INSERT INTO book (author, title, description) VALUES ($1, $2, $3)`

	if err := b.psqlDB.QueryRow(ctx, query, book.Author, book.Title, book.Description); err != nil {
		return err
	}

	return nil
}
