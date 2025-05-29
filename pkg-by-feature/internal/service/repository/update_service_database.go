package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
	"github.com/booscaaa/initializers/postgres"
)

type updateServiceDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.UpdateServiceDatabaseRepository.
func (repository *updateServiceDatabaseRepository) Execute(ctx context.Context, service model.Service) (model.Service, error) {
	query := `UPDATE service SET name = $2, description = $3, price = $4 WHERE id = $1 RETURNING id, name, description, price`
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		service.ID,
		service.Name,
		service.Description,
		service.Price,
	).StructScan(&service)
	if err != nil {
		return model.Service{}, err
	}
	return service, nil
}

func NewUpdateServiceDatabaseRepository(database postgres.Database) contract.UpdateServiceDatabaseRepository {
	return &updateServiceDatabaseRepository{
		database: database,
	}
}
