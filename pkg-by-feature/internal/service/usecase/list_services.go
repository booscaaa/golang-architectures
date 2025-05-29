package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
)

type listServicesUseCase struct {
	repository contract.ListServicesDatabaseRepository
}

// Execute implements contract.ListServicesUseCase.
func (usecase *listServicesUseCase) Execute(ctx context.Context) (model.ListServicesOutput, error) {
	services, err := usecase.repository.Execute(ctx)
	if err != nil {
		return model.ListServicesOutput{}, err
	}
	return model.ServicesToListServicesOutput(services), nil
}

func NewListServicesUseCase(repository contract.ListServicesDatabaseRepository) contract.ListServicesUseCase {
	return &listServicesUseCase{
		repository: repository,
	}
}
