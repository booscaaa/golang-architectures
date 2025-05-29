package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
)

type getProductUseCase struct {
	repository contract.GetProductDatabaseRepository
}

// Execute implements contract.GetProductUseCase.
func (usecase *getProductUseCase) Execute(ctx context.Context, id string) (model.GetProductOutput, error) {
	product, err := usecase.repository.Execute(ctx, id)
	if err != nil {
		return model.GetProductOutput{}, err
	}
	return product.ToGetProductOutput(), nil
}

func NewGetProductUseCase(repository contract.GetProductDatabaseRepository) contract.GetProductUseCase {
	return &getProductUseCase{
		repository: repository,
	}
}
