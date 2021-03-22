package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/Yangiboev/golang-postgres-monolith/storage/postgres"
	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
)

type StorageI interface {
	Category() repo.CategoryStorageI
	Product() repo.ProductStorageI
	Retailer() repo.RetailerStorageI
	Price() repo.PriceStorageI
}

type storagePg struct {
	db           *sqlx.DB
	categoryRepo repo.CategoryStorageI
	retailerRepo repo.RetailerStorageI
	productRepo  repo.ProductStorageI
	priceRepo    repo.PriceStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:           db,
		categoryRepo: postgres.NewCategoryRepo(db),
		productRepo:  postgres.NewProductRepo(db),
		retailerRepo: postgres.NewRetailerRepo(db),
		priceRepo:    postgres.NewPriceRepo(db),
	}
}

func (s storagePg) Category() repo.CategoryStorageI {
	return s.categoryRepo
}
func (s storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
func (s storagePg) Price() repo.PriceStorageI {
	return s.priceRepo
}
func (s storagePg) Retailer() repo.RetailerStorageI {
	return s.retailerRepo
}
