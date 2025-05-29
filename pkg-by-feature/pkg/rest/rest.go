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

	// Product controllers
	createProductController := di.NewCreateProductController(database)
	getProductController := di.NewGetProductController(database)
	updateProductController := di.NewUpdateProductController(database)
	deleteProductController := di.NewDeleteProductController(database)
	listProductsController := di.NewListProductsController(database)

	// Other controllers
	createBudgetController := di.NewCreateBudgetController(database)
	createServiceController := di.NewCreateServiceController(database)

	router := chi.NewRouter()

	// Client routes - CRUD complete
	router.Post("/clients", createClientController.Execute)
	router.Get("/clients", listClientsController.Execute)
	router.Get("/clients/{id}", getClientController.Execute)
	router.Put("/clients/{id}", updateClientController.Execute)
	router.Delete("/clients/{id}", deleteClientController.Execute)

	// Product routes - CRUD complete
	router.Post("/products", createProductController.Execute)
	router.Get("/products", listProductsController.Execute)
	router.Get("/products/{id}", getProductController.Execute)
	router.Put("/products/{id}", updateProductController.Execute)
	router.Delete("/products/{id}", deleteProductController.Execute)

	// Other routes
	router.Post("/budgets", createBudgetController.Execute)
	router.Post("/services", createServiceController.Execute)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", router)
}
