package di

import (
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/controller"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/repository"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/usecase"
	"github.com/booscaaa/initializers/postgres"
)

func NewCreateProductController(database postgres.Database) contract.CreateProductController {
	return controller.NewCreateProductController(
		usecase.NewCreateProductUseCase(
			repository.NewCreateProductDatabaseRepository(database),
		),
	)
}
