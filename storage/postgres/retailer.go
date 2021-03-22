package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
)

type retailerRepo struct {
	db *sqlx.DB
}

func NewRetailerRepo(db *sqlx.DB) repo.RetailerStorageI {
	return &retailerRepo{db: db}
}

func (rr *retailerRepo) Create(cl *repo.Retailer) (string, error) {
	insertQuery := `insert into retailers(
		id,
		name,
		website,
		description)
		values($1, $2, $3, $4)`

	_, err := rr.db.Exec(insertQuery, cl.Id, cl.Name, cl.Website, cl.Description)

	if err != nil {
		return "", err
	}
	return cl.Id, nil
}

func (rr *retailerRepo) Update(cl *repo.Retailer) (string, error) {
	updateQuery := `update retailers set
	name=$1,
	website=$2,
	description=$3
	where id=$4`

	_, err := rr.db.Exec(updateQuery, cl.Name, cl.Website, cl.Description, cl.Id)

	if err != nil {
		return "", err
	}
	return cl.Id, nil
}

func (rr *retailerRepo) Get(id string) (*repo.Retailer, error) {
	var retailer repo.Retailer

	query := `select 
		id,
		name,
		website,
		description
	from retailers where id=$1`

	row := rr.db.QueryRow(query, id)

	err := row.Scan(
		&retailer.Id,
		&retailer.Name,
		&retailer.Website,
		&retailer.Description,
	)

	if err != nil {
		return nil, err
	}

	return &retailer, nil
}

func (rr *retailerRepo) GetAll(name string) ([]*repo.Retailer, error) {
	var (
		retailers []*repo.Retailer
		filter    string
		args      = make(map[string]interface{})
	)
	if name != "" {
		filter += ` and name ilike '%' || :name || '%'`
		args["name"] = name
	}
	query := `select 
		id,
		name,
		website,
		description
	from retailers where 1=1 %s`

	rows, err := rr.db.NamedQuery(fmt.Sprintf(query, filter), args)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			retailer repo.Retailer
		)
		err = rows.Scan(
			&retailer.Id,
			&retailer.Name,
			&retailer.Website,
			&retailer.Description,
		)
		if err != nil {
			return nil, err
		}
		retailers = append(retailers, &retailer)
	}

	return retailers, nil
}

func (rr *retailerRepo) Delete(id string) error {
	return nil
}
