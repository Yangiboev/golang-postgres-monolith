package models

type CreateProductRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CategoryId  string `json:"category_id"`
}
type UpdateProductRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CategoryId  string `json:"category_id"`
}

type GetProductResponse struct {
	Success string     `json:"true"`
	Product GetProduct `json:"data"`
}

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CategoryId  string `json:"category_id"`
}
type GetProduct struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
}

type GetAllProductsResponse struct {
	Success  bool      `json:"success"`
	Products []Product `json:"data"`
}
