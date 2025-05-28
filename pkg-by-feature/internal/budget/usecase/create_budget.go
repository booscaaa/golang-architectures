package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/model"
)

type createBudgetUseCase struct {
	repository contract.CreateBudgetDatabaseRepository
}

// Execute implements contract.CreateBudgetUseCase.
func (usecase *createBudgetUseCase) Execute(ctx context.Context, input model.CreateBudgetInput) (model.CreateBudgetOutput, error) {
	budget, err := usecase.repository.Execute(ctx, input.ToBudget())
	if err != nil {
		return model.CreateBudgetOutput{}, err
	}
	return budget.ToCreateBudgetOutput(), nil
}

func NewCreateBudgetUseCase(repository contract.CreateBudgetDatabaseRepository) contract.CreateBudgetUseCase {
	return &createBudgetUseCase{
		repository: repository,
	}
}
