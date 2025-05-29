package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
)

type createServiceUseCase struct {
	repository contract.CreateServiceDatabaseRepository
}

// Execute implements contract.CreateServiceUseCase.
func (usecase *createServiceUseCase) Execute(ctx context.Context, input model.CreateServiceInput) (model.CreateServiceOutput, error) {
	service, err := usecase.repository.Execute(ctx, input.ToService())
	if err != nil {
		return model.CreateServiceOutput{}, err
	}
	return service.ToCreateServiceOutput(), nil
}

func NewCreateServiceUseCase(repository contract.CreateServiceDatabaseRepository) contract.CreateServiceUseCase {
	return &createServiceUseCase{
		repository: repository,
	}
}
