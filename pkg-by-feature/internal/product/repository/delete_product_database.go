package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/initializers/postgres"
)

type deleteProductDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.DeleteProductDatabaseRepository.
func (repository *deleteProductDatabaseRepository) Execute(ctx context.Context, id string) error {
	query := `DELETE FROM product WHERE id = $1`
	_, err := repository.database.ExecContext(
		ctx,
		query,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteProductDatabaseRepository(database postgres.Database) contract.DeleteProductDatabaseRepository {
	return &deleteProductDatabaseRepository{
		database: database,
	}
}
