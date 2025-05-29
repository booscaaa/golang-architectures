package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
)

type getServiceUseCase struct {
	repository contract.GetServiceDatabaseRepository
}

// Execute implements contract.GetServiceUseCase.
func (usecase *getServiceUseCase) Execute(ctx context.Context, id string) (model.GetServiceOutput, error) {
	service, err := usecase.repository.Execute(ctx, id)
	if err != nil {
		return model.GetServiceOutput{}, err
	}
	return service.ToGetServiceOutput(), nil
}

func NewGetServiceUseCase(repository contract.GetServiceDatabaseRepository) contract.GetServiceUseCase {
	return &getServiceUseCase{
		repository: repository,
	}
}
