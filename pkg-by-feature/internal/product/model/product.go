package model

type Product struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Price int    `db:"price"`
}

type CreateProductOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (product Product) ToCreateProductOutput() CreateProductOutput {
	return CreateProductOutput{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func (input CreateProductInput) ToProduct() Product {
	return Product{
		Name:  input.Name,
		Price: input.Price,
	}
}
