package repo

import (
	"errors"
)

type Category struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

var (
	ErrAlreadyExists = errors.New("Already exists")
)

type CategoryStorageI interface {
	Create(*Category) (string, error)
	Update(*Category) (string, error)
	Get(id string) (*Category, error)
	GetAll(name string) ([]*Category, error)
	Delete(id string) error
}
