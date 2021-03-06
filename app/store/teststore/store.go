package teststore

import (
	"github.com/mishon-pon-pon/http-rest-api/app/model"
	"github.com/mishon-pon-pon/http-rest-api/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
