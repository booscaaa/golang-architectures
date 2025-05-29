package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
)

type deleteServiceUseCase struct {
	repository contract.DeleteServiceDatabaseRepository
}

// Execute implements contract.DeleteServiceUseCase.
func (usecase *deleteServiceUseCase) Execute(ctx context.Context, id string) error {
	return usecase.repository.Execute(ctx, id)
}

func NewDeleteServiceUseCase(repository contract.DeleteServiceDatabaseRepository) contract.DeleteServiceUseCase {
	return &deleteServiceUseCase{
		repository: repository,
	}
}
