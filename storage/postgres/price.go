package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
)

type priceRepo struct {
	db *sqlx.DB
}

func NewPriceRepo(db *sqlx.DB) repo.PriceStorageI {
	return &priceRepo{db: db}
}

func (cm *priceRepo) Create(pr *repo.Price) (string, error) {
	insertQuery := `insert into product_prices(
		id,
		price,
		product_id,
		retailer_id
	)
		values($1, $2, $3, $4)`

	_, err := cm.db.Exec(insertQuery, pr.Id, pr.Price, pr.ProductId, pr.RetailerId)

	if err != nil {
		return "", err
	}
	return pr.Id, nil
}

func (cm *priceRepo) Update(pr *repo.Price) (string, error) {

	updateQuery := `update product_prices set
	price=$1,
	product_id=$2,
	retailer_id=$3 where id=$4`

	_, err := cm.db.Exec(updateQuery, pr.Price, pr.ProductId, pr.RetailerId, pr.Id)
	if err != nil {
		return "", err
	}
	return pr.Id, nil
}

func (cm *priceRepo) Get(id string) (*repo.Price, error) {
	var price repo.Price

	query := `select 
		id,
		price,
		product_id,
		retailer_id
	from product_prices where id=$1`

	row := cm.db.QueryRow(query, id)

	err := row.Scan(
		&price.Id,
		&price.Price,
		&price.ProductId,
		&price.RetailerId,
	)

	if err != nil {
		return nil, err
	}

	return &price, nil
}

func (cm *priceRepo) GetAll(price string) ([]*repo.Price, error) {
	var (
		productPrices []*repo.Price
		filter        string
		args          = make(map[string]interface{})
	)
	if price != "" {
		filter += ` and price ilike '%' || :price || '%'`
		args["price"] = price
	}
	query := `select 
		id,
		price,
		product_id,
		retailer_id
	from product_prices where 1=1 %s`

	rows, err := cm.db.NamedQuery(fmt.Sprintf(query, filter), args)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			price repo.Price
		)
		err = rows.Scan(
			&price.Id,
			&price.Price,
			&price.ProductId,
			&price.RetailerId,
		)
		if err != nil {
			return nil, err
		}
		productPrices = append(productPrices, &price)
	}

	return productPrices, nil
}

func (cm *priceRepo) Delete(id string) error {
	return nil
}
