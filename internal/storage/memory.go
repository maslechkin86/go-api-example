package storage

import (
	"go-api-example/internal/types"
	"sync"
)

type MemoryStorage struct {
	mu     sync.Mutex
	users  map[int]*types.User
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:  make(map[int]*types.User),
		nextID: 1,
	}
}

func (s *MemoryStorage) Get(id int) (*types.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *MemoryStorage) Create(name string) (*types.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &types.User{
		ID:   s.nextID,
		Name: name,
	}
	s.users[user.ID] = user
	s.nextID++
	return user, nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return ErrUserNotFound
	}
	delete(s.users, id)
	return nil
}

func (s *MemoryStorage) GetAll(offset int, limit int) ([]*types.User, int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userList := make([]*types.User, 0, len(s.users))
	for _, user := range s.users {
		userList = append(userList, user)
	}

	start := offset
	if start > len(userList) {
		return []*types.User{}, 0, nil
	}

	end := start + limit
	if end > len(userList) {
		end = len(userList)
	}

	return userList[start:end], len(s.users), nil
}
