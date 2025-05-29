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

func NewGetServiceController(database postgres.Database) contract.GetServiceController {
	return controller.NewGetServiceController(
		usecase.NewGetServiceUseCase(
			repository.NewGetServiceDatabaseRepository(database),
		),
	)
}

func NewUpdateServiceController(database postgres.Database) contract.UpdateServiceController {
	return controller.NewUpdateServiceController(
		usecase.NewUpdateServiceUseCase(
			repository.NewUpdateServiceDatabaseRepository(database),
		),
	)
}

func NewDeleteServiceController(database postgres.Database) contract.DeleteServiceController {
	return controller.NewDeleteServiceController(
		usecase.NewDeleteServiceUseCase(
			repository.NewDeleteServiceDatabaseRepository(database),
		),
	)
}

func NewListServicesController(database postgres.Database) contract.ListServicesController {
	return controller.NewListServicesController(
		usecase.NewListServicesUseCase(
			repository.NewListServicesDatabaseRepository(database),
		),
	)
}
