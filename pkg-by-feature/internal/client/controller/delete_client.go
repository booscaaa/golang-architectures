package controller

import (
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/go-chi/chi/v5"
)

type deleteClientController struct {
	deleteClientUseCase contract.DeleteClientUseCase
}

// Execute implements contract.DeleteClientController.
func (controller *deleteClientController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(response, "ID is required", http.StatusBadRequest)
		return
	}

	err := controller.deleteClientUseCase.Execute(ctx, id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func NewDeleteClientController(deleteClientUseCase contract.DeleteClientUseCase) contract.DeleteClientController {
	return &deleteClientController{
		deleteClientUseCase: deleteClientUseCase,
	}
}
