package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
	"github.com/go-chi/chi/v5"
)

type updateClientController struct {
	updateClientUseCase contract.UpdateClientUseCase
}

// Execute implements contract.UpdateClientController.
func (controller *updateClientController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(response, "ID is required", http.StatusBadRequest)
		return
	}

	var input model.UpdateClientInput
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.updateClientUseCase.Execute(ctx, id, input)
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

func NewUpdateClientController(updateClientUseCase contract.UpdateClientUseCase) contract.UpdateClientController {
	return &updateClientController{
		updateClientUseCase: updateClientUseCase,
	}
}
