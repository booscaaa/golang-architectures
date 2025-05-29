package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
)

type updateProductUseCase struct {
	repository contract.UpdateProductDatabaseRepository
}

// Execute implements contract.UpdateProductUseCase.
func (usecase *updateProductUseCase) Execute(ctx context.Context, id string, input model.UpdateProductInput) (model.UpdateProductOutput, error) {
	product, err := usecase.repository.Execute(ctx, input.ToProduct(id))
	if err != nil {
		return model.UpdateProductOutput{}, err
	}
	return product.ToUpdateProductOutput(), nil
}

func NewUpdateProductUseCase(repository contract.UpdateProductDatabaseRepository) contract.UpdateProductUseCase {
	return &updateProductUseCase{
		repository: repository,
	}
}
