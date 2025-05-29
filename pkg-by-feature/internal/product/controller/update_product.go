package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/model"
	"github.com/go-chi/chi/v5"
)

type updateProductController struct {
	updateProductUseCase contract.UpdateProductUseCase
}

// Execute implements contract.UpdateProductController.
func (controller *updateProductController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var input model.UpdateProductInput
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := controller.updateProductUseCase.Execute(ctx, id, input)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}

func NewUpdateProductController(updateProductUseCase contract.UpdateProductUseCase) contract.UpdateProductController {
	return &updateProductController{
		updateProductUseCase: updateProductUseCase,
	}
}
