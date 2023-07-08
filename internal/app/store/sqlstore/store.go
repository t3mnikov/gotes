package sqlstore

import (
    "database/sql"
    "github.com/t3mnikov/gotes/internal/app/store"
)

// DB Store component
type Store struct {
    db             *sql.DB
    userRepository *UserRepository
    goteRepository *GoteRepository
}

// Constructor of Store
func New(db *sql.DB) *Store {
    return &Store{
        db: db,
    }
}

// User repo
func (s *Store) User() store.UserRepository {
    if s.userRepository != nil {
        return s.userRepository
    }

    s.userRepository = &UserRepository{
        store: s,
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
    }

    return s.goteRepository
}
