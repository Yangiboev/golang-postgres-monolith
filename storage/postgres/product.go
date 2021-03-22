package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) repo.ProductStorageI {
	return &productRepo{db: db}
}

func (cm *productRepo) Create(cl *repo.Product) (string, error) {
	insertQuery := `insert into products(
		id,
		name,
		image,
		description,
		category_id)
		values($1, $2, $3, $4, $5)`

	_, err := cm.db.Exec(insertQuery, cl.Id, cl.Name, cl.Image, cl.Description, cl.CategoryId)

	if err != nil {
		return "", err
	}
	return cl.Id, nil
}

func (cm *productRepo) Update(cl *repo.Product) (string, error) {
	updateQuery := `update products set
	name=$1,
	image=$2,
	description=$3,
	category_id=$4
	where id=$5`

	_, err := cm.db.Exec(updateQuery, cl.Name, cl.Image, cl.Description, cl.CategoryId, cl.Id)

	if err != nil {
		return "", err
	}
	return cl.Id, nil
}

func (cm *productRepo) Get(id string) (*repo.Product, error) {
	var product repo.Product

	query := `select 
		id,
		name,
		image,
		description,
		category_id
	from products where id=$1`

	row := cm.db.QueryRow(query, id)

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Image,
		&product.Description,
		&product.CategoryId,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (cm *productRepo) GetAll(name string) ([]*repo.Product, error) {
	var (
		products []*repo.Product
		filter   string
		args     = make(map[string]interface{})
	)
	if name != "" {
		filter += ` and name ilike '%' || :name || '%'`
		args["name"] = name
	}
	query := `select 
		id,
		name,
		image,
		description,
		category_id
	from products where 1=1 %s`

	rows, err := cm.db.NamedQuery(fmt.Sprintf(query, filter), args)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			product repo.Product
		)
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Image,
			&product.Description,
			&product.CategoryId,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (cm *productRepo) Delete(id string) error {
	return nil
}
