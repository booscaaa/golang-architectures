package di

import (
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/controller"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/repository"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/usecase"
	"github.com/booscaaa/initializers/postgres"
)

func NewCreateServiceController(database postgres.Database) contract.CreateServiceController {
	return controller.NewCreateServiceController(
		usecase.NewCreateServiceUseCase(
			repository.NewCreateServiceDatabaseRepository(database),
		),
	)
}
