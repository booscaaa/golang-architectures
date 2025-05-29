package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
)

// Create Client interfaces
type CreateClientDatabaseRepository interface {
	Execute(context.Context, model.Client) (model.Client, error)
}

type CreateClientUseCase interface {
	Execute(context.Context, model.CreateClientInput) (model.CreateClientOutput, error)
}

type CreateClientController interface {
	Execute(http.ResponseWriter, *http.Request)
}

// Get Client interfaces
type GetClientDatabaseRepository interface {
	Execute(context.Context, string) (model.Client, error)
}

type GetClientUseCase interface {
	Execute(context.Context, string) (model.GetClientOutput, error)
}

type GetClientController interface {
	Execute(http.ResponseWriter, *http.Request)
}

// Update Client interfaces
type UpdateClientDatabaseRepository interface {
	Execute(context.Context, model.Client) (model.Client, error)
}

type UpdateClientUseCase interface {
	Execute(context.Context, string, model.UpdateClientInput) (model.UpdateClientOutput, error)
}

type UpdateClientController interface {
	Execute(http.ResponseWriter, *http.Request)
}

// Delete Client interfaces
type DeleteClientDatabaseRepository interface {
	Execute(context.Context, string) error
}

type DeleteClientUseCase interface {
	Execute(context.Context, string) error
}

type DeleteClientController interface {
	Execute(http.ResponseWriter, *http.Request)
}

// List Clients interfaces
type ListClientsDatabaseRepository interface {
	Execute(context.Context) ([]model.Client, error)
}

type ListClientsUseCase interface {
	Execute(context.Context) (model.ListClientsOutput, error)
}

type ListClientsController interface {
	Execute(http.ResponseWriter, *http.Request)
}
