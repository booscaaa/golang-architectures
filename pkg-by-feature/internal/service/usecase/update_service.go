package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
)

type updateServiceUseCase struct {
	repository contract.UpdateServiceDatabaseRepository
}

// Execute implements contract.UpdateServiceUseCase.
func (usecase *updateServiceUseCase) Execute(ctx context.Context, id string, input model.UpdateServiceInput) (model.UpdateServiceOutput, error) {
	service := input.ToService(id)
	updatedService, err := usecase.repository.Execute(ctx, service)
	if err != nil {
		return model.UpdateServiceOutput{}, err
	}
	return updatedService.ToUpdateServiceOutput(), nil
}

func NewUpdateServiceUseCase(repository contract.UpdateServiceDatabaseRepository) contract.UpdateServiceUseCase {
	return &updateServiceUseCase{
		repository: repository,
	}
}
