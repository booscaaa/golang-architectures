package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
	"github.com/booscaaa/initializers/postgres"
)

type getClientDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.GetClientDatabaseRepository.
func (repository *getClientDatabaseRepository) Execute(ctx context.Context, id string) (model.Client, error) {
	var client model.Client
	query := "SELECT id, name FROM client WHERE id = $1"
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		id,
	).StructScan(&client)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func NewGetClientDatabaseRepository(database postgres.Database) contract.GetClientDatabaseRepository {
	return &getClientDatabaseRepository{
		database: database,
	}
}
