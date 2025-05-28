package di

import (
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/controller"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/repository"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/usecase"
	"github.com/booscaaa/initializers/postgres"
)

func NewCreateClientController(database postgres.Database) contract.CreateClientController {
	return controller.NewCreateClientController(
		usecase.NewCreateClientUseCase(
			repository.NewCreateClientDatabaseRepository(database),
		),
	)
}
