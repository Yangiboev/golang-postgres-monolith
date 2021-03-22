package repo

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CategoryId  string `json:"category_id"`
}

type ProductStorageI interface {
	Create(*Product) (string, error)
	Update(*Product) (string, error)
	Get(id string) (*Product, error)
	GetAll(name string) ([]*Product, error)
	Delete(id string) error
}
