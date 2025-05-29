package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
	"github.com/go-chi/chi/v5"
)

type updateServiceController struct {
	updateServiceUseCase contract.UpdateServiceUseCase
}

// Execute implements contract.UpdateServiceController.
func (controller *updateServiceController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var input model.UpdateServiceInput
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := controller.updateServiceUseCase.Execute(ctx, id, input)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}

func NewUpdateServiceController(updateServiceUseCase contract.UpdateServiceUseCase) contract.UpdateServiceController {
	return &updateServiceController{
		updateServiceUseCase: updateServiceUseCase,
	}
}
