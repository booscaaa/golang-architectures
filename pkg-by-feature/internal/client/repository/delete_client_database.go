package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/initializers/postgres"
)

type deleteClientDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.DeleteClientDatabaseRepository.
func (repository *deleteClientDatabaseRepository) Execute(ctx context.Context, id string) error {
	query := "DELETE FROM client WHERE id = $1"
	_, err := repository.database.ExecContext(
		ctx,
		query,
		id,
	)
	return err
}

func NewDeleteClientDatabaseRepository(database postgres.Database) contract.DeleteClientDatabaseRepository {
	return &deleteClientDatabaseRepository{
		database: database,
	}
}
