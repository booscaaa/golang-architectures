package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
)

type createProductController struct {
	createProductUseCase contract.CreateProductUseCase
}

// Execute implements contract.CreateProductController.
func (controller *createProductController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var input model.CreateProductInput
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := controller.createProductUseCase.Execute(ctx, input)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func NewCreateProductController(createProductUseCase contract.CreateProductUseCase) contract.CreateProductController {
	return &createProductController{
		createProductUseCase: createProductUseCase,
	}
}
