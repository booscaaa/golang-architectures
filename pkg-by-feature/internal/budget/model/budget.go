package model

type Budget struct {
	ID        string `db:"id"`
	ClientID  string `db:"client_id"`
	ProductID string `db:"product_id"`
}

type CreateBudgetInput struct {
	ClientID  string `json:"client_id"`
	ProductID string `json:"product_id"`
}

type CreateBudgetOutput struct {
	ID        string `json:"id"`
	ClientID  string `json:"client_id"`
	ProductID string `json:"product_id"`
}

func (input *CreateBudgetInput) ToBudget() Budget {
	return Budget{
		ClientID:  input.ClientID,
		ProductID: input.ProductID,
	}
}

func (budget Budget) ToCreateBudgetOutput() CreateBudgetOutput {
	return CreateBudgetOutput{
		ID:        budget.ID,
		ClientID:  budget.ClientID,
		ProductID: budget.ProductID,
	}
}
