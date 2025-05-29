package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
	"github.com/booscaaa/initializers/postgres"
)

type getServiceDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.GetServiceDatabaseRepository.
func (repository *getServiceDatabaseRepository) Execute(ctx context.Context, id string) (model.Service, error) {
	var service model.Service
	query := `SELECT id, name, description, price FROM service WHERE id = $1`
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		id,
	).StructScan(&service)
	if err != nil {
		return model.Service{}, err
	}
	return service, nil
}

func NewGetServiceDatabaseRepository(database postgres.Database) contract.GetServiceDatabaseRepository {
	return &getServiceDatabaseRepository{
		database: database,
	}
}
