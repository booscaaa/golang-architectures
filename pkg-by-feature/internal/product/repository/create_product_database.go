package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
	"github.com/booscaaa/initializers/postgres"
)

type createProductDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.CreateProductDatabaseRepository.
func (repository *createProductDatabaseRepository) Execute(ctx context.Context, product model.Product) (model.Product, error) {
	var productCreated model.Product
	query := `INSERT INTO product (name, price) VALUES ($1, $2) RETURNING id, name, price`
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		product.Name,
		product.Price,
	).StructScan(&productCreated)
	if err != nil {
		return model.Product{}, err
	}
	return productCreated, nil
}

func NewCreateProductDatabaseRepository(database postgres.Database) contract.CreateProductDatabaseRepository {
	return &createProductDatabaseRepository{
		database: database,
	}
}
