package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
)

type listProductsController struct {
	listProductsUseCase contract.ListProductsUseCase
}

// Execute implements contract.ListProductsController.
func (controller *listProductsController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	output, err := controller.listProductsUseCase.Execute(ctx)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}

func NewListProductsController(listProductsUseCase contract.ListProductsUseCase) contract.ListProductsController {
	return &listProductsController{
		listProductsUseCase: listProductsUseCase,
	}
}
