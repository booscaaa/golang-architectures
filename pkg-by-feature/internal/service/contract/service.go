package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
)

type CreateServiceDatabaseRepository interface {
	Execute(context.Context, model.Service) (model.Service, error)
}

type CreateServiceUseCase interface {
	Execute(context.Context, model.CreateServiceInput) (model.CreateServiceOutput, error)
}

type CreateServiceController interface {
	Execute(http.ResponseWriter, *http.Request)
}
