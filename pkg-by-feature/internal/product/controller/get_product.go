package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/go-chi/chi/v5"
)

type getProductController struct {
	getProductUseCase contract.GetProductUseCase
}

// Execute implements contract.GetProductController.
func (controller *getProductController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := controller.getProductUseCase.Execute(ctx, id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}

func NewGetProductController(getProductUseCase contract.GetProductUseCase) contract.GetProductController {
	return &getProductController{
		getProductUseCase: getProductUseCase,
	}
}
