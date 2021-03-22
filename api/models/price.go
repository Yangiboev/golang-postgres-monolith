package models

type CreatePriceRequest struct {
	Price      int64  `json:"price" binding:"required"`
	ProductId  string `json:"product_id"`
	RetailerId string `json:"retailer_id"`
}
type UpdatePriceRequest struct {
	Price      int64  `json:"price"`
	ProductId  string `json:"product_id"`
	RetailerId string `json:"retailer_id"`
}

type GetPrice struct {
	Success string `json:"true"`
	Prices  Price  `json:"data"`
}

type Price struct {
	Id         string `json:"id"`
	Price      int64  `json:"price"`
	ProductId  string `json:"product_id"`
	RetailerId string `json:"retailer_id"`
}

type GetAllPricesResponse struct {
	Success bool       `json:"success"`
	Prices  []GetPrice `json:"data"`
}
