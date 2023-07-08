package teststore

import (
    "github.com/t3mnikov/gotes/internal/app/model"
    "github.com/t3mnikov/gotes/internal/app/store"
)

// User memory repository
type UserRepository struct {
    store *Store
    users map[int]*model.User
}

// Check User already exists
func (r UserRepository) EmailExists(email string) (bool, error) {
    for _, u := range r.users {
        if u.Email == email {
            return true, store.ErrEmailAlreadyExists
        }
    }

    return false, nil
}

// Create a User
func (r UserRepository) Create(u *model.User) error {
    err := u.Validate()
    if err != nil {
        return err
    }

    err = u.BeforeCreate()
    if err != nil {
        return err
    }

    u.ID = len(r.users) + 1
    r.users[u.ID] = u

    return nil
}

// FindByID Find User by ID
func (r UserRepository) FindByID(id int) (*model.User, error) {
    u, ok := r.users[id]
    if !ok {
        return nil, store.ErrRecordNotFound
    }

    return u, nil
}

// FindByEmail Find User by Email
func (r UserRepository) FindByEmail(email string) (*model.User, error) {
    for _, u := range r.users {
        if u.Email == email {
            return u, nil
        }
    }

    return nil, store.ErrRecordNotFound
}
