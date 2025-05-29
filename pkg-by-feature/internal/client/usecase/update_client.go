package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
)

type updateClientUseCase struct {
	repository contract.UpdateClientDatabaseRepository
}

// Execute implements contract.UpdateClientUseCase.
func (usecase *updateClientUseCase) Execute(ctx context.Context, id string, input model.UpdateClientInput) (model.UpdateClientOutput, error) {
	client, err := usecase.repository.Execute(ctx, input.ToClient(id))
	if err != nil {
		return model.UpdateClientOutput{}, err
	}
	return client.ToUpdateClientOutput(), nil
}

func NewUpdateClientUseCase(repository contract.UpdateClientDatabaseRepository) contract.UpdateClientUseCase {
	return &updateClientUseCase{
		repository: repository,
	}
}
