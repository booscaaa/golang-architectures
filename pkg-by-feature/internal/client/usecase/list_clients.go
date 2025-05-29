package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
)

type listClientsUseCase struct {
	repository contract.ListClientsDatabaseRepository
}

// Execute implements contract.ListClientsUseCase.
func (usecase *listClientsUseCase) Execute(ctx context.Context) (model.ListClientsOutput, error) {
	clients, err := usecase.repository.Execute(ctx)
	if err != nil {
		return model.ListClientsOutput{}, err
	}
	return model.ClientsToListClientsOutput(clients), nil
}

func NewListClientsUseCase(repository contract.ListClientsDatabaseRepository) contract.ListClientsUseCase {
	return &listClientsUseCase{
		repository: repository,
	}
}
