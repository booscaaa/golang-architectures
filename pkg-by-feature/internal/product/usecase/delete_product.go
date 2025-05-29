package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
)

type deleteProductUseCase struct {
	repository contract.DeleteProductDatabaseRepository
}

// Execute implements contract.DeleteProductUseCase.
func (usecase *deleteProductUseCase) Execute(ctx context.Context, id string) error {
	err := usecase.repository.Execute(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteProductUseCase(repository contract.DeleteProductDatabaseRepository) contract.DeleteProductUseCase {
	return &deleteProductUseCase{
		repository: repository,
	}
}
