package models

type CreateCategoryRequest struct {
	Name     string `json:"name" binding:"required"`
	ParentId string `json:"parent_id"`
}
type UpdateCategoryRequest struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type GetCategory struct {
	Success    string   `json:"true"`
	Categories Category `json:"data"`
}

type Category struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type GetAllCategoriesResponse struct {
	Success    bool       `json:"success"`
	Categories []Category `json:"data"`
}
