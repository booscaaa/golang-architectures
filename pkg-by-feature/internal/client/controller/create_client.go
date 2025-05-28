package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/client/model"
)

type createClientController struct {
	createClientUseCase contract.CreateClientUseCase
}

// Execute implements contract.CreateClientController.
func (controller *createClientController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var input model.CreateClientInput
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := controller.createClientUseCase.Execute(ctx, input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func NewCreateClientController(createClientUseCase contract.CreateClientUseCase) contract.CreateClientController {
	return &createClientController{
		createClientUseCase: createClientUseCase,
	}
}
