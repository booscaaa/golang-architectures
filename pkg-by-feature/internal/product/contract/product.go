package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
)

type CreateProductDatabaseRepository interface {
	Execute(context.Context, model.Product) (model.Product, error)
}

type CreateProductUseCase interface {
	Execute(context.Context, model.CreateProductInput) (model.CreateProductOutput, error)
}

type CreateProductController interface {
	Execute(http.ResponseWriter, *http.Request)
}
