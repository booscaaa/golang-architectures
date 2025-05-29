package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
	"github.com/booscaaa/initializers/postgres"
)

type listClientsDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.ListClientsDatabaseRepository.
func (repository *listClientsDatabaseRepository) Execute(ctx context.Context) ([]model.Client, error) {
	var clients []model.Client
	query := "SELECT id, name FROM client ORDER BY name"
	err := repository.database.SelectContext(
		ctx,
		&clients,
		query,
	)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func NewListClientsDatabaseRepository(database postgres.Database) contract.ListClientsDatabaseRepository {
	return &listClientsDatabaseRepository{
		database: database,
	}
}
