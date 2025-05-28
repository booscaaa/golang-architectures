package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
)

type createProductUseCase struct {
	repository contract.CreateProductDatabaseRepository
}

// Execute implements contract.CreateProductUseCase.
func (usecase *createProductUseCase) Execute(ctx context.Context, input model.CreateProductInput) (model.CreateProductOutput, error) {
	product, err := usecase.repository.Execute(ctx, input.ToProduct())
	if err != nil {
		return model.CreateProductOutput{}, err
	}
	return product.ToCreateProductOutput(), nil
}

func NewCreateProductUseCase(repository contract.CreateProductDatabaseRepository) contract.CreateProductUseCase {
	return &createProductUseCase{
		repository: repository,
	}
}
