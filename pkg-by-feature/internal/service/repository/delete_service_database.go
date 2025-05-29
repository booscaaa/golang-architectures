package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/initializers/postgres"
)

type deleteServiceDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.DeleteServiceDatabaseRepository.
func (repository *deleteServiceDatabaseRepository) Execute(ctx context.Context, id string) error {
	query := `DELETE FROM service WHERE id = $1`
	_, err := repository.database.ExecContext(ctx, query, id)
	return err
}

func NewDeleteServiceDatabaseRepository(database postgres.Database) contract.DeleteServiceDatabaseRepository {
	return &deleteServiceDatabaseRepository{
		database: database,
	}
}
