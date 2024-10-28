package storage

import (
	"go-api-example/internal/types"
)

type Storage interface {
	Get(ID int) (*types.User, error)
	GetAll(offset int, limit int) ([]*types.User, int, error)
	Create(name string) (*types.User, error)
	Delete(ID int) error
}
