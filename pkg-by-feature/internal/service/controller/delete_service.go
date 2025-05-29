package controller

import (
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/service/contract"
	"github.com/go-chi/chi/v5"
)

type deleteServiceController struct {
	deleteServiceUseCase contract.DeleteServiceUseCase
}

// Execute implements contract.DeleteServiceController.
func (controller *deleteServiceController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	id := chi.URLParam(request, "id")

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err := controller.deleteServiceUseCase.Execute(ctx, id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func NewDeleteServiceController(deleteServiceUseCase contract.DeleteServiceUseCase) contract.DeleteServiceController {
	return &deleteServiceController{
		deleteServiceUseCase: deleteServiceUseCase,
	}
}
