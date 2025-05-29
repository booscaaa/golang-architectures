package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
	"github.com/booscaaa/initializers/postgres"
)

type updateProductDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.UpdateProductDatabaseRepository.
func (repository *updateProductDatabaseRepository) Execute(ctx context.Context, product model.Product) (model.Product, error) {
	var productUpdated model.Product
	query := `UPDATE product SET name = $1, price = $2 WHERE id = $3 RETURNING id, name, price`
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		product.Name,
		product.Price,
		product.ID,
	).StructScan(&productUpdated)
	if err != nil {
		return model.Product{}, err
	}
	return productUpdated, nil
}

func NewUpdateProductDatabaseRepository(database postgres.Database) contract.UpdateProductDatabaseRepository {
	return &updateProductDatabaseRepository{
		database: database,
	}
}
