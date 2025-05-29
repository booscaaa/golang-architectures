package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
)

// Repository interfaces
type CreateServiceDatabaseRepository interface {
	Execute(context.Context, model.Service) (model.Service, error)
}

type GetServiceDatabaseRepository interface {
	Execute(context.Context, string) (model.Service, error)
}

type UpdateServiceDatabaseRepository interface {
	Execute(context.Context, model.Service) (model.Service, error)
}

type DeleteServiceDatabaseRepository interface {
	Execute(context.Context, string) error
}

type ListServicesDatabaseRepository interface {
	Execute(context.Context) ([]model.Service, error)
}

// UseCase interfaces
type CreateServiceUseCase interface {
	Execute(context.Context, model.CreateServiceInput) (model.CreateServiceOutput, error)
}

type GetServiceUseCase interface {
	Execute(context.Context, string) (model.GetServiceOutput, error)
}

type UpdateServiceUseCase interface {
	Execute(context.Context, string, model.UpdateServiceInput) (model.UpdateServiceOutput, error)
}

type DeleteServiceUseCase interface {
	Execute(context.Context, string) error
}

type ListServicesUseCase interface {
	Execute(context.Context) (model.ListServicesOutput, error)
}

// Controller interfaces
type CreateServiceController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type GetServiceController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpdateServiceController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type DeleteServiceController interface {
	Execute(http.ResponseWriter, *http.Request)
}

type ListServicesController interface {
	Execute(http.ResponseWriter, *http.Request)
}
