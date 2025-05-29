package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
)

// Repository interfaces
type CreateProductDatabaseRepository interface {
	Execute(context.Context, model.Product) (model.Product, error)
}

type GetProductDatabaseRepository interface {
	Execute(context.Context, string) (model.Product, error)
}

type UpdateProductDatabaseRepository interface {
	Execute(context.Context, model.Product) (model.Product, error)
}

type DeleteProductDatabaseRepository interface {
	Execute(context.Context, string) error
}

type ListProductsDatabaseRepository interface {
	Execute(context.Context) ([]model.Product, error)
}

// UseCase interfaces
type CreateProductUseCase interface {
	Execute(context.Context, model.CreateProductInput) (model.CreateProductOutput, error)
}

type GetProductUseCase interface {
	Execute(context.Context, string) (model.GetProductOutput, error)
}

type UpdateProductUseCase interface {
	Execute(context.Context, string, model.UpdateProductInput) (model.UpdateProductOutput, error)
}

type DeleteProductUseCase interface {
	Execute(context.Context, string) error
}

type ListProductsUseCase interface {
	Execute(context.Context) (model.ListProductsOutput, error)
}

// Controller interfaces
type CreateProductController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type GetProductController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpdateProductController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type DeleteProductController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type ListProductsController interface {
	Execute(http.ResponseWriter, *http.Request)
}
