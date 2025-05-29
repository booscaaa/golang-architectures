package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/go-chi/chi/v5"
)

type getServiceController struct {
	getServiceUseCase contract.GetServiceUseCase
}

// Execute implements contract.GetServiceController.
func (controller *getServiceController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := controller.getServiceUseCase.Execute(ctx, id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}

func NewGetServiceController(getServiceUseCase contract.GetServiceUseCase) contract.GetServiceController {
	return &getServiceController{
		getServiceUseCase: getServiceUseCase,
	}
}
