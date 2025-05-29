package rest

import (
	"fmt"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/pkg/di"
	"github.com/booscaaa/initializers/postgres"
	"github.com/go-chi/chi"
)

func Initialize(database postgres.Database) {
	createClientController := di.NewCreateClientController(database)
	createProductController := di.NewCreateProductController(database)
	createBudgetController := di.NewCreateBudgetController(database)
	createServiceController := di.NewCreateServiceController(database)

	router := chi.NewRouter()

	router.Post("/clients", createClientController.Execute)
	router.Post("/products", createProductController.Execute)
	router.Post("/budgets", createBudgetController.Execute)
	router.Post("/services", createServiceController.Execute)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", router)
}
