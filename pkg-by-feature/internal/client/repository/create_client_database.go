package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
	"github.com/booscaaa/initializers/postgres"
)

type createClientDatabaseRepository struct {
	database postgres.Database
}

// Create implements contract.CreateClientDatabaseRepository.
func (repository *createClientDatabaseRepository) Execute(ctx context.Context, client model.Client) (model.Client, error) {
	var clientCreated model.Client
	query := `INSERT INTO client (name) VALUES ($1) RETURNING id, name;`
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		client.Name,
	).StructScan(&clientCreated)
	if err != nil {
		return model.Client{}, err
	}
	return clientCreated, nil
}

func NewCreateClientDatabaseRepository(database postgres.Database) contract.CreateClientDatabaseRepository {
	return &createClientDatabaseRepository{
		database: database,
	}
}
