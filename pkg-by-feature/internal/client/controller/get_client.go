package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/go-chi/chi/v5"
)

type getClientController struct {
	getClientUseCase contract.GetClientUseCase
}

// Execute implements contract.GetClientController.
func (controller *getClientController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		http.Error(response, "ID is required", http.StatusBadRequest)
		return
	}

	output, err := controller.getClientUseCase.Execute(ctx, id)
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

func NewGetClientController(getClientUseCase contract.GetClientUseCase) contract.GetClientController {
	return &getClientController{
		getClientUseCase: getClientUseCase,
	}
}
