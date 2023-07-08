package teststore

import (
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store"
)

// Memory Store component
type Store struct {
    userRepository *UserRepository
    goteRepository *GoteRepository
}

// Constructor of Store
func New() *Store {
    return &Store{}
}

// User repo
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

// Gote repo
func (s *Store) Gote() store.GoteRepository {
    if s.goteRepository != nil {
        return s.goteRepository
    }

    s.goteRepository = &GoteRepository{
        store: s,
        gotes: make(map[int]*model.Gote),
    }

    return s.goteRepository
}
