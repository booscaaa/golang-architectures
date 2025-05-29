package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
	"github.com/booscaaa/initializers/postgres"
)

type listProductsDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.ListProductsDatabaseRepository.
func (repository *listProductsDatabaseRepository) Execute(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	query := `SELECT id, name, price FROM product ORDER BY name`
	err := repository.database.SelectContext(
		ctx,
		&products,
		query,
	)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func NewListProductsDatabaseRepository(database postgres.Database) contract.ListProductsDatabaseRepository {
	return &listProductsDatabaseRepository{
		database: database,
	}
}
