package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
)

type getClientUseCase struct {
	repository contract.GetClientDatabaseRepository
}

// Execute implements contract.GetClientUseCase.
func (usecase *getClientUseCase) Execute(ctx context.Context, id string) (model.GetClientOutput, error) {
	client, err := usecase.repository.Execute(ctx, id)
	if err != nil {
		return model.GetClientOutput{}, err
	}
	return client.ToGetClientOutput(), nil
}

func NewGetClientUseCase(repository contract.GetClientDatabaseRepository) contract.GetClientUseCase {
	return &getClientUseCase{
		repository: repository,
	}
}
