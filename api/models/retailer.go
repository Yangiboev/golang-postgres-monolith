package models

type CreateRetailerRequest struct {
	Name        string `json:"name"`
	Website     string `json:"website"`
	Description string `json:"description"`
}
type UpdateRetailerRequest struct {
	Name        string `json:"name"`
	Website     string `json:"website"`
	Description string `json:"description"`
}

type GetRetailerResponse struct {
	Success  string      `json:"true"`
	Retailer GetRetailer `json:"data"`
}

type Retailer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Website     string `json:"website"`
	Description string `json:"description"`
}
type GetRetailer struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Website     string   `json:"website"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
}

type GetAllRetailersResponse struct {
	Success   bool       `json:"success"`
	Retailers []Retailer `json:"data"`
}
