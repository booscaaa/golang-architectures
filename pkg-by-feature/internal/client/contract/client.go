package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
)

type CreateClientDatabaseRepository interface {
	Execute(context.Context, model.Client) (model.Client, error)
}

type CreateClientUseCase interface {
	Execute(context.Context, model.CreateClientInput) (model.CreateClientOutput, error)
}

type CreateClientController interface {
	Execute(http.ResponseWriter, *http.Request)
}
