package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
	"github.com/booscaaa/initializers/postgres"
)

type updateClientDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.UpdateClientDatabaseRepository.
func (repository *updateClientDatabaseRepository) Execute(ctx context.Context, client model.Client) (model.Client, error) {
	var clientUpdated model.Client
	query := "UPDATE client SET name = $2 WHERE id = $1 RETURNING *"
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		client.ID,
		client.Name,
	).StructScan(&clientUpdated)
	if err != nil {
		return model.Client{}, err
	}

	return clientUpdated, nil
}

func NewUpdateClientDatabaseRepository(database postgres.Database) contract.UpdateClientDatabaseRepository {
	return &updateClientDatabaseRepository{
		database: database,
	}
}
