package controller

import (
	"encoding/json"
	"fmt"
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
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.createServiceUseCase.Execute(ctx, input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(input)

	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewCreateServiceController(createServiceUseCase contract.CreateServiceUseCase) contract.CreateServiceController {
	return &createServiceController{
		createServiceUseCase: createServiceUseCase,
	}
}
