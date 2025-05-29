package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
	"github.com/booscaaa/initializers/postgres"
)

type listServicesDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.ListServicesDatabaseRepository.
func (repository *listServicesDatabaseRepository) Execute(ctx context.Context) ([]model.Service, error) {
	var services []model.Service
	query := `SELECT id, name, description, price FROM service ORDER BY name`
	err := repository.database.SelectContext(ctx, &services, query)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func NewListServicesDatabaseRepository(database postgres.Database) contract.ListServicesDatabaseRepository {
	return &listServicesDatabaseRepository{
		database: database,
	}
}
