package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/Yangiboev/golang-postgres-monolith/storage/repo"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) repo.CategoryStorageI {
	return &categoryRepo{db: db}
}

func (cm *categoryRepo) Create(cl *repo.Category) (string, error) {
	if cl.ParentID == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		cl.ParentID = id.String()
	}
	insertQuery := `insert into categories(
		id,
		name,
		parent_id)
		values($1, $2, $3)`

	_, err := cm.db.Exec(insertQuery, cl.Id, cl.Name, cl.ParentID)

	if err != nil {
		return "", err
	}
	return cl.Id, nil
}

func (cm *categoryRepo) Update(cl *repo.Category) (string, error) {
	if cl.ParentID == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		cl.ParentID = id.String()
	}
	updateQuery := `update categories set
		name=$1, parent_id=$2 where id=$3`

	_, err := cm.db.Exec(updateQuery, cl.Name, cl.ParentID, cl.Id)

	if err != nil {
		return "", err
	}
	return cl.Id, nil
}

func (cm *categoryRepo) Get(id string) (*repo.Category, error) {
	var category repo.Category

	query := `select 
		id,
		name,
		parent_id
	from categories where id=$1`

	row := cm.db.QueryRow(query, id)

	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.ParentID,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (cm *categoryRepo) GetAll(name string) ([]*repo.Category, error) {
	var (
		categories []*repo.Category
		filter     string
		args       = make(map[string]interface{})
	)
	if name != "" {
		filter += ` and name ilike '%' || :name || '%'`
		args["name"] = name
	}
	query := `select 
		id,
		name,
		parent_id
	from categories where 1=1 %s`

	rows, err := cm.db.NamedQuery(fmt.Sprintf(query, filter), args)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			category repo.Category
		)
		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.ParentID,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (cm *categoryRepo) Delete(id string) error {
	return nil
}
