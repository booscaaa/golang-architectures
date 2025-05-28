package di

import (
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/controller"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/repository"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/usecase"
	"github.com/booscaaa/initializers/postgres"
)

func NewCreateBudgetController(database postgres.Database) contract.CreateBudgetController {
	return controller.NewCreateBudgetController(
		usecase.NewCreateBudgetUseCase(
			repository.NewCreateBudgetDatabaseRepository(database),
		),
	)
}
