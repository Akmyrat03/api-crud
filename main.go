package main

import (
	"context"
	"log"

	"github.com/Akmyrat03/api/crud/internal/config"
	"github.com/Akmyrat03/api/crud/internal/controller"
	"github.com/Akmyrat03/api/crud/internal/repository"
	"github.com/Akmyrat03/api/crud/internal/usecase"
	"github.com/Akmyrat03/api/crud/pkg/connection"
	"github.com/gofiber/fiber/v2"
)

// crud = create read update delete
func main() {
	app := fiber.New()

	group := app.Group("/api/v1")

	psqlDB, err := connection.NewDBConnection(context.Background(), config)
	if err != nil {
		log.Printf("failed connect to db: %v\n", err)
	}

	bookRepository := repository.NewBookRepository(psqlDB)

	bookUseCase := usecase.NewBookUC(bookRepository)

	bookController := controller.NewBookController(bookUseCase)

	controller.MapBookRoutes(group, *bookController)

	cfg := config.LoadConfig()

	port := cfg.App.Port

	log.Printf("port = %v", port)

	app.Listen(port)
}
