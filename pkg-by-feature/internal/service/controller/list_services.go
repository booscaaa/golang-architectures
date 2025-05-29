package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
)

type listServicesController struct {
	listServicesUseCase contract.ListServicesUseCase
}

// Execute implements contract.ListServicesController.
func (controller *listServicesController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	output, err := controller.listServicesUseCase.Execute(ctx)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}

func NewListServicesController(listServicesUseCase contract.ListServicesUseCase) contract.ListServicesController {
	return &listServicesController{
		listServicesUseCase: listServicesUseCase,
	}
}
