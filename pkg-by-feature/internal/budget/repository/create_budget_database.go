package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/model"
	"github.com/booscaaa/initializers/postgres"
)

type createBudgetDatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.CreateBudgetDatabaseRepository.
func (repository *createBudgetDatabaseRepository) Execute(ctx context.Context, budget model.Budget) (model.Budget, error) {
	var budgetCreated model.Budget
	query := "INSERT INTO budget (client_id, product_id) VALUES ($1, $2) RETURNING *"
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		budget.ClientID,
		budget.ProductID,
	).StructScan(&budgetCreated)
	if err != nil {
		return model.Budget{}, err
	}

	return budgetCreated, nil
}

func NewCreateBudgetDatabaseRepository(database postgres.Database) contract.CreateBudgetDatabaseRepository {
	return &createBudgetDatabaseRepository{
		database: database,
	}
}
