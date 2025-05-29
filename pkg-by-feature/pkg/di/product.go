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

func NewGetProductController(database postgres.Database) contract.GetProductController {
	return controller.NewGetProductController(
		usecase.NewGetProductUseCase(
			repository.NewGetProductDatabaseRepository(database),
		),
	)
}

func NewUpdateProductController(database postgres.Database) contract.UpdateProductController {
	return controller.NewUpdateProductController(
		usecase.NewUpdateProductUseCase(
			repository.NewUpdateProductDatabaseRepository(database),
		),
	)
}

func NewDeleteProductController(database postgres.Database) contract.DeleteProductController {
	return controller.NewDeleteProductController(
		usecase.NewDeleteProductUseCase(
			repository.NewDeleteProductDatabaseRepository(database),
		),
	)
}

func NewListProductsController(database postgres.Database) contract.ListProductsController {
	return controller.NewListProductsController(
		usecase.NewListProductsUseCase(
			repository.NewListProductsDatabaseRepository(database),
		),
	)
}
