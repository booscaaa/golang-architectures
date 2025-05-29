package controller

import (
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/product/contract"
	"github.com/go-chi/chi/v5"
)

type deleteProductController struct {
	deleteProductUseCase contract.DeleteProductUseCase
}

// Execute implements contract.DeleteProductController.
func (controller *deleteProductController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err := controller.deleteProductUseCase.Execute(ctx, id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func NewDeleteProductController(deleteProductUseCase contract.DeleteProductUseCase) contract.DeleteProductController {
	return &deleteProductController{
		deleteProductUseCase: deleteProductUseCase,
	}
}
