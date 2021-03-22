package repo

type Price struct {
	Id         string `json:"id"`
	Price      int64  `json:"price"`
	ProductId  string `json:"product_id"`
	RetailerId string `json:"retailer_id"`
}

type PriceStorageI interface {
	Create(*Price) (string, error)
	Update(*Price) (string, error)
	Get(id string) (*Price, error)
	GetAll(name string) ([]*Price, error)
	Delete(id string) error
}
