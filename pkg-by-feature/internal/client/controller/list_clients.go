package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
)

type listClientsController struct {
	listClientsUseCase contract.ListClientsUseCase
}

// Execute implements contract.ListClientsController.
func (controller *listClientsController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	output, err := controller.listClientsUseCase.Execute(ctx)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewListClientsController(listClientsUseCase contract.ListClientsUseCase) contract.ListClientsController {
	return &listClientsController{
		listClientsUseCase: listClientsUseCase,
	}
}
