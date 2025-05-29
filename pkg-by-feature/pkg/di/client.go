package di

import (
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/controller"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/repository"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/usecase"
	"github.com/booscaaa/initializers/postgres"
)

// Create Client DI
func NewCreateClientController(database postgres.Database) contract.CreateClientController {
	return controller.NewCreateClientController(
		usecase.NewCreateClientUseCase(
			repository.NewCreateClientDatabaseRepository(database),
		),
	)
}

// Get Client DI
func NewGetClientController(database postgres.Database) contract.GetClientController {
	return controller.NewGetClientController(
		usecase.NewGetClientUseCase(
			repository.NewGetClientDatabaseRepository(database),
		),
	)
}

// Update Client DI
func NewUpdateClientController(database postgres.Database) contract.UpdateClientController {
	return controller.NewUpdateClientController(
		usecase.NewUpdateClientUseCase(
			repository.NewUpdateClientDatabaseRepository(database),
		),
	)
}

// Delete Client DI
func NewDeleteClientController(database postgres.Database) contract.DeleteClientController {
	return controller.NewDeleteClientController(
		usecase.NewDeleteClientUseCase(
			repository.NewDeleteClientDatabaseRepository(database),
		),
	)
}

// List Clients DI
func NewListClientsController(database postgres.Database) contract.ListClientsController {
	return controller.NewListClientsController(
		usecase.NewListClientsUseCase(
			repository.NewListClientsDatabaseRepository(database),
		),
	)
}
