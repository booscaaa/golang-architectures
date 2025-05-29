package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/model"
)

type createServiceController struct {
	createServiceUseCase contract.CreateServiceUseCase
}

// Execute implements contract.CreateServiceController.
func (controller *createServiceController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var input model.CreateServiceInput
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := controller.createServiceUseCase.Execute(ctx, input)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func NewCreateServiceController(createServiceUseCase contract.CreateServiceUseCase) contract.CreateServiceController {
	return &createServiceController{
		createServiceUseCase: createServiceUseCase,
	}
}
