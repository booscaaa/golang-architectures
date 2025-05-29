package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
	"github.com/booscaaa/initializers/postgres"
)

type getProductDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.GetProductDatabaseRepository.
func (repository *getProductDatabaseRepository) Execute(ctx context.Context, id string) (model.Product, error) {
	var product model.Product
	query := `SELECT id, name, price FROM product WHERE id = $1`
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		id,
	).StructScan(&product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func NewGetProductDatabaseRepository(database postgres.Database) contract.GetProductDatabaseRepository {
	return &getProductDatabaseRepository{
		database: database,
	}
}
