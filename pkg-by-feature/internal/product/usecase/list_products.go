package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
)

type listProductsUseCase struct {
	repository contract.ListProductsDatabaseRepository
}

// Execute implements contract.ListProductsUseCase.
func (usecase *listProductsUseCase) Execute(ctx context.Context) (model.ListProductsOutput, error) {
	products, err := usecase.repository.Execute(ctx)
	if err != nil {
		return model.ListProductsOutput{}, err
	}
	return model.ProductsToListProductsOutput(products), nil
}

func NewListProductsUseCase(repository contract.ListProductsDatabaseRepository) contract.ListProductsUseCase {
	return &listProductsUseCase{
		repository: repository,
	}
}
