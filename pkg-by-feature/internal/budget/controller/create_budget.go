package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/budget/model"
)

type createBudgetController struct {
	createBudgetUseCase contract.CreateBudgetUseCase
}

// Execute implements contract.CreateBudgetController.
func (controller *createBudgetController) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var input model.CreateBudgetInput
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.createBudgetUseCase.Execute(ctx, input)
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

func NewCreateBudgetController(createBudgetUseCase contract.CreateBudgetUseCase) contract.CreateBudgetController {
	return &createBudgetController{
		createBudgetUseCase: createBudgetUseCase,
	}
}
