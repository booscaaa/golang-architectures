package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
)

type deleteClientUseCase struct {
	repository contract.DeleteClientDatabaseRepository
}

// Execute implements contract.DeleteClientUseCase.
func (usecase *deleteClientUseCase) Execute(ctx context.Context, id string) error {
	return usecase.repository.Execute(ctx, id)
}

func NewDeleteClientUseCase(repository contract.DeleteClientDatabaseRepository) contract.DeleteClientUseCase {
	return &deleteClientUseCase{
		repository: repository,
	}
}
