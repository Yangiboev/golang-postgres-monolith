package repo

type Retailer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Website     string `json:"website"`
	Description string `json:"description"`
}

type RetailerStorageI interface {
	Create(*Retailer) (string, error)
	Update(*Retailer) (string, error)
	Get(id string) (*Retailer, error)
	GetAll(name string) ([]*Retailer, error)
	Delete(id string) error
}
