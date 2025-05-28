package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
)

type createClientUseCase struct {
	createClienteDatabaseRepository contract.CreateClientDatabaseRepository
}

// Execute implements contract.CreateClientUseCase.
func (usecase *createClientUseCase) Execute(ctx context.Context, input model.CreateClientInput) (model.CreateClientOutput, error) {
	client, err := usecase.createClienteDatabaseRepository.Execute(ctx, input.ToClient())
	if err != nil {
		return model.CreateClientOutput{}, err
	}
	return client.ToCreateClientOutput(), nil
}

func NewCreateClientUseCase(createClienteDatabaseRepository contract.CreateClientDatabaseRepository) contract.CreateClientUseCase {
	return &createClientUseCase{
		createClienteDatabaseRepository: createClienteDatabaseRepository,
	}
}
