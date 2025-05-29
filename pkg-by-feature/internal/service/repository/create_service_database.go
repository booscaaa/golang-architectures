package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
	"github.com/booscaaa/initializers/postgres"
)

type createServiceDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.CreateServiceDatabaseRepository.
func (repository *createServiceDatabaseRepository) Execute(ctx context.Context, service model.Service) (model.Service, error) {
	var serviceCreated model.Service
	query := "INSERT INTO service (name, description, price) VALUES ($1, $2, $3) RETURNING id, name, description, price"
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		service.Name,
		service.Description,
		service.Price,
	).StructScan(&serviceCreated)
	if err != nil {
		return model.Service{}, err
	}

	return serviceCreated, nil
}

func NewCreateServiceDatabaseRepository(database postgres.Database) contract.CreateServiceDatabaseRepository {
	return &createServiceDatabaseRepository{
		database: database,
	}
}
