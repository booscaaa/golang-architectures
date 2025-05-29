package rest

import (
	"fmt"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/pkg/di"
	"github.com/booscaaa/initializers/postgres"
	"github.com/go-chi/chi"
)

func Initialize(database postgres.Database) {
	// Client controllers
	createClientController := di.NewCreateClientController(database)
	getClientController := di.NewGetClientController(database)
	updateClientController := di.NewUpdateClientController(database)
	deleteClientController := di.NewDeleteClientController(database)
	listClientsController := di.NewListClientsController(database)

	// Other controllers
	createProductController := di.NewCreateProductController(database)
	createBudgetController := di.NewCreateBudgetController(database)
	createServiceController := di.NewCreateServiceController(database)

	router := chi.NewRouter()

	// Client routes - CRUD complete
	router.Post("/clients", createClientController.Execute)
	router.Get("/clients", listClientsController.Execute)
	router.Get("/clients/{id}", getClientController.Execute)
	router.Put("/clients/{id}", updateClientController.Execute)
	router.Delete("/clients/{id}", deleteClientController.Execute)

	// Other routes
	router.Post("/products", createProductController.Execute)
	router.Post("/budgets", createBudgetController.Execute)
	router.Post("/services", createServiceController.Execute)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", router)
}
