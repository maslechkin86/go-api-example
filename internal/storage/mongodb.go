package storage

import (
	"go-api-example/internal/types"
)

type MongoStorage struct{}

func NewMongoStorage() *MongoStorage {
	return &MongoStorage{}
}

func (s *MongoStorage) Get(_id int) (*types.User, error) {
	// implement
	return nil, nil
}

func (s *MongoStorage) Create(name string) (*types.User, error) {
	// implement
	user := &types.User{
		ID:   0,
		Name: name,
	}
	return user, nil
}

func (s *MongoStorage) Delete(id int) error {
	// implement
	return nil
}

func (s *MongoStorage) GetAll(offset int, limit int) ([]*types.User, int, error) {
	// implement
	return nil, 0, nil
}
