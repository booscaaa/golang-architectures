package model

type Product struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Price int    `db:"price"`
}

// Create Product types
type CreateProductOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Get Product types
type GetProductOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Update Product types
type UpdateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UpdateProductOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// List Products types
type ListProductsOutput struct {
	Products []ProductItem `json:"products"`
}

type ProductItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Conversion methods
func (product Product) ToCreateProductOutput() CreateProductOutput {
	return CreateProductOutput{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func (product Product) ToGetProductOutput() GetProductOutput {
	return GetProductOutput{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func (product Product) ToUpdateProductOutput() UpdateProductOutput {
	return UpdateProductOutput{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func (product Product) ToProductItem() ProductItem {
	return ProductItem{
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

func (input UpdateProductInput) ToProduct(id string) Product {
	return Product{
		ID:    id,
		Name:  input.Name,
		Price: input.Price,
	}
}

func ProductsToListProductsOutput(products []Product) ListProductsOutput {
	productItems := make([]ProductItem, len(products))
	for i, product := range products {
		productItems[i] = product.ToProductItem()
	}
	return ListProductsOutput{
		Products: productItems,
	}
}
