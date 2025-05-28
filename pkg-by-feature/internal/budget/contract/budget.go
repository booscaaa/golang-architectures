package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/model"
)

type CreateBudgetDatabaseRepository interface {
	Execute(context.Context, model.Budget) (model.Budget, error)
}

type CreateBudgetUseCase interface {
	Execute(context.Context, model.CreateBudgetInput) (model.CreateBudgetOutput, error)
}

type CreateBudgetController interface {
	Execute(http.ResponseWriter, *http.Request)
}
